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

type ScheduleJobClient struct {
	baseUrlFunc           clients.ClientBaseUrlFunc
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewScheduleJobClient creates an instance of ScheduleJobClient
func NewScheduleJobClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.ScheduleJobClient {
	return &ScheduleJobClient{
		baseUrlFunc:           clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// NewScheduleJobClientWithUrlCallback creates an instance of ScheduleJobClient with ClientBaseUrlFunc.
func NewScheduleJobClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.ScheduleJobClient {
	return &ScheduleJobClient{
		baseUrlFunc:           baseUrlFunc,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// Add adds new schedule jobs
func (client ScheduleJobClient) Add(ctx context.Context, reqs []requests.AddScheduleJobRequest) (
	res []dtoCommon.BaseWithIdResponse, err errors.IIOT) {
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiScheduleJobRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// Update updates schedule jobs
func (client ScheduleJobClient) Update(ctx context.Context, reqs []requests.UpdateScheduleJobRequest) (
	res []dtoCommon.BaseResponse, err errors.IIOT) {
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.PatchRequest(ctx, &res, baseUrl, common.ApiScheduleJobRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// AllScheduleJobs queries the schedule jobs with offset, limit
func (client ScheduleJobClient) AllScheduleJobs(ctx context.Context, labels []string, offset, limit int) (
	res responses.MultiScheduleJobsResponse, err errors.IIOT) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(common.Labels, strings.Join(labels, common.CommaSeparator))
	}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllScheduleJobRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// ScheduleJobByName queries the schedule job by name
func (client ScheduleJobClient) ScheduleJobByName(ctx context.Context, name string) (
	res responses.ScheduleJobResponse, err errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiScheduleJobRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
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

// DeleteScheduleJobByName deletes the schedule job by name
func (client ScheduleJobClient) DeleteScheduleJobByName(ctx context.Context, name string) (
	res dtoCommon.BaseResponse, err errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiScheduleJobRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
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

// TriggerScheduleJobByName triggers the schedule job by name
func (client ScheduleJobClient) TriggerScheduleJobByName(ctx context.Context, name string) (
	res dtoCommon.BaseResponse, err errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiTriggerScheduleJobRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, requestPath, nil, nil, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}
