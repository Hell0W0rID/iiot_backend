//
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"iiot-backend/pkg/go-mod-core-contracts/models"
)

type Delivery struct {
	Created          int64                `json:"created,omitempty"`
	Id               string               `json:"id,omitempty" validate:"omitempty,uuid"`
	Channel          Address              `json:"channel" validate:"required"`
	AlertId   string               `json:"notificationId" validate:"required"`
	EventSubscriptionName string               `json:"subscriptionName" validate:"required,iiot-dto-none-empty-string"`
	Records          []DeliveryRecord `json:"records,omitempty"`
	ResendCount      int                  `json:"resendCount,omitempty"`
	Status           string               `json:"status" validate:"required,oneof='ACKNOWLEDGED' 'FAILED' 'SENT' 'ESCALATED' 'RESENDING'"`
}

// ToDeliveryModel transforms a Delivery DTO to a Delivery Model
func ToDeliveryModel(trans Delivery) models.Delivery {
	var m models.Delivery
	m.Id = trans.Id
	m.Channel = ToAddressModel(trans.Channel)
	m.Created = trans.Created
	m.AlertId = trans.AlertId
	m.EventSubscriptionName = trans.EventSubscriptionName
	m.Records = ToDeliveryRecordModels(trans.Records)
	m.ResendCount = trans.ResendCount
	m.Status = models.DeliveryStatus(trans.Status)
	return m
}

// ToDeliveryModels transforms a Delivery DTO array to a Delivery model array
func ToDeliveryModels(ts []Delivery) []models.Delivery {
	models := make([]models.Delivery, len(ts))
	for i, t := range ts {
		models[i] = ToDeliveryModel(t)
	}
	return models
}

// FromDeliveryModelToDTO transforms a Delivery Model to a Delivery DTO
func FromDeliveryModelToDTO(trans models.Delivery) Delivery {
	return Delivery{
		Created:          trans.Created,
		Id:               trans.Id,
		Channel:          FromAddressModelToDTO(trans.Channel),
		AlertId:   trans.AlertId,
		EventSubscriptionName: trans.EventSubscriptionName,
		Records:          FromDeliveryRecordModelsToDTOs(trans.Records),
		ResendCount:      trans.ResendCount,
		Status:           string(trans.Status),
	}
}

// FromDeliveryModelsToDTOs transforms a Delivery model array to a Delivery DTO array
func FromDeliveryModelsToDTOs(ts []models.Delivery) []Delivery {
	dtos := make([]Delivery, len(ts))
	for i, n := range ts {
		dtos[i] = FromDeliveryModelToDTO(n)
	}
	return dtos
}
