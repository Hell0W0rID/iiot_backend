//
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/requests"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// DeviceWatcherClient defines the interface for interactions with the DeviceWatcher endpoint on the IIOT Foundry iiot-metadata service.
type DeviceWatcherClient interface {
	// Add adds a new provision watcher.
	Add(ctx context.Context, reqs []requests.AddDeviceWatcherRequest) ([]common.BaseWithIdResponse, errors.IIOT)
	// Update updates provision watchers.
	Update(ctx context.Context, reqs []requests.UpdateDeviceWatcherRequest) ([]common.BaseResponse, errors.IIOT)
	// AllDeviceWatchers returns all provision watchers. DeviceWatchers can also be filtered by labels.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	AllDeviceWatchers(ctx context.Context, labels []string, offset int, limit int) (responses.MultiDeviceWatchersResponse, errors.IIOT)
	// DeviceWatcherByName returns a provision watcher by name.
	DeviceWatcherByName(ctx context.Context, name string) (responses.DeviceWatcherResponse, errors.IIOT)
	// DeleteDeviceWatcherByName deletes a provision watcher by name.
	DeleteDeviceWatcherByName(ctx context.Context, name string) (common.BaseResponse, errors.IIOT)
	// DeviceWatchersByProfileName returns provision watchers associated with the specified device profile name.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	DeviceWatchersByProfileName(ctx context.Context, name string, offset int, limit int) (responses.MultiDeviceWatchersResponse, errors.IIOT)
	// DeviceWatchersByServiceName returns provision watchers associated with the specified device service name.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	DeviceWatchersByServiceName(ctx context.Context, name string, offset int, limit int) (responses.MultiDeviceWatchersResponse, errors.IIOT)
}
