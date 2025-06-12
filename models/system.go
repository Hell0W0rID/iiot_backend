package models

import (
	"time"
)

// ServiceInfo represents information about a service
type ServiceInfo struct {
	ServiceID   string                 `json:"serviceId" db:"service_id"`
	Name        string                 `json:"name" db:"name"`
	Version     string                 `json:"version" db:"version"`
	Status      string                 `json:"status" db:"status"`
	Host        string                 `json:"host" db:"host"`
	Port        int                    `json:"port" db:"port"`
	HealthCheck string                 `json:"healthCheck" db:"health_check"`
	Tags        []string               `json:"tags" db:"tags"`
	Meta        map[string]string      `json:"meta" db:"meta"`
	Created     time.Time              `json:"created" db:"created"`
	Modified    time.Time              `json:"modified" db:"modified"`
	LastSeen    time.Time              `json:"lastSeen" db:"last_seen"`
}

// ServiceHealth represents the health status of a service
type ServiceHealth struct {
	ServiceID    string    `json:"serviceId"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	LastCheck    time.Time `json:"lastCheck"`
	ResponseTime int64     `json:"responseTime"`
	Message      string    `json:"message"`
	Checks       []HealthCheck `json:"checks"`
}

// HealthCheck represents a health check result
type HealthCheck struct {
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	Message      string    `json:"message"`
	Timestamp    time.Time `json:"timestamp"`
	ResponseTime int64     `json:"responseTime"`
}

// ServiceOperation represents an operation to perform on a service
type ServiceOperation struct {
	ServiceID string `json:"serviceId" validate:"required"`
	Operation string `json:"operation" validate:"required"`
	Force     bool   `json:"force"`
}

// ServiceConfig represents service configuration
type ServiceConfig struct {
	ServiceID string                 `json:"serviceId" db:"service_id"`
	Config    map[string]interface{} `json:"config" db:"config"`
	Version   string                 `json:"version" db:"version"`
	Created   time.Time              `json:"created" db:"created"`
	Modified  time.Time              `json:"modified" db:"modified"`
}

// ServiceMetrics represents service metrics
type ServiceMetrics struct {
	ServiceID      string            `json:"serviceId"`
	Name           string            `json:"name"`
	CPUUsage       float64           `json:"cpuUsage"`
	MemoryUsage    float64           `json:"memoryUsage"`
	DiskUsage      float64           `json:"diskUsage"`
	NetworkIn      int64             `json:"networkIn"`
	NetworkOut     int64             `json:"networkOut"`
	RequestCount   int64             `json:"requestCount"`
	ErrorCount     int64             `json:"errorCount"`
	ResponseTime   float64           `json:"responseTime"`
	Uptime         int64             `json:"uptime"`
	CustomMetrics  map[string]float64 `json:"customMetrics"`
	Timestamp      time.Time         `json:"timestamp"`
}

// SystemInfo represents overall system information
type SystemInfo struct {
	Version     string                 `json:"version"`
	Services    int                    `json:"services"`
	ActiveServices int                 `json:"activeServices"`
	TotalDevices   int                 `json:"totalDevices"`
	ActiveDevices  int                 `json:"activeDevices"`
	TotalEvents    int64               `json:"totalEvents"`
	SystemMetrics  SystemMetrics       `json:"systemMetrics"`
	Timestamp      time.Time           `json:"timestamp"`
}

// SystemMetrics represents overall system metrics
type SystemMetrics struct {
	CPUUsage     float64 `json:"cpuUsage"`
	MemoryUsage  float64 `json:"memoryUsage"`
	DiskUsage    float64 `json:"diskUsage"`
	NetworkIn    int64   `json:"networkIn"`
	NetworkOut   int64   `json:"networkOut"`
	LoadAverage  float64 `json:"loadAverage"`
	Processes    int     `json:"processes"`
	Uptime       int64   `json:"uptime"`
}
