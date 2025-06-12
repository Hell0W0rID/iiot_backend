//
//
// SPDX-License-Identifier: Apache-2.0

package models

type DeviceHandler struct {
	DBTimestamp
	Id          string
	Name        string
	Description string
	Labels      []string
	ServiceAddress string
	ServiceState  ServiceState
	Properties  map[string]any
}
