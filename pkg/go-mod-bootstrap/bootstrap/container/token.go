/*******************************************************************************
 *******************************************************************************/

package container

import (
	"iiot-backend/pkg/go-mod-secrets/pkg/token/authtokenloader"

	"iiot-backend/pkg/go-mod-bootstrap/di"
)

//// FileIoPerformerInterfaceName contains the name of the fileioperformer.FileIoPerformer implementation in the DIC.
//var FileIoPerformerInterfaceName = di.TypeInstanceToName((*fileioperformer.FileIoPerformer)(nil))
//
//// FileIoPerformerFrom helper function queries the DIC and returns the fileioperformer.FileIoPerformer implementation.
//func FileIoPerformerFrom(get di.Get) fileioperformer.FileIoPerformer {
//	fileIo := get(FileIoPerformerInterfaceName)
//	if fileIo != nil {
//		return fileIo.(fileioperformer.FileIoPerformer)
//	}
//	return (fileioperformer.FileIoPerformer)(nil)
//}

// AuthTokenLoaderInterfaceName contains the name of the authtokenloader.AuthTokenLoader implementation in the DIC.
var AuthTokenLoaderInterfaceName = di.TypeInstanceToName((*authtokenloader.AuthTokenLoader)(nil))

// AuthTokenLoaderFrom helper function queries the DIC and returns the authtokenloader.AuthTokenLoader implementation.
func AuthTokenLoaderFrom(get di.Get) authtokenloader.AuthTokenLoader {
	loader, ok := get(AuthTokenLoaderInterfaceName).(authtokenloader.AuthTokenLoader)
	if !ok {
		return nil
	}

	return loader
}
