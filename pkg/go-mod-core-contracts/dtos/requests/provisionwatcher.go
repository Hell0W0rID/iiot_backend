//
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

	"iiot-backend/pkg/go-mod-core-contracts/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
	"iiot-backend/pkg/go-mod-core-contracts/models"
)

// AddDeviceWatcherRequest defines the Request Content for POST DeviceWatcher DTO.
type AddDeviceWatcherRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	DeviceWatcher      dtos.DeviceWatcher `json:"provisionWatcher"`
}

// Validate satisfies the Validator interface
func (pw *AddDeviceWatcherRequest) Validate() error {
	err := common.Validate(pw)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddDeviceWatcherRequest type
func (pw *AddDeviceWatcherRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		DeviceWatcher dtos.DeviceWatcher
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	if alias.DeviceWatcher.DiscoveredDevice.Properties == nil {
		alias.DeviceWatcher.DiscoveredDevice.Properties = make(map[string]any)
	}

	*pw = AddDeviceWatcherRequest(alias)

	// validate AddDeviceRequest DTO
	if err := pw.Validate(); err != nil {
		return err
	}
	return nil
}

// AddDeviceWatcherReqToDeviceWatcherModels transforms the AddDeviceWatcherRequest DTO array to the DeviceWatcher model array
func AddDeviceWatcherReqToDeviceWatcherModels(addRequests []AddDeviceWatcherRequest) (DeviceWatchers []models.DeviceWatcher) {
	for _, req := range addRequests {
		d := dtos.ToDeviceWatcherModel(req.DeviceWatcher)
		DeviceWatchers = append(DeviceWatchers, d)
	}
	return DeviceWatchers
}

// UpdateDeviceWatcherRequest defines the Request Content for PUT event as pushed DTO.
type UpdateDeviceWatcherRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	DeviceWatcher      dtos.UpdateDeviceWatcher `json:"provisionWatcher"`
}

// Validate satisfies the Validator interface
func (pw *UpdateDeviceWatcherRequest) Validate() error {
	err := common.Validate(pw)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateDeviceWatcherRequest type
func (pw *UpdateDeviceWatcherRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		DeviceWatcher dtos.UpdateDeviceWatcher
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*pw = UpdateDeviceWatcherRequest(alias)

	// validate UpdateDeviceRequest DTO
	if err := pw.Validate(); err != nil {
		return err
	}
	return nil
}

// ReplaceDeviceWatcherModelFieldsWithDTO replace existing DeviceWatcher's fields with DTO patch
func ReplaceDeviceWatcherModelFieldsWithDTO(pw *models.DeviceWatcher, patch dtos.UpdateDeviceWatcher) {
	if patch.Labels != nil {
		pw.Labels = patch.Labels
	}
	if patch.Identifiers != nil {
		pw.Identifiers = patch.Identifiers
	}
	if patch.BlockingIdentifiers != nil {
		pw.BlockingIdentifiers = patch.BlockingIdentifiers
	}
	if patch.ServiceState != nil {
		pw.ServiceState = models.ServiceState(*patch.ServiceState)
	}
	if patch.DiscoveredDevice.ProfileName != nil {
		pw.DiscoveredDevice.ProfileName = *patch.DiscoveredDevice.ProfileName
	}
	if patch.ServiceName != nil {
		pw.ServiceName = *patch.ServiceName
	}
	if patch.DiscoveredDevice.ServiceState != nil {
		pw.DiscoveredDevice.ServiceState = models.ServiceState(*patch.DiscoveredDevice.ServiceState)
	}
	if patch.DiscoveredDevice.AutoDataEvents != nil {
		pw.DiscoveredDevice.AutoDataEvents = dtos.ToAutoDataEventModels(patch.DiscoveredDevice.AutoDataEvents)
	}
	if patch.DiscoveredDevice.Properties != nil {
		pw.DiscoveredDevice.Properties = patch.DiscoveredDevice.Properties
	}
}

func NewAddDeviceWatcherRequest(dto dtos.DeviceWatcher) AddDeviceWatcherRequest {
	return AddDeviceWatcherRequest{
		BaseRequest:      dtoCommon.NewBaseRequest(),
		DeviceWatcher: dto,
	}
}

func NewUpdateDeviceWatcherRequest(dto dtos.UpdateDeviceWatcher) UpdateDeviceWatcherRequest {
	return UpdateDeviceWatcherRequest{
		BaseRequest:      dtoCommon.NewBaseRequest(),
		DeviceWatcher: dto,
	}
}
