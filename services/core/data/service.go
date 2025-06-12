package data

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

// Event methods
func (s *Service) GetEvents(filter models.EventFilter) ([]models.Event, error) {
	query := `
		SELECT id, device_name, profile_name, source_name, origin, tags, created, modified
		FROM events
		WHERE ($1 = '' OR device_name = $1)
		  AND ($2 = '' OR profile_name = $2)
		  AND ($3 = '' OR source_name = $3)
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

	rows, err := s.db.Query(query, filter.DeviceName, filter.ProfileName, filter.SourceName,
		start, end, filter.Limit, filter.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query events: %w", err)
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		var tagsJSON []byte

		err := rows.Scan(
			&event.ID, &event.DeviceName, &event.ProfileName, &event.SourceName,
			&event.Origin, &tagsJSON, &event.Created, &event.Modified,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan event: %w", err)
		}

		if len(tagsJSON) > 0 {
			json.Unmarshal(tagsJSON, &event.Tags)
		}

		// Get readings for this event
		readings, err := s.getReadingsByEventID(event.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get readings for event %s: %w", event.ID, err)
		}
		event.Readings = readings

		events = append(events, event)
	}

	return events, nil
}

func (s *Service) GetEventByID(id string) (*models.Event, error) {
	query := `
		SELECT id, device_name, profile_name, source_name, origin, tags, created, modified
		FROM events
		WHERE id = $1
	`

	var event models.Event
	var tagsJSON []byte

	err := s.db.QueryRow(query, id).Scan(
		&event.ID, &event.DeviceName, &event.ProfileName, &event.SourceName,
		&event.Origin, &tagsJSON, &event.Created, &event.Modified,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get event: %w", err)
	}

	if len(tagsJSON) > 0 {
		json.Unmarshal(tagsJSON, &event.Tags)
	}

	// Get readings for this event
	readings, err := s.getReadingsByEventID(event.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get readings for event: %w", err)
	}
	event.Readings = readings

	return &event, nil
}

func (s *Service) CreateEvent(req *models.EventRequest) (string, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Create event
	eventID := uuid.New().String()
	now := time.Now()
	origin := req.Origin
	if origin == 0 {
		origin = now.UnixNano()
	}

	tagsJSON, _ := json.Marshal(req.Tags)

	eventQuery := `
		INSERT INTO events (id, device_name, profile_name, source_name, origin, tags, created, modified)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err = tx.Exec(eventQuery, eventID, req.DeviceName, req.ProfileName, req.SourceName,
		origin, tagsJSON, now, now)
	if err != nil {
		return "", fmt.Errorf("failed to create event: %w", err)
	}

	// Create readings
	readingQuery := `
		INSERT INTO readings (id, event_id, device_name, resource_name, profile_name, value_type,
		                     value, binary_value, media_type, units, tags, origin, created, modified)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`

	for _, readingReq := range req.Readings {
		readingID := uuid.New().String()
		readingOrigin := readingReq.Origin
		if readingOrigin == 0 {
			readingOrigin = origin
		}

		readingTagsJSON, _ := json.Marshal(readingReq.Tags)

		_, err = tx.Exec(readingQuery, readingID, eventID, readingReq.DeviceName,
			readingReq.ResourceName, readingReq.ProfileName, readingReq.ValueType,
			readingReq.Value, readingReq.BinaryValue, readingReq.MediaType,
			readingReq.Units, readingTagsJSON, readingOrigin, now, now)
		if err != nil {
			return "", fmt.Errorf("failed to create reading: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return eventID, nil
}

func (s *Service) DeleteEvent(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Delete readings first
	_, err = tx.Exec("DELETE FROM readings WHERE event_id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete readings: %w", err)
	}

	// Delete event
	_, err = tx.Exec("DELETE FROM events WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete event: %w", err)
	}

	return tx.Commit()
}

func (s *Service) DeleteEventsByDevice(deviceName string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Delete readings first
	_, err = tx.Exec("DELETE FROM readings WHERE device_name = $1", deviceName)
	if err != nil {
		return fmt.Errorf("failed to delete readings: %w", err)
	}

	// Delete events
	_, err = tx.Exec("DELETE FROM events WHERE device_name = $1", deviceName)
	if err != nil {
		return fmt.Errorf("failed to delete events: %w", err)
	}

	return tx.Commit()
}

func (s *Service) DeleteEventsByAge(ageInMilliseconds int64) error {
	cutoffTime := time.Now().Add(-time.Duration(ageInMilliseconds) * time.Millisecond)

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Delete readings first
	_, err = tx.Exec("DELETE FROM readings WHERE created < $1", cutoffTime)
	if err != nil {
		return fmt.Errorf("failed to delete readings: %w", err)
	}

	// Delete events
	_, err = tx.Exec("DELETE FROM events WHERE created < $1", cutoffTime)
	if err != nil {
		return fmt.Errorf("failed to delete events: %w", err)
	}

	return tx.Commit()
}

// Reading methods
func (s *Service) GetReadings(filter models.ReadingFilter) ([]models.Reading, error) {
	query := `
		SELECT id, event_id, device_name, resource_name, profile_name, value_type,
		       value, binary_value, media_type, units, tags, origin, created, modified
		FROM readings
		WHERE ($1 = '' OR device_name = $1)
		  AND ($2 = '' OR resource_name = $2)
		  AND ($3 = '' OR profile_name = $3)
		  AND ($4 = '' OR value_type = $4)
		  AND ($5::timestamp IS NULL OR created >= $5)
		  AND ($6::timestamp IS NULL OR created <= $6)
		ORDER BY created DESC
		LIMIT $7 OFFSET $8
	`

	var start, end interface{}
	if !filter.Start.IsZero() {
		start = filter.Start
	}
	if !filter.End.IsZero() {
		end = filter.End
	}

	rows, err := s.db.Query(query, filter.DeviceName, filter.ResourceName, filter.ProfileName,
		filter.ValueType, start, end, filter.Limit, filter.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query readings: %w", err)
	}
	defer rows.Close()

	var readings []models.Reading
	for rows.Next() {
		var reading models.Reading
		var tagsJSON []byte

		err := rows.Scan(
			&reading.ID, &reading.EventID, &reading.DeviceName, &reading.ResourceName,
			&reading.ProfileName, &reading.ValueType, &reading.Value, &reading.BinaryValue,
			&reading.MediaType, &reading.Units, &tagsJSON, &reading.Origin,
			&reading.Created, &reading.Modified,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan reading: %w", err)
		}

		if len(tagsJSON) > 0 {
			json.Unmarshal(tagsJSON, &reading.Tags)
		}

		readings = append(readings, reading)
	}

	return readings, nil
}

func (s *Service) GetReadingByID(id string) (*models.Reading, error) {
	query := `
		SELECT id, event_id, device_name, resource_name, profile_name, value_type,
		       value, binary_value, media_type, units, tags, origin, created, modified
		FROM readings
		WHERE id = $1
	`

	var reading models.Reading
	var tagsJSON []byte

	err := s.db.QueryRow(query, id).Scan(
		&reading.ID, &reading.EventID, &reading.DeviceName, &reading.ResourceName,
		&reading.ProfileName, &reading.ValueType, &reading.Value, &reading.BinaryValue,
		&reading.MediaType, &reading.Units, &tagsJSON, &reading.Origin,
		&reading.Created, &reading.Modified,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get reading: %w", err)
	}

	if len(tagsJSON) > 0 {
		json.Unmarshal(tagsJSON, &reading.Tags)
	}

	return &reading, nil
}

func (s *Service) DeleteReading(id string) error {
	query := `DELETE FROM readings WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete reading: %w", err)
	}
	return nil
}

// Helper methods
func (s *Service) getReadingsByEventID(eventID string) ([]models.Reading, error) {
	query := `
		SELECT id, event_id, device_name, resource_name, profile_name, value_type,
		       value, binary_value, media_type, units, tags, origin, created, modified
		FROM readings
		WHERE event_id = $1
		ORDER BY created ASC
	`

	rows, err := s.db.Query(query, eventID)
	if err != nil {
		return nil, fmt.Errorf("failed to query readings: %w", err)
	}
	defer rows.Close()

	var readings []models.Reading
	for rows.Next() {
		var reading models.Reading
		var tagsJSON []byte

		err := rows.Scan(
			&reading.ID, &reading.EventID, &reading.DeviceName, &reading.ResourceName,
			&reading.ProfileName, &reading.ValueType, &reading.Value, &reading.BinaryValue,
			&reading.MediaType, &reading.Units, &tagsJSON, &reading.Origin,
			&reading.Created, &reading.Modified,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan reading: %w", err)
		}

		if len(tagsJSON) > 0 {
			json.Unmarshal(tagsJSON, &reading.Tags)
		}

		readings = append(readings, reading)
	}

	return readings, nil
}

// Count methods
func (s *Service) GetEventCount() (int64, error) {
	var count int64
	err := s.db.QueryRow("SELECT COUNT(*) FROM events").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get event count: %w", err)
	}
	return count, nil
}

func (s *Service) GetEventCountByDevice(deviceName string) (int64, error) {
	var count int64
	err := s.db.QueryRow("SELECT COUNT(*) FROM events WHERE device_name = $1", deviceName).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get event count by device: %w", err)
	}
	return count, nil
}

func (s *Service) GetReadingCount() (int64, error) {
	var count int64
	err := s.db.QueryRow("SELECT COUNT(*) FROM readings").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get reading count: %w", err)
	}
	return count, nil
}

func (s *Service) GetReadingCountByDevice(deviceName string) (int64, error) {
	var count int64
	err := s.db.QueryRow("SELECT COUNT(*) FROM readings WHERE device_name = $1", deviceName).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get reading count by device: %w", err)
	}
	return count, nil
}
