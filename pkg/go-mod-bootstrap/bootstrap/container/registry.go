/*******************************************************************************
 *******************************************************************************/

package container

import (
	"iiot-backend/pkg/go-mod-bootstrap/di"

	"iiot-backend/pkg/go-mod-registry/registry"
)

// RegistryClientInterfaceName contains the name of the registry.Client implementation in the DIC.
var RegistryClientInterfaceName = di.TypeInstanceToName((*registry.Client)(nil))

// RegistryFrom helper function queries the DIC and returns the registry.Client implementation.
func RegistryFrom(get di.Get) registry.Client {
	registryClient, ok := get(RegistryClientInterfaceName).(registry.Client)
	if !ok {
		return nil
	}

	return registryClient
}
