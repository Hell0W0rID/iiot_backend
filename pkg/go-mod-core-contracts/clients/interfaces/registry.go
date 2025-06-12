//
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"iiot-backend/pkg/go-mod-core-contracts/dtos/requests"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// RegistryClient defines the interface for interactions with the registry endpoint on the IIOT core-keeper service.
type RegistryClient interface {
	Register(context.Context, requests.AddRegistrationRequest) errors.IIOT
	UpdateRegister(context.Context, requests.AddRegistrationRequest) errors.IIOT
	RegistrationByServiceId(context.Context, string) (responses.RegistrationResponse, errors.IIOT)
	AllRegistry(context.Context, bool) (responses.MultiRegistrationsResponse, errors.IIOT)
	Deregister(context.Context, string) errors.IIOT
}
