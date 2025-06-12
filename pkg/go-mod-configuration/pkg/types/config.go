//
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

package types

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
)

const DefaultProtocol = "http"

// ServiceConfig defines the information need to connect to the Configuration service and optionally register the service
// for discovery and health checks
type ServiceConfig struct {
	// The Protocol that should be used to connect to the Configuration service. HTTP is used if not set.
	Protocol string
	// Host is the hostname or IP address of the Configuration service
	Host string
	// Port is the HTTP port of the Configuration service
	Port int
	// Type is the implementation type of the Configuration service, i.e. keeper
	Type string
	// BasePath is the base path with in the Configuration service where the your service's configuration is stored
	BasePath string
	// AuthInjector is an interface to obtain a JWT and secure transport for remote service calls
	AuthInjector interfaces.AuthenticationInjector
	// Optional contains all other properties of the configuration provider might use.
	// For example, it might need the message bus connection information to publish the config changes.
	Optional map[string]any
}

//
// A few helper functions for building URLs.
//

func (config ServiceConfig) GetUrl() string {
	return fmt.Sprintf("%s://%s:%v", config.GetProtocol(), config.Host, config.Port)
}

func (config *ServiceConfig) GetProtocol() string {
	if config.Protocol == "" {
		return "http"
	}

	return config.Protocol
}

func (config *ServiceConfig) PopulateFromUrl(providerUrl string) error {
	url, err := url.Parse(providerUrl)
	if err != nil {
		return fmt.Errorf("the format of Provider URL is incorrect (%s): %s", providerUrl, err.Error())
	}

	port, err := strconv.Atoi(url.Port())
	if err != nil {
		return fmt.Errorf("the port from Provider URL is incorrect (%s): %s", providerUrl, err.Error())
	}

	config.Host = url.Hostname()
	config.Port = port

	typeAndProtocol := strings.Split(url.Scheme, ".")

	// TODO: Enforce both Type and Protocol present for release V2.0.0
	// Support for default protocol is for backwards compatibility with Fuji Device Services.
	switch len(typeAndProtocol) {
	case 1:
		config.Type = typeAndProtocol[0]
		config.Protocol = DefaultProtocol
	case 2:
		config.Type = typeAndProtocol[0]
		config.Protocol = typeAndProtocol[1]
	default:
		return fmt.Errorf("the Type and Protocol spec from Provider URL is incorrect: %s", providerUrl)
	}

	return nil
}
