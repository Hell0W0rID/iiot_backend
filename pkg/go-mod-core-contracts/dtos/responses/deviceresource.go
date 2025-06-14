//
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
)

// DeviceResourceResponse defines the Response Content for GET DeviceResource DTOs.
type DeviceResourceResponse struct {
	common.BaseResponse `json:",inline"`
	Resource            dtos.DeviceResource `json:"resource"`
}

// NewDeviceResourceResponse creates deviceResource response DTO with required fields
func NewDeviceResourceResponse(requestId string, message string, statusCode int, resource dtos.DeviceResource) DeviceResourceResponse {
	return DeviceResourceResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Resource:     resource,
	}
}
