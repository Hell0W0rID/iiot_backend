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

// AddDeviceResourceRequest defines the Request Content for POST DeviceResource DTO.
type AddDeviceResourceRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProfileName           string              `json:"profileName" validate:"required,iiot-dto-none-empty-string"`
	Resource              dtos.DeviceResource `json:"resource"`
}

func (request AddDeviceResourceRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddDeviceResourceReques type
func (dr *AddDeviceResourceRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		ProfileName string
		Resource    dtos.DeviceResource
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*dr = AddDeviceResourceRequest(alias)

	if err := dr.Validate(); err != nil {
		return err
	}

	return nil
}

// UpdateDeviceResourceRequest defines the Request Content for PATCH DeviceResource DTO.
type UpdateDeviceResourceRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProfileName           string                    `json:"profileName" validate:"required,iiot-dto-none-empty-string"`
	Resource              dtos.UpdateDeviceResource `json:"resource"`
}

func (request UpdateDeviceResourceRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateDeviceResourceRequest type
func (dr *UpdateDeviceResourceRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		ProfileName string
		Resource    dtos.UpdateDeviceResource
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*dr = UpdateDeviceResourceRequest(alias)

	if err := dr.Validate(); err != nil {
		return err
	}

	return nil
}

// ReplaceDeviceResourceModelFieldsWithDTO replace existing DeviceResource's fields with DTO patch
func ReplaceDeviceResourceModelFieldsWithDTO(dr *models.DeviceResource, patch dtos.UpdateDeviceResource) {
	if patch.Description != nil {
		dr.Description = *patch.Description
	}
	if patch.IsHidden != nil {
		dr.IsHidden = *patch.IsHidden
	}
}
