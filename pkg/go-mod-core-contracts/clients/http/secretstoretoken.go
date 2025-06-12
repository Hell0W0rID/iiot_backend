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
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

type SecretStoreTokenClient struct {
	baseUrlFunc  clients.ClientBaseUrlFunc
	authInjector interfaces.AuthenticationInjector
}

// NewSecretStoreTokenClient creates an instance of SecretStoreTokenClient
func NewSecretStoreTokenClient(baseUrl string, authInjector interfaces.AuthenticationInjector) interfaces.SecretStoreTokenClient {
	return &SecretStoreTokenClient{
		baseUrlFunc:  clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector: authInjector,
	}
}

// NewSecretStoreTokenClientWithUrlCallback creates an instance of SecretStoreTokenClient with ClientBaseUrlFunc.
func NewSecretStoreTokenClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector) interfaces.AuthClient {
	return &AuthClient{
		baseUrlFunc:  baseUrlFunc,
		authInjector: authInjector,
	}
}

// RegenToken regenerates the secret store client token based on the specified entity id
func (ac *SecretStoreTokenClient) RegenToken(ctx context.Context, entityId string) (dtoCommon.BaseResponse, errors.IIOT) {
	var response dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(ac.baseUrlFunc)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}

	path := common.NewPathBuilder().SetPath(common.ApiTokenRoute).SetPath(common.EntityId).SetPath(entityId).BuildPath()
	err = utils.PutRequest(ctx, &response, baseUrl, path, nil, nil, ac.authInjector)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	return response, nil
}
