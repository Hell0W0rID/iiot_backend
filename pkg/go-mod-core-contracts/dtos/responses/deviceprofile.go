//
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
)

// DeviceTemplateResponse defines the Response Content for GET DeviceTemplate DTOs.
type DeviceTemplateResponse struct {
	common.BaseResponse `json:",inline"`
	Profile             dtos.DeviceTemplate `json:"profile"`
}

func NewDeviceTemplateResponse(requestId string, message string, statusCode int, deviceProfile dtos.DeviceTemplate) DeviceTemplateResponse {
	return DeviceTemplateResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Profile:      deviceProfile,
	}
}

// MultiDeviceTemplatesResponse defines the Response Content for GET multiple DeviceTemplate DTOs.
type MultiDeviceTemplatesResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Profiles                          []dtos.DeviceTemplate `json:"profiles"`
}

func NewMultiDeviceTemplatesResponse(requestId string, message string, statusCode int, totalCount uint32, deviceProfiles []dtos.DeviceTemplate) MultiDeviceTemplatesResponse {
	return MultiDeviceTemplatesResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Profiles:                   deviceProfiles,
	}
}

// MultiDeviceTemplateBasicInfoResponse defines the Response Content for GET multiple DeviceTemplateBasicInfo DTOs.
type MultiDeviceTemplateBasicInfoResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Profiles                          []dtos.DeviceTemplateBasicInfo `json:"profiles"`
}

func NewMultiDeviceTemplateBasicInfosResponse(requestId string, message string, statusCode int, totalCount uint32, deviceProfileBasicInfos []dtos.DeviceTemplateBasicInfo) MultiDeviceTemplateBasicInfoResponse {
	return MultiDeviceTemplateBasicInfoResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Profiles:                   deviceProfileBasicInfos,
	}
}
