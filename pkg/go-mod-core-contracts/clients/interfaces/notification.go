//
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/requests"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// AlertClient defines the interface for interactions with the Alert endpoint on the IIOT Foundry support-notifications service.
type AlertClient interface {
	// SendAlert sends new notifications.
	SendAlert(ctx context.Context, reqs []requests.AddAlertRequest) ([]common.BaseWithIdResponse, errors.IIOT)
	// AlertById query notification by id.
	AlertById(ctx context.Context, id string) (responses.AlertResponse, errors.IIOT)
	// DeleteAlertById deletes a notification by id.
	DeleteAlertById(ctx context.Context, id string) (common.BaseResponse, errors.IIOT)
	// AlertsByCategory queries notifications with category, offset, ack and limit
	AlertsByCategory(ctx context.Context, category string, offset int, limit int, ack string) (responses.MultiAlertsResponse, errors.IIOT)
	// AlertsByLabel queries notifications with label, offset, ack and limit
	AlertsByLabel(ctx context.Context, label string, offset int, limit int, ack string) (responses.MultiAlertsResponse, errors.IIOT)
	// AlertsByStatus queries notifications with status, offset, ack and limit
	AlertsByStatus(ctx context.Context, status string, offset int, limit int, ack string) (responses.MultiAlertsResponse, errors.IIOT)
	// AlertsByTimeRange query notifications with time range, offset, ack and limit
	AlertsByTimeRange(ctx context.Context, start, end int64, offset int, limit int, ack string) (responses.MultiAlertsResponse, errors.IIOT)
	// AlertsByEventSubscriptionName query notifications with subscriptionName, offset, ack and limit
	AlertsByEventSubscriptionName(ctx context.Context, subscriptionName string, offset int, limit int, ack string) (responses.MultiAlertsResponse, errors.IIOT)
	// CleanupAlertsByAge removes notifications that are older than age. And the corresponding transmissions will also be deleted.
	// Age is supposed in milliseconds since modified timestamp
	CleanupAlertsByAge(ctx context.Context, age int) (common.BaseResponse, errors.IIOT)
	// CleanupAlerts removes notifications and the corresponding transmissions.
	CleanupAlerts(ctx context.Context) (common.BaseResponse, errors.IIOT)
	// DeleteProcessedAlertsByAge removes processed notifications that are older than age. And the corresponding transmissions will also be deleted.
	// Age is supposed in milliseconds since modified timestamp
	// Please notice that this API is only for processed notifications (status = PROCESSED). If the deletion purpose includes each kind of notifications, please refer to cleanup API.
	DeleteProcessedAlertsByAge(ctx context.Context, age int) (common.BaseResponse, errors.IIOT)
	// AlertsByQueryConditions queries notifications with offset, limit, acknowledgement status, category and time range
	AlertsByQueryConditions(ctx context.Context, offset, limit int, ack string, conditionReq requests.GetAlertRequest) (responses.MultiAlertsResponse, errors.IIOT)
	// DeleteAlertByIds deletes notifications by ids
	DeleteAlertByIds(ctx context.Context, ids []string) (common.BaseResponse, errors.IIOT)
	// UpdateAlertAckStatusByIds updates existing notification's acknowledgement status
	UpdateAlertAckStatusByIds(ctx context.Context, ack bool, ids []string) (common.BaseResponse, errors.IIOT)
}
