/*******************************************************************************
 *******************************************************************************/

package secret

import (
	"encoding/json"
	"fmt"

	validation "iiot-backend/pkg/go-mod-core-contracts/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"github.com/hashicorp/go-multierror"
)

// ServiceSecrets contains the list of secrets to import into a service's SecretStore
type ServiceSecrets struct {
	Secrets []ServiceSecret `json:"secrets" validate:"required,gt=0,dive"`
}

// ServiceSecret contains the information about a service's secret to import into a service's SecretStore
type ServiceSecret struct {
	SecretName string                      `json:"secretName" validate:"iiot-dto-none-empty-string"`
	Imported   bool                        `json:"imported"`
	SecretData []common.SecretDataKeyValue `json:"secretData" validate:"required,dive"`
}

// MarshalJson marshal the service's secrets to JSON.
func (s *ServiceSecrets) MarshalJson() ([]byte, error) {
	return json.Marshal(s)
}

// UnmarshalServiceSecretsJson un-marshals the JSON containing the services list of secrets
func UnmarshalServiceSecretsJson(data []byte) (*ServiceSecrets, error) {
	secrets := &ServiceSecrets{}

	if err := json.Unmarshal(data, secrets); err != nil {
		return nil, err
	}

	if err := validation.Validate(secrets); err != nil {
		return nil, err
	}

	var validationErrs error

	// Since secretData len validation can't be specified to only validate when Imported=false, we have to do it manually here
	for _, secret := range secrets.Secrets {
		if !secret.Imported && len(secret.SecretData) == 0 {
			validationErrs = multierror.Append(validationErrs, fmt.Errorf("SecretData for '%s' must not be empty when Imported=false", secret.SecretName))
		}
	}

	if validationErrs != nil {
		return nil, validationErrs
	}

	return secrets, nil
}
