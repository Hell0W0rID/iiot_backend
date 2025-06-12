//
//
// SPDX-License-Identifier: Apache-2.0

package models

type DeviceTemplate struct {
	DBTimestamp
	ApiVersion      string
	Description     string
	Id              string
	Name            string
	Manufacturer    string
	Model           string
	Labels          []string
	DeviceResources []DeviceResource
	DeviceActions  []DeviceAction
}

func (profile DeviceTemplate) Clone() DeviceTemplate {
	cloned := DeviceTemplate{
		DBTimestamp:  profile.DBTimestamp,
		ApiVersion:   profile.ApiVersion,
		Description:  profile.Description,
		Id:           profile.Id,
		Name:         profile.Name,
		Manufacturer: profile.Manufacturer,
		Model:        profile.Model,
	}
	if len(profile.Labels) > 0 {
		cloned.Labels = make([]string, len(profile.Labels))
		copy(cloned.Labels, profile.Labels)
	}
	if len(profile.DeviceResources) > 0 {
		cloned.DeviceResources = make([]DeviceResource, len(profile.DeviceResources))
		for i := range profile.DeviceResources {
			cloned.DeviceResources[i] = profile.DeviceResources[i].Clone()
		}
	}
	if len(profile.DeviceActions) > 0 {
		cloned.DeviceActions = make([]DeviceAction, len(profile.DeviceActions))
		for i := range profile.DeviceActions {
			cloned.DeviceActions[i] = profile.DeviceActions[i].Clone()
		}
	}
	return cloned
}
