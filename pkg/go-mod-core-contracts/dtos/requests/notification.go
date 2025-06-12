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

// AddAlertRequest defines the Request Content for POST Alert DTO.
type AddAlertRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	Alert          dtos.Alert `json:"notification"`
}

// Validate satisfies the Validator interface
func (request AddAlertRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddAlertRequest type
func (request *AddAlertRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		Alert dtos.Alert
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*request = AddAlertRequest(alias)

	// validate AddAlertRequest DTO
	if err := request.Validate(); err != nil {
		return err
	}
	return nil
}

// AddAlertReqToAlertModels transforms the AddAlertRequest DTO array to the AddAlertRequest model array
func AddAlertReqToAlertModels(reqs []AddAlertRequest) (n []models.Alert) {
	for _, req := range reqs {
		d := dtos.ToAlertModel(req.Alert)
		n = append(n, d)
	}
	return n
}

func NewAddAlertRequest(dto dtos.Alert) AddAlertRequest {
	return AddAlertRequest{
		BaseRequest:  dtoCommon.NewBaseRequest(),
		Alert: dto,
	}
}

// GetAlertRequest defines the Request Content for GET Alert DTO.
type GetAlertRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	QueryCondition        AlertQueryCondition `json:"queryCondition"`
}

type AlertQueryCondition struct {
	Category []string `json:"category,omitempty"`
	Start    int64    `json:"start,omitempty"`
	End      int64    `json:"end,omitempty"`
}

// Validate satisfies the Validator interface
func (request GetAlertRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the GetAlertRequest type
func (request *GetAlertRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		QueryCondition AlertQueryCondition
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*request = GetAlertRequest(alias)

	// validate GetAlertRequest DTO
	if err := request.Validate(); err != nil {
		return err
	}
	return nil
}
