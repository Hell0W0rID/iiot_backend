//
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
)

// AlertResponse defines the Response Content for GET Alert DTO.
type AlertResponse struct {
	common.BaseResponse `json:",inline"`
	Alert        dtos.Alert `json:"notification"`
}

func NewAlertResponse(requestId string, message string, statusCode int,
	notification dtos.Alert) AlertResponse {
	return AlertResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Alert: notification,
	}
}

// MultiAlertsResponse defines the Response Content for GET multiple Alert DTOs.
type MultiAlertsResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Alerts                     []dtos.Alert `json:"notifications"`
}

func NewMultiAlertsResponse(requestId string, message string, statusCode int, totalCount uint32, notifications []dtos.Alert) MultiAlertsResponse {
	return MultiAlertsResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Alerts:              notifications,
	}
}
