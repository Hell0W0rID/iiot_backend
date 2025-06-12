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

type EventSubscriptionClient struct {
	baseUrlFunc           clients.ClientBaseUrlFunc
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewEventSubscriptionClient creates an instance of EventSubscriptionClient
func NewEventSubscriptionClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.EventSubscriptionClient {
	return &EventSubscriptionClient{
		baseUrlFunc:           clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// NewEventSubscriptionClientWithUrlCallback creates an instance of EventSubscriptionClient with ClientBaseUrlFunc.
func NewEventSubscriptionClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.EventSubscriptionClient {
	return &EventSubscriptionClient{
		baseUrlFunc:           baseUrlFunc,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// Add adds new subscriptions.
func (client *EventSubscriptionClient) Add(ctx context.Context, reqs []requests.AddEventSubscriptionRequest) (res []dtoCommon.BaseWithIdResponse, err errors.IIOT) {
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiEventSubscriptionRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// Update updates subscriptions.
func (client *EventSubscriptionClient) Update(ctx context.Context, reqs []requests.UpdateEventSubscriptionRequest) (res []dtoCommon.BaseResponse, err errors.IIOT) {
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.PatchRequest(ctx, &res, baseUrl, common.ApiEventSubscriptionRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// AllEventSubscriptions queries subscriptions with offset and limit
func (client *EventSubscriptionClient) AllEventSubscriptions(ctx context.Context, offset int, limit int) (res responses.MultiEventSubscriptionsResponse, err errors.IIOT) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, goErr := clients.GetBaseUrl(client.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonIIOTWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllEventSubscriptionRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// EventSubscriptionsByCategory queries subscriptions with category, offset and limit
func (client *EventSubscriptionClient) EventSubscriptionsByCategory(ctx context.Context, category string, offset int, limit int) (res responses.MultiEventSubscriptionsResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiEventSubscriptionRoute, common.Category, category)
	requestParams := url.Values{}
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

// EventSubscriptionsByLabel queries subscriptions with label, offset and limit
func (client *EventSubscriptionClient) EventSubscriptionsByLabel(ctx context.Context, label string, offset int, limit int) (res responses.MultiEventSubscriptionsResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiEventSubscriptionRoute, common.Label, label)
	requestParams := url.Values{}
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

// EventSubscriptionsByReceiver queries subscriptions with receiver, offset and limit
func (client *EventSubscriptionClient) EventSubscriptionsByReceiver(ctx context.Context, receiver string, offset int, limit int) (res responses.MultiEventSubscriptionsResponse, err errors.IIOT) {
	requestPath := path.Join(common.ApiEventSubscriptionRoute, common.Receiver, receiver)
	requestParams := url.Values{}
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

// EventSubscriptionByName query subscription by name.
func (client *EventSubscriptionClient) EventSubscriptionByName(ctx context.Context, name string) (res responses.EventSubscriptionResponse, err errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiEventSubscriptionRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
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

// DeleteEventSubscriptionByName deletes a subscription by name.
func (client *EventSubscriptionClient) DeleteEventSubscriptionByName(ctx context.Context, name string) (res dtoCommon.BaseResponse, err errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiEventSubscriptionRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
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
