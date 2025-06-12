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

type commonClient struct {
	baseUrl      string
	authInjector interfaces.AuthenticationInjector
}

// NewCommonClient creates an instance of CommonClient
func NewCommonClient(baseUrl string, authInjector interfaces.AuthenticationInjector) interfaces.CommonClient {
	return &commonClient{
		baseUrl:      baseUrl,
		authInjector: authInjector,
	}
}

func (cc *commonClient) Configuration(ctx context.Context) (dtoCommon.ConfigResponse, errors.IIOT) {
	cr := dtoCommon.ConfigResponse{}
	err := utils.GetRequest(ctx, &cr, cc.baseUrl, common.ApiConfigRoute, nil, cc.authInjector)
	if err != nil {
		return cr, errors.NewCommonIIOTWrapper(err)
	}
	return cr, nil
}

func (cc *commonClient) Ping(ctx context.Context) (dtoCommon.PingResponse, errors.IIOT) {
	pr := dtoCommon.PingResponse{}
	err := utils.GetRequest(ctx, &pr, cc.baseUrl, common.ApiPingRoute, nil, cc.authInjector)
	if err != nil {
		return pr, errors.NewCommonIIOTWrapper(err)
	}
	return pr, nil
}

func (cc *commonClient) Version(ctx context.Context) (dtoCommon.VersionResponse, errors.IIOT) {
	vr := dtoCommon.VersionResponse{}
	err := utils.GetRequest(ctx, &vr, cc.baseUrl, common.ApiVersionRoute, nil, cc.authInjector)
	if err != nil {
		return vr, errors.NewCommonIIOTWrapper(err)
	}
	return vr, nil
}

func (cc *commonClient) AddSecret(ctx context.Context, request dtoCommon.SecretRequest) (res dtoCommon.BaseResponse, err errors.IIOT) {
	err = utils.PostRequestWithRawData(ctx, &res, cc.baseUrl, common.ApiSecretRoute, nil, request, cc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}
