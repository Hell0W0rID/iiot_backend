/*******************************************************************************
 *******************************************************************************/

package pkg

import (
	"iiot-backend/pkg/go-mod-core-contracts/common"
)

// Defines the valid secret store providers.
const (
	CoreSecurityServiceName         = "iiot-core-security"
	ConfigFileName                 = "configuration.toml"
	ConfigDirectory                = "./res"
	ConfigDirEnv                   = "IIOT_CONF_DIR"
	SpiffeTokenProviderGetTokenAPI = common.ApiBase + "/gettoken" // nolint: gosec
)
