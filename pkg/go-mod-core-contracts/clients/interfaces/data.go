package interfaces

import (
	"context"
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// DataEventClient defines the interface for event operations
type DataEventClient interface {
	Add(ctx context.Context, reqs []dtos.DataEventRequest) ([]dtos.BaseWithIdResponse, errors.IIOTError)
	AllDataEvents(ctx context.Context, offset int, limit int) (dtos.MultiDataEventsResponse, errors.IIOTError)
	DataEventsByDeviceName(ctx context.Context, name string, offset int, limit int) (dtos.MultiDataEventsResponse, errors.IIOTError)
	DeleteByDeviceName(ctx context.Context, name string) (dtos.BaseResponse, errors.IIOTError)
	DataEventsByTimeRange(ctx context.Context, start int, end int, offset int, limit int) (dtos.MultiDataEventsResponse, errors.IIOTError)
	DataEventCount(ctx context.Context) (dtos.CountResponse, errors.IIOTError)
	DataEventCountByDeviceName(ctx context.Context, name string) (dtos.CountResponse, errors.IIOTError)
}

// MeasurementClient defines the interface for reading operations  
type MeasurementClient interface {
	AllMeasurements(ctx context.Context, offset int, limit int) (dtos.MultiMeasurementsResponse, errors.IIOTError)
	MeasurementsByDeviceName(ctx context.Context, name string, offset int, limit int) (dtos.MultiMeasurementsResponse, errors.IIOTError)
	MeasurementsByResourceName(ctx context.Context, name string, offset int, limit int) (dtos.MultiMeasurementsResponse, errors.IIOTError)
	MeasurementsByTimeRange(ctx context.Context, start int, end int, offset int, limit int) (dtos.MultiMeasurementsResponse, errors.IIOTError)
	MeasurementCount(ctx context.Context) (dtos.CountResponse, errors.IIOTError)
	MeasurementCountByDeviceName(ctx context.Context, name string) (dtos.CountResponse, errors.IIOTError)
}