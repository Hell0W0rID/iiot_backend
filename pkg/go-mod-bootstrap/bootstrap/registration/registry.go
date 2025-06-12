/*******************************************************************************
 *******************************************************************************/

package registration

import (
	"context"
	"errors"
	"fmt"

	"iiot-backend/pkg/go-mod-core-contracts/clients/logger"
	"iiot-backend/pkg/go-mod-core-contracts/common"

	registryTypes "iiot-backend/pkg/go-mod-registry/pkg/types"
	"iiot-backend/pkg/go-mod-registry/registry"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/container"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/interfaces"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/secret"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/startup"
	"iiot-backend/pkg/go-mod-bootstrap/config"
	"iiot-backend/pkg/go-mod-bootstrap/di"
)

// createRegistryClient creates and returns a registry.Client instance.
func createRegistryClient(
	serviceKey string,
	serviceConfig interfaces.Configuration,
	lc logger.LoggerClient,
	dic *di.Container) (registry.Client, error) {
	bootstrapConfig := serviceConfig.GetBootstrap()
	secretProvider := container.SecretProviderExtFrom(dic.Get)

	if len(bootstrapConfig.Registry.Host) == 0 || bootstrapConfig.Registry.Port == 0 || len(bootstrapConfig.Registry.Type) == 0 {
		return nil, errors.New("Registry configuration is empty or incomplete, missing common config? Use -cp or -cc flags for common config.")
	}

	registryConfig := registryTypes.Config{
		Host:            bootstrapConfig.Registry.Host,
		Port:            bootstrapConfig.Registry.Port,
		Type:            bootstrapConfig.Registry.Type,
		ServiceName:      serviceKey,
		ServiceHost:     bootstrapConfig.Service.Host,
		ServicePort:     bootstrapConfig.Service.Port,
		ServiceProtocol: config.DefaultHttpProtocol,
		CheckInterval:   bootstrapConfig.Service.HealthCheckInterval,
		CheckRoute:      common.ApiPingRoute,
		AuthInjector:    secret.NewJWTSecretProvider(secretProvider),
	}

	lc.Info(fmt.Sprintf("Using Registry (%s) from %s", registryConfig.Type, registryConfig.GetRegistryUrl()))

	return registry.NewRegistryClient(registryConfig)
}

// RegisterWithRegistry connects to the registry and registers the service with the Registry
func RegisterWithRegistry(
	ctx context.Context,
	startupTimer startup.Timer,
	config interfaces.Configuration,
	lc logger.LoggerClient,
	serviceKey string,
	dic *di.Container) (registry.Client, error) {

	var registryWithRegistry = func(registryClient registry.Client) error {
		if !registryClient.IsAlive() {
			return errors.New("registry is not available")
		}

		if err := registryClient.Register(); err != nil {
			return fmt.Errorf("could not register service with Registry: %v", err.Error())
		}

		return nil
	}

	registryClient, err := createRegistryClient(serviceKey, config, lc, dic)
	if err != nil {
		return nil, fmt.Errorf("createRegistryClient failed: %v", err.Error())
	}

	for startupTimer.HasNotElapsed() {
		if err := registryWithRegistry(registryClient); err != nil {
			lc.Warn(err.Error())
			select {
			case <-ctx.Done():
				return nil, errors.New("aborted RegisterWithRegistry()")
			default:
				startupTimer.SleepForInterval()
				continue
			}
		}
		return registryClient, nil
	}
	return nil, errors.New("unable to register with Registry in allotted time")
}
