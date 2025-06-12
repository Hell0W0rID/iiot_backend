package interfaces

import (
	"context"
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// IntervalClient defines the interface for interval operations
type IntervalClient interface {
	Add(ctx context.Context, reqs []dtos.IntervalRequest) ([]dtos.BaseWithIdResponse, errors.IIOTError)
	Update(ctx context.Context, reqs []dtos.UpdateIntervalRequest) ([]dtos.BaseResponse, errors.IIOTError)
	AllIntervals(ctx context.Context, offset int, limit int) (dtos.MultiIntervalsResponse, errors.IIOTError)
	IntervalByName(ctx context.Context, name string) (dtos.IntervalResponse, errors.IIOTError)
	DeleteIntervalByName(ctx context.Context, name string) (dtos.BaseResponse, errors.IIOTError)
}

// IntervalActionClient defines the interface for interval action operations
type IntervalActionClient interface {
	Add(ctx context.Context, reqs []dtos.IntervalActionRequest) ([]dtos.BaseWithIdResponse, errors.IIOTError)
	Update(ctx context.Context, reqs []dtos.UpdateIntervalActionRequest) ([]dtos.BaseResponse, errors.IIOTError)
	AllIntervalActions(ctx context.Context, offset int, limit int) (dtos.MultiIntervalActionsResponse, errors.IIOTError)
	IntervalActionsByTarget(ctx context.Context, target string, offset int, limit int) (dtos.MultiIntervalActionsResponse, errors.IIOTError)
	IntervalActionByName(ctx context.Context, name string) (dtos.IntervalActionResponse, errors.IIOTError)
	DeleteIntervalActionByName(ctx context.Context, name string) (dtos.BaseResponse, errors.IIOTError)
}