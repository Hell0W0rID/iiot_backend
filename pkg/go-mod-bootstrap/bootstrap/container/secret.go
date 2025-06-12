/*******************************************************************************
 *******************************************************************************/

package container

import (
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/interfaces"
	"iiot-backend/pkg/go-mod-bootstrap/di"
)

// SecretProviderName contains the name of the interfaces.SecretProvider implementation in the DIC.
var SecretProviderName = di.TypeInstanceToName((*interfaces.SecretProvider)(nil))

// SecretProviderFrom helper function queries the DIC and returns the interfaces.SecretProvider
// implementation.
func SecretProviderFrom(get di.Get) interfaces.SecretProvider {
	provider, ok := get(SecretProviderName).(interfaces.SecretProvider)
	if !ok {
		return nil
	}

	return provider
}

// SecretProviderExtName contains the name of the interfaces.SecretProviderExt implementation in the DIC.
var SecretProviderExtName = di.TypeInstanceToName((*interfaces.SecretProvider)(nil))

// SecretProviderExtFrom helper function queries the DIC and returns the interfaces.SecretProviderExt
// implementation.
func SecretProviderExtFrom(get di.Get) interfaces.SecretProviderExt {
	provider, ok := get(SecretProviderExtName).(interfaces.SecretProviderExt)
	if !ok {
		return nil
	}

	return provider
}
