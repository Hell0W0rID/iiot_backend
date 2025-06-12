/*******************************************************************************
 *******************************************************************************/

package container

import (
	"iiot-backend/pkg/go-mod-secrets/pkg/token/runtimetokenprovider"

	"iiot-backend/pkg/go-mod-bootstrap/di"
)

// RuntimeTokenProviderInterfaceName contains the name of the runtimetokenprovider.RuntimeTokenProvider implementation in the DIC.
var RuntimeTokenProviderInterfaceName = di.TypeInstanceToName((*runtimetokenprovider.RuntimeTokenProvider)(nil))

// RuntimeTokenProviderFrom helper function queries the DIC and returns the runtimetokenprovider.RuntimeTokenProvider implementation.
func RuntimeTokenProviderFrom(get di.Get) runtimetokenprovider.RuntimeTokenProvider {
	runtimeTokenProvider, ok := get(RuntimeTokenProviderInterfaceName).(runtimetokenprovider.RuntimeTokenProvider)
	if !ok {
		return nil
	}

	return runtimeTokenProvider
}
