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

// MeasurementClient defines the interface for interactions with the Measurement endpoint on the IIOT Foundry iiot-data service.
type MeasurementClient interface {
	// AllMeasurements returns all readings sorted in descending order of created time.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	AllMeasurements(ctx context.Context, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT)
	// MeasurementCount returns a count of all readings currently stored in the database.
	MeasurementCount(ctx context.Context) (common.CountResponse, errors.IIOT)
	// MeasurementCountByDeviceName returns a count of all readings currently stored in the database, sourced from the specified device.
	MeasurementCountByDeviceName(ctx context.Context, name string) (common.CountResponse, errors.IIOT)
	// MeasurementsByDeviceName returns a portion of the entire readings according to the device name, offset and limit parameters. Measurements are sorted in descending order of created time.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	MeasurementsByDeviceName(ctx context.Context, name string, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT)
	// MeasurementsByResourceName returns a portion of the entire readings according to the device resource name, offset and limit parameters. Measurements are sorted in descending order of created time.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	MeasurementsByResourceName(ctx context.Context, name string, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT)
	// MeasurementsByTimeRange returns readings between a given start and end date/time. Measurements are sorted in descending order of created time.
	// start, end: Unix timestamp, indicating the date/time range.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	MeasurementsByTimeRange(ctx context.Context, start, end int64, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT)
	// MeasurementsByResourceNameAndTimeRange returns readings by resource name and specified time range. Measurements are sorted in descending order of origin time.
	// start, end: Unix timestamp, indicating the date/time range
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	MeasurementsByResourceNameAndTimeRange(ctx context.Context, name string, start, end int64, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT)
	// MeasurementsByDeviceNameAndResourceName returns readings by device name and resource name. Measurements are sorted in descending order of origin time.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	MeasurementsByDeviceNameAndResourceName(ctx context.Context, deviceName, resourceName string, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT)
	// MeasurementsByDeviceNameAndResourceNameAndTimeRange returns readings by device name, resource name and specified time range. Measurements are sorted in descending order of origin time.
	// start, end: Unix timestamp, indicating the date/time range
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	MeasurementsByDeviceNameAndResourceNameAndTimeRange(ctx context.Context, deviceName, resourceName string, start, end int64, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT)
	// MeasurementsByDeviceNameAndResourceNamesAndTimeRange returns readings by device name, multiple resource names and specified time range. Measurements are sorted in descending order of origin time.
	// If none of resourceNames is specified, return all Measurements under specified deviceName and within specified time range
	// start, end: Unix timestamp, indicating the date/time range
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	MeasurementsByDeviceNameAndResourceNamesAndTimeRange(ctx context.Context, deviceName string, resourceNames []string, start, end int64, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT)
}
