//
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"iiot-backend/pkg/go-mod-core-contracts/models"
)

type DeviceWatcher struct {
	DBTimestamp         `json:",inline"`
	Id                  string              `json:"id,omitempty" yaml:"id,omitempty" validate:"omitempty,uuid"`
	Name                string              `json:"name" yaml:"name" validate:"required,iiot-dto-none-empty-string"`
	ServiceName         string              `json:"serviceName" yaml:"serviceName" validate:"required,iiot-dto-none-empty-string"`
	Labels              []string            `json:"labels,omitempty" yaml:"labels,omitempty"`
	Identifiers         map[string]string   `json:"identifiers" yaml:"identifiers" validate:"gt=0,dive,keys,required,endkeys,required"`
	BlockingIdentifiers map[string][]string `json:"blockingIdentifiers,omitempty" yaml:"blockingIdentifiers,omitempty"`
	ServiceState          string              `json:"adminState" yaml:"adminState" validate:"oneof='LOCKED' 'UNLOCKED'"`
	DiscoveredDevice    DiscoveredDevice    `json:"discoveredDevice" yaml:"discoveredDevice"`
}

type UpdateDeviceWatcher struct {
	Id                  *string                `json:"id" validate:"required_without=Name,iiot-dto-uuid"`
	Name                *string                `json:"name" validate:"required_without=Id,iiot-dto-none-empty-string"`
	ServiceName         *string                `json:"serviceName" validate:"omitempty,iiot-dto-none-empty-string"`
	Labels              []string               `json:"labels"`
	Identifiers         map[string]string      `json:"identifiers" validate:"omitempty,gt=0,dive,keys,required,endkeys,required"`
	BlockingIdentifiers map[string][]string    `json:"blockingIdentifiers"`
	ServiceState          *string                `json:"adminState" validate:"omitempty,oneof='LOCKED' 'UNLOCKED'"`
	DiscoveredDevice    UpdateDiscoveredDevice `json:"discoveredDevice"`
}

// ToDeviceWatcherModel transforms the DeviceWatcher DTO to the DeviceWatcher model
func ToDeviceWatcherModel(dto DeviceWatcher) models.DeviceWatcher {
	return models.DeviceWatcher{
		DBTimestamp:         models.DBTimestamp(dto.DBTimestamp),
		Id:                  dto.Id,
		Name:                dto.Name,
		ServiceName:         dto.ServiceName,
		Labels:              dto.Labels,
		Identifiers:         dto.Identifiers,
		BlockingIdentifiers: dto.BlockingIdentifiers,
		ServiceState:          models.ServiceState(dto.ServiceState),
		DiscoveredDevice:    ToDiscoveredDeviceModel(dto.DiscoveredDevice),
	}
}

// FromDeviceWatcherModelToDTO transforms the DeviceWatcher Model to the DeviceWatcher DTO
func FromDeviceWatcherModelToDTO(pw models.DeviceWatcher) DeviceWatcher {
	return DeviceWatcher{
		DBTimestamp:         DBTimestamp(pw.DBTimestamp),
		Id:                  pw.Id,
		Name:                pw.Name,
		ServiceName:         pw.ServiceName,
		Labels:              pw.Labels,
		Identifiers:         pw.Identifiers,
		BlockingIdentifiers: pw.BlockingIdentifiers,
		ServiceState:          string(pw.ServiceState),
		DiscoveredDevice:    FromDiscoveredDeviceModelToDTO(pw.DiscoveredDevice),
	}
}

// FromDeviceWatcherModelToUpdateDTO transforms the DeviceWatcher Model to the UpdateDeviceWatcher DTO
func FromDeviceWatcherModelToUpdateDTO(pw models.DeviceWatcher) UpdateDeviceWatcher {
	adminState := string(pw.ServiceState)
	dto := UpdateDeviceWatcher{
		Id:                  &pw.Id,
		Name:                &pw.Name,
		ServiceName:         &pw.ServiceName,
		Labels:              pw.Labels,
		Identifiers:         pw.Identifiers,
		BlockingIdentifiers: pw.BlockingIdentifiers,
		ServiceState:          &adminState,
		DiscoveredDevice:    FromDiscoveredDeviceModelToUpdateDTO(pw.DiscoveredDevice),
	}
	return dto
}
