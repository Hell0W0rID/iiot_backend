package models

import (
	"time"
)

// DeviceService represents a device service in the system
type DeviceService struct {
	ID          string            `json:"id" db:"id"`
	Name        string            `json:"name" db:"name"`
	Description string            `json:"description" db:"description"`
	BaseAddress string            `json:"baseAddress" db:"base_address"`
	AdminState  string            `json:"adminState" db:"admin_state"`
	Labels      []string          `json:"labels" db:"labels"`
	Created     time.Time         `json:"created" db:"created"`
	Modified    time.Time         `json:"modified" db:"modified"`
}

// Device represents a device in the system
type Device struct {
	ID             string            `json:"id" db:"id"`
	Name           string            `json:"name" db:"name"`
	Description    string            `json:"description" db:"description"`
	AdminState     string            `json:"adminState" db:"admin_state"`
	OperatingState string            `json:"operatingState" db:"operating_state"`
	Protocols      map[string]interface{} `json:"protocols" db:"protocols"`
	Labels         []string          `json:"labels" db:"labels"`
	Location       string            `json:"location" db:"location"`
	ServiceName    string            `json:"serviceName" db:"service_name"`
	ProfileName    string            `json:"profileName" db:"profile_name"`
	AutoEvents     []AutoEvent       `json:"autoEvents" db:"auto_events"`
	Tags           map[string]string `json:"tags" db:"tags"`
	Properties     map[string]string `json:"properties" db:"properties"`
	Created        time.Time         `json:"created" db:"created"`
	Modified       time.Time         `json:"modified" db:"modified"`
}

// DeviceProfile represents a device profile
type DeviceProfile struct {
	ID                string                 `json:"id" db:"id"`
	Name              string                 `json:"name" db:"name"`
	Description       string                 `json:"description" db:"description"`
	Manufacturer      string                 `json:"manufacturer" db:"manufacturer"`
	Model             string                 `json:"model" db:"model"`
	Labels            []string               `json:"labels" db:"labels"`
	DeviceResources   []DeviceResource       `json:"deviceResources" db:"device_resources"`
	DeviceCommands    []DeviceCommand        `json:"deviceCommands" db:"device_commands"`
	CoreCommands      []CoreCommand          `json:"coreCommands" db:"core_commands"`
	Created           time.Time              `json:"created" db:"created"`
	Modified          time.Time              `json:"modified" db:"modified"`
}

// DeviceResource represents a device resource
type DeviceResource struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Tag         string                 `json:"tag"`
	Properties  ResourceProperties     `json:"properties"`
	Attributes  map[string]interface{} `json:"attributes"`
}

// ResourceProperties represents resource properties
type ResourceProperties struct {
	ValueType    string `json:"valueType"`
	ReadWrite    string `json:"readWrite"`
	Units        string `json:"units"`
	Minimum      string `json:"minimum"`
	Maximum      string `json:"maximum"`
	DefaultValue string `json:"defaultValue"`
	Mask         string `json:"mask"`
	Shift        string `json:"shift"`
	Scale        string `json:"scale"`
	Offset       string `json:"offset"`
	Base         string `json:"base"`
	Assertion    string `json:"assertion"`
	MediaType    string `json:"mediaType"`
}

// DeviceCommand represents a device command
type DeviceCommand struct {
	Name               string                    `json:"name"`
	IsHidden           bool                      `json:"isHidden"`
	ReadWrite          string                    `json:"readWrite"`
	ResourceOperations []ResourceOperation       `json:"resourceOperations"`
}

// ResourceOperation represents a resource operation
type ResourceOperation struct {
	DeviceResource string                 `json:"deviceResource"`
	DefaultValue   string                 `json:"defaultValue"`
	Mappings       map[string]string      `json:"mappings"`
}

// CoreCommand represents a core command
type CoreCommand struct {
	Name      string `json:"name"`
	Get       bool   `json:"get"`
	Set       bool   `json:"set"`
	Path      string `json:"path"`
	URL       string `json:"url"`
	Parameters []Parameter `json:"parameters"`
}

// Parameter represents a command parameter
type Parameter struct {
	ResourceName   string `json:"resourceName"`
	ValueType      string `json:"valueType"`
}

// AutoEvent represents an auto event configuration
type AutoEvent struct {
	Interval   string `json:"interval"`
	OnChange   bool   `json:"onChange"`
	SourceName string `json:"sourceName"`
}

// ProvisionWatcher represents a provision watcher
type ProvisionWatcher struct {
	ID                  string            `json:"id" db:"id"`
	Name                string            `json:"name" db:"name"`
	Labels              []string          `json:"labels" db:"labels"`
	Identifiers         map[string]string `json:"identifiers" db:"identifiers"`
	BlockingIdentifiers map[string][]string `json:"blockingIdentifiers" db:"blocking_identifiers"`
	ProfileName         string            `json:"profileName" db:"profile_name"`
	ServiceName         string            `json:"serviceName" db:"service_name"`
	AdminState          string            `json:"adminState" db:"admin_state"`
	AutoEvents          []AutoEvent       `json:"autoEvents" db:"auto_events"`
	Created             time.Time         `json:"created" db:"created"`
	Modified            time.Time         `json:"modified" db:"modified"`
}
