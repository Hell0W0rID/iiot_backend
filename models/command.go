package models

import (
	"time"
)

// CommandRequest represents a command request to a device
type CommandRequest struct {
	DeviceID    string                 `json:"deviceId" validate:"required"`
	CommandName string                 `json:"commandName" validate:"required"`
	Parameters  map[string]interface{} `json:"parameters"`
	PushEvent   bool                   `json:"pushEvent"`
	ReturnEvent bool                   `json:"returnEvent"`
}

// CommandResponse represents the response from a device command
type CommandResponse struct {
	DeviceID    string                 `json:"deviceId"`
	CommandName string                 `json:"commandName"`
	StatusCode  int                    `json:"statusCode"`
	Event       *Event                 `json:"event,omitempty"`
	Response    map[string]interface{} `json:"response,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Timestamp   time.Time              `json:"timestamp"`
}

// DeviceCommandInfo represents command information for a device
type DeviceCommandInfo struct {
	DeviceID     string        `json:"deviceId"`
	DeviceName   string        `json:"deviceName"`
	ProfileName  string        `json:"profileName"`
	Commands     []CommandInfo `json:"commands"`
}

// CommandInfo represents information about a specific command
type CommandInfo struct {
	Name      string `json:"name"`
	Get       bool   `json:"get"`
	Set       bool   `json:"set"`
	Path      string `json:"path"`
	URL       string `json:"url"`
	Parameters []Parameter `json:"parameters"`
}

// SetCommandRequest represents a SET command request
type SetCommandRequest struct {
	DeviceID    string                 `json:"deviceId" validate:"required"`
	CommandName string                 `json:"commandName" validate:"required"`
	Settings    map[string]interface{} `json:"settings" validate:"required"`
}

// GetCommandRequest represents a GET command request
type GetCommandRequest struct {
	DeviceID    string `json:"deviceId" validate:"required"`
	CommandName string `json:"commandName" validate:"required"`
	DS_PushEvent bool  `json:"ds-pushevent"`
	DS_ReturnEvent bool `json:"ds-returnevent"`
}
