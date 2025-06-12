//
//
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"encoding/json"

	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

type EventSubscription struct {
	DBTimestamp
	Categories     []string
	Labels         []string
	Channels       []Address
	Description    string
	Id             string
	Receiver       string
	Name           string
	ResendLimit    int
	ResendInterval string
	ServiceState     ServiceState
}

// ChannelType controls the range of values which constitute valid delivery types for channels
type ChannelType string

func (subscription *EventSubscription) UnmarshalJSON(b []byte) error {
	var alias struct {
		DBTimestamp
		Categories     []string
		Labels         []string
		Channels       []interface{}
		Description    string
		Id             string
		Receiver       string
		Name           string
		ResendLimit    int
		ResendInterval string
		ServiceState     ServiceState
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal intervalAction.", err)
	}
	channels := make([]Address, len(alias.Channels))
	for i, c := range alias.Channels {
		address, err := instantiateAddress(c)
		if err != nil {
			return errors.NewCommonIIOTWrapper(err)
		}
		channels[i] = address
	}

	*subscription = EventSubscription{
		DBTimestamp:    alias.DBTimestamp,
		Categories:     alias.Categories,
		Labels:         alias.Labels,
		Description:    alias.Description,
		Id:             alias.Id,
		Receiver:       alias.Receiver,
		Name:           alias.Name,
		ResendLimit:    alias.ResendLimit,
		ResendInterval: alias.ResendInterval,
		Channels:       channels,
		ServiceState:     alias.ServiceState,
	}
	return nil
}
