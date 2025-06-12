//
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "iiot-backend/pkg/go-mod-core-contracts/models"

type DiscoveredDevice struct {
	ProfileName string         `json:"profileName" yaml:"profileName" validate:"len=0|iiot-dto-none-empty-string"`
	ServiceState  string         `json:"adminState" yaml:"adminState" validate:"oneof='LOCKED' 'UNLOCKED'"`
	AutoDataEvents  []AutoDataEvent    `json:"autoDataEvents,omitempty" yaml:"autoDataEvents,omitempty" validate:"dive"`
	Properties  map[string]any `json:"properties" yaml:"properties"`
}

type UpdateDiscoveredDevice struct {
	ProfileName *string        `json:"profileName" validate:"omitempty,len=0|iiot-dto-none-empty-string"`
	ServiceState  *string        `json:"adminState" validate:"omitempty,oneof='LOCKED' 'UNLOCKED'"`
	AutoDataEvents  []AutoDataEvent    `json:"autoDataEvents" validate:"dive"`
	Properties  map[string]any `json:"properties"`
}

func ToDiscoveredDeviceModel(dto DiscoveredDevice) models.DiscoveredDevice {
	m := models.DiscoveredDevice{
		ProfileName: dto.ProfileName,
		ServiceState:  models.ServiceState(dto.ServiceState),
		AutoDataEvents:  ToAutoDataEventModels(dto.AutoDataEvents),
		Properties:  dto.Properties,
	}
	if m.Properties == nil {
		m.Properties = make(map[string]any)
	}
	return m
}

func FromDiscoveredDeviceModelToDTO(d models.DiscoveredDevice) DiscoveredDevice {
	dto := DiscoveredDevice{
		ProfileName: d.ProfileName,
		ServiceState:  string(d.ServiceState),
		AutoDataEvents:  FromAutoDataEventModelsToDTOs(d.AutoDataEvents),
		Properties:  d.Properties,
	}
	if dto.Properties == nil {
		dto.Properties = make(map[string]any)
	}
	return dto
}

func FromDiscoveredDeviceModelToUpdateDTO(d models.DiscoveredDevice) UpdateDiscoveredDevice {
	adminState := string(d.ServiceState)
	return UpdateDiscoveredDevice{
		ProfileName: &d.ProfileName,
		ServiceState:  &adminState,
		AutoDataEvents:  FromAutoDataEventModelsToDTOs(d.AutoDataEvents),
		Properties:  d.Properties,
	}
}
