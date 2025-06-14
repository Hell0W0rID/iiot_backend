//
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
)

// DeviceCoreCommandResponse defines the Response Content for GET DeviceCoreCommand DTO.
type DeviceCoreCommandResponse struct {
	common.BaseResponse `json:",inline"`
	DeviceCoreCommand   dtos.DeviceCoreCommand `json:"deviceCoreCommand"`
}

func NewDeviceCoreCommandResponse(requestId string, message string, statusCode int, deviceCoreCommand dtos.DeviceCoreCommand) DeviceCoreCommandResponse {
	return DeviceCoreCommandResponse{
		BaseResponse:      common.NewBaseResponse(requestId, message, statusCode),
		DeviceCoreCommand: deviceCoreCommand,
	}
}

// MultiDeviceCoreCommandsResponse defines the Response Content for GET multiple DeviceCoreCommand DTOs.
type MultiDeviceCoreCommandsResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	DeviceCoreCommands                []dtos.DeviceCoreCommand `json:"deviceCoreCommands"`
}

func NewMultiDeviceCoreCommandsResponse(requestId string, message string, statusCode int, totalCount uint32, commands []dtos.DeviceCoreCommand) MultiDeviceCoreCommandsResponse {
	return MultiDeviceCoreCommandsResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		DeviceCoreCommands:         commands,
	}
}
