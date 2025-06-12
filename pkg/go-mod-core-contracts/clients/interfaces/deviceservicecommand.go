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

// DeviceHandlerCommandClient defines the interface for interactions with the device command endpoints on the IIOT Foundry device services.
type DeviceHandlerCommandClient interface {
	// GetCommand invokes device service's command API for issuing get(read) command
	GetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string) (*responses.DataEventResponse, errors.IIOT)
	// SetCommand invokes device service's command API for issuing set(write) command
	SetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string, settings map[string]string) (common.BaseResponse, errors.IIOT)
	// SetCommandWithObject invokes device service's set command API and the settings supports object value type
	SetCommandWithObject(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string, settings map[string]interface{}) (common.BaseResponse, errors.IIOT)
	// Discovery invokes device service's discovery API
	Discovery(ctx context.Context, baseUrl string) (common.BaseResponse, errors.IIOT)
	// ProfileScan sends an HTTP POST request to the device service's profile scan API endpoint.
	ProfileScan(ctx context.Context, baseUrl string, req requests.ProfileScanRequest) (common.BaseResponse, errors.IIOT)
	// StopDeviceDiscovery invokes device service's stop device discovery API
	StopDeviceDiscovery(ctx context.Context, baseUrl string, requestId string, queryParams map[string]string) (common.BaseResponse, errors.IIOT)
	// StopProfileScan invokes device service's stop profile scan API
	StopProfileScan(ctx context.Context, baseUrl string, deviceName string, queryParams map[string]string) (common.BaseResponse, errors.IIOT)
}
