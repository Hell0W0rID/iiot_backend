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

// AddKeyDataRequest defines the Request Content for POST Key DTO.
type AddKeyDataRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	KeyData               dtos.KeyData `json:"keyData"`
}

// Validate satisfies the Validator interface
func (a *AddKeyDataRequest) Validate() error {
	err := common.Validate(a)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddUserRequest type
func (a *AddKeyDataRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		KeyData dtos.KeyData
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*a = AddKeyDataRequest(alias)
	if err := a.Validate(); err != nil {
		return err
	}
	return nil
}
