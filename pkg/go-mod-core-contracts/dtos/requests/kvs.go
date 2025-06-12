//
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
	"iiot-backend/pkg/go-mod-core-contracts/models"
)

// UpdateKeysRequest defines the Request Content for PUT Key DTO.
type UpdateKeysRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	Value                 any `json:"value,omitempty"`
}

// Validate checks if the fields are valid of the UpdateKeysRequest struct
func (u UpdateKeysRequest) Validate() errors.IIOT {
	// check if Value field is nil
	if u.Value == nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "the value field is undefined", nil)
	}
	// check if Value field is an empty map
	if v, ok := u.Value.(map[string]any); ok {
		if len(v) == 0 {
			return errors.NewCommonIIOT(errors.KindContractInvalid, "the value field is an empty object", nil)
		}
	}

	return nil
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateKeysRequest type
func (u *UpdateKeysRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		Value any
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*u = UpdateKeysRequest(alias)

	if err := u.Validate(); err != nil {
		return err
	}

	return nil
}

// UpdateKeysReqToKVModels transforms the UpdateKeysRequest DTO to the KV model
func UpdateKeysReqToKVModels(req UpdateKeysRequest, key string) models.KVS {
	var kv models.KVS
	kv.Value = req.Value
	kv.Key = key

	return kv
}
