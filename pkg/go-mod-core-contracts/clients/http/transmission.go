//
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"path"
	"strconv"

	"iiot-backend/pkg/go-mod-core-contracts/clients/http/utils"
	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
	"iiot-backend/pkg/go-mod-core-contracts/common"
	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

type DeliveryClient struct {
	baseUrl               string
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewDeliveryClient creates an instance of DeliveryClient
func NewDeliveryClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.DeliveryClient {
	return &DeliveryClient{
		baseUrl:               baseUrl,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// DeliveryById query transmission by id.
func (client *DeliveryClient) DeliveryById(ctx context.Context, id string) (res responses.DeliveryResponse, err errors.IIOT) {
	path := path.Join(common.ApiDeliveryRoute, common.Id, id)
	err = utils.GetRequest(ctx, &res, client.baseUrl, path, nil, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeliverysByTimeRange query transmissions with time range, offset and limit
func (client *DeliveryClient) DeliverysByTimeRange(ctx context.Context, start int64, end int64, offset int, limit int) (res responses.MultiDeliverysResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiDeliveryRoute, common.Start, strconv.FormatInt(start, 10), common.End, strconv.FormatInt(end, 10))
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// AllDeliverys query transmissions with offset and limit
func (client *DeliveryClient) AllDeliverys(ctx context.Context, offset int, limit int) (res responses.MultiDeliverysResponse, err errors.IIOT) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, common.ApiAllDeliveryRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeliverysByStatus queries transmissions with status, offset and limit
func (client *DeliveryClient) DeliverysByStatus(ctx context.Context, status string, offset int, limit int) (res responses.MultiDeliverysResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiDeliveryRoute, common.Status, status)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeleteProcessedDeliverysByAge deletes the processed transmissions if the current timestamp minus their created timestamp is less than the age parameter.
func (client *DeliveryClient) DeleteProcessedDeliverysByAge(ctx context.Context, age int) (res dtoCommon.BaseResponse, err errors.IIOT) {
	path := path.Join(common.ApiDeliveryRoute, common.Age, strconv.Itoa(age))
	err = utils.DeleteRequest(ctx, &res, client.baseUrl, path, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeliverysByEventSubscriptionName query transmissions with subscriptionName, offset and limit
func (client *DeliveryClient) DeliverysByEventSubscriptionName(ctx context.Context, subscriptionName string, offset int, limit int) (res responses.MultiDeliverysResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiDeliveryRoute, common.EventSubscription, common.Name, subscriptionName)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeliverysByAlertId query transmissions with notification id, offset and limit
func (client *DeliveryClient) DeliverysByAlertId(ctx context.Context, id string, offset int, limit int) (res responses.MultiDeliverysResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiDeliveryRoute, common.Alert, common.Id, id)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}
