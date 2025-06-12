//
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/requests"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// AuthClient defines the interface for interactions with the auth API endpoint on the security-proxy-auth service.
type AuthClient interface {
	// AddKey adds the JWT signing or verification key
	AddKey(ctx context.Context, req requests.AddKeyDataRequest) (common.BaseResponse, errors.IIOT)
	// VerificationKeyByIssuer returns the JWT verification key by the specified issuer
	VerificationKeyByIssuer(ctx context.Context, issuer string) (res responses.KeyDataResponse, err errors.IIOT)
}
