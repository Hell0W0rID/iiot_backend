//
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
)

// DeliveryResponse defines the Response Content for GET Delivery DTO.
type DeliveryResponse struct {
	common.BaseResponse `json:",inline"`
	Delivery        dtos.Delivery `json:"transmission"`
}

func NewDeliveryResponse(requestId string, message string, statusCode int,
	transmission dtos.Delivery) DeliveryResponse {
	return DeliveryResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Delivery: transmission,
	}
}

// MultiDeliverysResponse defines the Response Content for GET multiple Delivery DTOs.
type MultiDeliverysResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Deliverys                     []dtos.Delivery `json:"transmissions"`
}

func NewMultiDeliverysResponse(requestId string, message string, statusCode int, totalCount uint32, transmissions []dtos.Delivery) MultiDeliverysResponse {
	return MultiDeliverysResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Deliverys:              transmissions,
	}
}
