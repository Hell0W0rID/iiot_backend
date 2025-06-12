//
//
// SPDX-License-Identifier: Apache-2.0

package models

import "maps"

type Device struct {
	DBTimestamp
	Id             string
	Name           string
	Parent         string
	Description    string
	ServiceState     ServiceState
	OperatingState OperatingState
	Protocols      map[string]ProtocolProperties
	Labels         []string
	Location       interface{}
	ServiceName    string
	ProfileName    string
	AutoDataEvents     []AutoDataEvent
	Tags           map[string]any
	Properties     map[string]any
}

// ProtocolProperties contains the device connection information in key/value pair
type ProtocolProperties map[string]any

func (p ProtocolProperties) Clone() ProtocolProperties {
	cloned := make(map[string]any)
	maps.Copy(cloned, p)
	return cloned
}

// ServiceState controls the range of values which constitute valid administrative states for a device
type ServiceState string

// AssignServiceState provides a default value "UNLOCKED" to ServiceState
func AssignServiceState(dtoServiceState string) ServiceState {
	if dtoServiceState == "" {
		return ServiceState(Unlocked)
	}
	return ServiceState(dtoServiceState)
}

// OperatingState is an indication of the operations of the device.
type OperatingState string

func (device Device) Clone() Device {
	cloned := Device{
		DBTimestamp:    device.DBTimestamp,
		Id:             device.Id,
		Name:           device.Name,
		Parent:         device.Parent,
		Description:    device.Description,
		ServiceState:     device.ServiceState,
		OperatingState: device.OperatingState,
		Location:       device.Location,
		ServiceName:    device.ServiceName,
		ProfileName:    device.ProfileName,
	}
	if len(device.Protocols) > 0 {
		cloned.Protocols = make(map[string]ProtocolProperties)
		for k, v := range device.Protocols {
			cloned.Protocols[k] = v.Clone()
		}
	}
	if len(device.Labels) > 0 {
		cloned.Labels = make([]string, len(device.Labels))
		copy(cloned.Labels, device.Labels)
	}
	if len(device.AutoDataEvents) > 0 {
		cloned.AutoDataEvents = make([]AutoDataEvent, len(device.AutoDataEvents))
		copy(cloned.AutoDataEvents, device.AutoDataEvents)
	}
	if len(device.Tags) > 0 {
		cloned.Tags = make(map[string]any)
		maps.Copy(cloned.Tags, device.Tags)
	}
	if len(device.Properties) > 0 {
		cloned.Properties = make(map[string]any)
		maps.Copy(cloned.Properties, device.Properties)
	}
	return cloned
}
