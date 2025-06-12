//
//
// SPDX-License-Identifier: Apache-2.0

package responses

import "iiot-backend/pkg/go-mod-core-contracts/dtos/common"

type UnitsOfMeasureResponse struct {
	common.BaseResponse `json:",inline"`
	Uom                 any `json:"uom"`
}

func NewUnitsOfMeasureResponse(requestId string, message string, statusCode int, uom any) UnitsOfMeasureResponse {
	return UnitsOfMeasureResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Uom:          uom,
	}
}
