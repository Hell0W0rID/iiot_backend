//
//
// SPDX-License-Identifier: Apache-2.0

package models

type AutoDataEvent struct {
	Interval          string
	OnChange          bool
	OnChangeThreshold float64
	SourceName        string
	Retention         Retention
}

type Retention struct {
	MaxCap   int64
	MinCap   int64
	Duration string
}
