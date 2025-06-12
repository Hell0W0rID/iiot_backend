//
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"encoding/xml"
	"time"

	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/models"

	"github.com/google/uuid"
)

type DataEvent struct {
	common.Versionable `json:",inline"`
	Id                 string        `json:"id" validate:"required,uuid"`
	DeviceName         string        `json:"deviceName" validate:"required,iiot-dto-none-empty-string"`
	ProfileName        string        `json:"profileName" validate:"required,iiot-dto-none-empty-string"`
	SourceName         string        `json:"sourceName" validate:"required,iiot-dto-none-empty-string"`
	Origin             int64         `json:"origin" validate:"required"`
	Measurements           []BaseMeasurement `json:"readings" validate:"gt=0,dive,required"`
	Tags               Tags          `json:"tags,omitempty"`
}

// NewDataEvent creates and returns an initialized DataEvent with no Measurements
func NewDataEvent(profileName, deviceName, sourceName string) DataEvent {
	return DataEvent{
		Versionable: common.NewVersionable(),
		Id:          uuid.NewString(),
		DeviceName:  deviceName,
		ProfileName: profileName,
		SourceName:  sourceName,
		Origin:      time.Now().UnixNano(),
	}
}

// FromDataEventModelToDTO transforms the DataEvent Model to the DataEvent DTO
func FromDataEventModelToDTO(event models.DataEvent) DataEvent {
	var readings []BaseMeasurement
	for _, reading := range event.Measurements {
		readings = append(readings, FromMeasurementModelToDTO(reading))
	}

	tags := make(map[string]interface{})
	for tag, value := range event.Tags {
		tags[tag] = value
	}

	return DataEvent{
		Versionable: common.NewVersionable(),
		Id:          event.Id,
		DeviceName:  event.DeviceName,
		ProfileName: event.ProfileName,
		SourceName:  event.SourceName,
		Origin:      event.Origin,
		Measurements:    readings,
		Tags:        tags,
	}
}

// AddSimpleMeasurement adds a simple reading to the DataEvent
func (e *DataEvent) AddSimpleMeasurement(resourceName string, valueType string, value interface{}) error {
	reading, err := NewSimpleMeasurement(e.ProfileName, e.DeviceName, resourceName, valueType, value)
	if err != nil {
		return err
	}
	e.Measurements = append(e.Measurements, reading)
	return nil
}

// AddBinaryMeasurement adds a binary reading to the DataEvent
func (e *DataEvent) AddBinaryMeasurement(resourceName string, binaryValue []byte, mediaType string) {
	e.Measurements = append(e.Measurements, NewBinaryMeasurement(e.ProfileName, e.DeviceName, resourceName, binaryValue, mediaType))
}

// AddObjectMeasurement adds a object reading to the DataEvent
func (e *DataEvent) AddObjectMeasurement(resourceName string, objectValue interface{}) {
	e.Measurements = append(e.Measurements, NewObjectMeasurement(e.ProfileName, e.DeviceName, resourceName, objectValue))
}

// AddNullMeasurement adds a simple reading with null value to the DataEvent
func (e *DataEvent) AddNullMeasurement(resourceName string, valueType string) {
	e.Measurements = append(e.Measurements, NewNullMeasurement(e.ProfileName, e.DeviceName, resourceName, valueType))
}

// ToXML provides a XML representation of the DataEvent as a string
func (e *DataEvent) ToXML() (string, error) {
	eventXml, err := xml.Marshal(e)
	if err != nil {
		return "", err
	}

	return string(eventXml), nil
}
