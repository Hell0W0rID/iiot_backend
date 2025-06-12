//
//
// SPDX-License-Identifier: Apache-2.0

package models

type DataEvent struct {
	Id          string
	DeviceName  string
	ProfileName string
	SourceName  string
	Origin      int64
	Measurements    []Measurement
	Tags        map[string]interface{}
}
