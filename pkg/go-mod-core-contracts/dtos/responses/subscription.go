//
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
)

// EventSubscriptionResponse defines the EventSubscription Content for GET EventSubscription DTOs.
type EventSubscriptionResponse struct {
	common.BaseResponse `json:",inline"`
	EventSubscription        dtos.EventSubscription `json:"subscription"`
}

func NewEventSubscriptionResponse(requestId string, message string, statusCode int,
	subscription dtos.EventSubscription) EventSubscriptionResponse {
	return EventSubscriptionResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		EventSubscription: subscription,
	}
}

// MultiEventSubscriptionsResponse defines the EventSubscription Content for GET multiple EventSubscription DTOs.
type MultiEventSubscriptionsResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	EventSubscriptions                     []dtos.EventSubscription `json:"subscriptions"`
}

func NewMultiEventSubscriptionsResponse(requestId string, message string, statusCode int, totalCount uint32, subscriptions []dtos.EventSubscription) MultiEventSubscriptionsResponse {
	return MultiEventSubscriptionsResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		EventSubscriptions:              subscriptions,
	}
}
