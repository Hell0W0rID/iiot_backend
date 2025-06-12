//
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"path"
	"strconv"

	"iiot-backend/pkg/go-mod-core-contracts/clients"
	"iiot-backend/pkg/go-mod-core-contracts/clients/http/utils"
	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
	"iiot-backend/pkg/go-mod-core-contracts/common"
	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/requests"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

type eventClient struct {
	baseUrlFunc           clients.ClientBaseUrlFunc
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewDataEventClient creates an instance of DataEventClient
func NewDataEventClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.DataEventClient {
	return &eventClient{
		baseUrlFunc:           clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// NewDataEventClientWithUrlCallback creates an instance of DataEventClient with ClientBaseUrlFunc.
func NewDataEventClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.DataEventClient {
	return &eventClient{
		baseUrlFunc:           baseUrlFunc,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

func (ec *eventClient) Add(ctx context.Context, serviceName string, req requests.AddDataEventRequest) (
	dtoCommon.BaseWithIdResponse, errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(ec.enableNameFieldEscape).
		SetPath(common.ApiDataEventRoute).SetNameFieldPath(serviceName).SetNameFieldPath(req.DataEvent.ProfileName).SetNameFieldPath(req.DataEvent.DeviceName).SetNameFieldPath(req.DataEvent.SourceName).BuildPath()
	var br dtoCommon.BaseWithIdResponse

	bytes, encoding, err := req.Encode()
	if err != nil {
		return br, errors.NewCommonIIOTWrapper(err)
	}
	baseUrl, err := clients.GetBaseUrl(ec.baseUrlFunc)
	if err != nil {
		return br, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.PostRequest(ctx, &br, baseUrl, requestPath, bytes, encoding, ec.authInjector)
	if err != nil {
		return br, errors.NewCommonIIOTWrapper(err)
	}
	return br, nil
}

func (ec *eventClient) AllDataEvents(ctx context.Context, offset, limit int) (responses.MultiDataEventsResponse, errors.IIOT) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiDataEventsResponse{}
	baseUrl, err := clients.GetBaseUrl(ec.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllDataEventRoute, requestParams, ec.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) DataEventCount(ctx context.Context) (dtoCommon.CountResponse, errors.IIOT) {
	res := dtoCommon.CountResponse{}
	baseUrl, err := clients.GetBaseUrl(ec.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiDataEventCountRoute, nil, ec.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) DataEventCountByDeviceName(ctx context.Context, name string) (dtoCommon.CountResponse, errors.IIOT) {
	requestPath := path.Join(common.ApiDataEventCountRoute, common.Device, common.Name, name)
	res := dtoCommon.CountResponse{}
	baseUrl, err := clients.GetBaseUrl(ec.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, nil, ec.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) DataEventsByDeviceName(ctx context.Context, name string, offset, limit int) (
	responses.MultiDataEventsResponse, errors.IIOT) {
	requestPath := path.Join(common.ApiDataEventRoute, common.Device, common.Name, name)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiDataEventsResponse{}
	baseUrl, err := clients.GetBaseUrl(ec.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, ec.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) DeleteByDeviceName(ctx context.Context, name string) (dtoCommon.BaseResponse, errors.IIOT) {
	requestPath := path.Join(common.ApiDataEventRoute, common.Device, common.Name, name)
	res := dtoCommon.BaseResponse{}
	baseUrl, err := clients.GetBaseUrl(ec.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.DeleteRequest(ctx, &res, baseUrl, requestPath, ec.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) DataEventsByTimeRange(ctx context.Context, start, end int64, offset, limit int) (
	responses.MultiDataEventsResponse, errors.IIOT) {
	requestPath := path.Join(common.ApiDataEventRoute, common.Start, strconv.FormatInt(start, 10), common.End, strconv.FormatInt(end, 10))
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiDataEventsResponse{}
	baseUrl, err := clients.GetBaseUrl(ec.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, ec.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) DeleteByAge(ctx context.Context, age int) (dtoCommon.BaseResponse, errors.IIOT) {
	requestPath := path.Join(common.ApiDataEventRoute, common.Age, strconv.Itoa(age))
	res := dtoCommon.BaseResponse{}
	baseUrl, err := clients.GetBaseUrl(ec.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.DeleteRequest(ctx, &res, baseUrl, requestPath, ec.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) DeleteById(ctx context.Context, id string) (dtoCommon.BaseResponse, errors.IIOT) {
	requestPath := path.Join(common.ApiDataEventRoute, common.Id, id)
	res := dtoCommon.BaseResponse{}
	baseUrl, err := clients.GetBaseUrl(ec.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.DeleteRequest(ctx, &res, baseUrl, requestPath, ec.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}
