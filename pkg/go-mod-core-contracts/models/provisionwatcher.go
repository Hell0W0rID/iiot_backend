//
//
// SPDX-License-Identifier: Apache-2.0

package models

import "maps"

type DeviceWatcher struct {
	DBTimestamp
	Id                  string
	Name                string
	ServiceName         string
	Labels              []string
	Identifiers         map[string]string
	BlockingIdentifiers map[string][]string
	ServiceState          ServiceState
	DiscoveredDevice    DiscoveredDevice
}

func (pw DeviceWatcher) Clone() DeviceWatcher {
	cloned := DeviceWatcher{
		DBTimestamp:      pw.DBTimestamp,
		Id:               pw.Id,
		Name:             pw.Name,
		ServiceName:      pw.ServiceName,
		ServiceState:       pw.ServiceState,
		DiscoveredDevice: pw.DiscoveredDevice.Clone(),
	}
	if len(pw.Labels) > 0 {
		cloned.Labels = make([]string, len(pw.Labels))
		copy(cloned.Labels, pw.Labels)
	}
	if len(pw.Identifiers) > 0 {
		cloned.Identifiers = make(map[string]string)
		maps.Copy(cloned.Identifiers, pw.Identifiers)
	}
	if len(pw.BlockingIdentifiers) > 0 {
		cloned.BlockingIdentifiers = make(map[string][]string)
		for k, v := range pw.BlockingIdentifiers {
			cloned.BlockingIdentifiers[k] = make([]string, len(v))
			copy(cloned.BlockingIdentifiers[k], v)
		}
	}
	return cloned
}
