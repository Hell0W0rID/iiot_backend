//
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"errors"
	"strings"

	"iiot-backend/pkg/go-mod-core-contracts/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	iiotErrors "iiot-backend/pkg/go-mod-core-contracts/errors"
	"iiot-backend/pkg/go-mod-core-contracts/models"
)

// DeviceTemplateRequest defines the Request Content for POST DeviceTemplate DTO.
type DeviceTemplateRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	Profile               dtos.DeviceTemplate `json:"profile"`
}

// Validate satisfies the Validator interface
func (dp DeviceTemplateRequest) Validate() error {
	err := common.Validate(dp)
	if err != nil {
		// The DeviceTemplateBasicInfo is the internal struct in Golang programming, not in the Profile model,
		// so it should be hidden from the error messages.
		err = errors.New(strings.ReplaceAll(err.Error(), ".DeviceTemplateBasicInfo", ""))
		return iiotErrors.NewCommonIIOT(iiotErrors.KindContractInvalid, "", err)
	}
	return dp.Profile.Validate()
}

// UnmarshalJSON implements the Unmarshaler interface for the DeviceTemplateRequest type
func (dp *DeviceTemplateRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		Profile dtos.DeviceTemplate
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return iiotErrors.NewCommonIIOT(iiotErrors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*dp = DeviceTemplateRequest(alias)

	// validate DeviceTemplateRequest DTO
	if err := dp.Validate(); err != nil {
		return err
	}

	// Normalize resource's value type
	for i, resource := range dp.Profile.DeviceResources {
		valueType, err := common.NormalizeValueType(resource.Properties.ValueType)
		if err != nil {
			return iiotErrors.NewCommonIIOTWrapper(err)
		}
		dp.Profile.DeviceResources[i].Properties.ValueType = valueType
	}
	return nil
}

// DeviceTemplateReqToDeviceTemplateModel transforms the DeviceTemplateRequest DTO to the DeviceTemplate model
func DeviceTemplateReqToDeviceTemplateModel(addReq DeviceTemplateRequest) (DeviceTemplates models.DeviceTemplate) {
	return dtos.ToDeviceTemplateModel(addReq.Profile)
}

// DeviceTemplateReqToDeviceTemplateModels transforms the DeviceTemplateRequest DTO array to the DeviceTemplate model array
func DeviceTemplateReqToDeviceTemplateModels(addRequests []DeviceTemplateRequest) (DeviceTemplates []models.DeviceTemplate) {
	for _, req := range addRequests {
		dp := DeviceTemplateReqToDeviceTemplateModel(req)
		DeviceTemplates = append(DeviceTemplates, dp)
	}
	return DeviceTemplates
}

func NewDeviceTemplateRequest(dto dtos.DeviceTemplate) DeviceTemplateRequest {
	return DeviceTemplateRequest{
		BaseRequest: dtoCommon.NewBaseRequest(),
		Profile:     dto,
	}
}
