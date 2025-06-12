/*******************************************************************************
 *******************************************************************************/

package container

import (
	"iiot-backend/pkg/go-mod-core-contracts/clients/logger"

	"iiot-backend/pkg/go-mod-bootstrap/di"
)

// LoggerClientInterfaceName contains the name of the logger.LoggerClient implementation in the DIC.
var LoggerClientInterfaceName = di.TypeInstanceToName((*logger.LoggerClient)(nil))

// LoggerClientFrom helper function queries the DIC and returns the logger.loggingClient implementation.
func LoggerClientFrom(get di.Get) logger.LoggerClient {
	loggingClient, ok := get(LoggerClientInterfaceName).(logger.LoggerClient)
	if !ok {
		return nil
	}

	return loggingClient
}
