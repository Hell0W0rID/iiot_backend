//
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
)

// DeviceWatcherResponse defines the Response Content for GET DeviceWatcher DTOs.
type DeviceWatcherResponse struct {
	common.BaseResponse `json:",inline"`
	DeviceWatcher    dtos.DeviceWatcher `json:"provisionWatcher"`
}

func NewDeviceWatcherResponse(requestId string, message string, statusCode int, pw dtos.DeviceWatcher) DeviceWatcherResponse {
	return DeviceWatcherResponse{
		BaseResponse:     common.NewBaseResponse(requestId, message, statusCode),
		DeviceWatcher: pw,
	}
}

// MultiDeviceWatchersResponse defines the Response Content for GET multiple DeviceWatcher DTOs.
type MultiDeviceWatchersResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	DeviceWatchers                 []dtos.DeviceWatcher `json:"provisionWatchers"`
}

func NewMultiDeviceWatchersResponse(requestId string, message string, statusCode int, totalCount uint32, pws []dtos.DeviceWatcher) MultiDeviceWatchersResponse {
	return MultiDeviceWatchersResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		DeviceWatchers:          pws,
	}
}
