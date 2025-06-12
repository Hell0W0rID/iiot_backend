package container

import (
	"github.com/edgexfoundry/go-mod-bootstrap/v3/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"

	"iiot-backend/pkg/common"
)

// ConfigurationName contains the name of data service's config structure
var ConfigurationName = common.ConfigStemCore + common.CoreDataServiceKey

// Container contains the dependencies for the data service
var Container = &di.Container{}

// ConfigurationFrom helper function to get the Configuration from the DIC
func ConfigurationFrom(get di.Get) *common.ConfigurationStruct {
	config, ok := get(ConfigurationName).(*common.ConfigurationStruct)
	if !ok {
		return &common.ConfigurationStruct{}
	}
	return config
}

// LoggingClientFrom helper function to get the logger from the DIC
func LoggingClientFrom(get di.Get) logger.LoggingClient {
	lc, ok := get(di.LoggingClientInterfaceName).(logger.LoggingClient)
	if !ok {
		return logger.NewMockClient()
	}
	return lc
}