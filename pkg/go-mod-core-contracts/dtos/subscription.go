//
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"iiot-backend/pkg/go-mod-core-contracts/models"
)

type EventSubscription struct {
	DBTimestamp    `json:",inline"`
	Id             string    `json:"id,omitempty" validate:"omitempty,uuid"`
	Name           string    `json:"name" validate:"required,iiot-dto-none-empty-string"`
	Channels       []Address `json:"channels" validate:"required,gt=0,dive"`
	Receiver       string    `json:"receiver" validate:"required,iiot-dto-none-empty-string,iiot-dto-rfc3986-unreserved-chars"`
	Categories     []string  `json:"categories,omitempty" validate:"required_without=Labels,omitempty,gt=0,dive,iiot-dto-none-empty-string,iiot-dto-rfc3986-unreserved-chars"`
	Labels         []string  `json:"labels,omitempty" validate:"required_without=Categories,omitempty,gt=0,dive,iiot-dto-none-empty-string,iiot-dto-rfc3986-unreserved-chars"`
	Description    string    `json:"description,omitempty"`
	ResendLimit    int       `json:"resendLimit,omitempty"`
	ResendInterval string    `json:"resendInterval,omitempty" validate:"omitempty,iiot-dto-duration"`
	ServiceState     string    `json:"adminState" validate:"oneof='LOCKED' 'UNLOCKED'"`
}

type UpdateEventSubscription struct {
	Id             *string   `json:"id" validate:"required_without=Name,iiot-dto-uuid"`
	Name           *string   `json:"name" validate:"required_without=Id,iiot-dto-none-empty-string"`
	Channels       []Address `json:"channels" validate:"omitempty,gt=0,dive"`
	Receiver       *string   `json:"receiver" validate:"omitempty,iiot-dto-none-empty-string,iiot-dto-rfc3986-unreserved-chars"`
	Categories     []string  `json:"categories" validate:"omitempty,dive,gt=0,iiot-dto-none-empty-string,iiot-dto-rfc3986-unreserved-chars"`
	Labels         []string  `json:"labels" validate:"omitempty,dive,iiot-dto-none-empty-string,iiot-dto-rfc3986-unreserved-chars"`
	Description    *string   `json:"description"`
	ResendLimit    *int      `json:"resendLimit"`
	ResendInterval *string   `json:"resendInterval" validate:"omitempty,iiot-dto-duration"`
	ServiceState     *string   `json:"adminState" validate:"omitempty,oneof='LOCKED' 'UNLOCKED'"`
}

// ToEventSubscriptionModel transforms the EventSubscription DTO to the EventSubscription Model
func ToEventSubscriptionModel(s EventSubscription) models.EventSubscription {
	var m models.EventSubscription
	m.Categories = s.Categories
	m.Labels = s.Labels
	m.Channels = ToAddressModels(s.Channels)
	m.DBTimestamp = models.DBTimestamp(s.DBTimestamp)
	m.Description = s.Description
	m.Id = s.Id
	m.Receiver = s.Receiver
	m.Name = s.Name
	m.ResendLimit = s.ResendLimit
	m.ResendInterval = s.ResendInterval
	m.ServiceState = models.ServiceState(s.ServiceState)
	return m
}

// ToEventSubscriptionModels transforms the EventSubscription DTO array to the EventSubscription model array
func ToEventSubscriptionModels(subs []EventSubscription) []models.EventSubscription {
	models := make([]models.EventSubscription, len(subs))
	for i, s := range subs {
		models[i] = ToEventSubscriptionModel(s)
	}
	return models
}

// FromEventSubscriptionModelToDTO transforms the EventSubscription Model to the EventSubscription DTO
func FromEventSubscriptionModelToDTO(s models.EventSubscription) EventSubscription {
	return EventSubscription{
		DBTimestamp:    DBTimestamp(s.DBTimestamp),
		Categories:     s.Categories,
		Labels:         s.Labels,
		Channels:       FromAddressModelsToDTOs(s.Channels),
		Description:    s.Description,
		Id:             s.Id,
		Receiver:       s.Receiver,
		Name:           s.Name,
		ResendLimit:    s.ResendLimit,
		ResendInterval: s.ResendInterval,
		ServiceState:     string(s.ServiceState),
	}
}

// FromEventSubscriptionModelsToDTOs transforms the EventSubscription model array to the EventSubscription DTO array
func FromEventSubscriptionModelsToDTOs(subscruptions []models.EventSubscription) []EventSubscription {
	dtos := make([]EventSubscription, len(subscruptions))
	for i, s := range subscruptions {
		dtos[i] = FromEventSubscriptionModelToDTO(s)
	}
	return dtos
}
