//
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"

	"iiot-backend/pkg/go-mod-core-contracts/clients"
	"iiot-backend/pkg/go-mod-core-contracts/clients/http/utils"
	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
	"iiot-backend/pkg/go-mod-core-contracts/common"
	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/requests"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

type AuthClient struct {
	baseUrlFunc  clients.ClientBaseUrlFunc
	authInjector interfaces.AuthenticationInjector
}

// NewAuthClient creates an instance of AuthClient
func NewAuthClient(baseUrl string, authInjector interfaces.AuthenticationInjector) interfaces.AuthClient {
	return &AuthClient{
		baseUrlFunc:  clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector: authInjector,
	}
}

// NewAuthClientWithUrlCallback creates an instance of AuthClient with ClientBaseUrlFunc.
func NewAuthClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector) interfaces.AuthClient {
	return &AuthClient{
		baseUrlFunc:  baseUrlFunc,
		authInjector: authInjector,
	}
}

// AddKey adds new key
func (ac *AuthClient) AddKey(ctx context.Context, req requests.AddKeyDataRequest) (dtoCommon.BaseResponse, errors.IIOT) {
	var response dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(ac.baseUrlFunc)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.PostRequestWithRawData(ctx, &response, baseUrl, common.ApiKeyRoute, nil, req, ac.authInjector)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	return response, nil
}

func (ac *AuthClient) VerificationKeyByIssuer(ctx context.Context, issuer string) (res responses.KeyDataResponse, err errors.IIOT) {
	path := common.NewPathBuilder().SetPath(common.ApiKeyRoute).SetPath(common.VerificationKeyType).SetPath(common.Issuer).SetNameFieldPath(issuer).BuildPath()
	baseUrl, goErr := clients.GetBaseUrl(ac.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, path, nil, ac.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}
