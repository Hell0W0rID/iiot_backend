//
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

package configuration

import (
	"fmt"

	"iiot-backend/pkg/go-mod-configuration/internal/pkg/keeper"
	"iiot-backend/pkg/go-mod-configuration/pkg/types"
)

func NewConfigurationClient(config types.ServiceConfig) (Client, error) {

	if config.Host == "" || config.Port == 0 {
		return nil, fmt.Errorf("unable to create Configuration Client: Configuration service host and/or port or serviceKey not set")
	}

	switch config.Type {
	case "keeper":
		client := keeper.NewKeeperClient(config)
		return client, nil
	default:
		return nil, fmt.Errorf("unknown configuration client type '%s' requested", config.Type)
	}
}
