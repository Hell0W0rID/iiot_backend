package management

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"runtime"
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

// Service management methods
func (s *Service) GetAllServices() ([]models.ServiceInfo, error) {
	query := `
		SELECT service_id, name, version, status, host, port, health_check, tags, meta, created, modified, last_seen
		FROM services
		ORDER BY created DESC
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
	_, err := s.db.Exec(query, serviceID)
	if err != nil {
		return fmt.Errorf("failed to deregister service: %w", err)
	}
	return nil
}

func (s *Service) UpdateServiceStatus(serviceID, status string) error {
	query := `UPDATE services SET status = $2, modified = $3, last_seen = $3 WHERE service_id = $1`
	_, err := s.db.Exec(query, serviceID, status, time.Now())
	if err != nil {
		return fmt.Errorf("failed to update service status: %w", err)
	}
	return nil
}

func (s *Service) GetServiceHealth(serviceID string) (*models.ServiceHealth, error) {
	service, err := s.GetServiceByID(serviceID)
	if err != nil {
		return nil, err
	}

	// Create health status based on service information
	health := &models.ServiceHealth{
		ServiceID:    serviceID,
		Name:         service.Name,
		Status:       service.Status,
		LastCheck:    service.LastSeen,
		ResponseTime: 0, // Would be calculated from actual health checks
		Message:      fmt.Sprintf("Service %s is %s", service.Name, service.Status),
		Checks:       []models.HealthCheck{},
	}

	// Add basic health checks
	health.Checks = append(health.Checks, models.HealthCheck{
		Name:      "service_status",
		Status:    service.Status,
		Message:   fmt.Sprintf("Service status: %s", service.Status),
		Timestamp: service.LastSeen,
	})

	return health, nil
}

func (s *Service) GetAllServiceHealth() ([]models.ServiceHealth, error) {
	services, err := s.GetAllServices()
	if err != nil {
		return nil, err
	}

	var healthStatuses []models.ServiceHealth
	for _, service := range services {
		health, err := s.GetServiceHealth(service.ServiceID)
		if err != nil {
			continue // Skip services with health check errors
		}
		healthStatuses = append(healthStatuses, *health)
	}

	return healthStatuses, nil
}

func (s *Service) PerformOperation(operation *models.ServiceOperation) (map[string]interface{}, error) {
	service, err := s.GetServiceByID(operation.ServiceID)
	if err != nil {
		return nil, fmt.Errorf("service not found: %w", err)
	}

	result := make(map[string]interface{})
	result["serviceId"] = operation.ServiceID
	result["serviceName"] = service.Name
	result["operation"] = operation.Operation
	result["timestamp"] = time.Now()

	switch operation.Operation {
	case "start":
		err = s.UpdateServiceStatus(operation.ServiceID, "UP")
		result["status"] = "UP"
		result["message"] = "Service started successfully"
	case "stop":
		err = s.UpdateServiceStatus(operation.ServiceID, "DOWN")
		result["status"] = "DOWN"
		result["message"] = "Service stopped successfully"
	case "restart":
		// First stop, then start
		err = s.UpdateServiceStatus(operation.ServiceID, "DOWN")
		if err == nil {
			time.Sleep(100 * time.Millisecond) // Brief pause
			err = s.UpdateServiceStatus(operation.ServiceID, "UP")
		}
		result["status"] = "UP"
		result["message"] = "Service restarted successfully"
	case "health-check":
		health, healthErr := s.GetServiceHealth(operation.ServiceID)
		if healthErr != nil {
			err = healthErr
		} else {
			result["health"] = health
			result["message"] = "Health check completed"
		}
	default:
		return nil, fmt.Errorf("unsupported operation: %s", operation.Operation)
	}

	if err != nil {
		result["error"] = err.Error()
		result["success"] = false
	} else {
		result["success"] = true
	}

	return result, err
}

func (s *Service) GetServiceConfig(serviceID string) (*models.ServiceConfig, error) {
	query := `
		SELECT service_id, config, version, created, modified
		FROM service_configs
		WHERE service_id = $1
		ORDER BY created DESC
		LIMIT 1
	`

	var config models.ServiceConfig
	var configJSON []byte

	err := s.db.QueryRow(query, serviceID).Scan(
		&config.ServiceID, &configJSON, &config.Version,
		&config.Created, &config.Modified,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get service config: %w", err)
	}

	if len(configJSON) > 0 {
		json.Unmarshal(configJSON, &config.Config)
	}

	return &config, nil
}

func (s *Service) UpdateServiceConfig(serviceID string, configData map[string]interface{}) error {
	configJSON, _ := json.Marshal(configData)

	// Check if config exists
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM service_configs WHERE service_id = $1)", serviceID).Scan(&exists)
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

func (s *Service) GetServiceMetrics(serviceID string) (*models.ServiceMetrics, error) {
	service, err := s.GetServiceByID(serviceID)
	if err != nil {
		return nil, err
	}

	// In a real implementation, this would collect actual metrics
	// For now, return mock metrics
	metrics := &models.ServiceMetrics{
		ServiceID:      serviceID,
		Name:           service.Name,
		CPUUsage:       0.0,    // Would be collected from system
		MemoryUsage:    0.0,    // Would be collected from system
		DiskUsage:      0.0,    // Would be collected from system
		NetworkIn:      0,      // Would be collected from system
		NetworkOut:     0,      // Would be collected from system
		RequestCount:   0,      // Would be collected from application metrics
		ErrorCount:     0,      // Would be collected from application metrics
		ResponseTime:   0.0,    // Would be calculated from request logs
		Uptime:         0,      // Would be calculated from service start time
		CustomMetrics:  make(map[string]float64),
		Timestamp:      time.Now(),
	}

	return metrics, nil
}

func (s *Service) GetSystemInfo() (*models.SystemInfo, error) {
	// Get system statistics
	var totalServices, activeServices, totalDevices, activeDevices int
	var totalEvents int64

	// Count services
	err := s.db.QueryRow("SELECT COUNT(*) FROM services").Scan(&totalServices)
	if err != nil {
		return nil, fmt.Errorf("failed to count services: %w", err)
	}

	err = s.db.QueryRow("SELECT COUNT(*) FROM services WHERE status = 'UP'").Scan(&activeServices)
	if err != nil {
		return nil, fmt.Errorf("failed to count active services: %w", err)
	}

	// Count devices
	err = s.db.QueryRow("SELECT COUNT(*) FROM devices").Scan(&totalDevices)
	if err != nil {
		// If devices table doesn't exist, set to 0
		totalDevices = 0
	}

	err = s.db.QueryRow("SELECT COUNT(*) FROM devices WHERE admin_state = 'UNLOCKED' AND operating_state = 'UP'").Scan(&activeDevices)
	if err != nil {
		activeDevices = 0
	}

	// Count events
	err = s.db.QueryRow("SELECT COUNT(*) FROM events").Scan(&totalEvents)
	if err != nil {
		totalEvents = 0
	}

	systemMetrics, err := s.GetSystemMetrics()
	if err != nil {
		// If we can't get metrics, create empty metrics
		systemMetrics = &models.SystemMetrics{}
	}

	systemInfo := &models.SystemInfo{
		Version:        "1.0.0",
		Services:       totalServices,
		ActiveServices: activeServices,
		TotalDevices:   totalDevices,
		ActiveDevices:  activeDevices,
		TotalEvents:    totalEvents,
		SystemMetrics:  *systemMetrics,
		Timestamp:      time.Now(),
	}

	return systemInfo, nil
}

func (s *Service) GetSystemMetrics() (*models.SystemMetrics, error) {
	// Get system metrics using runtime package
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	metrics := &models.SystemMetrics{
		CPUUsage:    0.0, // Would require additional library to get CPU usage
		MemoryUsage: float64(memStats.Alloc) / float64(memStats.Sys) * 100,
		DiskUsage:   0.0, // Would require additional library to get disk usage
		NetworkIn:   0,   // Would require additional library to get network stats
		NetworkOut:  0,   // Would require additional library to get network stats
		LoadAverage: 0.0, // Would require additional library to get load average
		Processes:   runtime.NumGoroutine(),
		Uptime:      0, // Would be calculated from service start time
	}

	return metrics, nil
}
