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
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

type readingClient struct {
	baseUrlFunc           clients.ClientBaseUrlFunc
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewMeasurementClient creates an instance of MeasurementClient
func NewMeasurementClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.MeasurementClient {
	return &readingClient{
		baseUrlFunc:           clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// NewMeasurementClientWithUrlCallback creates an instance of MeasurementClient with ClientBaseUrlFunc.
func NewMeasurementClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.MeasurementClient {
	return &readingClient{
		baseUrlFunc:           baseUrlFunc,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

func (rc readingClient) AllMeasurements(ctx context.Context, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiMeasurementsResponse{}
	baseUrl, err := clients.GetBaseUrl(rc.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllMeasurementRoute, requestParams, rc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (rc readingClient) MeasurementCount(ctx context.Context) (dtoCommon.CountResponse, errors.IIOT) {
	res := dtoCommon.CountResponse{}
	baseUrl, err := clients.GetBaseUrl(rc.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiMeasurementCountRoute, nil, rc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (rc readingClient) MeasurementCountByDeviceName(ctx context.Context, name string) (dtoCommon.CountResponse, errors.IIOT) {
	requestPath := path.Join(common.ApiMeasurementCountRoute, common.Device, common.Name, name)
	res := dtoCommon.CountResponse{}
	baseUrl, err := clients.GetBaseUrl(rc.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, nil, rc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (rc readingClient) MeasurementsByDeviceName(ctx context.Context, name string, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT) {
	requestPath := path.Join(common.ApiMeasurementRoute, common.Device, common.Name, name)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiMeasurementsResponse{}
	baseUrl, err := clients.GetBaseUrl(rc.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, rc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (rc readingClient) MeasurementsByResourceName(ctx context.Context, name string, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(rc.enableNameFieldEscape).
		SetPath(common.ApiMeasurementRoute).SetPath(common.ResourceName).SetNameFieldPath(name).BuildPath()
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiMeasurementsResponse{}
	baseUrl, err := clients.GetBaseUrl(rc.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, rc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (rc readingClient) MeasurementsByTimeRange(ctx context.Context, start, end int64, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT) {
	requestPath := path.Join(common.ApiMeasurementRoute, common.Start, strconv.FormatInt(start, 10), common.End, strconv.FormatInt(end, 10))
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiMeasurementsResponse{}
	baseUrl, err := clients.GetBaseUrl(rc.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, rc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// MeasurementsByResourceNameAndTimeRange returns readings by resource name and specified time range. Measurements are sorted in descending order of origin time.
func (rc readingClient) MeasurementsByResourceNameAndTimeRange(ctx context.Context, name string, start, end int64, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(rc.enableNameFieldEscape).
		SetPath(common.ApiMeasurementRoute).SetPath(common.ResourceName).SetNameFieldPath(name).SetPath(common.Start).SetPath(strconv.FormatInt(start, 10)).SetPath(common.End).SetPath(strconv.FormatInt(end, 10)).BuildPath()
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiMeasurementsResponse{}
	baseUrl, err := clients.GetBaseUrl(rc.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, rc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (rc readingClient) MeasurementsByDeviceNameAndResourceName(ctx context.Context, deviceName, resourceName string, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(rc.enableNameFieldEscape).
		SetPath(common.ApiMeasurementRoute).SetPath(common.Device).SetPath(common.Name).SetNameFieldPath(deviceName).SetPath(common.ResourceName).SetNameFieldPath(resourceName).BuildPath()
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiMeasurementsResponse{}
	baseUrl, err := clients.GetBaseUrl(rc.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, rc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil

}

func (rc readingClient) MeasurementsByDeviceNameAndResourceNameAndTimeRange(ctx context.Context, deviceName, resourceName string, start, end int64, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(rc.enableNameFieldEscape).
		SetPath(common.ApiMeasurementRoute).SetPath(common.Device).SetPath(common.Name).SetNameFieldPath(deviceName).SetPath(common.ResourceName).SetNameFieldPath(resourceName).
		SetPath(common.Start).SetPath(strconv.FormatInt(start, 10)).SetPath(common.End).SetPath(strconv.FormatInt(end, 10)).BuildPath()
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiMeasurementsResponse{}
	baseUrl, err := clients.GetBaseUrl(rc.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, rc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

func (rc readingClient) MeasurementsByDeviceNameAndResourceNamesAndTimeRange(ctx context.Context, deviceName string, resourceNames []string, start, end int64, offset, limit int) (responses.MultiMeasurementsResponse, errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(rc.enableNameFieldEscape).
		SetPath(common.ApiMeasurementRoute).SetPath(common.Device).SetPath(common.Name).SetNameFieldPath(deviceName).
		SetPath(common.Start).SetPath(strconv.FormatInt(start, 10)).SetPath(common.End).SetPath(strconv.FormatInt(end, 10)).BuildPath()
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	var queryPayload map[string]interface{}
	if len(resourceNames) > 0 { // gosimple S1009: len(nil slice) == 0
		queryPayload = make(map[string]interface{}, 1)
		queryPayload[common.ResourceNames] = resourceNames
	}
	res := responses.MultiMeasurementsResponse{}
	baseUrl, err := clients.GetBaseUrl(rc.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	err = utils.GetRequestWithBodyRawData(ctx, &res, baseUrl, requestPath, requestParams, queryPayload, rc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}
