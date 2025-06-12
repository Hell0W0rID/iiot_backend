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

// AddDeviceActionRequest defines the Request Content for POST DeviceAction DTO.
type AddDeviceActionRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProfileName           string             `json:"profileName" validate:"required,iiot-dto-none-empty-string"`
	DeviceAction         dtos.DeviceAction `json:"deviceCommand"`
}

// Validate satisfies the Validator interface
func (request AddDeviceActionRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddDeviceActionRequest type
func (dc *AddDeviceActionRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		ProfileName   string
		DeviceAction dtos.DeviceAction
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*dc = AddDeviceActionRequest(alias)

	if err := dc.Validate(); err != nil {
		return err
	}

	return nil
}

// UpdateDeviceActionRequest defines the Request Content for PATCH DeviceAction DTO.
type UpdateDeviceActionRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProfileName           string                   `json:"profileName" validate:"required,iiot-dto-none-empty-string"`
	DeviceAction         dtos.UpdateDeviceAction `json:"deviceCommand"`
}

// Validate satisfies the Validator interface
func (request UpdateDeviceActionRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateDeviceActionRequest type
func (dc *UpdateDeviceActionRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		ProfileName   string
		DeviceAction dtos.UpdateDeviceAction
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*dc = UpdateDeviceActionRequest(alias)

	if err := dc.Validate(); err != nil {
		return err
	}

	return nil
}

// ReplaceDeviceActionModelFieldsWithDTO replace existing DeviceAction's fields with DTO patch
func ReplaceDeviceActionModelFieldsWithDTO(dc *models.DeviceAction, patch dtos.UpdateDeviceAction) {
	if patch.IsHidden != nil {
		dc.IsHidden = *patch.IsHidden
	}
}
