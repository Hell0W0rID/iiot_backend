//
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
)

// DeviceResponse defines the Response Content for GET Device DTOs.
type DeviceResponse struct {
	common.BaseResponse `json:",inline"`
	Device              dtos.Device `json:"device"`
}

func NewDeviceResponse(requestId string, message string, statusCode int, device dtos.Device) DeviceResponse {
	return DeviceResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Device:       device,
	}
}

// MultiDevicesResponse defines the Response Content for GET multiple Device DTOs.
type MultiDevicesResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Devices                           []dtos.Device `json:"devices"`
}

func NewMultiDevicesResponse(requestId string, message string, statusCode int, totalCount uint32, devices []dtos.Device) MultiDevicesResponse {
	return MultiDevicesResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Devices:                    devices,
	}
}
