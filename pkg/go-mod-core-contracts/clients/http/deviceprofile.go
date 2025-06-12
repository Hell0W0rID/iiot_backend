//
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"strconv"
	"strings"
	"sync"

	"iiot-backend/pkg/go-mod-core-contracts/clients"
	"iiot-backend/pkg/go-mod-core-contracts/clients/http/utils"
	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
	"iiot-backend/pkg/go-mod-core-contracts/common"
	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/requests"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

type DeviceTemplateClient struct {
	baseUrlFunc           clients.ClientBaseUrlFunc
	authInjector          interfaces.AuthenticationInjector
	resourcesCache        map[string]responses.DeviceResourceResponse
	mux                   sync.RWMutex
	enableNameFieldEscape bool
}

// NewDeviceTemplateClient creates an instance of DeviceTemplateClient
func NewDeviceTemplateClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.DeviceTemplateClient {
	return &DeviceTemplateClient{
		baseUrlFunc:           clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector:          authInjector,
		resourcesCache:        make(map[string]responses.DeviceResourceResponse),
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// NewDeviceTemplateClientWithUrlCallback creates an instance of DeviceTemplateClient with ClientBaseUrlFunc.
func NewDeviceTemplateClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.DeviceTemplateClient {
	return &DeviceTemplateClient{
		baseUrlFunc:           baseUrlFunc,
		authInjector:          authInjector,
		resourcesCache:        make(map[string]responses.DeviceResourceResponse),
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// Add adds new device profile
func (client *DeviceTemplateClient) Add(ctx context.Context, reqs []requests.DeviceTemplateRequest) ([]dtoCommon.BaseWithIdResponse, errors.IIOT) {
	var res []dtoCommon.BaseWithIdResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiDeviceTemplateRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// Update updates device profile
func (client *DeviceTemplateClient) Update(ctx context.Context, reqs []requests.DeviceTemplateRequest) ([]dtoCommon.BaseResponse, errors.IIOT) {
	var res []dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.PutRequest(ctx, &res, baseUrl, common.ApiDeviceTemplateRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// AddByYaml adds new device profile by uploading a yaml file
func (client *DeviceTemplateClient) AddByYaml(ctx context.Context, yamlFilePath string) (dtoCommon.BaseWithIdResponse, errors.IIOT) {
	var res dtoCommon.BaseWithIdResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.PostByFileRequest(ctx, &res, baseUrl, common.ApiDeviceTemplateUploadFileRoute, yamlFilePath, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// UpdateByYaml updates device profile by uploading a yaml file
func (client *DeviceTemplateClient) UpdateByYaml(ctx context.Context, yamlFilePath string) (dtoCommon.BaseResponse, errors.IIOT) {
	var res dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.PutByFileRequest(ctx, &res, baseUrl, common.ApiDeviceTemplateUploadFileRoute, yamlFilePath, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeleteByName deletes the device profile by name
func (client *DeviceTemplateClient) DeleteByName(ctx context.Context, name string) (dtoCommon.BaseResponse, errors.IIOT) {
	var response dtoCommon.BaseResponse
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiDeviceTemplateRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.DeleteRequest(ctx, &response, baseUrl, requestPath, client.authInjector)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	return response, nil
}

// DeviceTemplateByName queries the device profile by name
func (client *DeviceTemplateClient) DeviceTemplateByName(ctx context.Context, name string) (res responses.DeviceTemplateResponse, iiotError errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiDeviceTemplateRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, nil, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// AllDeviceTemplates queries the device profiles with offset, and limit
func (client *DeviceTemplateClient) AllDeviceTemplates(ctx context.Context, labels []string, offset int, limit int) (res responses.MultiDeviceTemplatesResponse, iiotError errors.IIOT) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(common.Labels, strings.Join(labels, common.CommaSeparator))
	}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllDeviceTemplateRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// AllDeviceTemplateBasicInfos queries the device profile basic infos with offset, and limit
func (client *DeviceTemplateClient) AllDeviceTemplateBasicInfos(ctx context.Context, labels []string, offset int, limit int) (res responses.MultiDeviceTemplateBasicInfoResponse, iiotError errors.IIOT) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(common.Labels, strings.Join(labels, common.CommaSeparator))
	}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllDeviceTemplateBasicInfoRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeviceTemplatesByModel queries the device profiles with offset, limit and model
func (client *DeviceTemplateClient) DeviceTemplatesByModel(ctx context.Context, model string, offset int, limit int) (res responses.MultiDeviceTemplatesResponse, iiotError errors.IIOT) {
	requestPath := path.Join(common.ApiDeviceTemplateRoute, common.Model, model)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeviceTemplatesByManufacturer queries the device profiles with offset, limit and manufacturer
func (client *DeviceTemplateClient) DeviceTemplatesByManufacturer(ctx context.Context, manufacturer string, offset int, limit int) (res responses.MultiDeviceTemplatesResponse, iiotError errors.IIOT) {
	requestPath := path.Join(common.ApiDeviceTemplateRoute, common.Manufacturer, manufacturer)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeviceTemplatesByManufacturerAndModel queries the device profiles with offset, limit, manufacturer and model
func (client *DeviceTemplateClient) DeviceTemplatesByManufacturerAndModel(ctx context.Context, manufacturer string, model string, offset int, limit int) (res responses.MultiDeviceTemplatesResponse, iiotError errors.IIOT) {
	requestPath := path.Join(common.ApiDeviceTemplateRoute, common.Manufacturer, manufacturer, common.Model, model)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeviceResourceByProfileNameAndResourceName queries the device resource by profileName and resourceName
func (client *DeviceTemplateClient) DeviceResourceByProfileNameAndResourceName(ctx context.Context, profileName string, resourceName string) (res responses.DeviceResourceResponse, iiotError errors.IIOT) {
	resourceMapKey := fmt.Sprintf("%s:%s", profileName, resourceName)
	res, exists := client.resourceByMapKey(resourceMapKey)
	if exists {
		return res, nil
	}
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiDeviceResourceRoute).SetPath(common.Profile).SetNameFieldPath(profileName).SetPath(common.Resource).SetNameFieldPath(resourceName).BuildPath()
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, nil, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	client.setResourceWithMapKey(res, resourceMapKey)
	return res, nil
}

func (client *DeviceTemplateClient) resourceByMapKey(key string) (res responses.DeviceResourceResponse, exists bool) {
	client.mux.RLock()
	defer client.mux.RUnlock()
	res, exists = client.resourcesCache[key]
	return
}

func (client *DeviceTemplateClient) setResourceWithMapKey(res responses.DeviceResourceResponse, key string) {
	client.mux.Lock()
	defer client.mux.Unlock()
	client.resourcesCache[key] = res
}

func (client *DeviceTemplateClient) CleanResourcesCache() {
	client.mux.Lock()
	defer client.mux.Unlock()
	client.resourcesCache = make(map[string]responses.DeviceResourceResponse)
}

// UpdateDeviceTemplateBasicInfo updates existing profile's basic info
func (client *DeviceTemplateClient) UpdateDeviceTemplateBasicInfo(ctx context.Context, reqs []requests.DeviceTemplateBasicInfoRequest) ([]dtoCommon.BaseResponse, errors.IIOT) {
	var res []dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.PatchRequest(ctx, &res, baseUrl, common.ApiDeviceTemplateBasicInfoRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// AddDeviceTemplateResource adds new device resource to an existing profile
func (client *DeviceTemplateClient) AddDeviceTemplateResource(ctx context.Context, reqs []requests.AddDeviceResourceRequest) ([]dtoCommon.BaseResponse, errors.IIOT) {
	var res []dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiDeviceTemplateResourceRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// UpdateDeviceTemplateResource updates existing device resource
func (client *DeviceTemplateClient) UpdateDeviceTemplateResource(ctx context.Context, reqs []requests.UpdateDeviceResourceRequest) ([]dtoCommon.BaseResponse, errors.IIOT) {
	var res []dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.PatchRequest(ctx, &res, baseUrl, common.ApiDeviceTemplateResourceRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeleteDeviceResourceByName deletes device resource by name
func (client *DeviceTemplateClient) DeleteDeviceResourceByName(ctx context.Context, profileName string, resourceName string) (dtoCommon.BaseResponse, errors.IIOT) {
	var response dtoCommon.BaseResponse
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiDeviceTemplateRoute).SetPath(common.Name).SetNameFieldPath(profileName).SetPath(common.Resource).SetNameFieldPath(resourceName).BuildPath()
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.DeleteRequest(ctx, &response, baseUrl, requestPath, client.authInjector)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	return response, nil
}

// AddDeviceTemplateDeviceAction adds new device command to an existing profile
func (client *DeviceTemplateClient) AddDeviceTemplateDeviceAction(ctx context.Context, reqs []requests.AddDeviceActionRequest) ([]dtoCommon.BaseResponse, errors.IIOT) {
	var res []dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiDeviceTemplateDeviceActionRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// UpdateDeviceTemplateDeviceAction updates existing device command
func (client *DeviceTemplateClient) UpdateDeviceTemplateDeviceAction(ctx context.Context, reqs []requests.UpdateDeviceActionRequest) ([]dtoCommon.BaseResponse, errors.IIOT) {
	var res []dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.PatchRequest(ctx, &res, baseUrl, common.ApiDeviceTemplateDeviceActionRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeleteDeviceActionByName deletes device command by name
func (client *DeviceTemplateClient) DeleteDeviceActionByName(ctx context.Context, profileName string, commandName string) (dtoCommon.BaseResponse, errors.IIOT) {
	var response dtoCommon.BaseResponse
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiDeviceTemplateRoute).SetPath(common.Name).SetNameFieldPath(profileName).SetPath(common.DeviceAction).SetNameFieldPath(commandName).BuildPath()
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.DeleteRequest(ctx, &response, baseUrl, requestPath, client.authInjector)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	return response, nil
}
