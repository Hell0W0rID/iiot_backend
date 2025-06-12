/*******************************************************************************
 *******************************************************************************/

package handlers

import (
	"context"
	"strings"
	"sync"

	"iiot-backend/pkg/go-mod-messaging/messaging"
	"iiot-backend/pkg/go-mod-messaging/pkg/types"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/container"
	boostrapMessaging "iiot-backend/pkg/go-mod-bootstrap/bootstrap/messaging"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/startup"
	"iiot-backend/pkg/go-mod-bootstrap/config"
	"iiot-backend/pkg/go-mod-bootstrap/di"
)

// MessagingBootstrapHandler fulfills the BootstrapHandler contract.  If creates and initializes the Messaging client
// and adds it to the DIC
func MessagingBootstrapHandler(ctx context.Context, wg *sync.WaitGroup, startupTimer startup.Timer, dic *di.Container) bool {
	lc := container.LoggerClientFrom(dic.Get)
	configuration := container.ConfigurationFrom(dic.Get)

	messageBus := configuration.GetBootstrap().MessageBus
	if messageBus.Disabled {
		lc.Info("MessageBus is disabled in configuration, skipping setup.")
		return true
	}

	if len(messageBus.Host) == 0 || messageBus.Port == 0 || len(messageBus.Protocol) == 0 || len(messageBus.Type) == 0 {
		lc.Error("MessageBus configuration is incomplete, missing common config? Use -cp or -cc flags for common config.")
		return false
	}

	// Make sure the MessageBus password is not leaked into the Service Config that can be retrieved via the /config endpoint
	messageBusInfo := deepCopy(*messageBus)

	if len(messageBusInfo.AuthMode) > 0 &&
		!strings.EqualFold(strings.TrimSpace(messageBusInfo.AuthMode), boostrapMessaging.AuthModeNone) {
		if err := boostrapMessaging.SetOptionsAuthData(&messageBusInfo, lc, dic); err != nil {
			lc.Errorf("setting the MessageBus auth options failed: %v", err)
			return false
		}
	}

	msgClient, err := messaging.NewMessageClient(
		types.MessageBusConfig{
			Broker: types.HostInfo{
				Host:     messageBusInfo.Host,
				Port:     messageBusInfo.Port,
				Protocol: messageBusInfo.Protocol,
			},
			Type:     messageBusInfo.Type,
			Optional: messageBusInfo.Optional,
		})

	if err != nil {
		lc.Errorf("Failed to create MessageClient: %v", err)
		return false
	}

	for startupTimer.HasNotElapsed() {
		select {
		case <-ctx.Done():
			return false
		default:
			err = msgClient.Connect()
			if err != nil {
				lc.Warnf("Unable to connect MessageBus: %s", err.Error())
				startupTimer.SleepForInterval()
				continue
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				<-ctx.Done()
				if msgClient != nil {
					_ = msgClient.Disconnect()
				}
				lc.Infof("Disconnected from MessageBus")
			}()

			dic.Update(di.ServiceConstructorMap{
				container.MessagingClientName: func(get di.Get) interface{} {
					return msgClient
				},
			})

			lc.Infof(
				"Connected to %s Message Bus @ %s://%s:%d with AuthMode='%s'",
				messageBusInfo.Type,
				messageBusInfo.Protocol,
				messageBusInfo.Host,
				messageBusInfo.Port,
				messageBusInfo.AuthMode)

			return true
		}
	}

	lc.Error("Connecting to MessageBus time out")
	return false
}

func deepCopy(target config.MessageBusInfo) config.MessageBusInfo {
	result := config.MessageBusInfo{
		Disabled:        target.Disabled,
		Type:            target.Type,
		Protocol:        target.Protocol,
		Host:            target.Host,
		Port:            target.Port,
		AuthMode:        target.AuthMode,
		SecretName:      target.SecretName,
		BaseTopicPrefix: target.BaseTopicPrefix,
	}

	result.Optional = make(map[string]string)
	for key, value := range target.Optional {
		result.Optional[key] = value
	}

	return result
}
