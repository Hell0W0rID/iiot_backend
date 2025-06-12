//
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"

	"iiot-backend/pkg/go-mod-core-contracts/clients/http/utils"
	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
	"iiot-backend/pkg/go-mod-core-contracts/common"
	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

type generalClient struct {
	baseUrl      string
	authInjector interfaces.AuthenticationInjector
}

func NewGeneralClient(baseUrl string, authInjector interfaces.AuthenticationInjector) interfaces.GeneralClient {
	return &generalClient{
		baseUrl:      baseUrl,
		authInjector: authInjector,
	}
}

func (g *generalClient) FetchConfiguration(ctx context.Context) (res dtoCommon.ConfigResponse, err errors.IIOT) {
	err = utils.GetRequest(ctx, &res, g.baseUrl, common.ApiConfigRoute, nil, g.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}

	return res, nil
}
