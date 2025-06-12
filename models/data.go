package models

import (
	"time"
)

// Event represents an event from a device
type Event struct {
	ID          string    `json:"id" db:"id"`
	DeviceName  string    `json:"deviceName" db:"device_name"`
	ProfileName string    `json:"profileName" db:"profile_name"`
	SourceName  string    `json:"sourceName" db:"source_name"`
	Origin      int64     `json:"origin" db:"origin"`
	Tags        map[string]string `json:"tags" db:"tags"`
	Readings    []Reading `json:"readings" db:"-"`
	Created     time.Time `json:"created" db:"created"`
	Modified    time.Time `json:"modified" db:"modified"`
}

// Reading represents a reading from a device
type Reading struct {
	ID           string    `json:"id" db:"id"`
	EventID      string    `json:"eventId" db:"event_id"`
	DeviceName   string    `json:"deviceName" db:"device_name"`
	ResourceName string    `json:"resourceName" db:"resource_name"`
	ProfileName  string    `json:"profileName" db:"profile_name"`
	ValueType    string    `json:"valueType" db:"value_type"`
	Value        string    `json:"value" db:"value"`
	BinaryValue  []byte    `json:"binaryValue" db:"binary_value"`
	MediaType    string    `json:"mediaType" db:"media_type"`
	Units        string    `json:"units" db:"units"`
	Tags         map[string]string `json:"tags" db:"tags"`
	Origin       int64     `json:"origin" db:"origin"`
	Created      time.Time `json:"created" db:"created"`
	Modified     time.Time `json:"modified" db:"modified"`
}

// EventRequest represents a request to create an event
type EventRequest struct {
	DeviceName  string    `json:"deviceName" validate:"required"`
	ProfileName string    `json:"profileName" validate:"required"`
	SourceName  string    `json:"sourceName" validate:"required"`
	Origin      int64     `json:"origin"`
	Tags        map[string]string `json:"tags"`
	Readings    []ReadingRequest `json:"readings" validate:"required,min=1"`
}

// ReadingRequest represents a request to create a reading
type ReadingRequest struct {
	DeviceName   string    `json:"deviceName" validate:"required"`
	ResourceName string    `json:"resourceName" validate:"required"`
	ProfileName  string    `json:"profileName" validate:"required"`
	ValueType    string    `json:"valueType" validate:"required"`
	Value        string    `json:"value"`
	BinaryValue  []byte    `json:"binaryValue"`
	MediaType    string    `json:"mediaType"`
	Units        string    `json:"units"`
	Tags         map[string]string `json:"tags"`
	Origin       int64     `json:"origin"`
}

// EventFilter represents filters for querying events
type EventFilter struct {
	DeviceName  string    `json:"deviceName"`
	ProfileName string    `json:"profileName"`
	SourceName  string    `json:"sourceName"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Limit       int       `json:"limit"`
	Offset      int       `json:"offset"`
}

// ReadingFilter represents filters for querying readings
type ReadingFilter struct {
	DeviceName   string    `json:"deviceName"`
	ResourceName string    `json:"resourceName"`
	ProfileName  string    `json:"profileName"`
	ValueType    string    `json:"valueType"`
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	Limit        int       `json:"limit"`
	Offset       int       `json:"offset"`
}
