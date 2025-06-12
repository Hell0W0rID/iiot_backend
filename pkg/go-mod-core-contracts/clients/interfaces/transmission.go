//
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// DeliveryClient defines the interface for interactions with the Delivery endpoint on the IIOT Foundry support-notifications service.
type DeliveryClient interface {
	// DeliveryById query transmission by id.
	DeliveryById(ctx context.Context, id string) (responses.DeliveryResponse, errors.IIOT)
	// DeliverysByTimeRange query transmissions with time range, offset and limit
	DeliverysByTimeRange(ctx context.Context, start, end int64, offset int, limit int) (responses.MultiDeliverysResponse, errors.IIOT)
	// AllDeliverys query transmissions with offset and limit
	AllDeliverys(ctx context.Context, offset int, limit int) (responses.MultiDeliverysResponse, errors.IIOT)
	// DeliverysByStatus queries transmissions with status, offset and limit
	DeliverysByStatus(ctx context.Context, status string, offset int, limit int) (responses.MultiDeliverysResponse, errors.IIOT)
	// DeleteProcessedDeliverysByAge deletes the processed transmissions if the current timestamp minus their created timestamp is less than the age parameter.
	DeleteProcessedDeliverysByAge(ctx context.Context, age int) (common.BaseResponse, errors.IIOT)
	// DeliverysByEventSubscriptionName query transmissions with subscriptionName, offset and limit
	DeliverysByEventSubscriptionName(ctx context.Context, subscriptionName string, offset int, limit int) (responses.MultiDeliverysResponse, errors.IIOT)
	// DeliverysByAlertId query transmissions with notification id, offset and limit
	DeliverysByAlertId(ctx context.Context, id string, offset int, limit int) (responses.MultiDeliverysResponse, errors.IIOT)
}
