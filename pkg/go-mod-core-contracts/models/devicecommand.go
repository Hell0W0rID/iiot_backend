//
//
// SPDX-License-Identifier: Apache-2.0

package models

import "maps"

type DeviceAction struct {
	Name               string
	IsHidden           bool
	ReadWrite          string
	ResourceOperations []ResourceOperation
	Tags               map[string]any
}

func (dc DeviceAction) Clone() DeviceAction {
	cloned := DeviceAction{
		Name:      dc.Name,
		IsHidden:  dc.IsHidden,
		ReadWrite: dc.ReadWrite,
	}
	if len(dc.ResourceOperations) > 0 {
		cloned.ResourceOperations = make([]ResourceOperation, len(dc.ResourceOperations))
		for i, op := range dc.ResourceOperations {
			cloned.ResourceOperations[i] = op.Clone()
		}
	}
	if len(dc.Tags) > 0 {
		cloned.Tags = make(map[string]any)
		maps.Copy(cloned.Tags, dc.Tags)
	}
	return cloned
}
