//
//
// SPDX-License-Identifier: Apache-2.0

package models

import "maps"

type DiscoveredDevice struct {
	ProfileName string
	ServiceState  ServiceState
	AutoDataEvents  []AutoDataEvent
	Properties  map[string]any
}

func (d DiscoveredDevice) Clone() DiscoveredDevice {
	cloned := DiscoveredDevice{
		ProfileName: d.ProfileName,
		ServiceState:  d.ServiceState,
	}
	if len(d.AutoDataEvents) > 0 {
		cloned.AutoDataEvents = make([]AutoDataEvent, len(d.AutoDataEvents))
		copy(cloned.AutoDataEvents, d.AutoDataEvents)
	}
	if len(d.Properties) > 0 {
		cloned.Properties = make(map[string]any)
		maps.Copy(cloned.Properties, d.Properties)
	}
	return cloned
}
