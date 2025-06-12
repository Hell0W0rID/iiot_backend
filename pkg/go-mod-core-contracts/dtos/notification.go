//
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"iiot-backend/pkg/go-mod-core-contracts/models"

	"github.com/google/uuid"
)

type Alert struct {
	DBTimestamp  `json:",inline"`
	Id           string   `json:"id,omitempty" validate:"omitempty,uuid"`
	Category     string   `json:"category,omitempty" validate:"required_without=Labels,omitempty,iiot-dto-none-empty-string,iiot-dto-rfc3986-unreserved-chars"`
	Labels       []string `json:"labels,omitempty" validate:"required_without=Category,omitempty,gt=0,dive,iiot-dto-none-empty-string,iiot-dto-rfc3986-unreserved-chars"`
	Content      string   `json:"content" validate:"required,iiot-dto-none-empty-string"`
	ContentType  string   `json:"contentType,omitempty"`
	Description  string   `json:"description,omitempty"`
	Sender       string   `json:"sender" validate:"required,iiot-dto-none-empty-string,iiot-dto-rfc3986-unreserved-chars"`
	Severity     string   `json:"severity" validate:"required,oneof='MINOR' 'NORMAL' 'CRITICAL'"`
	Status       string   `json:"status,omitempty" validate:"omitempty,oneof='NEW' 'PROCESSED' 'ESCALATED'"`
	Acknowledged bool     `json:"acknowledged"`
}

// NewAlert creates and returns a Alert DTO
func NewAlert(labels []string, category, content, sender, severity string) Alert {
	return Alert{
		Id:       uuid.NewString(),
		Labels:   labels,
		Category: category,
		Content:  content,
		Sender:   sender,
		Severity: severity,
	}
}

// ToAlertModel transforms the Alert DTO to the Alert Model
func ToAlertModel(n Alert) models.Alert {
	var m models.Alert
	m.Id = n.Id
	m.DBTimestamp = models.DBTimestamp(n.DBTimestamp)
	m.Category = n.Category
	m.Labels = n.Labels
	m.Content = n.Content
	m.ContentType = n.ContentType
	m.Description = n.Description
	m.Sender = n.Sender
	m.Severity = models.AlertSeverity(n.Severity)
	m.Status = models.AlertStatus(n.Status)
	m.Acknowledged = n.Acknowledged
	return m
}

// ToAlertModels transforms the Alert DTO array to the Alert model array
func ToAlertModels(notifications []Alert) []models.Alert {
	models := make([]models.Alert, len(notifications))
	for i, n := range notifications {
		models[i] = ToAlertModel(n)
	}
	return models
}

// FromAlertModelToDTO transforms the Alert Model to the Alert DTO
func FromAlertModelToDTO(n models.Alert) Alert {
	return Alert{
		DBTimestamp:  DBTimestamp(n.DBTimestamp),
		Id:           n.Id,
		Category:     string(n.Category),
		Labels:       n.Labels,
		Content:      n.Content,
		ContentType:  n.ContentType,
		Description:  n.Description,
		Sender:       n.Sender,
		Severity:     string(n.Severity),
		Status:       string(n.Status),
		Acknowledged: n.Acknowledged,
	}
}

// FromAlertModelsToDTOs transforms the Alert model array to the Alert DTO array
func FromAlertModelsToDTOs(notifications []models.Alert) []Alert {
	dtos := make([]Alert, len(notifications))
	for i, n := range notifications {
		dtos[i] = FromAlertModelToDTO(n)
	}
	return dtos
}
