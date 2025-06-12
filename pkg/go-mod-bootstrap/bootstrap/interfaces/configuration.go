/*******************************************************************************
 *******************************************************************************/

package interfaces

import "iiot-backend/pkg/go-mod-bootstrap/config"

// UpdatableConfig interface allows service to have their custom configuration populated from configuration stored
// in the Configuration Provider (aka keeper). A service using custom configuration must implement this interface
// on the custom configuration, even if not using Configuration Provider. If not using the Configuration Provider
// it can have dummy implementations of this interface.
type UpdatableConfig interface {
	// UpdateFromRaw converts configuration received from the Configuration Provider to a service-specific
	// configuration struct which is then used to overwrite the service's existing configuration struct.
	UpdateFromRaw(rawConfig interface{}) bool
}

// WritableConfig allows service to listen for changes from the Configuration Provider and have the configuration updated
// when the changes occur
type WritableConfig interface {
	// UpdateWritableFromRaw converts updated configuration received from the Configuration Provider to a
	// service-specific struct that is being watched for changes by the Configuration Provider.
	// The changes are used to overwrite the service's existing configuration's watched struct.
	UpdateWritableFromRaw(rawWritableConfig interface{}) bool
}

// Configuration interface provides an abstraction around a configuration struct.
type Configuration interface {
	// These two interfaces have been separated out for use in the custom configuration capability for
	// App and Device services
	UpdatableConfig
	WritableConfig

	// EmptyWritablePtr returns a pointer to a service-specific empty WritableInfo struct.  It is used by the bootstrap to
	// provide the appropriate structure to registry.Client's WatchForChanges().
	EmptyWritablePtr() interface{}

	// GetBootstrap returns the configuration elements required by the bootstrap.
	GetBootstrap() config.BootstrapConfiguration

	// GetLogLevel returns the current ConfigurationStruct's log level.
	GetLogLevel() string

	// GetRegistryInfo gets the config.RegistryInfo field from the ConfigurationStruct.
	GetRegistryInfo() config.RegistryInfo

	// GetInsecureSecrets gets the config.InsecureSecrets field from the ConfigurationStruct.
	GetInsecureSecrets() config.InsecureSecrets

	// GetTelemetryInfo gets the config.Telemetry section from the ConfigurationStruct
	GetTelemetryInfo() *config.TelemetryInfo

	// GetWritablePtr gets the config.WritablePtr section from the ConfigurationStruct
	GetWritablePtr() any
}
