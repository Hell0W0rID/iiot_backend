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

// DeviceTemplateBasicInfoRequest defines the Request Content for PATCH UpdateDeviceTemplateBasicInfo DTO.
type DeviceTemplateBasicInfoRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	BasicInfo             dtos.UpdateDeviceTemplateBasicInfo `json:"basicinfo"`
}

// Validate satisfies the Validator interface
func (d DeviceTemplateBasicInfoRequest) Validate() error {
	err := common.Validate(d)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateDeviceRequest type
func (d *DeviceTemplateBasicInfoRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		BasicInfo dtos.UpdateDeviceTemplateBasicInfo
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*d = DeviceTemplateBasicInfoRequest(alias)

	// validate DeviceTemplateBasicInfoRequest DTO
	if err := d.Validate(); err != nil {
		return err
	}
	return nil
}

// ReplaceDeviceTemplateModelBasicInfoFieldsWithDTO replace existing DeviceTemplate's basic info fields with DTO patch
func ReplaceDeviceTemplateModelBasicInfoFieldsWithDTO(dp *models.DeviceTemplate, patch dtos.UpdateDeviceTemplateBasicInfo) {
	if patch.Description != nil {
		dp.Description = *patch.Description
	}
	if patch.Manufacturer != nil {
		dp.Manufacturer = *patch.Manufacturer
	}
	if patch.Model != nil {
		dp.Model = *patch.Model
	}
	if patch.Labels != nil {
		dp.Labels = patch.Labels
	}
}
