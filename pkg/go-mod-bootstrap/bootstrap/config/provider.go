/*******************************************************************************
 *******************************************************************************/
package config

import (
	"iiot-backend/pkg/go-mod-configuration/pkg/types"
	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/environment"
)

// ProviderInfo encapsulates the usage of the Configuration Provider information
type ProviderInfo struct {
	serviceConfig types.ServiceConfig
}

// NewProviderInfo creates a new ProviderInfo and initializes it
func NewProviderInfo(envVars *environment.Variables, providerUrl string) (*ProviderInfo, error) {
	var err error
	configProviderInfo := ProviderInfo{}

	// initialize config provider configuration for URL set in commandline options
	if providerUrl != "" {
		if err = configProviderInfo.serviceConfig.PopulateFromUrl(providerUrl); err != nil {
			return nil, err
		}
	}

	// override file-based configuration with Variables variables.
	configProviderInfo.serviceConfig, err = envVars.OverrideConfigProviderInfo(configProviderInfo.serviceConfig)
	if err != nil {
		return nil, err
	}

	return &configProviderInfo, nil
}

// UseProvider returns whether the Configuration Provider should be used or not.
func (config *ProviderInfo) UseProvider() bool {
	return config.serviceConfig.Host != ""
}

// SetHost sets the host name for the Configuration Provider.
func (config *ProviderInfo) SetHost(host string) {
	config.serviceConfig.Host = host
}

// ServiceConfig returns service configuration for the Configuration Provider
func (config *ProviderInfo) ServiceConfig() types.ServiceConfig {
	return config.serviceConfig
}

// SetAuthInjector sets the Authentication Injector for the Configuration Provider
func (config *ProviderInfo) SetAuthInjector(authInjector interfaces.AuthenticationInjector) {
	config.serviceConfig.AuthInjector = authInjector
}
