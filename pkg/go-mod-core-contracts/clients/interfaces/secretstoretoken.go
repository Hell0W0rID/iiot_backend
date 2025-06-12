//
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// SecretStoreTokenClient defines the interface for interactions with the API endpoint on the security-secretstore-setup service.
type SecretStoreTokenClient interface {
	// RegenToken regenerates the secret store client token based on the specified entity id
	RegenToken(ctx context.Context, entityId string) (dtoCommon.BaseResponse, errors.IIOT)
}
