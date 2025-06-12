//
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
)

// MeasurementResponse defines the Response Content for GET reading DTO.
type MeasurementResponse struct {
	common.BaseResponse `json:",inline"`
	Measurement             dtos.BaseMeasurement `json:"reading"`
}

// MultiMeasurementsResponse defines the Response Content for GET multiple reading DTO.
type MultiMeasurementsResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Measurements                          []dtos.BaseMeasurement `json:"readings"`
}

func NewMeasurementResponse(requestId string, message string, statusCode int, reading dtos.BaseMeasurement) MeasurementResponse {
	return MeasurementResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Measurement:      reading,
	}
}

func NewMultiMeasurementsResponse(requestId string, message string, statusCode int, totalCount uint32, readings []dtos.BaseMeasurement) MultiMeasurementsResponse {
	return MultiMeasurementsResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Measurements:                   readings,
	}
}
