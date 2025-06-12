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
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

type ScheduleActionRecordClient struct {
	baseUrlFunc           clients.ClientBaseUrlFunc
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewScheduleActionRecordClient creates an instance of ScheduleActionRecordClient
func NewScheduleActionRecordClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.ScheduleActionRecordClient {
	return &ScheduleActionRecordClient{
		baseUrlFunc:           clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// NewScheduleActionRecordClientWithUrlCallback creates an instance of ScheduleActionRecordClient with ClientBaseUrlFunc.
func NewScheduleActionRecordClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.ScheduleActionRecordClient {
	return &ScheduleActionRecordClient{
		baseUrlFunc:           baseUrlFunc,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// AllScheduleActionRecords query schedule action records with start, end, offset, and limit
func (client *ScheduleActionRecordClient) AllScheduleActionRecords(ctx context.Context, start, end int64, offset, limit int) (res responses.MultiScheduleActionRecordsResponse, err errors.IIOT) {
	requestParams := url.Values{}
	requestParams.Set(common.Start, strconv.FormatInt(start, 10))
	requestParams.Set(common.End, strconv.FormatInt(end, 10))
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllScheduleActionRecordRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// LatestScheduleActionRecordsByJobName query the latest schedule action records by job name
func (client *ScheduleActionRecordClient) LatestScheduleActionRecordsByJobName(ctx context.Context, jobName string) (res responses.MultiScheduleActionRecordsResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiScheduleActionRecordRoute, common.Latest, common.Job, common.Name, jobName)
	requestParams := url.Values{}
	requestParams.Set(common.Name, jobName)
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

// ScheduleActionRecordsByStatus queries schedule action records with status, start, end, offset, and limit
func (client *ScheduleActionRecordClient) ScheduleActionRecordsByStatus(ctx context.Context, status string, start, end int64, offset, limit int) (res responses.MultiScheduleActionRecordsResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiScheduleActionRecordRoute, common.Status, status)
	requestParams := url.Values{}
	requestParams.Set(common.Start, strconv.FormatInt(start, 10))
	requestParams.Set(common.End, strconv.FormatInt(end, 10))
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
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

// ScheduleActionRecordsByJobName queries schedule action records with jobName, start, end, offset, and limit
func (client *ScheduleActionRecordClient) ScheduleActionRecordsByJobName(ctx context.Context, jobName string, start, end int64, offset, limit int) (res responses.MultiScheduleActionRecordsResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiScheduleActionRecordRoute, common.Job, common.Name, jobName)
	requestParams := url.Values{}
	requestParams.Set(common.Start, strconv.FormatInt(start, 10))
	requestParams.Set(common.End, strconv.FormatInt(end, 10))
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
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

// ScheduleActionRecordsByJobNameAndStatus queries schedule action records with jobName, status, start, end, offset, and limit
func (client *ScheduleActionRecordClient) ScheduleActionRecordsByJobNameAndStatus(ctx context.Context, jobName, status string, start, end int64, offset, limit int) (res responses.MultiScheduleActionRecordsResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiScheduleActionRecordRoute, common.Job, common.Name, jobName, common.Status, status)
	requestParams := url.Values{}
	requestParams.Set(common.Start, strconv.FormatInt(start, 10))
	requestParams.Set(common.End, strconv.FormatInt(end, 10))
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
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
