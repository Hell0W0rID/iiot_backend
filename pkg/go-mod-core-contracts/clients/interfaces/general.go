//
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

type GeneralClient interface {
	// FetchConfiguration obtains configuration information from the target service.
	FetchConfiguration(ctx context.Context) (common.ConfigResponse, errors.IIOT)
}
