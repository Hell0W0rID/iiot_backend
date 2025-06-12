package models

import (
	"time"
)

// Notification represents a notification in the system
type Notification struct {
	ID          string            `json:"id" db:"id"`
	Slug        string            `json:"slug" db:"slug"`
	Sender      string            `json:"sender" db:"sender"`
	Category    string            `json:"category" db:"category"`
	Severity    string            `json:"severity" db:"severity"`
	Content     string            `json:"content" db:"content"`
	Description string            `json:"description" db:"description"`
	Status      string            `json:"status" db:"status"`
	Labels      []string          `json:"labels" db:"labels"`
	ContentType string            `json:"contentType" db:"content_type"`
	Created     time.Time         `json:"created" db:"created"`
	Modified    time.Time         `json:"modified" db:"modified"`
}

// Subscription represents a notification subscription
type Subscription struct {
	ID           string            `json:"id" db:"id"`
	Name         string            `json:"name" db:"name"`
	Slug         string            `json:"slug" db:"slug"`
	Description  string            `json:"description" db:"description"`
	Receiver     string            `json:"receiver" db:"receiver"`
	SubscribedCategories []string `json:"subscribedCategories" db:"subscribed_categories"`
	SubscribedLabels     []string `json:"subscribedLabels" db:"subscribed_labels"`
	Channels     []Channel         `json:"channels" db:"channels"`
	ResendLimit  int               `json:"resendLimit" db:"resend_limit"`
	ResendInterval string          `json:"resendInterval" db:"resend_interval"`
	AdminState   string            `json:"adminState" db:"admin_state"`
	Created      time.Time         `json:"created" db:"created"`
	Modified     time.Time         `json:"modified" db:"modified"`
}

// Channel represents a notification channel
type Channel struct {
	Type           string            `json:"type"`
	EmailAddresses []string          `json:"emailAddresses,omitempty"`
	MailServerHost string            `json:"mailServerHost,omitempty"`
	MailServerPort int               `json:"mailServerPort,omitempty"`
	MailServerUsername string        `json:"mailServerUsername,omitempty"`
	MailServerPassword string        `json:"mailServerPassword,omitempty"`
	URL            string            `json:"url,omitempty"`
	HTTPMethod     string            `json:"httpMethod,omitempty"`
	HTTPHeaders    map[string]string `json:"httpHeaders,omitempty"`
}

// NotificationRequest represents a request to create a notification
type NotificationRequest struct {
	Slug        string   `json:"slug" validate:"required"`
	Sender      string   `json:"sender" validate:"required"`
	Category    string   `json:"category" validate:"required"`
	Severity    string   `json:"severity" validate:"required"`
	Content     string   `json:"content" validate:"required"`
	Description string   `json:"description"`
	Labels      []string `json:"labels"`
	ContentType string   `json:"contentType"`
}

// SubscriptionRequest represents a request to create/update a subscription
type SubscriptionRequest struct {
	Name         string    `json:"name" validate:"required"`
	Slug         string    `json:"slug" validate:"required"`
	Description  string    `json:"description"`
	Receiver     string    `json:"receiver" validate:"required"`
	SubscribedCategories []string `json:"subscribedCategories"`
	SubscribedLabels     []string `json:"subscribedLabels"`
	Channels     []Channel `json:"channels" validate:"required"`
	ResendLimit  int       `json:"resendLimit"`
	ResendInterval string  `json:"resendInterval"`
	AdminState   string    `json:"adminState"`
}

// NotificationFilter represents filters for querying notifications
type NotificationFilter struct {
	Category string    `json:"category"`
	Severity string    `json:"severity"`
	Status   string    `json:"status"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Limit    int       `json:"limit"`
	Offset   int       `json:"offset"`
}
