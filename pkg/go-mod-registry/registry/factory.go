//
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

package registry

import (
	"fmt"

	"iiot-backend/pkg/go-mod-registry/internal/pkg/keeper"
	"iiot-backend/pkg/go-mod-registry/pkg/types"
)

func NewRegistryClient(registryConfig types.Config) (Client, error) {

	if registryConfig.Host == "" || registryConfig.Port == 0 {
		return nil, fmt.Errorf("unable to create RegistryClient: registry host and/or port or serviceKey not set")
	}

	switch registryConfig.Type {
	case "keeper":
		registryClient, err := keeper.NewKeeperClient(registryConfig)
		return registryClient, err
	default:
		return nil, fmt.Errorf("unknown registry type '%s' requested", registryConfig.Type)
	}
}
