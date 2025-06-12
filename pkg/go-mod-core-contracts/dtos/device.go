//
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"iiot-backend/pkg/go-mod-core-contracts/models"
)

type Device struct {
	DBTimestamp    `json:",inline"`
	Id             string                        `json:"id,omitempty" yaml:"id,omitempty" validate:"omitempty,uuid"`
	Name           string                        `json:"name" yaml:"name" validate:"required,iiot-dto-none-empty-string"`
	Parent         string                        `json:"parent,omitempty" yaml:"parent,omitempty"`
	Description    string                        `json:"description,omitempty" yaml:"description,omitempty"`
	ServiceState     string                        `json:"adminState" yaml:"adminState" validate:"oneof='LOCKED' 'UNLOCKED'"`
	OperatingState string                        `json:"operatingState" yaml:"operatingState" validate:"oneof='UP' 'DOWN' 'UNKNOWN'"`
	Labels         []string                      `json:"labels,omitempty" yaml:"labels,omitempty"`
	Location       interface{}                   `json:"location,omitempty" yaml:"location,omitempty"`
	ServiceName    string                        `json:"serviceName" yaml:"serviceName" validate:"required,iiot-dto-none-empty-string"`
	ProfileName    string                        `json:"profileName,omitempty" yaml:"profileName,omitempty"`
	AutoDataEvents     []AutoDataEvent                   `json:"autoDataEvents,omitempty" yaml:"autoDataEvents,omitempty" validate:"dive"`
	Protocols      map[string]ProtocolProperties `json:"protocols" yaml:"protocols" validate:"required"`
	Tags           map[string]any                `json:"tags,omitempty" yaml:"tags,omitempty"`
	Properties     map[string]any                `json:"properties" yaml:"properties"`
}

type UpdateDevice struct {
	Id             *string                       `json:"id" validate:"required_without=Name,iiot-dto-uuid"`
	Name           *string                       `json:"name" validate:"required_without=Id,iiot-dto-none-empty-string"`
	Parent         *string                       `json:"parent,omitempty" yaml:"parent,omitempty"`
	Description    *string                       `json:"description" validate:"omitempty"`
	ServiceState     *string                       `json:"adminState" validate:"omitempty,oneof='LOCKED' 'UNLOCKED'"`
	OperatingState *string                       `json:"operatingState" validate:"omitempty,oneof='UP' 'DOWN' 'UNKNOWN'"`
	ServiceName    *string                       `json:"serviceName" validate:"omitempty,iiot-dto-none-empty-string"`
	ProfileName    *string                       `json:"profileName" validate:"omitempty"`
	Labels         []string                      `json:"labels"`
	Location       interface{}                   `json:"location"`
	AutoDataEvents     []AutoDataEvent                   `json:"autoDataEvents" validate:"dive"`
	Protocols      map[string]ProtocolProperties `json:"protocols" validate:"omitempty"`
	Tags           map[string]any                `json:"tags"`
	Properties     map[string]any                `json:"properties"`
}

// ToDeviceModel transforms the Device DTO to the Device Model
func ToDeviceModel(dto Device) models.Device {
	var d models.Device
	d.Id = dto.Id
	d.Name = dto.Name
	d.Parent = dto.Parent
	d.Description = dto.Description
	d.ServiceName = dto.ServiceName
	d.ProfileName = dto.ProfileName
	d.ServiceState = models.ServiceState(dto.ServiceState)
	d.OperatingState = models.OperatingState(dto.OperatingState)
	d.Labels = dto.Labels
	d.Location = dto.Location
	d.AutoDataEvents = ToAutoDataEventModels(dto.AutoDataEvents)
	d.Protocols = ToProtocolModels(dto.Protocols)
	d.Tags = dto.Tags
	d.Properties = dto.Properties
	if d.Properties == nil {
		d.Properties = make(map[string]any)
	}
	return d
}

// FromDeviceModelToDTO transforms the Device Model to the Device DTO
func FromDeviceModelToDTO(d models.Device) Device {
	var dto Device
	dto.DBTimestamp = DBTimestamp(d.DBTimestamp)
	dto.Id = d.Id
	dto.Name = d.Name
	dto.Parent = d.Parent
	dto.Description = d.Description
	dto.ServiceName = d.ServiceName
	dto.ProfileName = d.ProfileName
	dto.ServiceState = string(d.ServiceState)
	dto.OperatingState = string(d.OperatingState)
	dto.Labels = d.Labels
	dto.Location = d.Location
	dto.AutoDataEvents = FromAutoDataEventModelsToDTOs(d.AutoDataEvents)
	dto.Protocols = FromProtocolModelsToDTOs(d.Protocols)
	dto.Tags = d.Tags
	dto.Properties = d.Properties
	if dto.Properties == nil {
		dto.Properties = make(map[string]any)
	}
	return dto
}

// FromDeviceModelToUpdateDTO transforms the Device Model to the UpdateDevice DTO
func FromDeviceModelToUpdateDTO(d models.Device) UpdateDevice {
	adminState := string(d.ServiceState)
	operatingState := string(d.OperatingState)
	dto := UpdateDevice{
		Id:             &d.Id,
		Name:           &d.Name,
		Parent:         &d.Parent,
		Description:    &d.Description,
		ServiceState:     &adminState,
		OperatingState: &operatingState,
		ServiceName:    &d.ServiceName,
		ProfileName:    &d.ProfileName,
		Location:       d.Location,
		AutoDataEvents:     FromAutoDataEventModelsToDTOs(d.AutoDataEvents),
		Protocols:      FromProtocolModelsToDTOs(d.Protocols),
		Labels:         d.Labels,
		Tags:           d.Tags,
		Properties:     d.Properties,
	}
	return dto
}
