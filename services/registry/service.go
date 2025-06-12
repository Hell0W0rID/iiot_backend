package registry

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
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

func (s *Service) RegisterService(serviceInfo *models.ServiceInfo) (string, error) {
	if serviceInfo.ServiceID == "" {
		serviceInfo.ServiceID = uuid.New().String()
	}
	
	now := time.Now()
	serviceInfo.Created = now
	serviceInfo.Modified = now
	serviceInfo.LastSeen = now

	if serviceInfo.Status == "" {
		serviceInfo.Status = "UP"
	}

	// Set default health check endpoint if not provided
	if serviceInfo.HealthCheck == "" {
		serviceInfo.HealthCheck = fmt.Sprintf("http://%s:%d/health", serviceInfo.Host, serviceInfo.Port)
	}

	// Marshal JSON fields
	tagsJSON, _ := json.Marshal(serviceInfo.Tags)
	metaJSON, _ := json.Marshal(serviceInfo.Meta)

	query := `
		INSERT INTO services (service_id, name, version, status, host, port, health_check, tags, meta, created, modified, last_seen)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		ON CONFLICT (service_id) DO UPDATE SET
			name = EXCLUDED.name,
			version = EXCLUDED.version,
			status = EXCLUDED.status,
			host = EXCLUDED.host,
			port = EXCLUDED.port,
			health_check = EXCLUDED.health_check,
			tags = EXCLUDED.tags,
			meta = EXCLUDED.meta,
			modified = EXCLUDED.modified,
			last_seen = EXCLUDED.last_seen
	`

	_, err := s.db.Exec(query, serviceInfo.ServiceID, serviceInfo.Name, serviceInfo.Version,
		serviceInfo.Status, serviceInfo.Host, serviceInfo.Port, serviceInfo.HealthCheck,
		tagsJSON, metaJSON, serviceInfo.Created, serviceInfo.Modified, serviceInfo.LastSeen)
	if err != nil {
		return "", fmt.Errorf("failed to register service: %w", err)
	}

	return serviceInfo.ServiceID, nil
}

func (s *Service) DeregisterService(serviceID string) error {
	query := `DELETE FROM services WHERE service_id = $1`
	result, err := s.db.Exec(query, serviceID)
	if err != nil {
		return fmt.Errorf("failed to deregister service: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("service with ID %s not found", serviceID)
	}

	return nil
}

func (s *Service) GetAllServices() ([]models.ServiceInfo, error) {
	query := `
		SELECT service_id, name, version, status, host, port, health_check, tags, meta, created, modified, last_seen
		FROM services
		ORDER BY name, created DESC
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query services: %w", err)
	}
	defer rows.Close()

	var services []models.ServiceInfo
	for rows.Next() {
		var service models.ServiceInfo
		var tagsJSON, metaJSON []byte

		err := rows.Scan(
			&service.ServiceID, &service.Name, &service.Version, &service.Status,
			&service.Host, &service.Port, &service.HealthCheck, &tagsJSON,
			&metaJSON, &service.Created, &service.Modified, &service.LastSeen,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan service: %w", err)
		}

		// Unmarshal JSON fields
		if len(tagsJSON) > 0 {
			json.Unmarshal(tagsJSON, &service.Tags)
		}
		if len(metaJSON) > 0 {
			json.Unmarshal(metaJSON, &service.Meta)
		}

		services = append(services, service)
	}

	return services, nil
}

func (s *Service) GetServiceByID(serviceID string) (*models.ServiceInfo, error) {
	query := `
		SELECT service_id, name, version, status, host, port, health_check, tags, meta, created, modified, last_seen
		FROM services
		WHERE service_id = $1
	`

	var service models.ServiceInfo
	var tagsJSON, metaJSON []byte

	err := s.db.QueryRow(query, serviceID).Scan(
		&service.ServiceID, &service.Name, &service.Version, &service.Status,
		&service.Host, &service.Port, &service.HealthCheck, &tagsJSON,
		&metaJSON, &service.Created, &service.Modified, &service.LastSeen,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get service: %w", err)
	}

	// Unmarshal JSON fields
	if len(tagsJSON) > 0 {
		json.Unmarshal(tagsJSON, &service.Tags)
	}
	if len(metaJSON) > 0 {
		json.Unmarshal(metaJSON, &service.Meta)
	}

	return &service, nil
}

func (s *Service) GetServiceByName(name string) (*models.ServiceInfo, error) {
	query := `
		SELECT service_id, name, version, status, host, port, health_check, tags, meta, created, modified, last_seen
		FROM services
		WHERE name = $1
		ORDER BY created DESC
		LIMIT 1
	`

	var service models.ServiceInfo
	var tagsJSON, metaJSON []byte

	err := s.db.QueryRow(query, name).Scan(
		&service.ServiceID, &service.Name, &service.Version, &service.Status,
		&service.Host, &service.Port, &service.HealthCheck, &tagsJSON,
		&metaJSON, &service.Created, &service.Modified, &service.LastSeen,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get service: %w", err)
	}

	// Unmarshal JSON fields
	if len(tagsJSON) > 0 {
		json.Unmarshal(tagsJSON, &service.Tags)
	}
	if len(metaJSON) > 0 {
		json.Unmarshal(metaJSON, &service.Meta)
	}

	return &service, nil
}

func (s *Service) UpdateServiceHealth(serviceID, status, message string) error {
	query := `UPDATE services SET status = $2, modified = $3, last_seen = $3 WHERE service_id = $1`
	result, err := s.db.Exec(query, serviceID, status, time.Now())
	if err != nil {
		return fmt.Errorf("failed to update service health: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("service with ID %s not found", serviceID)
	}

	return nil
}

func (s *Service) GetServiceHealth(serviceID string) (*models.ServiceHealth, error) {
	service, err := s.GetServiceByID(serviceID)
	if err != nil {
		return nil, err
	}

	// Calculate response time based on last seen (mock calculation)
	responseTime := time.Since(service.LastSeen).Milliseconds()
	if responseTime < 0 {
		responseTime = 0
	}

	health := &models.ServiceHealth{
		ServiceID:    serviceID,
		Name:         service.Name,
		Status:       service.Status,
		LastCheck:    service.LastSeen,
		ResponseTime: responseTime,
		Message:      fmt.Sprintf("Service %s is %s", service.Name, service.Status),
		Checks: []models.HealthCheck{
			{
				Name:         "service_status",
				Status:       service.Status,
				Message:      fmt.Sprintf("Service status: %s", service.Status),
				Timestamp:    service.LastSeen,
				ResponseTime: responseTime,
			},
		},
	}

	return health, nil
}

func (s *Service) GetHealthyServices() ([]models.ServiceInfo, error) {
	query := `
		SELECT service_id, name, version, status, host, port, health_check, tags, meta, created, modified, last_seen
		FROM services
		WHERE status = 'UP'
		ORDER BY name, created DESC
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query healthy services: %w", err)
	}
	defer rows.Close()

	var services []models.ServiceInfo
	for rows.Next() {
		var service models.ServiceInfo
		var tagsJSON, metaJSON []byte

		err := rows.Scan(
			&service.ServiceID, &service.Name, &service.Version, &service.Status,
			&service.Host, &service.Port, &service.HealthCheck, &tagsJSON,
			&metaJSON, &service.Created, &service.Modified, &service.LastSeen,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan service: %w", err)
		}

		// Unmarshal JSON fields
		if len(tagsJSON) > 0 {
			json.Unmarshal(tagsJSON, &service.Tags)
		}
		if len(metaJSON) > 0 {
			json.Unmarshal(metaJSON, &service.Meta)
		}

		services = append(services, service)
	}

	return services, nil
}

func (s *Service) DiscoverServices(serviceType, tag string) ([]models.ServiceInfo, error) {
	var query string
	var args []interface{}

	if serviceType != "" && tag != "" {
		// Search by both service type (in meta) and tag
		query = `
			SELECT service_id, name, version, status, host, port, health_check, tags, meta, created, modified, last_seen
			FROM services
			WHERE status = 'UP' 
			  AND (meta->>'type' = $1 OR name ILIKE $2)
			  AND tags::text ILIKE $3
			ORDER BY name, created DESC
		`
		args = []interface{}{serviceType, "%" + serviceType + "%", "%" + tag + "%"}
	} else if serviceType != "" {
		// Search by service type only
		query = `
			SELECT service_id, name, version, status, host, port, health_check, tags, meta, created, modified, last_seen
			FROM services
			WHERE status = 'UP' 
			  AND (meta->>'type' = $1 OR name ILIKE $2)
			ORDER BY name, created DESC
		`
		args = []interface{}{serviceType, "%" + serviceType + "%"}
	} else if tag != "" {
		// Search by tag only
		query = `
			SELECT service_id, name, version, status, host, port, health_check, tags, meta, created, modified, last_seen
			FROM services
			WHERE status = 'UP' 
			  AND tags::text ILIKE $1
			ORDER BY name, created DESC
		`
		args = []interface{}{"%" + tag + "%"}
	} else {
		// Return all healthy services
		return s.GetHealthyServices()
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to discover services: %w", err)
	}
	defer rows.Close()

	var services []models.ServiceInfo
	for rows.Next() {
		var service models.ServiceInfo
		var tagsJSON, metaJSON []byte

		err := rows.Scan(
			&service.ServiceID, &service.Name, &service.Version, &service.Status,
			&service.Host, &service.Port, &service.HealthCheck, &tagsJSON,
			&metaJSON, &service.Created, &service.Modified, &service.LastSeen,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan service: %w", err)
		}

		// Unmarshal JSON fields
		if len(tagsJSON) > 0 {
			json.Unmarshal(tagsJSON, &service.Tags)
		}
		if len(metaJSON) > 0 {
			json.Unmarshal(metaJSON, &service.Meta)
		}

		// Additional filtering for tags if using JSON-based search
		if tag != "" && len(service.Tags) > 0 {
			found := false
			for _, serviceTag := range service.Tags {
				if strings.Contains(strings.ToLower(serviceTag), strings.ToLower(tag)) {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		services = append(services, service)
	}

	return services, nil
}

func (s *Service) GetServiceConfig(serviceID string) (map[string]interface{}, error) {
	query := `
		SELECT config
		FROM service_configs
		WHERE service_id = $1
		ORDER BY created DESC
		LIMIT 1
	`

	var configJSON []byte
	err := s.db.QueryRow(query, serviceID).Scan(&configJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to get service config: %w", err)
	}

	var config map[string]interface{}
	if len(configJSON) > 0 {
		err = json.Unmarshal(configJSON, &config)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal config: %w", err)
		}
	}

	return config, nil
}

func (s *Service) UpdateServiceConfig(serviceID string, config map[string]interface{}) error {
	configJSON, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Check if config exists
	var exists bool
	err = s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM service_configs WHERE service_id = $1)", serviceID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check config existence: %w", err)
	}

	now := time.Now()

	if exists {
		query := `UPDATE service_configs SET config = $2, modified = $3 WHERE service_id = $1`
		_, err = s.db.Exec(query, serviceID, configJSON, now)
	} else {
		query := `INSERT INTO service_configs (service_id, config, version, created, modified) VALUES ($1, $2, $3, $4, $5)`
		_, err = s.db.Exec(query, serviceID, configJSON, "1.0", now, now)
	}

	if err != nil {
		return fmt.Errorf("failed to update service config: %w", err)
	}

	return nil
}
