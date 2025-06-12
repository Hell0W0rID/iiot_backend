//
//
// SPDX-License-Identifier: Apache-2.0

package models

type Alert struct {
	DBTimestamp
	Category     string
	Content      string
	ContentType  string
	Description  string
	Id           string
	Labels       []string
	Sender       string
	Severity     AlertSeverity
	Status       AlertStatus
	Acknowledged bool
}

// AlertSeverity indicates the level of severity for the notification.
type AlertSeverity string

// AlertStatus indicates the current processing status of the notification.
type AlertStatus string
