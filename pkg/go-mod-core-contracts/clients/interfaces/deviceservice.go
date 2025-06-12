package interfaces

import (
	"context"

	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/requests"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// DeviceHandlerClient defines the interface for interactions with the Device Service endpoint on the IIOT Foundry iiot-metadata service.
type DeviceHandlerClient interface {
	// Add adds new device services.
	Add(ctx context.Context, reqs []requests.AddDeviceHandlerRequest) ([]common.BaseWithIdResponse, errors.IIOT)
	// Update updates device services.
	Update(ctx context.Context, reqs []requests.UpdateDeviceHandlerRequest) ([]common.BaseResponse, errors.IIOT)
	// AllDeviceHandlers returns all device services. Device services can also be filtered by labels.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	AllDeviceHandlers(ctx context.Context, labels []string, offset int, limit int) (responses.MultiDeviceHandlersResponse, errors.IIOT)
	// DeviceHandlerByName returns a device service by name.
	DeviceHandlerByName(ctx context.Context, name string) (responses.DeviceHandlerResponse, errors.IIOT)
	// DeleteByName deletes a device service by name.
	DeleteByName(ctx context.Context, name string) (common.BaseResponse, errors.IIOT)
}
