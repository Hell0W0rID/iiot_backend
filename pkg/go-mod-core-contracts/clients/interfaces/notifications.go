package interfaces

import (
	"context"
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// AlertClient defines the interface for notification operations
type AlertClient interface {
	SendAlert(ctx context.Context, reqs []dtos.AlertRequest) ([]dtos.BaseWithIdResponse, errors.IIOTError)
	AlertsByCategory(ctx context.Context, category string, offset int, limit int) (dtos.MultiAlertsResponse, errors.IIOTError)
	AlertsByLabel(ctx context.Context, label string, offset int, limit int) (dtos.MultiAlertsResponse, errors.IIOTError)
	AlertsByStatus(ctx context.Context, status string, offset int, limit int) (dtos.MultiAlertsResponse, errors.IIOTError)
	AlertsByTimeRange(ctx context.Context, start int, end int, offset int, limit int) (dtos.MultiAlertsResponse, errors.IIOTError)
	DeleteAlertsByAge(ctx context.Context, age int) (dtos.BaseResponse, errors.IIOTError)
	DeleteProcessedAlertsByAge(ctx context.Context, age int) (dtos.BaseResponse, errors.IIOTError)
	CleanupAlerts(ctx context.Context) (dtos.BaseResponse, errors.IIOTError)
}

// EventSubscriptionClient defines the interface for subscription operations
type EventSubscriptionClient interface {
	Add(ctx context.Context, reqs []dtos.EventSubscriptionRequest) ([]dtos.BaseWithIdResponse, errors.IIOTError)
	Update(ctx context.Context, reqs []dtos.EventSubscriptionRequest) ([]dtos.BaseResponse, errors.IIOTError)
	AllEventSubscriptions(ctx context.Context, offset int, limit int) (dtos.MultiEventSubscriptionsResponse, errors.IIOTError)
	EventSubscriptionsByCategory(ctx context.Context, category string, offset int, limit int) (dtos.MultiEventSubscriptionsResponse, errors.IIOTError)
	EventSubscriptionsByLabel(ctx context.Context, label string, offset int, limit int) (dtos.MultiEventSubscriptionsResponse, errors.IIOTError)
	EventSubscriptionByName(ctx context.Context, name string) (dtos.EventSubscriptionResponse, errors.IIOTError)
	DeleteEventSubscriptionByName(ctx context.Context, name string) (dtos.BaseResponse, errors.IIOTError)
}

// DeliveryClient defines the interface for transmission operations
type DeliveryClient interface {
	AllDeliverys(ctx context.Context, offset int, limit int) (dtos.MultiDeliverysResponse, errors.IIOTError)
	DeliverysByTimeRange(ctx context.Context, start int, end int, offset int, limit int) (dtos.MultiDeliverysResponse, errors.IIOTError)
	DeliverysByStatus(ctx context.Context, status string, offset int, limit int) (dtos.MultiDeliverysResponse, errors.IIOTError)
	DeleteProcessedDeliverysByAge(ctx context.Context, age int) (dtos.BaseResponse, errors.IIOTError)
}