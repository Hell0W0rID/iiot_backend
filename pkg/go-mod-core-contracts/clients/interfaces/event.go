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

// DataEventClient defines the interface for interactions with the DataEvent endpoint on the IIOT Foundry iiot-data service.
type DataEventClient interface {
	// Add adds new event.
	Add(ctx context.Context, serviceName string, req requests.AddDataEventRequest) (common.BaseWithIdResponse, errors.IIOT)
	// AllDataEvents returns all events sorted in descending order of created time.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	AllDataEvents(ctx context.Context, offset, limit int) (responses.MultiDataEventsResponse, errors.IIOT)
	// DataEventCount returns a count of all of events currently stored in the database.
	DataEventCount(ctx context.Context) (common.CountResponse, errors.IIOT)
	// DataEventCountByDeviceName returns a count of all of events currently stored in the database, sourced from the specified device.
	DataEventCountByDeviceName(ctx context.Context, name string) (common.CountResponse, errors.IIOT)
	// DataEventsByDeviceName returns a portion of the entire events according to the device name, offset and limit parameters. DataEvents are sorted in descending order of created time.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	DataEventsByDeviceName(ctx context.Context, name string, offset, limit int) (responses.MultiDataEventsResponse, errors.IIOT)
	// DeleteByDeviceName deletes all events for the specified device.
	DeleteByDeviceName(ctx context.Context, name string) (common.BaseResponse, errors.IIOT)
	// DataEventsByTimeRange returns events between a given start and end date/time. DataEvents are sorted in descending order of created time.
	// start, end: Unix timestamp, indicating the date/time range.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	DataEventsByTimeRange(ctx context.Context, start, end int64, offset, limit int) (responses.MultiDataEventsResponse, errors.IIOT)
	// DeleteByAge deletes events that are older than the given age. Age is supposed in milliseconds from created timestamp.
	DeleteByAge(ctx context.Context, age int) (common.BaseResponse, errors.IIOT)
	// DeleteById deletes an event by its id
	DeleteById(ctx context.Context, id string) (common.BaseResponse, errors.IIOT)
}
