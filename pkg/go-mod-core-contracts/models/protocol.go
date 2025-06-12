// SPDX-License-Identifier: Apache-2.0

package models

// Protocol - from IIOT official repository models
type Protocol struct {
	Name       string                 `json:"name" yaml:"name" validate:"required"`
	Properties map[string]interface{} `json:"properties,omitempty" yaml:"properties,omitempty"`
}

// NewProtocol creates protocol with required fields
func NewProtocol(name string) Protocol {
	return Protocol{
		Name:       name,
		Properties: make(map[string]interface{}),
	}
}