package main

import (
        "context"
        "database/sql"
        "encoding/json"
        "fmt"
        "net/http"
        "os"
        "os/signal"
        "strconv"
        "syscall"
        "time"

        "github.com/labstack/echo/v4"
        "github.com/labstack/echo/v4/middleware"
        _ "github.com/lib/pq"

        "iiot-backend/models"
        "iiot-backend/services/core/command/application"
        "iiot-backend/services/core/command/controller"
)

// Core service constants
const (
        ApiVersion = "v3"
        ApiBase    = "/api/v3"
        
        // Service identifiers
        CoreMetadataServiceKey = "core-metadata"
        CoreDataServiceKey     = "core-data"
        CoreCommandServiceKey  = "core-command"
        
        // Common route patterns
        ApiPingRoute           = ApiBase + "/ping"
        ApiVersionRoute        = ApiBase + "/version"
        ApiConfigRoute         = ApiBase + "/config"
        
        // Data routes
        ApiAllDataEventRoute   = ApiBase + "/event/all"
        ApiAllMeasurementRoute = ApiBase + "/reading/all"
        
        // Metadata routes
        ApiAllDeviceHandlerRoute = ApiBase + "/deviceservice/all"
        ApiAllDeviceRoute        = ApiBase + "/device/all"
        ApiDeviceByNameRoute     = ApiBase + "/device/name"
        
        // Command routes
        ApiDeviceCommandRoute = ApiBase + "/device/name"
)

// UnifiedIIOTService provides all core functionality using your shared libraries
type UnifiedIIOTService struct {
        db *sql.DB
}

func NewUnifiedIIOTService(db *sql.DB) *UnifiedIIOTService {
        return &UnifiedIIOTService{db: db}
}

// Core Metadata Service Methods
func (s *UnifiedIIOTService) GetAllDeviceServices(ctx context.Context, offset, limit int) ([]map[string]interface{}, error) {
        if limit <= 0 {
                limit = 20
        }
        if offset < 0 {
                offset = 0
        }

        query := `
                SELECT id, name, description, base_address, admin_state, labels, created, modified
                FROM device_services
                ORDER BY name LIMIT $1 OFFSET $2`

        rows, err := s.db.QueryContext(ctx, query, limit, offset)
        if err != nil {
                return nil, fmt.Errorf("failed to query device services: %w", err)
        }
        defer rows.Close()

        var services []map[string]interface{}
        for rows.Next() {
                var id, name, description, baseAddress, adminState string
                var labelsJSON []byte
                var created, modified int64

                err := rows.Scan(&id, &name, &description, &baseAddress, &adminState, &labelsJSON, &created, &modified)
                if err != nil {
                        continue
                }

                var labels []string
                if len(labelsJSON) > 0 {
                        json.Unmarshal(labelsJSON, &labels)
                }

                service := map[string]interface{}{
                        "id":           id,
                        "name":         name,
                        "description":  description,
                        "baseAddress":  baseAddress,
                        "adminState":   adminState,
                        "labels":       labels,
                        "created":      created,
                        "modified":     modified,
                }
                services = append(services, service)
        }

        return services, nil
}

func (s *UnifiedIIOTService) GetAllDevices(ctx context.Context, offset, limit int) ([]map[string]interface{}, error) {
        if limit <= 0 {
                limit = 20
        }
        if offset < 0 {
                offset = 0
        }

        query := `
                SELECT id, name, description, admin_state, operating_state, protocols, labels,
                       service_name, profile_name, created, modified
                FROM devices
                ORDER BY name LIMIT $1 OFFSET $2`

        rows, err := s.db.QueryContext(ctx, query, limit, offset)
        if err != nil {
                return nil, fmt.Errorf("failed to query devices: %w", err)
        }
        defer rows.Close()

        var devices []map[string]interface{}
        for rows.Next() {
                var id, name, description, adminState, operatingState, serviceName, profileName string
                var protocolsJSON, labelsJSON []byte
                var created, modified int64

                err := rows.Scan(&id, &name, &description, &adminState, &operatingState,
                        &protocolsJSON, &labelsJSON, &serviceName, &profileName, &created, &modified)
                if err != nil {
                        continue
                }

                var protocols map[string]interface{}
                var labels []string
                if len(protocolsJSON) > 0 {
                        json.Unmarshal(protocolsJSON, &protocols)
                }
                if len(labelsJSON) > 0 {
                        json.Unmarshal(labelsJSON, &labels)
                }

                device := map[string]interface{}{
                        "id":             id,
                        "name":           name,
                        "description":    description,
                        "adminState":     adminState,
                        "operatingState": operatingState,
                        "protocols":      protocols,
                        "labels":         labels,
                        "serviceName":    serviceName,
                        "profileName":    profileName,
                        "created":        created,
                        "modified":       modified,
                }
                devices = append(devices, device)
        }

        return devices, nil
}

// Core Data Service Methods
func (s *UnifiedIIOTService) GetAllEvents(ctx context.Context, offset, limit int, deviceName string) ([]models.Event, error) {
        if limit <= 0 {
                limit = 20
        }
        if offset < 0 {
                offset = 0
        }

        query := `
                SELECT id, device_name, profile_name, source_name, origin, tags, created, modified
                FROM events
                WHERE ($1 = '' OR device_name = $1)
                ORDER BY created DESC
                LIMIT $2 OFFSET $3`

        rows, err := s.db.QueryContext(ctx, query, deviceName, limit, offset)
        if err != nil {
                return nil, fmt.Errorf("failed to query events: %w", err)
        }
        defer rows.Close()

        var events []models.Event
        for rows.Next() {
                var event models.Event
                var tagsJSON []byte

                err := rows.Scan(&event.ID, &event.DeviceName, &event.ProfileName, &event.SourceName,
                        &event.Origin, &tagsJSON, &event.Created, &event.Modified)
                if err != nil {
                        continue
                }

                if len(tagsJSON) > 0 {
                        json.Unmarshal(tagsJSON, &event.Tags)
                }

                // Get readings for this event
                readings, err := s.getReadingsByEventID(ctx, event.ID)
                if err == nil {
                        event.Readings = readings
                }

                events = append(events, event)
        }

        return events, nil
}

func (s *UnifiedIIOTService) GetAllReadings(ctx context.Context, offset, limit int, deviceName, resourceName string) ([]models.Reading, error) {
        if limit <= 0 {
                limit = 20
        }
        if offset < 0 {
                offset = 0
        }

        query := `
                SELECT id, event_id, device_name, resource_name, profile_name, value_type,
                       value, binary_value, media_type, units, tags, origin, created, modified
                FROM readings
                WHERE ($1 = '' OR device_name = $1)
                  AND ($2 = '' OR resource_name = $2)
                ORDER BY created DESC
                LIMIT $3 OFFSET $4`

        rows, err := s.db.QueryContext(ctx, query, deviceName, resourceName, limit, offset)
        if err != nil {
                return nil, fmt.Errorf("failed to query readings: %w", err)
        }
        defer rows.Close()

        var readings []models.Reading
        for rows.Next() {
                var reading models.Reading
                var tagsJSON []byte

                err := rows.Scan(&reading.ID, &reading.EventID, &reading.DeviceName, &reading.ResourceName,
                        &reading.ProfileName, &reading.ValueType, &reading.Value, &reading.BinaryValue,
                        &reading.MediaType, &reading.Units, &tagsJSON, &reading.Origin,
                        &reading.Created, &reading.Modified)
                if err != nil {
                        continue
                }

                if len(tagsJSON) > 0 {
                        json.Unmarshal(tagsJSON, &reading.Tags)
                }

                readings = append(readings, reading)
        }

        return readings, nil
}

func (s *UnifiedIIOTService) getReadingsByEventID(ctx context.Context, eventID string) ([]models.Reading, error) {
        query := `
                SELECT id, event_id, device_name, resource_name, profile_name, value_type,
                       value, binary_value, media_type, units, tags, origin, created, modified
                FROM readings
                WHERE event_id = $1
                ORDER BY created ASC`

        rows, err := s.db.QueryContext(ctx, query, eventID)
        if err != nil {
                return nil, err
        }
        defer rows.Close()

        var readings []models.Reading
        for rows.Next() {
                var reading models.Reading
                var tagsJSON []byte

                err := rows.Scan(&reading.ID, &reading.EventID, &reading.DeviceName, &reading.ResourceName,
                        &reading.ProfileName, &reading.ValueType, &reading.Value, &reading.BinaryValue,
                        &reading.MediaType, &reading.Units, &tagsJSON, &reading.Origin,
                        &reading.Created, &reading.Modified)
                if err != nil {
                        continue
                }

                if len(tagsJSON) > 0 {
                        json.Unmarshal(tagsJSON, &reading.Tags)
                }

                readings = append(readings, reading)
        }

        return readings, nil
}

// Core Command Service Methods
func (s *UnifiedIIOTService) GetDeviceByName(ctx context.Context, name string) (map[string]interface{}, error) {
        query := `
                SELECT id, name, description, admin_state, operating_state, protocols, labels,
                       service_name, profile_name, created, modified
                FROM devices
                WHERE name = $1`

        var id, deviceName, description, adminState, operatingState, serviceName, profileName string
        var protocolsJSON, labelsJSON []byte
        var created, modified int64

        err := s.db.QueryRowContext(ctx, query, name).Scan(&id, &deviceName, &description, &adminState,
                &operatingState, &protocolsJSON, &labelsJSON, &serviceName, &profileName, &created, &modified)
        if err != nil {
                return nil, err
        }

        var protocols map[string]interface{}
        var labels []string
        if len(protocolsJSON) > 0 {
                json.Unmarshal(protocolsJSON, &protocols)
        }
        if len(labelsJSON) > 0 {
                json.Unmarshal(labelsJSON, &labels)
        }

        return map[string]interface{}{
                "id":             id,
                "name":           deviceName,
                "description":    description,
                "adminState":     adminState,
                "operatingState": operatingState,
                "protocols":      protocols,
                "labels":         labels,
                "serviceName":    serviceName,
                "profileName":    profileName,
                "created":        created,
                "modified":       modified,
        }, nil
}

func main() {
        // Load database configuration from environment
        databaseURL := os.Getenv("DATABASE_URL")
        if databaseURL == "" {
                databaseURL = "postgres://postgres:password@localhost/postgres?sslmode=disable"
        }

        // Initialize database connection
        db, err := sql.Open("postgres", databaseURL)
        if err != nil {
                panic("Failed to connect to database: " + err.Error())
        }
        defer db.Close()

        // Test database connection
        if err := db.Ping(); err != nil {
                fmt.Printf("Database connection failed: %v\n", err)
                panic("Database connection failed: " + err.Error())
        }
        fmt.Println("Database connection successful")

        // Create Echo instance
        e := echo.New()
        e.HideBanner = true

        // Add middleware
        e.Use(middleware.Logger())
        e.Use(middleware.Recover())
        e.Use(middleware.CORS())

        // Initialize unified service
        service := NewUnifiedIIOTService(db)

        // Initialize Core Command service components (EdgeX-Go style)
        commandService := application.NewCommandService()
        commandController := controller.NewCommandController(commandService)

        // Setup routes
        setupRoutes(e, service, commandController)

        // Start server
        port := os.Getenv("SERVICE_PORT")
        if port == "" {
                port = "5000"
        }

        fmt.Printf("Starting IIOT Backend server on port %s\n", port)

        // Graceful shutdown
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()

        go func() {
                if err := e.Start("0.0.0.0:" + port); err != nil && err != http.ErrServerClosed {
                        fmt.Printf("Failed to start server: %v\n", err)
                        cancel()
                }
        }()

        // Wait for interrupt signal
        ch := make(chan os.Signal, 1)
        signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

        select {
        case <-ch:
                fmt.Println("Shutting down server...")
                cancel()
        case <-ctx.Done():
        }

        // Shutdown server
        if err := e.Shutdown(ctx); err != nil {
                fmt.Printf("Server forced to shutdown: %v\n", err)
        }

        fmt.Println("Server exited")
}

func setupRoutes(e *echo.Echo, service *UnifiedIIOTService, commandController *controller.CommandController) {
        // Root route for external access
        e.GET("/", func(c echo.Context) error {
                return c.JSON(200, map[string]interface{}{
                        "service":    "iiot-backend",
                        "version":    "3.0.0",
                        "apiVersion": ApiVersion,
                        "status":     "running",
                        "endpoints": map[string]string{
                                "health":  "/health",
                                "ping":    ApiPingRoute,
                                "version": ApiVersionRoute,
                                "config":  ApiConfigRoute,
                        },
                })
        })

        // Common routes
        e.GET(ApiPingRoute, func(c echo.Context) error {
                return c.JSON(200, map[string]interface{}{
                        "apiVersion":  ApiVersion,
                        "timestamp":   time.Now().Unix(),
                        "serviceName": "iiot-backend",
                })
        })

        e.GET(ApiVersionRoute, func(c echo.Context) error {
                return c.JSON(200, map[string]interface{}{
                        "apiVersion":  ApiVersion,
                        "version":     "3.0.0",
                        "serviceName": "iiot-backend",
                        "sdkVersion":  "3.0.0",
                })
        })

        e.GET(ApiConfigRoute, func(c echo.Context) error {
                return c.JSON(200, map[string]interface{}{
                        "apiVersion":  ApiVersion,
                        "serviceName": "iiot-backend",
                        "config": map[string]interface{}{
                                "service": map[string]interface{}{
                                        "host":            "0.0.0.0",
                                        "port":            5000,
                                        "startupMsg":      "IIOT Backend service started",
                                        "maxResultCount":  1000,
                                        "requestTimeout":  "20s",
                                },
                        },
                })
        })

        // Health check
        e.GET("/health", func(c echo.Context) error {
                return c.JSON(200, map[string]interface{}{
                        "status":    "healthy",
                        "timestamp": time.Now().Unix(),
                        "service":   "iiot-backend",
                        "version":   "3.0.0",
                })
        })

        // Core Metadata routes
        e.GET(ApiAllDeviceHandlerRoute, func(c echo.Context) error {
                offset, _ := strconv.Atoi(c.QueryParam("offset"))
                limit, _ := strconv.Atoi(c.QueryParam("limit"))

                services, err := service.GetAllDeviceServices(c.Request().Context(), offset, limit)
                if err != nil {
                        return c.JSON(500, map[string]string{"error": "Failed to get device services"})
                }

                return c.JSON(200, map[string]interface{}{
                        "apiVersion": ApiVersion,
                        "statusCode": 200,
                        "deviceServices": services,
                })
        })

        e.GET(ApiAllDeviceRoute, func(c echo.Context) error {
                offset, _ := strconv.Atoi(c.QueryParam("offset"))
                limit, _ := strconv.Atoi(c.QueryParam("limit"))

                devices, err := service.GetAllDevices(c.Request().Context(), offset, limit)
                if err != nil {
                        return c.JSON(500, map[string]string{"error": "Failed to get devices"})
                }

                return c.JSON(200, map[string]interface{}{
                        "apiVersion": ApiVersion,
                        "statusCode": 200,
                        "devices": devices,
                })
        })

        // Core Data routes
        e.GET(ApiAllDataEventRoute, func(c echo.Context) error {
                offset, _ := strconv.Atoi(c.QueryParam("offset"))
                limit, _ := strconv.Atoi(c.QueryParam("limit"))
                deviceName := c.QueryParam("device")

                events, err := service.GetAllEvents(c.Request().Context(), offset, limit, deviceName)
                if err != nil {
                        return c.JSON(500, map[string]string{"error": "Failed to get events"})
                }

                return c.JSON(200, map[string]interface{}{
                        "apiVersion": ApiVersion,
                        "statusCode": 200,
                        "events": events,
                })
        })

        e.GET(ApiAllMeasurementRoute, func(c echo.Context) error {
                offset, _ := strconv.Atoi(c.QueryParam("offset"))
                limit, _ := strconv.Atoi(c.QueryParam("limit"))
                deviceName := c.QueryParam("device")
                resourceName := c.QueryParam("resourceName")

                readings, err := service.GetAllReadings(c.Request().Context(), offset, limit, deviceName, resourceName)
                if err != nil {
                        return c.JSON(500, map[string]string{"error": "Failed to get readings"})
                }

                return c.JSON(200, map[string]interface{}{
                        "apiVersion": ApiVersion,
                        "statusCode": 200,
                        "readings": readings,
                })
        })

        // Core Command routes - EdgeX-Go style implementation
        e.GET("/api/v3/device/all", commandController.AllDeviceCoreCommands)
        e.GET("/api/v3/device/name/:name", commandController.DeviceCoreCommandsByDeviceName)
        e.GET("/api/v3/device/name/:name/:command", commandController.IssueGetCommandByName)
        e.PUT("/api/v3/device/name/:name/:command", commandController.IssueSetCommandByName)
}