/*******************************************************************************
 *******************************************************************************/

package container

import (
	"iiot-backend/pkg/go-mod-configuration/configuration"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/interfaces"
	"iiot-backend/pkg/go-mod-bootstrap/di"
)

// ConfigurationInterfaceName contains the name of the interfaces.Configuration implementation in the DIC.
var ConfigurationInterfaceName = di.TypeInstanceToName((*interfaces.Configuration)(nil))

// ConfigurationFrom helper function queries the DIC and returns the interfaces.Configuration implementation.
func ConfigurationFrom(get di.Get) interfaces.Configuration {
	configuration, ok := get(ConfigurationInterfaceName).(interfaces.Configuration)
	if !ok {
		return nil
	}

	return configuration
}

// ConfigClientInterfaceName contains the name of the configuration.Client implementation in the DIC.
var ConfigClientInterfaceName = di.TypeInstanceToName((*configuration.Client)(nil))

// ConfigClientFrom helper function queries the DIC and returns the configuration.Client implementation.
func ConfigClientFrom(get di.Get) configuration.Client {
	client, ok := get(ConfigClientInterfaceName).(configuration.Client)
	if !ok {
		return nil
	}

	return client
}
