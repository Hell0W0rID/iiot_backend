//
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

package registry

import (
	"iiot-backend/pkg/go-mod-registry/pkg/types"
)

type Client interface {
	// Registers the current service with Registry for discover and health check
	Register() error

	// Un-registers the current service with Registry for discover and health check
	Unregister() error

	// Registers a
	RegisterCheck(id string, name string, notes string, url string, interval string) error

	// Simply checks if Registry is up and running at the configured URL
	IsAlive() bool

	// Gets the service endpoint information for the target ID from the Registry
	GetServiceEndpoint(serviceId string) (types.ServiceEndpoint, error)

	// Gets all the service endpoints information from the Registry
	GetAllServiceEndpoints() ([]types.ServiceEndpoint, error)

	// Checks with the Registry if the target service is available, i.e. registered and healthy
	IsServiceAvailable(serviceId string) (bool, error)
}
