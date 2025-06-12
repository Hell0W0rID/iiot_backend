// Package metadata provides core metadata service functionality following EdgeX Foundry architecture
package metadata

import (
        "context"
        "database/sql"
        "encoding/json"
        "fmt"
        "net/http"
        "strconv"
        "time"

        "github.com/google/uuid"
        "github.com/labstack/echo/v4"
        "github.com/lib/pq"
)

// Working service implementation
type WorkingMetadataService struct {
        db *sql.DB
}

func NewWorkingMetadataService(db *sql.DB) *WorkingMetadataService {
        return &WorkingMetadataService{db: db}
}

// Device Service operations
func (s *WorkingMetadataService) AddDeviceService(ctx context.Context, req DeviceService) (string, EdgeXError) {
        if req.Name == "" {
                return "", EdgeXError{Code: http.StatusBadRequest, Message: "device service name is required"}
        }
        if req.BaseAddress == "" {
                return "", EdgeXError{Code: http.StatusBadRequest, Message: "device service base address is required"}
        }
        
        if req.AdminState == "" {
                req.AdminState = "UNLOCKED"
        }
        
        id := uuid.New().String()
        now := time.Now().Unix()
        
        labelsJSON, _ := json.Marshal(req.Labels)
        if labelsJSON == nil {
                labelsJSON = []byte("[]")
        }

        query := `
                INSERT INTO device_services (id, name, description, base_address, admin_state, labels, created, modified)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
        
        _, err := s.db.ExecContext(ctx, query, id, req.Name, req.Description, req.BaseAddress, 
                req.AdminState, labelsJSON, now, now)
        if err != nil {
                if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
                        return "", EdgeXError{Code: http.StatusConflict, Message: fmt.Sprintf("device service %s already exists", req.Name)}
                }
                return "", EdgeXError{Code: http.StatusInternalServerError, Message: "failed to add device service"}
        }
        
        return id, EdgeXError{}
}

func (s *WorkingMetadataService) GetDeviceServiceByName(ctx context.Context, name string) (DeviceService, EdgeXError) {
        if name == "" {
                return DeviceService{}, EdgeXError{Code: http.StatusBadRequest, Message: "device service name is required"}
        }
        
        var ds DeviceService
        var labelsJSON []byte
        
        query := `
                SELECT id, name, description, base_address, admin_state, labels, created, modified
                FROM device_services 
                WHERE name = $1`
        
        err := s.db.QueryRowContext(ctx, query, name).Scan(
                &ds.Id, &ds.Name, &ds.Description, &ds.BaseAddress, &ds.AdminState, 
                &labelsJSON, &ds.Created, &ds.Modified)
        if err != nil {
                if err == sql.ErrNoRows {
                        return DeviceService{}, EdgeXError{Code: http.StatusNotFound, Message: fmt.Sprintf("device service %s not found", name)}
                }
                return DeviceService{}, EdgeXError{Code: http.StatusInternalServerError, Message: "failed to get device service"}
        }
        
        if err := json.Unmarshal(labelsJSON, &ds.Labels); err != nil {
                ds.Labels = []string{}
        }
        
        return ds, EdgeXError{}
}

func (s *WorkingMetadataService) GetAllDeviceServices(ctx context.Context, offset, limit int) ([]DeviceService, uint32, EdgeXError) {
        if limit <= 0 {
                limit = 20
        }
        if limit > 1000 {
                limit = 1000
        }
        if offset < 0 {
                offset = 0
        }
        
        var services []DeviceService
        var totalCount uint32
        
        // Get total count
        countQuery := `SELECT COUNT(*) FROM device_services`
        err := s.db.QueryRowContext(ctx, countQuery).Scan(&totalCount)
        if err != nil {
                return nil, 0, EdgeXError{Code: http.StatusInternalServerError, Message: "failed to get total count"}
        }
        
        // Get paginated results
        query := `
                SELECT id, name, description, base_address, admin_state, labels, created, modified
                FROM device_services
                ORDER BY name LIMIT $1 OFFSET $2`
        
        rows, err := s.db.QueryContext(ctx, query, limit, offset)
        if err != nil {
                return nil, 0, EdgeXError{Code: http.StatusInternalServerError, Message: "failed to query device services"}
        }
        defer rows.Close()
        
        for rows.Next() {
                var ds DeviceService
                var labelsJSON []byte
                
                err := rows.Scan(&ds.Id, &ds.Name, &ds.Description, &ds.BaseAddress, 
                        &ds.AdminState, &labelsJSON, &ds.Created, &ds.Modified)
                if err != nil {
                        return nil, 0, EdgeXError{Code: http.StatusInternalServerError, Message: "failed to scan device service"}
                }
                
                if err := json.Unmarshal(labelsJSON, &ds.Labels); err != nil {
                        ds.Labels = []string{}
                }
                
                services = append(services, ds)
        }
        
        return services, totalCount, EdgeXError{}
}

func (s *WorkingMetadataService) DeleteDeviceServiceByName(ctx context.Context, name string) EdgeXError {
        if name == "" {
                return EdgeXError{Code: http.StatusBadRequest, Message: "device service name is required"}
        }
        
        query := `DELETE FROM device_services WHERE name = $1`
        result, err := s.db.ExecContext(ctx, query, name)
        if err != nil {
                return EdgeXError{Code: http.StatusInternalServerError, Message: "failed to delete device service"}
        }
        
        rowsAffected, err := result.RowsAffected()
        if err != nil {
                return EdgeXError{Code: http.StatusInternalServerError, Message: "failed to get rows affected"}
        }
        
        if rowsAffected == 0 {
                return EdgeXError{Code: http.StatusNotFound, Message: fmt.Sprintf("device service %s not found", name)}
        }
        
        return EdgeXError{}
}

// Device operations
func (s *WorkingMetadataService) AddDevice(ctx context.Context, req Device) (string, EdgeXError) {
        if req.Name == "" {
                return "", EdgeXError{Code: http.StatusBadRequest, Message: "device name is required"}
        }
        if req.ServiceName == "" {
                return "", EdgeXError{Code: http.StatusBadRequest, Message: "device service name is required"}
        }
        if req.ProfileName == "" {
                return "", EdgeXError{Code: http.StatusBadRequest, Message: "device profile name is required"}
        }
        
        if req.AdminState == "" {
                req.AdminState = "UNLOCKED"
        }
        if req.OperatingState == "" {
                req.OperatingState = "UP"
        }
        
        // Verify service exists
        _, err := s.GetDeviceServiceByName(ctx, req.ServiceName)
        if err.Code != 0 {
                return "", EdgeXError{Code: http.StatusBadRequest, Message: fmt.Sprintf("device service %s not found", req.ServiceName)}
        }
        
        id := uuid.New().String()
        now := time.Now().Unix()
        
        labelsJSON, _ := json.Marshal(req.Labels)
        protocolsJSON, _ := json.Marshal(req.Protocols)
        autoEventsJSON, _ := json.Marshal(req.AutoEvents)
        locationJSON, _ := json.Marshal(req.Location)
        
        if labelsJSON == nil {
                labelsJSON = []byte("[]")
        }
        if protocolsJSON == nil {
                protocolsJSON = []byte("{}")
        }
        if autoEventsJSON == nil {
                autoEventsJSON = []byte("[]")
        }
        if locationJSON == nil {
                locationJSON = []byte("{}")
        }

        query := `
                INSERT INTO devices (id, name, description, admin_state, operating_state, protocols, 
                        labels, location, service_name, profile_name, auto_events, tags, properties, created, modified)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`
        
        _, err2 := s.db.ExecContext(ctx, query, id, req.Name, req.Description, 
                req.AdminState, req.OperatingState, protocolsJSON, labelsJSON, 
                locationJSON, req.ServiceName, req.ProfileName, autoEventsJSON, 
                []byte("{}"), []byte("{}"), now, now)
        if err2 != nil {
                if pqErr, ok := err2.(*pq.Error); ok && pqErr.Code == "23505" {
                        return "", EdgeXError{Code: http.StatusConflict, Message: fmt.Sprintf("device %s already exists", req.Name)}
                }
                return "", EdgeXError{Code: http.StatusInternalServerError, Message: "failed to add device"}
        }
        
        return id, EdgeXError{}
}

func (s *WorkingMetadataService) GetDeviceByName(ctx context.Context, name string) (Device, EdgeXError) {
        if name == "" {
                return Device{}, EdgeXError{Code: http.StatusBadRequest, Message: "device name is required"}
        }
        
        var device Device
        var labelsJSON, protocolsJSON, autoEventsJSON, locationJSON []byte
        
        query := `
                SELECT id, name, description, admin_state, operating_state, protocols, labels, 
                        location, service_name, profile_name, auto_events, created, modified
                FROM devices 
                WHERE name = $1`
        
        err := s.db.QueryRowContext(ctx, query, name).Scan(
                &device.Id, &device.Name, &device.Description, &device.AdminState, 
                &device.OperatingState, &protocolsJSON, &labelsJSON, &locationJSON,
                &device.ServiceName, &device.ProfileName, &autoEventsJSON, 
                &device.Created, &device.Modified)
        if err != nil {
                if err == sql.ErrNoRows {
                        return Device{}, EdgeXError{Code: http.StatusNotFound, Message: fmt.Sprintf("device %s not found", name)}
                }
                return Device{}, EdgeXError{Code: http.StatusInternalServerError, Message: "failed to get device"}
        }
        
        // Unmarshal JSON fields
        json.Unmarshal(labelsJSON, &device.Labels)
        json.Unmarshal(protocolsJSON, &device.Protocols)
        json.Unmarshal(autoEventsJSON, &device.AutoEvents)
        json.Unmarshal(locationJSON, &device.Location)
        
        return device, EdgeXError{}
}

func (s *WorkingMetadataService) GetAllDevices(ctx context.Context, offset, limit int) ([]Device, uint32, EdgeXError) {
        if limit <= 0 {
                limit = 20
        }
        if limit > 1000 {
                limit = 1000
        }
        if offset < 0 {
                offset = 0
        }
        
        var devices []Device
        var totalCount uint32
        
        // Get total count
        countQuery := `SELECT COUNT(*) FROM devices`
        err := s.db.QueryRowContext(ctx, countQuery).Scan(&totalCount)
        if err != nil {
                return nil, 0, EdgeXError{Code: http.StatusInternalServerError, Message: "failed to get total count"}
        }
        
        // Get paginated results
        query := `
                SELECT id, name, description, admin_state, operating_state, protocols, labels, 
                        location, service_name, profile_name, auto_events, created, modified
                FROM devices
                ORDER BY name LIMIT $1 OFFSET $2`
        
        rows, err := s.db.QueryContext(ctx, query, limit, offset)
        if err != nil {
                return nil, 0, EdgeXError{Code: http.StatusInternalServerError, Message: "failed to query devices"}
        }
        defer rows.Close()
        
        for rows.Next() {
                var device Device
                var labelsJSON, protocolsJSON, autoEventsJSON, locationJSON []byte
                
                err := rows.Scan(&device.Id, &device.Name, &device.Description, &device.AdminState,
                        &device.OperatingState, &protocolsJSON, &labelsJSON, &locationJSON,
                        &device.ServiceName, &device.ProfileName, &autoEventsJSON,
                        &device.Created, &device.Modified)
                if err != nil {
                        return nil, 0, EdgeXError{Code: http.StatusInternalServerError, Message: "failed to scan device"}
                }
                
                // Unmarshal JSON fields
                json.Unmarshal(labelsJSON, &device.Labels)
                json.Unmarshal(protocolsJSON, &device.Protocols)
                json.Unmarshal(autoEventsJSON, &device.AutoEvents)
                json.Unmarshal(locationJSON, &device.Location)
                
                devices = append(devices, device)
        }
        
        return devices, totalCount, EdgeXError{}
}

func (s *WorkingMetadataService) DeleteDeviceByName(ctx context.Context, name string) EdgeXError {
        if name == "" {
                return EdgeXError{Code: http.StatusBadRequest, Message: "device name is required"}
        }
        
        query := `DELETE FROM devices WHERE name = $1`
        result, err := s.db.ExecContext(ctx, query, name)
        if err != nil {
                return EdgeXError{Code: http.StatusInternalServerError, Message: "failed to delete device"}
        }
        
        rowsAffected, err := result.RowsAffected()
        if err != nil {
                return EdgeXError{Code: http.StatusInternalServerError, Message: "failed to get rows affected"}
        }
        
        if rowsAffected == 0 {
                return EdgeXError{Code: http.StatusNotFound, Message: fmt.Sprintf("device %s not found", name)}
        }
        
        return EdgeXError{}
}

// Working handler implementation
type WorkingHandler struct {
        service *WorkingMetadataService
}

func NewWorkingHandler(service *WorkingMetadataService) *WorkingHandler {
        return &WorkingHandler{service: service}
}

// Common endpoints
func (h *WorkingHandler) Ping(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}

func (h *WorkingHandler) Version(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]interface{}{
                "version":     "3.0.0",
                "serviceName": "iiot-metadata",
        })
}

// Device Service endpoints
func (h *WorkingHandler) AddDeviceService(c echo.Context) error {
        var req AddDeviceServiceRequest
        if err := c.Bind(&req); err != nil {
                return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
        }
        
        if err := req.Validate(); err != nil {
                return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
        }
        
        id, edgeErr := h.service.AddDeviceService(c.Request().Context(), req.DeviceService)
        if edgeErr.Code != 0 {
                return c.JSON(edgeErr.Code, map[string]string{"error": edgeErr.Message})
        }
        
        response := map[string]interface{}{
                "apiVersion": "v3",
                "statusCode": http.StatusCreated,
                "id":         id,
        }
        return c.JSON(http.StatusCreated, response)
}

func (h *WorkingHandler) GetDeviceServiceByName(c echo.Context) error {
        name := c.Param("name")
        
        service, edgeErr := h.service.GetDeviceServiceByName(c.Request().Context(), name)
        if edgeErr.Code != 0 {
                return c.JSON(edgeErr.Code, map[string]string{"error": edgeErr.Message})
        }
        
        response := map[string]interface{}{
                "apiVersion":    "v3",
                "statusCode":    http.StatusOK,
                "deviceService": service,
        }
        return c.JSON(http.StatusOK, response)
}

func (h *WorkingHandler) GetAllDeviceServices(c echo.Context) error {
        offset, _ := strconv.Atoi(c.QueryParam("offset"))
        limit, _ := strconv.Atoi(c.QueryParam("limit"))
        
        services, totalCount, edgeErr := h.service.GetAllDeviceServices(c.Request().Context(), offset, limit)
        if edgeErr.Code != 0 {
                return c.JSON(edgeErr.Code, map[string]string{"error": edgeErr.Message})
        }
        
        response := map[string]interface{}{
                "apiVersion":     "v3",
                "statusCode":     http.StatusOK,
                "totalCount":     totalCount,
                "deviceServices": services,
        }
        return c.JSON(http.StatusOK, response)
}

func (h *WorkingHandler) DeleteDeviceServiceByName(c echo.Context) error {
        name := c.Param("name")
        
        edgeErr := h.service.DeleteDeviceServiceByName(c.Request().Context(), name)
        if edgeErr.Code != 0 {
                return c.JSON(edgeErr.Code, map[string]string{"error": edgeErr.Message})
        }
        
        response := map[string]interface{}{
                "apiVersion": "v3",
                "statusCode": http.StatusOK,
                "message":    "Device service deleted successfully",
        }
        return c.JSON(http.StatusOK, response)
}

// Device endpoints
func (h *WorkingHandler) AddDevice(c echo.Context) error {
        var req AddDeviceRequest
        if err := c.Bind(&req); err != nil {
                return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
        }
        
        if err := req.Validate(); err != nil {
                return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
        }
        
        id, edgeErr := h.service.AddDevice(c.Request().Context(), req.Device)
        if edgeErr.Code != 0 {
                return c.JSON(edgeErr.Code, map[string]string{"error": edgeErr.Message})
        }
        
        response := map[string]interface{}{
                "apiVersion": "v3",
                "statusCode": http.StatusCreated,
                "id":         id,
        }
        return c.JSON(http.StatusCreated, response)
}

func (h *WorkingHandler) GetDeviceByName(c echo.Context) error {
        name := c.Param("name")
        
        device, edgeErr := h.service.GetDeviceByName(c.Request().Context(), name)
        if edgeErr.Code != 0 {
                return c.JSON(edgeErr.Code, map[string]string{"error": edgeErr.Message})
        }
        
        response := map[string]interface{}{
                "apiVersion": "v3",
                "statusCode": http.StatusOK,
                "device":     device,
        }
        return c.JSON(http.StatusOK, response)
}

func (h *WorkingHandler) GetAllDevices(c echo.Context) error {
        offset, _ := strconv.Atoi(c.QueryParam("offset"))
        limit, _ := strconv.Atoi(c.QueryParam("limit"))
        
        devices, totalCount, edgeErr := h.service.GetAllDevices(c.Request().Context(), offset, limit)
        if edgeErr.Code != 0 {
                return c.JSON(edgeErr.Code, map[string]string{"error": edgeErr.Message})
        }
        
        response := map[string]interface{}{
                "apiVersion": "v3",
                "statusCode": http.StatusOK,
                "totalCount": totalCount,
                "devices":    devices,
        }
        return c.JSON(http.StatusOK, response)
}

func (h *WorkingHandler) DeleteDeviceByName(c echo.Context) error {
        name := c.Param("name")
        
        edgeErr := h.service.DeleteDeviceByName(c.Request().Context(), name)
        if edgeErr.Code != 0 {
                return c.JSON(edgeErr.Code, map[string]string{"error": edgeErr.Message})
        }
        
        response := map[string]interface{}{
                "apiVersion": "v3",
                "statusCode": http.StatusOK,
                "message":    "Device deleted successfully",
        }
        return c.JSON(http.StatusOK, response)
}

// Route registration
func RegisterWorkingEdgeXRoutes(g *echo.Group, service *WorkingMetadataService) {
        handler := NewWorkingHandler(service)

        // Common endpoints
        g.GET("/ping", handler.Ping)
        g.GET("/version", handler.Version)

        // Device Service endpoints
        g.POST("/deviceservice", handler.AddDeviceService)
        g.GET("/deviceservice/all", handler.GetAllDeviceServices)
        g.GET("/deviceservice/name/:name", handler.GetDeviceServiceByName)
        g.DELETE("/deviceservice/name/:name", handler.DeleteDeviceServiceByName)

        // Device endpoints
        g.POST("/device", handler.AddDevice)
        g.GET("/device/all", handler.GetAllDevices)
        g.GET("/device/name/:name", handler.GetDeviceByName)
        g.DELETE("/device/name/:name", handler.DeleteDeviceByName)
}