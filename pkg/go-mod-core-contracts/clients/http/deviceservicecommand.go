//
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"encoding/json"
	"net/url"

	"iiot-backend/pkg/go-mod-core-contracts/clients/http/utils"
	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
	"iiot-backend/pkg/go-mod-core-contracts/common"
	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/requests"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"

	"github.com/fxamacker/cbor/v2"
)

type deviceServiceCommandClient struct {
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewDeviceHandlerCommandClient creates an instance of deviceServiceCommandClient
func NewDeviceHandlerCommandClient(authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.DeviceHandlerCommandClient {
	return &deviceServiceCommandClient{
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// GetCommand sends HTTP request to execute the Get command
func (client *deviceServiceCommandClient) GetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string) (*responses.DataEventResponse, errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiDeviceRoute).SetPath(common.Name).SetNameFieldPath(deviceName).SetNameFieldPath(commandName).BuildPath()
	params, err := url.ParseQuery(queryParams)
	if err != nil {
		return nil, errors.NewCommonIIOTWrapper(err)
	}
	res, contentType, edgeXerr := utils.GetRequestAndReturnBinaryRes(ctx, baseUrl, requestPath, params, client.authInjector)
	if edgeXerr != nil {
		return nil, errors.NewCommonIIOTWrapper(edgeXerr)
	}
	// If execute GetCommand with dsReturnDataEvent query parameter 'no', there will be no content returned in the http response.
	// So we can use the nil pointer to indicate that the HTTP response content is empty
	if len(res) == 0 {
		return nil, nil
	}
	response := &responses.DataEventResponse{}
	if contentType == common.ContentTypeCBOR {
		if err = cbor.Unmarshal(res, response); err != nil {
			return nil, errors.NewCommonIIOT(errors.KindContractInvalid, "failed to decode the cbor response", err)
		}
	} else {
		if err = json.Unmarshal(res, response); err != nil {
			return nil, errors.NewCommonIIOT(errors.KindContractInvalid, "failed to decode the json response", err)
		}
	}
	return response, nil
}

// SetCommand sends HTTP request to execute the Set command
func (client *deviceServiceCommandClient) SetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string, settings map[string]string) (dtoCommon.BaseResponse, errors.IIOT) {
	var response dtoCommon.BaseResponse
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiDeviceRoute).SetPath(common.Name).SetNameFieldPath(deviceName).SetNameFieldPath(commandName).BuildPath()
	params, err := url.ParseQuery(queryParams)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.PutRequest(ctx, &response, baseUrl, requestPath, params, settings, client.authInjector)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	return response, nil
}

// SetCommandWithObject invokes device service's set command API and the settings supports object value type
func (client *deviceServiceCommandClient) SetCommandWithObject(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string, settings map[string]interface{}) (dtoCommon.BaseResponse, errors.IIOT) {
	var response dtoCommon.BaseResponse
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiDeviceRoute).SetPath(common.Name).SetNameFieldPath(deviceName).SetNameFieldPath(commandName).BuildPath()
	params, err := url.ParseQuery(queryParams)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.PutRequest(ctx, &response, baseUrl, requestPath, params, settings, client.authInjector)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	return response, nil
}

func (client *deviceServiceCommandClient) Discovery(ctx context.Context, baseUrl string) (dtoCommon.BaseResponse, errors.IIOT) {
	var response dtoCommon.BaseResponse
	err := utils.PostRequest(ctx, &response, baseUrl, common.ApiDiscoveryRoute, nil, "", client.authInjector)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	return response, nil
}

// ProfileScan sends an HTTP POST request to the device service's profile scan API endpoint.
func (client *deviceServiceCommandClient) ProfileScan(ctx context.Context, baseUrl string, req requests.ProfileScanRequest) (dtoCommon.BaseResponse, errors.IIOT) {
	var response dtoCommon.BaseResponse
	err := utils.PostRequestWithRawData(ctx, &response, baseUrl, common.ApiProfileScanRoute, nil, req, client.authInjector)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	return response, nil
}

func (client *deviceServiceCommandClient) StopDeviceDiscovery(ctx context.Context, baseUrl string, requestId string, queryParams map[string]string) (dtoCommon.BaseResponse, errors.IIOT) {
	requestPath := common.ApiDiscoveryRoute
	if len(requestId) != 0 {
		requestPath = common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
			SetPath(common.ApiDiscoveryRoute).SetPath(common.RequestId).SetNameFieldPath(requestId).BuildPath()
	}
	response := dtoCommon.BaseResponse{}
	params := url.Values{}
	for k, v := range queryParams {
		params.Set(k, v)
	}
	err := utils.DeleteRequestWithParams(ctx, &response, baseUrl, requestPath, params, client.authInjector)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	return response, nil
}

func (client *deviceServiceCommandClient) StopProfileScan(ctx context.Context, baseUrl string, deviceName string, queryParams map[string]string) (dtoCommon.BaseResponse, errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiProfileScanRoute).SetPath(common.Device).SetPath(common.Name).SetNameFieldPath(deviceName).BuildPath()
	response := dtoCommon.BaseResponse{}
	params := url.Values{}
	for k, v := range queryParams {
		params.Set(k, v)
	}
	err := utils.DeleteRequestWithParams(ctx, &response, baseUrl, requestPath, params, client.authInjector)
	if err != nil {
		return response, errors.NewCommonIIOTWrapper(err)
	}
	return response, nil
}
