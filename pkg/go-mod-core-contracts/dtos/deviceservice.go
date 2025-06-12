//
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"iiot-backend/pkg/go-mod-core-contracts/models"
)

type DeviceHandler struct {
	DBTimestamp `json:",inline"`
	Id          string         `json:"id,omitempty" validate:"omitempty,uuid"`
	Name        string         `json:"name" validate:"required,iiot-dto-none-empty-string"`
	Description string         `json:"description,omitempty"`
	Labels      []string       `json:"labels,omitempty"`
	ServiceAddress string         `json:"baseAddress" validate:"required,uri"`
	ServiceState  string         `json:"adminState" validate:"oneof='LOCKED' 'UNLOCKED'"`
	Properties  map[string]any `json:"properties" yaml:"properties"`
}

type UpdateDeviceHandler struct {
	Id          *string        `json:"id" validate:"required_without=Name,iiot-dto-uuid"`
	Name        *string        `json:"name" validate:"required_without=Id,iiot-dto-none-empty-string"`
	Description *string        `json:"description"`
	ServiceAddress *string        `json:"baseAddress" validate:"omitempty,uri"`
	Labels      []string       `json:"labels"`
	ServiceState  *string        `json:"adminState" validate:"omitempty,oneof='LOCKED' 'UNLOCKED'"`
	Properties  map[string]any `json:"properties"`
}

// ToDeviceHandlerModel transforms the DeviceHandler DTO to the DeviceHandler Model
func ToDeviceHandlerModel(dto DeviceHandler) models.DeviceHandler {
	var ds models.DeviceHandler
	ds.Id = dto.Id
	ds.Name = dto.Name
	ds.Description = dto.Description
	ds.ServiceAddress = dto.ServiceAddress
	ds.Labels = dto.Labels
	ds.ServiceState = models.ServiceState(dto.ServiceState)
	ds.Properties = dto.Properties
	if ds.Properties == nil {
		ds.Properties = make(map[string]any)
	}
	return ds
}

// FromDeviceHandlerModelToDTO transforms the DeviceHandler Model to the DeviceHandler DTO
func FromDeviceHandlerModelToDTO(ds models.DeviceHandler) DeviceHandler {
	var dto DeviceHandler
	dto.DBTimestamp = DBTimestamp(ds.DBTimestamp)
	dto.Id = ds.Id
	dto.Name = ds.Name
	dto.Description = ds.Description
	dto.ServiceAddress = ds.ServiceAddress
	dto.Labels = ds.Labels
	dto.ServiceState = string(ds.ServiceState)
	dto.Properties = ds.Properties
	if dto.Properties == nil {
		dto.Properties = make(map[string]any)
	}
	return dto
}

// FromDeviceHandlerModelToUpdateDTO transforms the DeviceHandler Model to the UpdateDeviceHandler DTO
func FromDeviceHandlerModelToUpdateDTO(ds models.DeviceHandler) UpdateDeviceHandler {
	adminState := string(ds.ServiceState)
	dto := UpdateDeviceHandler{
		Id:          &ds.Id,
		Name:        &ds.Name,
		Description: &ds.Description,
		Labels:      ds.Labels,
		ServiceAddress: &ds.ServiceAddress,
		ServiceState:  &adminState,
		Properties:  ds.Properties,
	}
	return dto
}
