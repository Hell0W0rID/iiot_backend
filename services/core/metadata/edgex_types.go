package metadata

import "time"

// EdgeX-compatible DTOs and types for core-metadata service
type DeviceService struct {
        Id          string            `json:"id,omitempty"`
        Name        string            `json:"name" validate:"required"`
        Description string            `json:"description,omitempty"`
        BaseAddress string            `json:"baseAddress" validate:"required"`
        AdminState  string            `json:"adminState"`
        Labels      []string          `json:"labels,omitempty"`
        Created     int64             `json:"created,omitempty"`
        Modified    int64             `json:"modified,omitempty"`
}

type Device struct {
        Id             string                 `json:"id,omitempty"`
        Name           string                 `json:"name" validate:"required"`
        Description    string                 `json:"description,omitempty"`
        AdminState     string                 `json:"adminState"`
        OperatingState string                 `json:"operatingState"`
        ServiceName    string                 `json:"serviceName" validate:"required"`
        ProfileName    string                 `json:"profileName" validate:"required"`
        Labels         []string               `json:"labels,omitempty"`
        Location       map[string]interface{} `json:"location,omitempty"`
        Protocols      map[string]interface{} `json:"protocols,omitempty"`
        AutoEvents     []interface{}          `json:"autoEvents,omitempty"`
        Notify         bool                   `json:"notify"`
        Created        int64                  `json:"created,omitempty"`
        Modified       int64                  `json:"modified,omitempty"`
}

type DeviceProfile struct {
        Id               string                   `json:"id,omitempty"`
        Name             string                   `json:"name" validate:"required"`
        Description      string                   `json:"description,omitempty"`
        Manufacturer     string                   `json:"manufacturer,omitempty"`
        Model            string                   `json:"model,omitempty"`
        Labels           []string                 `json:"labels,omitempty"`
        DeviceResources  []map[string]interface{} `json:"deviceResources,omitempty"`
        DeviceCommands   []map[string]interface{} `json:"deviceCommands,omitempty"`
        CoreCommands     []map[string]interface{} `json:"coreCommands,omitempty"`
        Created          int64                    `json:"created,omitempty"`
        Modified         int64                    `json:"modified,omitempty"`
}

type ProvisionWatcher struct {
        Id                   string                 `json:"id,omitempty"`
        Name                 string                 `json:"name" validate:"required"`
        Description          string                 `json:"description,omitempty"`
        Labels               []string               `json:"labels,omitempty"`
        Identifiers          map[string]interface{} `json:"identifiers,omitempty"`
        BlockingIdentifiers  map[string]interface{} `json:"blockingIdentifiers,omitempty"`
        ProfileName          string                 `json:"profileName" validate:"required"`
        ServiceName          string                 `json:"serviceName" validate:"required"`
        AdminState           string                 `json:"adminState"`
        AutoEvents           []interface{}          `json:"autoEvents,omitempty"`
        DiscoveryProperties  map[string]interface{} `json:"discoveryProperties,omitempty"`
        Created              int64                  `json:"created,omitempty"`
        Modified             int64                  `json:"modified,omitempty"`
}

// Request DTOs
type AddDeviceServiceRequest struct {
        RequestId     string        `json:"requestId,omitempty"`
        DeviceService DeviceService `json:"deviceService"`
}

func (r *AddDeviceServiceRequest) Validate() error {
        if r.DeviceService.Name == "" {
                return &ValidationError{Message: "deviceService name is required"}
        }
        if r.DeviceService.BaseAddress == "" {
                return &ValidationError{Message: "deviceService baseAddress is required"}
        }
        return nil
}

type UpdateDeviceServiceRequest struct {
        RequestId     string              `json:"requestId,omitempty"`
        DeviceService UpdateDeviceService `json:"deviceService"`
}

func (r *UpdateDeviceServiceRequest) Validate() error {
        if r.DeviceService.Name == nil || *r.DeviceService.Name == "" {
                return &ValidationError{Message: "deviceService name is required for update"}
        }
        return nil
}

type UpdateDeviceService struct {
        Name        *string   `json:"name,omitempty"`
        Description *string   `json:"description,omitempty"`
        BaseAddress *string   `json:"baseAddress,omitempty"`
        AdminState  *string   `json:"adminState,omitempty"`
        Labels      *[]string `json:"labels,omitempty"`
}

type AddDeviceRequest struct {
        RequestId string `json:"requestId,omitempty"`
        Device    Device `json:"device"`
}

func (r *AddDeviceRequest) Validate() error {
        if r.Device.Name == "" {
                return &ValidationError{Message: "device name is required"}
        }
        if r.Device.ServiceName == "" {
                return &ValidationError{Message: "device serviceName is required"}
        }
        if r.Device.ProfileName == "" {
                return &ValidationError{Message: "device profileName is required"}
        }
        return nil
}

type UpdateDeviceRequest struct {
        RequestId string       `json:"requestId,omitempty"`
        Device    UpdateDevice `json:"device"`
}

func (r *UpdateDeviceRequest) Validate() error {
        if r.Device.Name == nil || *r.Device.Name == "" {
                return &ValidationError{Message: "device name is required for update"}
        }
        return nil
}

type UpdateDevice struct {
        Name           *string                 `json:"name,omitempty"`
        Description    *string                 `json:"description,omitempty"`
        AdminState     *string                 `json:"adminState,omitempty"`
        OperatingState *string                 `json:"operatingState,omitempty"`
        ServiceName    *string                 `json:"serviceName,omitempty"`
        ProfileName    *string                 `json:"profileName,omitempty"`
        Labels         *[]string               `json:"labels,omitempty"`
        Location       *map[string]interface{} `json:"location,omitempty"`
        Protocols      *map[string]interface{} `json:"protocols,omitempty"`
        AutoEvents     *[]interface{}          `json:"autoEvents,omitempty"`
        Notify         *bool                   `json:"notify,omitempty"`
}

// Response DTOs
type BaseResponse struct {
        ApiVersion string `json:"apiVersion"`
        RequestId  string `json:"requestId,omitempty"`
        Message    string `json:"message,omitempty"`
        StatusCode int    `json:"statusCode"`
}

type BaseWithIdResponse struct {
        BaseResponse
        Id string `json:"id"`
}

type DeviceServiceResponse struct {
        BaseResponse
        DeviceService DeviceService `json:"deviceService"`
}

type MultiDeviceServicesResponse struct {
        BaseResponse
        TotalCount     uint32          `json:"totalCount"`
        DeviceServices []DeviceService `json:"deviceServices"`
}

type DeviceResponse struct {
        BaseResponse
        Device Device `json:"device"`
}

type MultiDevicesResponse struct {
        BaseResponse
        TotalCount uint32   `json:"totalCount"`
        Devices    []Device `json:"devices"`
}

// Error types
type ValidationError struct {
        Message string
}

func (e *ValidationError) Error() string {
        return e.Message
}

type EdgeXError struct {
        Code    int
        Message string
}

func (e EdgeXError) Error() string {
        return e.Message
}

// Helper functions
func NewBaseResponse(requestId, message string, statusCode int) BaseResponse {
        return BaseResponse{
                ApiVersion: "v3",
                RequestId:  requestId,
                Message:    message,
                StatusCode: statusCode,
        }
}

func NewBaseWithIdResponse(requestId, message string, statusCode int, id string) BaseWithIdResponse {
        return BaseWithIdResponse{
                BaseResponse: NewBaseResponse(requestId, message, statusCode),
                Id:           id,
        }
}

func NewDeviceServiceResponse(requestId, message string, statusCode int, deviceService DeviceService) DeviceServiceResponse {
        return DeviceServiceResponse{
                BaseResponse:  NewBaseResponse(requestId, message, statusCode),
                DeviceService: deviceService,
        }
}

func NewMultiDeviceServicesResponse(requestId, message string, statusCode int, totalCount uint32, deviceServices []DeviceService) MultiDeviceServicesResponse {
        return MultiDeviceServicesResponse{
                BaseResponse:   NewBaseResponse(requestId, message, statusCode),
                TotalCount:     totalCount,
                DeviceServices: deviceServices,
        }
}

func NewDeviceResponse(requestId, message string, statusCode int, device Device) DeviceResponse {
        return DeviceResponse{
                BaseResponse: NewBaseResponse(requestId, message, statusCode),
                Device:       device,
        }
}

func NewMultiDevicesResponse(requestId, message string, statusCode int, totalCount uint32, devices []Device) MultiDevicesResponse {
        return MultiDevicesResponse{
                BaseResponse: NewBaseResponse(requestId, message, statusCode),
                TotalCount:   totalCount,
                Devices:      devices,
        }
}

func NewCommonEdgeX(code int, message string, err error) EdgeXError {
        msg := message
        if err != nil {
                msg = msg + ": " + err.Error()
        }
        return EdgeXError{Code: code, Message: msg}
}

func NewEntityDoesNotExistError(entityType, entityName string) EdgeXError {
        return EdgeXError{
                Code:    404,
                Message: entityType + " '" + entityName + "' not found",
        }
}

// Convert time.Time to EdgeX timestamp (milliseconds since epoch)
func TimeToEdgeXTimestamp(t time.Time) int64 {
        return t.UnixNano() / int64(time.Millisecond)
}