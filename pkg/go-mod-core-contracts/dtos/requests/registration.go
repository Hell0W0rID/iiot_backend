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
)

// AddRegistrationRequest defines the Request Content for POST Registration DTO.
type AddRegistrationRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	Registration          dtos.Registration `json:"registration"`
}

// Validate satisfies the Validator interface
func (r *AddRegistrationRequest) Validate() error {
	err := common.Validate(r)
	if err != nil {
		return errors.NewCommonIIOTWrapper(err)
	}
	err = r.Registration.Validate()
	if err != nil {
		return errors.NewCommonIIOTWrapper(err)
	}
	return nil
}

// UnmarshalJSON implements the Unmarshaler interface for the AddRegistrationRequest type
func (r *AddRegistrationRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		Registration dtos.Registration
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*r = AddRegistrationRequest(alias)

	// validate AddRegistrationRequest DTO
	if err := r.Validate(); err != nil {
		return errors.NewCommonIIOTWrapper(err)
	}
	return nil
}
