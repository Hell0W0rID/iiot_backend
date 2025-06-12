package interfaces

import (
	"context"
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// DeviceClient defines the interface for device operations
type DeviceClient interface {
	Add(ctx context.Context, reqs []dtos.DeviceRequest) ([]dtos.BaseWithIdResponse, errors.IIOTError)
	Update(ctx context.Context, reqs []dtos.UpdateDeviceRequest) ([]dtos.BaseResponse, errors.IIOTError)
	AllDevices(ctx context.Context, labels []string, offset int, limit int) (dtos.MultiDevicesResponse, errors.IIOTError)
	DeviceByName(ctx context.Context, name string) (dtos.DeviceResponse, errors.IIOTError)
	DeleteDeviceByName(ctx context.Context, name string) (dtos.BaseResponse, errors.IIOTError)
	DevicesByServiceName(ctx context.Context, name string, offset int, limit int) (dtos.MultiDevicesResponse, errors.IIOTError)
	DevicesByProfileName(ctx context.Context, name string, offset int, limit int) (dtos.MultiDevicesResponse, errors.IIOTError)
}

// DeviceHandlerClient defines the interface for device service operations
type DeviceHandlerClient interface {
	Add(ctx context.Context, reqs []dtos.DeviceHandlerRequest) ([]dtos.BaseWithIdResponse, errors.IIOTError)
	Update(ctx context.Context, reqs []dtos.UpdateDeviceHandlerRequest) ([]dtos.BaseResponse, errors.IIOTError)
	AllDeviceHandlers(ctx context.Context, labels []string, offset int, limit int) (dtos.MultiDeviceHandlersResponse, errors.IIOTError)
	DeviceHandlerByName(ctx context.Context, name string) (dtos.DeviceHandlerResponse, errors.IIOTError)
	DeleteDeviceHandlerByName(ctx context.Context, name string) (dtos.BaseResponse, errors.IIOTError)
}

// DeviceTemplateClient defines the interface for device profile operations
type DeviceTemplateClient interface {
	Add(ctx context.Context, reqs []dtos.DeviceTemplateRequest) ([]dtos.BaseWithIdResponse, errors.IIOTError)
	Update(ctx context.Context, reqs []dtos.UpdateDeviceTemplateRequest) ([]dtos.BaseResponse, errors.IIOTError)
	AllDeviceTemplates(ctx context.Context, labels []string, offset int, limit int) (dtos.MultiDeviceTemplatesResponse, errors.IIOTError)
	DeviceTemplateByName(ctx context.Context, name string) (dtos.DeviceTemplateResponse, errors.IIOTError)
	DeleteDeviceTemplateByName(ctx context.Context, name string) (dtos.BaseResponse, errors.IIOTError)
}

// DeviceWatcherClient defines the interface for provision watcher operations
type DeviceWatcherClient interface {
	Add(ctx context.Context, reqs []dtos.DeviceWatcherRequest) ([]dtos.BaseWithIdResponse, errors.IIOTError)
	Update(ctx context.Context, reqs []dtos.UpdateDeviceWatcherRequest) ([]dtos.BaseResponse, errors.IIOTError)
	AllDeviceWatchers(ctx context.Context, labels []string, offset int, limit int) (dtos.MultiDeviceWatchersResponse, errors.IIOTError)
	DeviceWatcherByName(ctx context.Context, name string) (dtos.DeviceWatcherResponse, errors.IIOTError)
	DeleteDeviceWatcherByName(ctx context.Context, name string) (dtos.BaseResponse, errors.IIOTError)
}