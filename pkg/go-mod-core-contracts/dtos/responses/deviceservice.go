//
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
)

// DeviceHandlerResponse defines the Response Content for GET DeviceHandler DTOs.
type DeviceHandlerResponse struct {
	common.BaseResponse `json:",inline"`
	Service             dtos.DeviceHandler `json:"service"`
}

func NewDeviceHandlerResponse(requestId string, message string, statusCode int, deviceService dtos.DeviceHandler) DeviceHandlerResponse {
	return DeviceHandlerResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Service:      deviceService,
	}
}

// MultiDeviceHandlersResponse defines the Response Content for GET multiple DeviceHandler DTOs.
type MultiDeviceHandlersResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Services                          []dtos.DeviceHandler `json:"services"`
}

func NewMultiDeviceHandlersResponse(requestId string, message string, statusCode int, totalCount uint32, deviceServices []dtos.DeviceHandler) MultiDeviceHandlersResponse {
	return MultiDeviceHandlersResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Services:                   deviceServices,
	}
}
