/*******************************************************************************
 *******************************************************************************/

package container

import (
	"iiot-backend/pkg/go-mod-bootstrap/di"
)

type DevRemoteMode struct {
	InDevMode    bool
	InRemoteMode bool
}

// DevRemoteModeName contains the name of the DevRemoteMode struct in the DIC.
var DevRemoteModeName = di.TypeInstanceToName((*DevRemoteMode)(nil))

// DevRemoteModeFrom helper function queries the DIC and returns the Dev and Remotes mode flags.
func DevRemoteModeFrom(get di.Get) DevRemoteMode {
	devOrRemoteMode, ok := get(DevRemoteModeName).(*DevRemoteMode)
	if !ok {
		return DevRemoteMode{}
	}

	return *devOrRemoteMode
}
