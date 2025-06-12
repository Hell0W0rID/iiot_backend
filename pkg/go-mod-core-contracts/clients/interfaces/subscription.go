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

// EventSubscriptionClient defines the interface for interactions with the EventSubscription endpoint on the IIOT Foundry support-notifications service.
type EventSubscriptionClient interface {
	// Add adds new subscriptions.
	Add(ctx context.Context, reqs []requests.AddEventSubscriptionRequest) ([]common.BaseWithIdResponse, errors.IIOT)
	// Update updates subscriptions.
	Update(ctx context.Context, reqs []requests.UpdateEventSubscriptionRequest) ([]common.BaseResponse, errors.IIOT)
	// AllEventSubscriptions queries subscriptions with offset and limit
	AllEventSubscriptions(ctx context.Context, offset int, limit int) (responses.MultiEventSubscriptionsResponse, errors.IIOT)
	// EventSubscriptionsByCategory queries subscriptions with category, offset and limit
	EventSubscriptionsByCategory(ctx context.Context, category string, offset int, limit int) (responses.MultiEventSubscriptionsResponse, errors.IIOT)
	// EventSubscriptionsByLabel queries subscriptions with label, offset and limit
	EventSubscriptionsByLabel(ctx context.Context, label string, offset int, limit int) (responses.MultiEventSubscriptionsResponse, errors.IIOT)
	// EventSubscriptionsByReceiver queries subscriptions with receiver, offset and limit
	EventSubscriptionsByReceiver(ctx context.Context, receiver string, offset int, limit int) (responses.MultiEventSubscriptionsResponse, errors.IIOT)
	// EventSubscriptionByName query subscription by name.
	EventSubscriptionByName(ctx context.Context, name string) (responses.EventSubscriptionResponse, errors.IIOT)
	// DeleteEventSubscriptionByName deletes a subscription by name.
	DeleteEventSubscriptionByName(ctx context.Context, name string) (common.BaseResponse, errors.IIOT)
}
