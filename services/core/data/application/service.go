package application

import (
	"context"
	"net/http"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/models"

	"iiot-backend/services/core/data/infrastructure/interfaces"
)

// DataService contains references to dependencies required by the service
type DataService struct {
	dbClient interfaces.DBClient
}

// NewDataService creates a new instance of DataService
func NewDataService() *DataService {
	return &DataService{}
}

// AddEvent adds a new event
func (ds *DataService) AddEvent(ctx context.Context, addEventRequest requests.AddEventRequest) (responses.BaseResponse, errors.EdgeX) {
	event := dtos.ToEventModel(addEventRequest.Event)
	
	// Validate event
	if event.DeviceName == "" {
		return responses.BaseResponse{}, errors.NewCommonEdgeX(errors.KindContractInvalid, "device name is required", nil)
	}

	// Store event in database
	id, err := ds.dbClient.AddEvent(event)
	if err != nil {
		return responses.BaseResponse{}, errors.NewCommonEdgeXWrapper(err)
	}

	return responses.NewBaseResponse("", "", http.StatusCreated, id), nil
}

// AllEvents returns events according to the offset, limit, and device name
func (ds *DataService) AllEvents(ctx context.Context, offset int, limit int, deviceName string) (responses.MultipleEventResponse, errors.EdgeX) {
	var events []models.Event
	var totalCount uint32
	var err errors.EdgeX

	if deviceName != "" {
		events, totalCount, err = ds.dbClient.EventsByDeviceName(offset, limit, deviceName)
	} else {
		events, totalCount, err = ds.dbClient.AllEvents(offset, limit)
	}

	if err != nil {
		return responses.MultipleEventResponse{}, errors.NewCommonEdgeXWrapper(err)
	}

	eventDTOs := make([]dtos.Event, len(events))
	for i, event := range events {
		eventDTOs[i] = dtos.FromEventModelToDTO(event)
	}

	return responses.NewMultipleEventResponse("", "", http.StatusOK, totalCount, eventDTOs), nil
}

// EventById returns an event by its database generated id
func (ds *DataService) EventById(ctx context.Context, id string) (responses.EventResponse, errors.EdgeX) {
	if id == "" {
		return responses.EventResponse{}, errors.NewCommonEdgeX(errors.KindContractInvalid, "event ID is required", nil)
	}

	event, err := ds.dbClient.EventById(id)
	if err != nil {
		return responses.EventResponse{}, errors.NewCommonEdgeXWrapper(err)
	}

	eventDTO := dtos.FromEventModelToDTO(event)
	return responses.NewEventResponse("", "", http.StatusOK, eventDTO), nil
}

// DeleteEventById deletes an event by its database generated id
func (ds *DataService) DeleteEventById(ctx context.Context, id string) (responses.BaseResponse, errors.EdgeX) {
	if id == "" {
		return responses.BaseResponse{}, errors.NewCommonEdgeX(errors.KindContractInvalid, "event ID is required", nil)
	}

	err := ds.dbClient.DeleteEventById(id)
	if err != nil {
		return responses.BaseResponse{}, errors.NewCommonEdgeXWrapper(err)
	}

	return responses.NewBaseResponse("", "", http.StatusOK, "Event deleted"), nil
}

// EventsByTimeRange returns events according to the specified time range, offset, and limit
func (ds *DataService) EventsByTimeRange(ctx context.Context, start int64, end int64, offset int, limit int) (responses.MultipleEventResponse, errors.EdgeX) {
	events, totalCount, err := ds.dbClient.EventsByTimeRange(start, end, offset, limit)
	if err != nil {
		return responses.MultipleEventResponse{}, errors.NewCommonEdgeXWrapper(err)
	}

	eventDTOs := make([]dtos.Event, len(events))
	for i, event := range events {
		eventDTOs[i] = dtos.FromEventModelToDTO(event)
	}

	return responses.NewMultipleEventResponse("", "", http.StatusOK, totalCount, eventDTOs), nil
}

// AddReading adds a new reading
func (ds *DataService) AddReading(ctx context.Context, addReadingRequest requests.AddReadingRequest) (responses.BaseResponse, errors.EdgeX) {
	reading := dtos.ToReadingModel(addReadingRequest.Reading)
	
	// Validate reading
	if reading.DeviceName == "" {
		return responses.BaseResponse{}, errors.NewCommonEdgeX(errors.KindContractInvalid, "device name is required", nil)
	}
	if reading.ResourceName == "" {
		return responses.BaseResponse{}, errors.NewCommonEdgeX(errors.KindContractInvalid, "resource name is required", nil)
	}

	// Store reading in database
	id, err := ds.dbClient.AddReading(reading)
	if err != nil {
		return responses.BaseResponse{}, errors.NewCommonEdgeXWrapper(err)
	}

	return responses.NewBaseResponse("", "", http.StatusCreated, id), nil
}

// ReadingsByDeviceName returns readings by device name according to the offset and limit
func (ds *DataService) ReadingsByDeviceName(ctx context.Context, deviceName string, offset int, limit int) (responses.MultipleReadingResponse, errors.EdgeX) {
	if deviceName == "" {
		return responses.MultipleReadingResponse{}, errors.NewCommonEdgeX(errors.KindContractInvalid, "device name is required", nil)
	}

	readings, totalCount, err := ds.dbClient.ReadingsByDeviceName(offset, limit, deviceName)
	if err != nil {
		return responses.MultipleReadingResponse{}, errors.NewCommonEdgeXWrapper(err)
	}

	readingDTOs := make([]dtos.BaseReading, len(readings))
	for i, reading := range readings {
		readingDTOs[i] = dtos.FromReadingModelToDTO(reading)
	}

	return responses.NewMultipleReadingResponse("", "", http.StatusOK, totalCount, readingDTOs), nil
}

// ReadingsByResourceName returns readings by resource name according to the offset and limit
func (ds *DataService) ReadingsByResourceName(ctx context.Context, resourceName string, offset int, limit int) (responses.MultipleReadingResponse, errors.EdgeX) {
	if resourceName == "" {
		return responses.MultipleReadingResponse{}, errors.NewCommonEdgeX(errors.KindContractInvalid, "resource name is required", nil)
	}

	readings, totalCount, err := ds.dbClient.ReadingsByResourceName(offset, limit, resourceName)
	if err != nil {
		return responses.MultipleReadingResponse{}, errors.NewCommonEdgeXWrapper(err)
	}

	readingDTOs := make([]dtos.BaseReading, len(readings))
	for i, reading := range readings {
		readingDTOs[i] = dtos.FromReadingModelToDTO(reading)
	}

	return responses.NewMultipleReadingResponse("", "", http.StatusOK, totalCount, readingDTOs), nil
}