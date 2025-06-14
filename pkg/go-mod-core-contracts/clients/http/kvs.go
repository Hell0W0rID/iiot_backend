//
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"strconv"

	"iiot-backend/pkg/go-mod-core-contracts/clients/http/utils"
	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
	"iiot-backend/pkg/go-mod-core-contracts/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/requests"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// KVSClient is the REST client for invoking the key-value APIs(/kvs/*) from Core Keeper
type KVSClient struct {
	baseUrl      string
	authInjector interfaces.AuthenticationInjector
}

// NewKVSClient creates an instance of KVSClient
func NewKVSClient(baseUrl string, authInjector interfaces.AuthenticationInjector) interfaces.KVSClient {
	return &KVSClient{
		baseUrl:      baseUrl,
		authInjector: authInjector,
	}
}

// UpdateValuesByKey updates values of the specified key and the child keys defined in the request payload.
// If no key exists at the given path, the key(s) will be created.
func (kc KVSClient) UpdateValuesByKey(ctx context.Context, key string, flatten bool, req requests.UpdateKeysRequest) (res responses.KeysResponse, err errors.IIOT) {
	path := utils.EscapeAndJoinPath(common.ApiKVSRoute, common.Key, key)
	queryParams := url.Values{}
	queryParams.Set(common.Flatten, strconv.FormatBool(flatten))
	err = utils.PutRequest(ctx, &res, kc.baseUrl, path, queryParams, req, kc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// ValuesByKey returns the values of the specified key prefix.
func (kc KVSClient) ValuesByKey(ctx context.Context, key string) (res responses.MultiKeyValueResponse, err errors.IIOT) {
	path := utils.EscapeAndJoinPath(common.ApiKVSRoute, common.Key, key)
	queryParams := url.Values{}
	queryParams.Set(common.Plaintext, common.ValueTrue)
	err = utils.GetRequest(ctx, &res, kc.baseUrl, path, queryParams, kc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// ListKeys returns the list of the keys with the specified key prefix.
func (kc KVSClient) ListKeys(ctx context.Context, key string) (res responses.KeysResponse, err errors.IIOT) {
	path := utils.EscapeAndJoinPath(common.ApiKVSRoute, common.Key, key)
	queryParams := url.Values{}
	queryParams.Set(common.KeyOnly, common.ValueTrue)
	err = utils.GetRequest(ctx, &res, kc.baseUrl, path, queryParams, kc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeleteKey deletes the specified key.
func (kc KVSClient) DeleteKey(ctx context.Context, key string) (res responses.KeysResponse, err errors.IIOT) {
	path := utils.EscapeAndJoinPath(common.ApiKVSRoute, common.Key, key)
	err = utils.DeleteRequest(ctx, &res, kc.baseUrl, path, kc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeleteKeysByPrefix deletes all keys with the specified prefix.
func (kc KVSClient) DeleteKeysByPrefix(ctx context.Context, key string) (res responses.KeysResponse, err errors.IIOT) {
	path := utils.EscapeAndJoinPath(common.ApiKVSRoute, common.Key, key)
	queryParams := url.Values{}
	queryParams.Set("prefixMatch", common.ValueTrue)
	err = utils.DeleteRequestWithParams(ctx, &res, kc.baseUrl, path, queryParams, kc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}
