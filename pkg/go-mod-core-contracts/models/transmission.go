//
//
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"encoding/json"

	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

type Delivery struct {
	Created          int64
	Id               string
	Channel          Address
	AlertId   string
	EventSubscriptionName string
	Records          []DeliveryRecord
	ResendCount      int
	Status           DeliveryStatus
}

func (trans *Delivery) UnmarshalJSON(b []byte) error {
	var alias struct {
		Created          int64
		Id               string
		Channel          interface{}
		AlertId   string
		EventSubscriptionName string
		Records          []DeliveryRecord
		ResendCount      int
		Status           DeliveryStatus
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal transmission.", err)
	}

	channel, err := instantiateAddress(alias.Channel)
	if err != nil {
		return errors.NewCommonIIOTWrapper(err)
	}

	*trans = Delivery{
		Created:          alias.Created,
		Id:               alias.Id,
		Channel:          channel,
		AlertId:   alias.AlertId,
		EventSubscriptionName: alias.EventSubscriptionName,
		Records:          alias.Records,
		ResendCount:      alias.ResendCount,
		Status:           alias.Status,
	}
	return nil
}

// NewDelivery create transmission model with required fields
func NewDelivery(subscriptionName string, channel Address, notificationId string) Delivery {
	return Delivery{
		EventSubscriptionName: subscriptionName,
		Channel:          channel,
		AlertId:   notificationId,
	}
}

// DeliveryStatus indicates the most recent success/failure of a given transmission attempt.
type DeliveryStatus string
