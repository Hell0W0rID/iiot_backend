/*******************************************************************************
 *******************************************************************************/

package handlers

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"iiot-backend/pkg/go-mod-core-contracts/clients"
	httpClients "iiot-backend/pkg/go-mod-core-contracts/clients/http"
	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
	"iiot-backend/pkg/go-mod-core-contracts/clients/logger"
	"iiot-backend/pkg/go-mod-core-contracts/common"
	clientsMessaging "iiot-backend/pkg/go-mod-messaging/clients"
	"iiot-backend/pkg/go-mod-registry/pkg/types"
	"iiot-backend/pkg/go-mod-registry/registry"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/config"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/container"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/secret"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/startup"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/zerotrust"
	"iiot-backend/pkg/go-mod-bootstrap/di"
)

// ClientsBootstrap contains data to boostrap the configured clients
type ClientsBootstrap struct {
	registry registry.Client
}

// NewClientsBootstrap is a factory method that returns the initialized "ClientsBootstrap" receiver struct.
func NewClientsBootstrap() *ClientsBootstrap {
	return &ClientsBootstrap{}
}

// BootstrapHandler fulfills the BootstrapHandler contract.
// It creates instances of each of the IIOT clients that are in the service's configuration and place them in the DIC.
// If the registry is enabled it will be used to get the URL for client otherwise it will use configuration for the url.
// This handler will fail if an unknown client is specified.
func (cb *ClientsBootstrap) BootstrapHandler(
	_ context.Context,
	_ *sync.WaitGroup,
	startupTimer startup.Timer,
	dic *di.Container) bool {

	lc := container.LoggerClientFrom(dic.Get)
	cfg := container.ConfigurationFrom(dic.Get)
	cb.registry = container.RegistryFrom(dic.Get)

	if cfg.GetBootstrap().Clients != nil {
		for serviceKey, serviceInfo := range *cfg.GetBootstrap().Clients {
			var urlFunc clients.ClientBaseUrlFunc

			sp := container.SecretProviderExtFrom(dic.Get)
			jwtSecretProvider := secret.NewJWTSecretProvider(sp)
			if serviceInfo.SecurityOptions[config.SecurityModeKey] == zerotrust.ZeroTrustMode {
				sp.EnableZeroTrust()
			}
			if rt, transpErr := zerotrust.HttpTransportFromClient(sp, serviceInfo, lc); transpErr != nil {
				lc.Errorf("could not obtain an http client for use with zero trust provider: %v", transpErr)
				return false
			} else {
				sp.SetHttpTransport(rt) //only need to set the transport when using SecretProviderExt
				sp.SetFallbackDialer(&net.Dialer{})
			}

			if !serviceInfo.UseMessageBus {
				mode := container.DevRemoteModeFrom(dic.Get)
				if cb.registry == nil || mode.InDevMode || mode.InRemoteMode {
					lc.Infof("Using REST for '%s' clients @ %s", serviceKey, serviceInfo.Url())
					urlFunc = clients.GetDefaultClientBaseUrlFunc(serviceInfo.Url())
				} else {
					lc.Infof("Using ClientBaseUrlFunc for '%s' clients", serviceKey)
					urlFunc = cb.clientUrlFunc(serviceKey, lc)
				}
			}

			switch serviceKey {
			case common.CoreDataServiceName:
				dic.Update(di.ServiceConstructorMap{
					container.DataEventClientName: func(get di.Get) interface{} {
						return httpClients.NewDataEventClientWithUrlCallback(urlFunc, jwtSecretProvider, cfg.GetBootstrap().Service.EnableNameFieldEscape)
					},
					container.MeasurementClientName: func(get di.Get) interface{} {
						return httpClients.NewMeasurementClientWithUrlCallback(urlFunc, jwtSecretProvider, cfg.GetBootstrap().Service.EnableNameFieldEscape)
					},
				})
			case common.CoreMetaDataServiceName:
				dic.Update(di.ServiceConstructorMap{
					container.DeviceClientName: func(get di.Get) interface{} {
						return httpClients.NewDeviceClientWithUrlCallback(urlFunc, jwtSecretProvider, cfg.GetBootstrap().Service.EnableNameFieldEscape)
					},
					container.DeviceHandlerClientName: func(get di.Get) interface{} {
						return httpClients.NewDeviceHandlerClientWithUrlCallback(urlFunc, jwtSecretProvider, cfg.GetBootstrap().Service.EnableNameFieldEscape)
					},
					container.DeviceTemplateClientName: func(get di.Get) interface{} {
						return httpClients.NewDeviceTemplateClientWithUrlCallback(urlFunc, jwtSecretProvider, cfg.GetBootstrap().Service.EnableNameFieldEscape)
					},
					container.DeviceWatcherClientName: func(get di.Get) interface{} {
						return httpClients.NewDeviceWatcherClientWithUrlCallback(urlFunc, jwtSecretProvider, cfg.GetBootstrap().Service.EnableNameFieldEscape)
					},
				})

			case common.CoreCommandServiceName:
				var client interfaces.CommandClient

				if serviceInfo.UseMessageBus {
					// TODO: Move following outside loop when multiple messaging based clients exist
					messageClient := container.MessagingClientFrom(dic.Get)
					if messageClient == nil {
						lc.Errorf("Unable to create Command client using MessageBus: %s", "MessageBus Client was not created")
						return false
					}

					if len(cfg.GetBootstrap().Service.RequestTimeout) == 0 {
						lc.Error("Service.RequestTimeout found empty in service's configuration, missing common config? Use -cp or -cc flags for common config")
						return false
					}

					// TODO: Move following outside loop when multiple messaging based clients exist
					timeout, err := time.ParseDuration(cfg.GetBootstrap().Service.RequestTimeout)
					if err != nil {
						lc.Errorf("Unable to parse Service.RequestTimeout as a time duration: %v", err)
						return false
					}

					baseTopic := cfg.GetBootstrap().MessageBus.GetBaseTopicPrefix()
					if cfg.GetBootstrap().Service.EnableNameFieldEscape {
						client = clientsMessaging.NewCommandClientWithNameFieldEscape(messageClient, baseTopic, timeout)
					} else {
						client = clientsMessaging.NewCommandClient(messageClient, baseTopic, timeout)
					}

					lc.Infof("Using messaging for '%s' clients", serviceKey)
				} else {
					client = httpClients.NewCommandClientWithUrlCallback(urlFunc, jwtSecretProvider, cfg.GetBootstrap().Service.EnableNameFieldEscape)
				}

				dic.Update(di.ServiceConstructorMap{
					container.CommandClientName: func(get di.Get) interface{} {
						return client
					},
				})

			case common.SupportAlertsServiceName:
				dic.Update(di.ServiceConstructorMap{
					container.AlertClientName: func(get di.Get) interface{} {
						return httpClients.NewAlertClientWithUrlCallback(urlFunc, jwtSecretProvider, cfg.GetBootstrap().Service.EnableNameFieldEscape)
					},
					container.EventSubscriptionClientName: func(get di.Get) interface{} {
						return httpClients.NewEventSubscriptionClientWithUrlCallback(urlFunc, jwtSecretProvider, cfg.GetBootstrap().Service.EnableNameFieldEscape)
					},
				})

			case common.SupportSchedulerServiceName:
				dic.Update(di.ServiceConstructorMap{
					container.ScheduleJobClientName: func(get di.Get) interface{} {
						return httpClients.NewScheduleJobClientWithUrlCallback(urlFunc, jwtSecretProvider, cfg.GetBootstrap().Service.EnableNameFieldEscape)
					},
					container.ScheduleActionRecordClientName: func(get di.Get) interface{} {
						return httpClients.NewScheduleActionRecordClientWithUrlCallback(urlFunc, jwtSecretProvider, cfg.GetBootstrap().Service.EnableNameFieldEscape)
					},
				})

			case common.SecurityProxyAuthServiceName:
				dic.Update(di.ServiceConstructorMap{
					container.SecurityProxyAuthClientName: func(get di.Get) interface{} {
						return httpClients.NewAuthClientWithUrlCallback(urlFunc, jwtSecretProvider)
					},
				})

			default:

			}
		}
	}
	return true
}

func (cb *ClientsBootstrap) clientUrlFunc(serviceKey string, lc logger.LoggerClient) clients.ClientBaseUrlFunc {
	return func() (string, error) {
		var err error
		var endpoint types.ServiceEndpoint

		endpoint, err = cb.registry.GetServiceEndpoint(serviceKey)
		if err != nil {
			return "", fmt.Errorf("unable to Get service endpoint for '%s': %s", serviceKey, err.Error())
		}

		url := fmt.Sprintf("http://%s:%v", endpoint.Host, endpoint.Port)

		lc.Infof("Using registry for URL for '%s': %s", serviceKey, url)

		return url, nil
	}
}
