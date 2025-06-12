//
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"iiot-backend/pkg/go-mod-core-contracts/clients"
	"iiot-backend/pkg/go-mod-core-contracts/clients/http/utils"
	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
	"iiot-backend/pkg/go-mod-core-contracts/common"
	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/requests"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

type DeviceWatcherClient struct {
	baseUrlFunc           clients.ClientBaseUrlFunc
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewDeviceWatcherClient creates an instance of DeviceWatcherClient
func NewDeviceWatcherClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.DeviceWatcherClient {
	return &DeviceWatcherClient{
		baseUrlFunc:           clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// NewDeviceWatcherClientWithUrlCallback creates an instance of DeviceWatcherClient with ClientBaseUrlFunc.
func NewDeviceWatcherClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.DeviceWatcherClient {
	return &DeviceWatcherClient{
		baseUrlFunc:           baseUrlFunc,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

func (pwc DeviceWatcherClient) Add(ctx context.Context, reqs []requests.AddDeviceWatcherRequest) (res []dtoCommon.BaseWithIdResponse, err errors.IIOT) {
	baseUrl, goErr := clients.GetBaseUrl(pwc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiDeviceWatcherRoute, nil, reqs, pwc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}

	return
}

func (pwc DeviceWatcherClient) Update(ctx context.Context, reqs []requests.UpdateDeviceWatcherRequest) (res []dtoCommon.BaseResponse, err errors.IIOT) {
	baseUrl, goErr := clients.GetBaseUrl(pwc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.PatchRequest(ctx, &res, baseUrl, common.ApiDeviceWatcherRoute, nil, reqs, pwc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}

	return
}

func (pwc DeviceWatcherClient) AllDeviceWatchers(ctx context.Context, labels []string, offset int, limit int) (res responses.MultiDeviceWatchersResponse, err errors.IIOT) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(common.Labels, strings.Join(labels, common.CommaSeparator))
	}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, goErr := clients.GetBaseUrl(pwc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllDeviceWatcherRoute, requestParams, pwc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}

	return
}

func (pwc DeviceWatcherClient) DeviceWatcherByName(ctx context.Context, name string) (res responses.DeviceWatcherResponse, err errors.IIOT) {
	path := common.NewPathBuilder().EnableNameFieldEscape(pwc.enableNameFieldEscape).
		SetPath(common.ApiDeviceWatcherRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, goErr := clients.GetBaseUrl(pwc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, path, nil, pwc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}

	return
}

func (pwc DeviceWatcherClient) DeleteDeviceWatcherByName(ctx context.Context, name string) (res dtoCommon.BaseResponse, err errors.IIOT) {
	path := common.NewPathBuilder().EnableNameFieldEscape(pwc.enableNameFieldEscape).
		SetPath(common.ApiDeviceWatcherRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, goErr := clients.GetBaseUrl(pwc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.DeleteRequest(ctx, &res, baseUrl, path, pwc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}

	return
}

func (pwc DeviceWatcherClient) DeviceWatchersByProfileName(ctx context.Context, name string, offset int, limit int) (res responses.MultiDeviceWatchersResponse, err errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(pwc.enableNameFieldEscape).
		SetPath(common.ApiDeviceWatcherRoute).SetPath(common.Profile).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, goErr := clients.GetBaseUrl(pwc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, pwc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}

	return
}

func (pwc DeviceWatcherClient) DeviceWatchersByServiceName(ctx context.Context, name string, offset int, limit int) (res responses.MultiDeviceWatchersResponse, err errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(pwc.enableNameFieldEscape).
		SetPath(common.ApiDeviceWatcherRoute).SetPath(common.Service).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, goErr := clients.GetBaseUrl(pwc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, pwc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}

	return
}
