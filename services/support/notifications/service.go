package notifications

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"iiot-backend/models"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

// Notification methods
func (s *Service) GetNotifications(filter models.NotificationFilter) ([]models.Notification, error) {
	query := `
		SELECT id, slug, sender, category, severity, content, description, status, 
		       labels, content_type, created, modified
		FROM notifications
		WHERE ($1 = '' OR category = $1)
		  AND ($2 = '' OR severity = $2)
		  AND ($3 = '' OR status = $3)
		  AND ($4::timestamp IS NULL OR created >= $4)
		  AND ($5::timestamp IS NULL OR created <= $5)
		ORDER BY created DESC
		LIMIT $6 OFFSET $7
	`

	var start, end interface{}
	if !filter.Start.IsZero() {
		start = filter.Start
	}
	if !filter.End.IsZero() {
		end = filter.End
	}

	rows, err := s.db.Query(query, filter.Category, filter.Severity, filter.Status,
		start, end, filter.Limit, filter.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query notifications: %w", err)
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var notification models.Notification
		var labelsJSON []byte

		err := rows.Scan(
			&notification.ID, &notification.Slug, &notification.Sender,
			&notification.Category, &notification.Severity, &notification.Content,
			&notification.Description, &notification.Status, &labelsJSON,
			&notification.ContentType, &notification.Created, &notification.Modified,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan notification: %w", err)
		}

		if len(labelsJSON) > 0 {
			json.Unmarshal(labelsJSON, &notification.Labels)
		}

		notifications = append(notifications, notification)
	}

	return notifications, nil
}

func (s *Service) GetNotificationByID(id string) (*models.Notification, error) {
	query := `
		SELECT id, slug, sender, category, severity, content, description, status,
		       labels, content_type, created, modified
		FROM notifications
		WHERE id = $1
	`

	var notification models.Notification
	var labelsJSON []byte

	err := s.db.QueryRow(query, id).Scan(
		&notification.ID, &notification.Slug, &notification.Sender,
		&notification.Category, &notification.Severity, &notification.Content,
		&notification.Description, &notification.Status, &labelsJSON,
		&notification.ContentType, &notification.Created, &notification.Modified,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get notification: %w", err)
	}

	if len(labelsJSON) > 0 {
		json.Unmarshal(labelsJSON, &notification.Labels)
	}

	return &notification, nil
}

func (s *Service) GetNotificationBySlug(slug string) (*models.Notification, error) {
	query := `
		SELECT id, slug, sender, category, severity, content, description, status,
		       labels, content_type, created, modified
		FROM notifications
		WHERE slug = $1
	`

	var notification models.Notification
	var labelsJSON []byte

	err := s.db.QueryRow(query, slug).Scan(
		&notification.ID, &notification.Slug, &notification.Sender,
		&notification.Category, &notification.Severity, &notification.Content,
		&notification.Description, &notification.Status, &labelsJSON,
		&notification.ContentType, &notification.Created, &notification.Modified,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get notification: %w", err)
	}

	if len(labelsJSON) > 0 {
		json.Unmarshal(labelsJSON, &notification.Labels)
	}

	return &notification, nil
}

func (s *Service) CreateNotification(req *models.NotificationRequest) (string, error) {
	notification := &models.Notification{
		ID:          uuid.New().String(),
		Slug:        req.Slug,
		Sender:      req.Sender,
		Category:    req.Category,
		Severity:    req.Severity,
		Content:     req.Content,
		Description: req.Description,
		Status:      "NEW", // Default status
		Labels:      req.Labels,
		ContentType: req.ContentType,
		Created:     time.Now(),
		Modified:    time.Now(),
	}

	if notification.ContentType == "" {
		notification.ContentType = "text/plain"
	}

	labelsJSON, _ := json.Marshal(notification.Labels)

	query := `
		INSERT INTO notifications (id, slug, sender, category, severity, content, description,
		                         status, labels, content_type, created, modified)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	_, err := s.db.Exec(query, notification.ID, notification.Slug, notification.Sender,
		notification.Category, notification.Severity, notification.Content, notification.Description,
		notification.Status, labelsJSON, notification.ContentType, notification.Created, notification.Modified)
	if err != nil {
		return "", fmt.Errorf("failed to create notification: %w", err)
	}

	// Process subscriptions for this notification
	go s.processNotificationSubscriptions(notification)

	return notification.ID, nil
}

func (s *Service) UpdateNotification(id string, req *models.NotificationRequest) error {
	labelsJSON, _ := json.Marshal(req.Labels)

	contentType := req.ContentType
	if contentType == "" {
		contentType = "text/plain"
	}

	query := `
		UPDATE notifications 
		SET slug = $2, sender = $3, category = $4, severity = $5, content = $6,
		    description = $7, labels = $8, content_type = $9, modified = $10
		WHERE id = $1
	`

	_, err := s.db.Exec(query, id, req.Slug, req.Sender, req.Category, req.Severity,
		req.Content, req.Description, labelsJSON, contentType, time.Now())
	if err != nil {
		return fmt.Errorf("failed to update notification: %w", err)
	}

	return nil
}

func (s *Service) DeleteNotification(id string) error {
	query := `DELETE FROM notifications WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete notification: %w", err)
	}
	return nil
}

func (s *Service) DeleteNotificationBySlug(slug string) error {
	query := `DELETE FROM notifications WHERE slug = $1`
	_, err := s.db.Exec(query, slug)
	if err != nil {
		return fmt.Errorf("failed to delete notification: %w", err)
	}
	return nil
}

func (s *Service) CleanupNotifications(ageInMilliseconds int64) (int64, error) {
	cutoffTime := time.Now().Add(-time.Duration(ageInMilliseconds) * time.Millisecond)

	query := `DELETE FROM notifications WHERE created < $1`
	result, err := s.db.Exec(query, cutoffTime)
	if err != nil {
		return 0, fmt.Errorf("failed to cleanup notifications: %w", err)
	}

	count, _ := result.RowsAffected()
	return count, nil
}

// Subscription methods
func (s *Service) GetAllSubscriptions(limit, offset int) ([]models.Subscription, error) {
	query := `
		SELECT id, name, slug, description, receiver, subscribed_categories, subscribed_labels,
		       channels, resend_limit, resend_interval, admin_state, created, modified
		FROM subscriptions
		ORDER BY created DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query subscriptions: %w", err)
	}
	defer rows.Close()

	var subscriptions []models.Subscription
	for rows.Next() {
		var subscription models.Subscription
		var subscribedCategoriesJSON, subscribedLabelsJSON, channelsJSON []byte

		err := rows.Scan(
			&subscription.ID, &subscription.Name, &subscription.Slug, &subscription.Description,
			&subscription.Receiver, &subscribedCategoriesJSON, &subscribedLabelsJSON,
			&channelsJSON, &subscription.ResendLimit, &subscription.ResendInterval,
			&subscription.AdminState, &subscription.Created, &subscription.Modified,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan subscription: %w", err)
		}

		// Unmarshal JSON fields
		if len(subscribedCategoriesJSON) > 0 {
			json.Unmarshal(subscribedCategoriesJSON, &subscription.SubscribedCategories)
		}
		if len(subscribedLabelsJSON) > 0 {
			json.Unmarshal(subscribedLabelsJSON, &subscription.SubscribedLabels)
		}
		if len(channelsJSON) > 0 {
			json.Unmarshal(channelsJSON, &subscription.Channels)
		}

		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions, nil
}

func (s *Service) GetSubscriptionByID(id string) (*models.Subscription, error) {
	query := `
		SELECT id, name, slug, description, receiver, subscribed_categories, subscribed_labels,
		       channels, resend_limit, resend_interval, admin_state, created, modified
		FROM subscriptions
		WHERE id = $1
	`

	var subscription models.Subscription
	var subscribedCategoriesJSON, subscribedLabelsJSON, channelsJSON []byte

	err := s.db.QueryRow(query, id).Scan(
		&subscription.ID, &subscription.Name, &subscription.Slug, &subscription.Description,
		&subscription.Receiver, &subscribedCategoriesJSON, &subscribedLabelsJSON,
		&channelsJSON, &subscription.ResendLimit, &subscription.ResendInterval,
		&subscription.AdminState, &subscription.Created, &subscription.Modified,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription: %w", err)
	}

	// Unmarshal JSON fields
	if len(subscribedCategoriesJSON) > 0 {
		json.Unmarshal(subscribedCategoriesJSON, &subscription.SubscribedCategories)
	}
	if len(subscribedLabelsJSON) > 0 {
		json.Unmarshal(subscribedLabelsJSON, &subscription.SubscribedLabels)
	}
	if len(channelsJSON) > 0 {
		json.Unmarshal(channelsJSON, &subscription.Channels)
	}

	return &subscription, nil
}

func (s *Service) GetSubscriptionByName(name string) (*models.Subscription, error) {
	query := `
		SELECT id, name, slug, description, receiver, subscribed_categories, subscribed_labels,
		       channels, resend_limit, resend_interval, admin_state, created, modified
		FROM subscriptions
		WHERE name = $1
	`

	var subscription models.Subscription
	var subscribedCategoriesJSON, subscribedLabelsJSON, channelsJSON []byte

	err := s.db.QueryRow(query, name).Scan(
		&subscription.ID, &subscription.Name, &subscription.Slug, &subscription.Description,
		&subscription.Receiver, &subscribedCategoriesJSON, &subscribedLabelsJSON,
		&channelsJSON, &subscription.ResendLimit, &subscription.ResendInterval,
		&subscription.AdminState, &subscription.Created, &subscription.Modified,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription: %w", err)
	}

	// Unmarshal JSON fields
	if len(subscribedCategoriesJSON) > 0 {
		json.Unmarshal(subscribedCategoriesJSON, &subscription.SubscribedCategories)
	}
	if len(subscribedLabelsJSON) > 0 {
		json.Unmarshal(subscribedLabelsJSON, &subscription.SubscribedLabels)
	}
	if len(channelsJSON) > 0 {
		json.Unmarshal(channelsJSON, &subscription.Channels)
	}

	return &subscription, nil
}

func (s *Service) CreateSubscription(req *models.SubscriptionRequest) (string, error) {
	subscription := &models.Subscription{
		ID:                   uuid.New().String(),
		Name:                 req.Name,
		Slug:                 req.Slug,
		Description:          req.Description,
		Receiver:             req.Receiver,
		SubscribedCategories: req.SubscribedCategories,
		SubscribedLabels:     req.SubscribedLabels,
		Channels:             req.Channels,
		ResendLimit:          req.ResendLimit,
		ResendInterval:       req.ResendInterval,
		AdminState:           req.AdminState,
		Created:              time.Now(),
		Modified:             time.Now(),
	}

	if subscription.AdminState == "" {
		subscription.AdminState = "UNLOCKED"
	}

	// Marshal JSON fields
	subscribedCategoriesJSON, _ := json.Marshal(subscription.SubscribedCategories)
	subscribedLabelsJSON, _ := json.Marshal(subscription.SubscribedLabels)
	channelsJSON, _ := json.Marshal(subscription.Channels)

	query := `
		INSERT INTO subscriptions (id, name, slug, description, receiver, subscribed_categories,
		                         subscribed_labels, channels, resend_limit, resend_interval,
		                         admin_state, created, modified)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	_, err := s.db.Exec(query, subscription.ID, subscription.Name, subscription.Slug,
		subscription.Description, subscription.Receiver, subscribedCategoriesJSON,
		subscribedLabelsJSON, channelsJSON, subscription.ResendLimit, subscription.ResendInterval,
		subscription.AdminState, subscription.Created, subscription.Modified)
	if err != nil {
		return "", fmt.Errorf("failed to create subscription: %w", err)
	}

	return subscription.ID, nil
}

func (s *Service) UpdateSubscription(id string, req *models.SubscriptionRequest) error {
	// Marshal JSON fields
	subscribedCategoriesJSON, _ := json.Marshal(req.SubscribedCategories)
	subscribedLabelsJSON, _ := json.Marshal(req.SubscribedLabels)
	channelsJSON, _ := json.Marshal(req.Channels)

	adminState := req.AdminState
	if adminState == "" {
		adminState = "UNLOCKED"
	}

	query := `
		UPDATE subscriptions 
		SET name = $2, slug = $3, description = $4, receiver = $5, subscribed_categories = $6,
		    subscribed_labels = $7, channels = $8, resend_limit = $9, resend_interval = $10,
		    admin_state = $11, modified = $12
		WHERE id = $1
	`

	_, err := s.db.Exec(query, id, req.Name, req.Slug, req.Description, req.Receiver,
		subscribedCategoriesJSON, subscribedLabelsJSON, channelsJSON, req.ResendLimit,
		req.ResendInterval, adminState, time.Now())
	if err != nil {
		return fmt.Errorf("failed to update subscription: %w", err)
	}

	return nil
}

func (s *Service) DeleteSubscription(id string) error {
	query := `DELETE FROM subscriptions WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete subscription: %w", err)
	}
	return nil
}

func (s *Service) TransmitNotification(id string) error {
	notification, err := s.GetNotificationByID(id)
	if err != nil {
		return err
	}

	// Update notification status to indicate transmission attempt
	query := `UPDATE notifications SET status = $1, modified = $2 WHERE id = $3`
	_, err = s.db.Exec(query, "PROCESSED", time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to update notification status: %w", err)
	}

	// Process subscriptions for this notification
	go s.processNotificationSubscriptions(notification)

	return nil
}

// Helper methods
func (s *Service) processNotificationSubscriptions(notification *models.Notification) {
	subscriptions, err := s.GetAllSubscriptions(1000, 0) // Get all subscriptions
	if err != nil {
		return
	}

	for _, subscription := range subscriptions {
		if s.matchesSubscription(notification, &subscription) {
			// In a real implementation, this would send the notification
			// through the configured channels (email, webhook, etc.)
			// For now, we just log that it would be transmitted
		}
	}
}

func (s *Service) matchesSubscription(notification *models.Notification, subscription *models.Subscription) bool {
	// Check if notification matches subscription criteria
	if subscription.AdminState != "UNLOCKED" {
		return false
	}

	// Check categories
	if len(subscription.SubscribedCategories) > 0 {
		found := false
		for _, category := range subscription.SubscribedCategories {
			if category == notification.Category {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	// Check labels
	if len(subscription.SubscribedLabels) > 0 {
		found := false
		for _, subLabel := range subscription.SubscribedLabels {
			for _, notifLabel := range notification.Labels {
				if subLabel == notifLabel {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
