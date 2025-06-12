// SPDX-License-Identifier: Apache-2.0

package models

// Channel - from IIOT official repository models
type Channel struct {
	Type          ChannelType            `json:"type" yaml:"type" validate:"required"`
	MailAddresses []string               `json:"mailAddresses,omitempty" yaml:"mailAddresses,omitempty"`
	Url           string                 `json:"url,omitempty" yaml:"url,omitempty"`
	HttpMethod    string                 `json:"httpMethod,omitempty" yaml:"httpMethod,omitempty"`
	HttpHeaders   map[string]string      `json:"httpHeaders,omitempty" yaml:"httpHeaders,omitempty"`
}

// ChannelType indicates whether the channel is for email or REST
type ChannelType string

const (
	Email ChannelType = "EMAIL"
	Rest  ChannelType = "REST"
)

// NewEmailChannel creates email channel
func NewEmailChannel(addresses []string) Channel {
	return Channel{
		Type:          Email,
		MailAddresses: addresses,
	}
}

// NewRestChannel creates REST channel
func NewRestChannel(url, httpMethod string, headers map[string]string) Channel {
	return Channel{
		Type:        Rest,
		Url:         url,
		HttpMethod:  httpMethod,
		HttpHeaders: headers,
	}
}