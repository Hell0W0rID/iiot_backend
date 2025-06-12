//
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"path"
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

type AlertClient struct {
	baseUrlFunc           clients.ClientBaseUrlFunc
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewAlertClient creates an instance of AlertClient
func NewAlertClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.AlertClient {
	return &AlertClient{
		baseUrlFunc:           clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// NewAlertClientWithUrlCallback creates an instance of AlertClient with ClientBaseUrlFunc.
func NewAlertClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.AlertClient {
	return &AlertClient{
		baseUrlFunc:           baseUrlFunc,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// SendAlert sends new notifications.
func (client *AlertClient) SendAlert(ctx context.Context, reqs []requests.AddAlertRequest) (res []dtoCommon.BaseWithIdResponse, err errors.IIOT) {
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiAlertRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// AlertById query notification by id.
func (client *AlertClient) AlertById(ctx context.Context, id string) (res responses.AlertResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiAlertRoute, common.Id, id)
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, nil, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeleteAlertById deletes a notification by id.
func (client *AlertClient) DeleteAlertById(ctx context.Context, id string) (res dtoCommon.BaseResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiAlertRoute, common.Id, id)
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.DeleteRequest(ctx, &res, baseUrl, requestPath, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// AlertsByCategory queries notifications with category, offset and limit
func (client *AlertClient) AlertsByCategory(ctx context.Context, category string, offset int, limit int, ack string) (res responses.MultiAlertsResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiAlertRoute, common.Category, category)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	requestParams.Set(common.Ack, ack)
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// AlertsByLabel queries notifications with label, offset and limit
func (client *AlertClient) AlertsByLabel(ctx context.Context, label string, offset int, limit int, ack string) (res responses.MultiAlertsResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiAlertRoute, common.Label, label)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	requestParams.Set(common.Ack, ack)
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// AlertsByStatus queries notifications with status, offset and limit
func (client *AlertClient) AlertsByStatus(ctx context.Context, status string, offset int, limit int, ack string) (res responses.MultiAlertsResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiAlertRoute, common.Status, status)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	requestParams.Set(common.Ack, ack)
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// AlertsByTimeRange query notifications with time range, offset and limit
func (client *AlertClient) AlertsByTimeRange(ctx context.Context, start int64, end int64, offset int, limit int, ack string) (res responses.MultiAlertsResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiAlertRoute, common.Start, strconv.FormatInt(start, 10), common.End, strconv.FormatInt(end, 10))
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	requestParams.Set(common.Ack, ack)
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// AlertsByEventSubscriptionName query notifications with subscriptionName, offset and limit
func (client *AlertClient) AlertsByEventSubscriptionName(ctx context.Context, subscriptionName string, offset int, limit int, ack string) (res responses.MultiAlertsResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiAlertRoute, common.EventSubscription, common.Name, subscriptionName)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	requestParams.Set(common.Ack, ack)
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// CleanupAlertsByAge removes notifications that are older than age. And the corresponding transmissions will also be deleted.
// Age is supposed in milliseconds since modified timestamp
func (client *AlertClient) CleanupAlertsByAge(ctx context.Context, age int) (res dtoCommon.BaseResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiAlertCleanupRoute, common.Age, strconv.Itoa(age))
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.DeleteRequest(ctx, &res, baseUrl, requestPath, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// CleanupAlerts removes notifications and the corresponding transmissions.
func (client *AlertClient) CleanupAlerts(ctx context.Context) (res dtoCommon.BaseResponse, err errors.IIOT) {
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.DeleteRequest(ctx, &res, baseUrl, common.ApiAlertCleanupRoute, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeleteProcessedAlertsByAge removes processed notifications that are older than age. And the corresponding transmissions will also be deleted.
// Age is supposed in milliseconds since modified timestamp
// Please notice that this API is only for processed notifications (status = PROCESSED). If the deletion purpose includes each kind of notifications, please refer to cleanup API.
func (client *AlertClient) DeleteProcessedAlertsByAge(ctx context.Context, age int) (res dtoCommon.BaseResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiAlertRoute, common.Age, strconv.Itoa(age))
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.DeleteRequest(ctx, &res, baseUrl, requestPath, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// AlertsByQueryConditions queries notifications with offset, limit, acknowledgement status, category and time range
func (client *AlertClient) AlertsByQueryConditions(ctx context.Context, offset, limit int, ack string, conditionReq requests.GetAlertRequest) (res responses.MultiAlertsResponse, err errors.IIOT) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	requestParams.Set(common.Ack, ack)
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequestWithBodyRawData(ctx, &res, baseUrl, common.ApiAlertRoute, requestParams, conditionReq, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// DeleteAlertByIds deletes notifications by ids
func (client *AlertClient) DeleteAlertByIds(ctx context.Context, ids []string) (res dtoCommon.BaseResponse, err errors.IIOT) {
	requestPath := utils.EscapeAndJoinPath(common.ApiAlertRoute, common.Ids, strings.Join(ids, common.CommaSeparator))
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.DeleteRequest(ctx, &res, baseUrl, requestPath, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// UpdateAlertAckStatusByIds updates existing notification's acknowledgement status
func (client *AlertClient) UpdateAlertAckStatusByIds(ctx context.Context, ack bool, ids []string) (res dtoCommon.BaseResponse, err errors.IIOT) {
	pathAck := common.Unacknowledge
	if ack {
		pathAck = common.Acknowledge
	}
	requestPath := utils.EscapeAndJoinPath(common.ApiAlertRoute, pathAck, common.Ids, strings.Join(ids, common.CommaSeparator))
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.PutRequest(ctx, &res, baseUrl, requestPath, nil, nil, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}
