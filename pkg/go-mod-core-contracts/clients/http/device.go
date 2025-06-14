//
// Copyright (C) 2020-2021 Unknown author
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

type DeviceClient struct {
	baseUrlFunc           clients.ClientBaseUrlFunc
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewDeviceClient creates an instance of DeviceClient
func NewDeviceClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.DeviceClient {
	return &DeviceClient{
		baseUrlFunc:           clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// NewDeviceClientWithUrlCallback creates an instance of DeviceClient with ClientBaseUrlFunc.
func NewDeviceClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.DeviceClient {
	return &DeviceClient{
		baseUrlFunc:           baseUrlFunc,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

func (dc DeviceClient) Add(ctx context.Context, reqs []requests.AddDeviceRequest) (res []dtoCommon.BaseWithIdResponse, err errors.IIOT) {
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiDeviceRoute, nil, reqs, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) AddWithQueryParams(ctx context.Context, reqs []requests.AddDeviceRequest, queryParams map[string]string) (res []dtoCommon.BaseWithIdResponse, err errors.IIOT) {
	requestParams := url.Values{}
	for k, v := range queryParams {
		requestParams.Set(k, v)
	}
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiDeviceRoute, requestParams, reqs, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) Update(ctx context.Context, reqs []requests.UpdateDeviceRequest) (res []dtoCommon.BaseResponse, err errors.IIOT) {
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.PatchRequest(ctx, &res, baseUrl, common.ApiDeviceRoute, nil, reqs, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) UpdateWithQueryParams(ctx context.Context, reqs []requests.UpdateDeviceRequest, queryParams map[string]string) (res []dtoCommon.BaseResponse, err errors.IIOT) {
	requestParams := url.Values{}
	for k, v := range queryParams {
		requestParams.Set(k, v)
	}
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.PatchRequest(ctx, &res, baseUrl, common.ApiDeviceRoute, requestParams, reqs, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) AllDevices(ctx context.Context, labels []string, offset int, limit int) (res responses.MultiDevicesResponse, err errors.IIOT) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(common.Labels, strings.Join(labels, common.CommaSeparator))
	}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllDeviceRoute, requestParams, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) AllDevicesWithChildren(ctx context.Context, parent string, maxLevels uint, labels []string, offset int, limit int) (res responses.MultiDevicesResponse, err errors.IIOT) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(common.Labels, strings.Join(labels, common.CommaSeparator))
	}
	requestParams.Set(common.DescendantsOf, parent)
	requestParams.Set(common.MaxLevels, strconv.FormatUint(uint64(maxLevels), 10))
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllDeviceRoute, requestParams, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) DeviceNameExists(ctx context.Context, name string) (res dtoCommon.BaseResponse, err errors.IIOT) {
	path := common.NewPathBuilder().EnableNameFieldEscape(dc.enableNameFieldEscape).
		SetPath(common.ApiDeviceRoute).SetPath(common.Check).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, path, nil, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) DeviceByName(ctx context.Context, name string) (res responses.DeviceResponse, err errors.IIOT) {
	path := common.NewPathBuilder().EnableNameFieldEscape(dc.enableNameFieldEscape).
		SetPath(common.ApiDeviceRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, path, nil, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) DeleteDeviceByName(ctx context.Context, name string) (res dtoCommon.BaseResponse, err errors.IIOT) {
	path := common.NewPathBuilder().EnableNameFieldEscape(dc.enableNameFieldEscape).
		SetPath(common.ApiDeviceRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.DeleteRequest(ctx, &res, baseUrl, path, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) DevicesByProfileName(ctx context.Context, name string, offset int, limit int) (res responses.MultiDevicesResponse, err errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(dc.enableNameFieldEscape).
		SetPath(common.ApiDeviceRoute).SetPath(common.Profile).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) DevicesByServiceName(ctx context.Context, name string, offset int, limit int) (res responses.MultiDevicesResponse, err errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(dc.enableNameFieldEscape).
		SetPath(common.ApiDeviceRoute).SetPath(common.Service).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}
