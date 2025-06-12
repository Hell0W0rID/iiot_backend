//
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

package common

import (
	"encoding/json"

	"iiot-backend/pkg/go-mod-core-contracts/common"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// SecretDataKeyValue is a key/value pair to be stored in the Secret Store as part of the Secret Data
type SecretDataKeyValue struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// SecretRequest is the request DTO for storing supplied secret at a given SecretName in the Secret Store
type SecretRequest struct {
	BaseRequest `json:",inline"`
	SecretName  string               `json:"secretName" validate:"required"`
	SecretData  []SecretDataKeyValue `json:"secretData" validate:"required,gt=0,dive"`
}

// Validate satisfies the Validator interface
func (sr *SecretRequest) Validate() error {
	err := common.Validate(sr)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the SecretRequest type
func (sr *SecretRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		BaseRequest
		SecretName string
		SecretData []SecretDataKeyValue
	}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal SecretRequest body as JSON.", err)
	}

	*sr = SecretRequest(alias)

	// validate SecretRequest DTO
	if err := sr.Validate(); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "SecretRequest validation failed.", err)
	}
	return nil
}

func NewSecretRequest(secretName string, secretData []SecretDataKeyValue) SecretRequest {
	return SecretRequest{
		BaseRequest: NewBaseRequest(),
		SecretName:  secretName,
		SecretData:  secretData,
	}
}
