package scheduler

import (
	"database/sql"
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

// Interval methods
func (s *Service) GetAllIntervals(limit, offset int) ([]models.Interval, error) {
	query := `
		SELECT id, name, start_time, end_time, interval_time, run_once, created, modified
		FROM intervals
		ORDER BY created DESC
		LIMIT $1 OFFSET $2
	`
	
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query intervals: %w", err)
	}
	defer rows.Close()

	var intervals []models.Interval
	for rows.Next() {
		var interval models.Interval
		err := rows.Scan(
			&interval.ID, &interval.Name, &interval.Start, &interval.End,
			&interval.Interval, &interval.RunOnce, &interval.Created, &interval.Modified,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan interval: %w", err)
		}
		intervals = append(intervals, interval)
	}

	return intervals, nil
}

func (s *Service) GetIntervalByID(id string) (*models.Interval, error) {
	query := `
		SELECT id, name, start_time, end_time, interval_time, run_once, created, modified
		FROM intervals
		WHERE id = $1
	`
	
	var interval models.Interval
	err := s.db.QueryRow(query, id).Scan(
		&interval.ID, &interval.Name, &interval.Start, &interval.End,
		&interval.Interval, &interval.RunOnce, &interval.Created, &interval.Modified,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get interval: %w", err)
	}

	return &interval, nil
}

func (s *Service) GetIntervalByName(name string) (*models.Interval, error) {
	query := `
		SELECT id, name, start_time, end_time, interval_time, run_once, created, modified
		FROM intervals
		WHERE name = $1
	`
	
	var interval models.Interval
	err := s.db.QueryRow(query, name).Scan(
		&interval.ID, &interval.Name, &interval.Start, &interval.End,
		&interval.Interval, &interval.RunOnce, &interval.Created, &interval.Modified,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get interval: %w", err)
	}

	return &interval, nil
}

func (s *Service) CreateInterval(req *models.IntervalRequest) (string, error) {
	interval := &models.Interval{
		ID:       uuid.New().String(),
		Name:     req.Name,
		Start:    req.Start,
		End:      req.End,
		Interval: req.Interval,
		RunOnce:  req.RunOnce,
		Created:  time.Now(),
		Modified: time.Now(),
	}

	query := `
		INSERT INTO intervals (id, name, start_time, end_time, interval_time, run_once, created, modified)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	
	_, err := s.db.Exec(query, interval.ID, interval.Name, interval.Start, interval.End,
		interval.Interval, interval.RunOnce, interval.Created, interval.Modified)
	if err != nil {
		return "", fmt.Errorf("failed to create interval: %w", err)
	}

	return interval.ID, nil
}

func (s *Service) UpdateInterval(id string, req *models.IntervalRequest) error {
	query := `
		UPDATE intervals 
		SET name = $2, start_time = $3, end_time = $4, interval_time = $5, 
		    run_once = $6, modified = $7
		WHERE id = $1
	`
	
	_, err := s.db.Exec(query, id, req.Name, req.Start, req.End,
		req.Interval, req.RunOnce, time.Now())
	if err != nil {
		return fmt.Errorf("failed to update interval: %w", err)
	}

	return nil
}

func (s *Service) DeleteInterval(id string) error {
	// First delete associated interval actions
	_, err := s.db.Exec("DELETE FROM interval_actions WHERE interval_name = (SELECT name FROM intervals WHERE id = $1)", id)
	if err != nil {
		return fmt.Errorf("failed to delete associated interval actions: %w", err)
	}

	// Then delete the interval
	query := `DELETE FROM intervals WHERE id = $1`
	_, err = s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete interval: %w", err)
	}
	return nil
}

// Interval Action methods
func (s *Service) GetIntervalActions(intervalName string, limit, offset int) ([]models.IntervalAction, error) {
	query := `
		SELECT id, name, interval_name, protocol, host, port, path, parameters,
		       http_method, address, publisher, target, user, password, topic, created, modified
		FROM interval_actions
		WHERE ($1 = '' OR interval_name = $1)
		ORDER BY created DESC
		LIMIT $2 OFFSET $3
	`
	
	rows, err := s.db.Query(query, intervalName, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query interval actions: %w", err)
	}
	defer rows.Close()

	var actions []models.IntervalAction
	for rows.Next() {
		var action models.IntervalAction
		err := rows.Scan(
			&action.ID, &action.Name, &action.IntervalName, &action.Protocol,
			&action.Host, &action.Port, &action.Path, &action.Parameters,
			&action.HTTPMethod, &action.Address, &action.Publisher, &action.Target,
			&action.User, &action.Password, &action.Topic, &action.Created, &action.Modified,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan interval action: %w", err)
		}
		actions = append(actions, action)
	}

	return actions, nil
}

func (s *Service) GetIntervalActionByID(id string) (*models.IntervalAction, error) {
	query := `
		SELECT id, name, interval_name, protocol, host, port, path, parameters,
		       http_method, address, publisher, target, user, password, topic, created, modified
		FROM interval_actions
		WHERE id = $1
	`
	
	var action models.IntervalAction
	err := s.db.QueryRow(query, id).Scan(
		&action.ID, &action.Name, &action.IntervalName, &action.Protocol,
		&action.Host, &action.Port, &action.Path, &action.Parameters,
		&action.HTTPMethod, &action.Address, &action.Publisher, &action.Target,
		&action.User, &action.Password, &action.Topic, &action.Created, &action.Modified,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get interval action: %w", err)
	}

	return &action, nil
}

func (s *Service) GetIntervalActionByName(name string) (*models.IntervalAction, error) {
	query := `
		SELECT id, name, interval_name, protocol, host, port, path, parameters,
		       http_method, address, publisher, target, user, password, topic, created, modified
		FROM interval_actions
		WHERE name = $1
	`
	
	var action models.IntervalAction
	err := s.db.QueryRow(query, name).Scan(
		&action.ID, &action.Name, &action.IntervalName, &action.Protocol,
		&action.Host, &action.Port, &action.Path, &action.Parameters,
		&action.HTTPMethod, &action.Address, &action.Publisher, &action.Target,
		&action.User, &action.Password, &action.Topic, &action.Created, &action.Modified,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get interval action: %w", err)
	}

	return &action, nil
}

func (s *Service) CreateIntervalAction(req *models.IntervalActionRequest) (string, error) {
	// Verify that the interval exists
	_, err := s.GetIntervalByName(req.IntervalName)
	if err != nil {
		return "", fmt.Errorf("interval %s not found: %w", req.IntervalName, err)
	}

	action := &models.IntervalAction{
		ID:           uuid.New().String(),
		Name:         req.Name,
		IntervalName: req.IntervalName,
		Protocol:     req.Protocol,
		Host:         req.Host,
		Port:         req.Port,
		Path:         req.Path,
		Parameters:   req.Parameters,
		HTTPMethod:   req.HTTPMethod,
		Address:      req.Address,
		Publisher:    req.Publisher,
		Target:       req.Target,
		User:         req.User,
		Password:     req.Password,
		Topic:        req.Topic,
		Created:      time.Now(),
		Modified:     time.Now(),
	}

	query := `
		INSERT INTO interval_actions (id, name, interval_name, protocol, host, port, path,
		                            parameters, http_method, address, publisher, target,
		                            user, password, topic, created, modified)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
	`
	
	_, err = s.db.Exec(query, action.ID, action.Name, action.IntervalName, action.Protocol,
		action.Host, action.Port, action.Path, action.Parameters, action.HTTPMethod,
		action.Address, action.Publisher, action.Target, action.User, action.Password,
		action.Topic, action.Created, action.Modified)
	if err != nil {
		return "", fmt.Errorf("failed to create interval action: %w", err)
	}

	return action.ID, nil
}

func (s *Service) UpdateIntervalAction(id string, req *models.IntervalActionRequest) error {
	// Verify that the interval exists
	_, err := s.GetIntervalByName(req.IntervalName)
	if err != nil {
		return fmt.Errorf("interval %s not found: %w", req.IntervalName, err)
	}

	query := `
		UPDATE interval_actions 
		SET name = $2, interval_name = $3, protocol = $4, host = $5, port = $6, path = $7,
		    parameters = $8, http_method = $9, address = $10, publisher = $11, target = $12,
		    user = $13, password = $14, topic = $15, modified = $16
		WHERE id = $1
	`
	
	_, err = s.db.Exec(query, id, req.Name, req.IntervalName, req.Protocol, req.Host,
		req.Port, req.Path, req.Parameters, req.HTTPMethod, req.Address, req.Publisher,
		req.Target, req.User, req.Password, req.Topic, time.Now())
	if err != nil {
		return fmt.Errorf("failed to update interval action: %w", err)
	}

	return nil
}

func (s *Service) DeleteIntervalAction(id string) error {
	query := `DELETE FROM interval_actions WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete interval action: %w", err)
	}
	return nil
}

// Schedule Status methods
func (s *Service) GetScheduleStatus(intervalName string) (*models.ScheduleStatus, error) {
	// Get interval information
	interval, err := s.GetIntervalByName(intervalName)
	if err != nil {
		return nil, err
	}

	// For now, return a mock status - in a real implementation this would
	// track actual execution status
	status := &models.ScheduleStatus{
		IntervalName:   intervalName,
		LastExecution:  time.Now().Add(-time.Hour),
		NextExecution:  time.Now().Add(time.Hour),
		ExecutionCount: 10, // Mock value
		Status:         "active",
		Message:        fmt.Sprintf("Interval %s is running normally", intervalName),
	}

	// Calculate next execution based on interval
	if interval.RunOnce {
		status.Status = "completed"
		status.Message = "One-time interval completed"
	}

	return status, nil
}

func (s *Service) GetAllScheduleStatuses() ([]models.ScheduleStatus, error) {
	// Get all intervals
	intervals, err := s.GetAllIntervals(1000, 0) // Get all intervals
	if err != nil {
		return nil, err
	}

	var statuses []models.ScheduleStatus
	for _, interval := range intervals {
		status, err := s.GetScheduleStatus(interval.Name)
		if err != nil {
			// Continue with other intervals if one fails
			continue
		}
		statuses = append(statuses, *status)
	}

	return statuses, nil
}
