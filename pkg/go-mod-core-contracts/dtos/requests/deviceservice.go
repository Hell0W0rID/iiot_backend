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

// AddDeviceHandlerRequest defines the Request Content for POST DeviceHandler DTO.
type AddDeviceHandlerRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	Service               dtos.DeviceHandler `json:"service"`
}

// Validate satisfies the Validator interface
func (ds AddDeviceHandlerRequest) Validate() error {
	err := common.Validate(ds)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddDeviceHandlerRequest type
func (ds *AddDeviceHandlerRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		Service dtos.DeviceHandler
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	if alias.Service.Properties == nil {
		alias.Service.Properties = make(map[string]any)
	}

	*ds = AddDeviceHandlerRequest(alias)

	// validate AddDeviceHandlerRequest DTO
	if err := ds.Validate(); err != nil {
		return err
	}
	return nil
}

// AddDeviceHandlerReqToDeviceHandlerModels transforms the AddDeviceHandlerRequest DTO array to the DeviceHandler model array
func AddDeviceHandlerReqToDeviceHandlerModels(addRequests []AddDeviceHandlerRequest) (DeviceHandlers []models.DeviceHandler) {
	for _, req := range addRequests {
		ds := dtos.ToDeviceHandlerModel(req.Service)
		DeviceHandlers = append(DeviceHandlers, ds)
	}
	return DeviceHandlers
}

// UpdateDeviceHandlerRequest defines the Request Content for PUT event as pushed DTO.
type UpdateDeviceHandlerRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	Service               dtos.UpdateDeviceHandler `json:"service"`
}

// Validate satisfies the Validator interface
func (ds UpdateDeviceHandlerRequest) Validate() error {
	err := common.Validate(ds)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateDeviceHandlerRequest type
func (ds *UpdateDeviceHandlerRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		Service dtos.UpdateDeviceHandler
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*ds = UpdateDeviceHandlerRequest(alias)

	// validate UpdateDeviceHandlerRequest DTO
	if err := ds.Validate(); err != nil {
		return err
	}
	return nil
}

// ReplaceDeviceHandlerModelFieldsWithDTO replace existing DeviceHandler's fields with DTO patch
func ReplaceDeviceHandlerModelFieldsWithDTO(ds *models.DeviceHandler, patch dtos.UpdateDeviceHandler) {
	if patch.Description != nil {
		ds.Description = *patch.Description
	}
	if patch.ServiceState != nil {
		ds.ServiceState = models.ServiceState(*patch.ServiceState)
	}
	if patch.Labels != nil {
		ds.Labels = patch.Labels
	}
	if patch.ServiceAddress != nil {
		ds.ServiceAddress = *patch.ServiceAddress
	}
	if patch.Properties != nil {
		ds.Properties = patch.Properties
	}
}

func NewAddDeviceHandlerRequest(dto dtos.DeviceHandler) AddDeviceHandlerRequest {
	return AddDeviceHandlerRequest{
		BaseRequest: dtoCommon.NewBaseRequest(),
		Service:     dto,
	}
}

func NewUpdateDeviceHandlerRequest(dto dtos.UpdateDeviceHandler) UpdateDeviceHandlerRequest {
	return UpdateDeviceHandlerRequest{
		BaseRequest: dtoCommon.NewBaseRequest(),
		Service:     dto,
	}
}
