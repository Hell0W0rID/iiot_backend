//
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"strconv"

	"iiot-backend/pkg/go-mod-core-contracts/clients/http/utils"
	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
	"iiot-backend/pkg/go-mod-core-contracts/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/requests"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

var emptyResponse any

// RegistryClient is the REST client for invoking the registry APIs(/registry/*) from Core Keeper
type registryClient struct {
	baseUrl               string
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewRegistryClient creates an instance of RegistryClient
func NewRegistryClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.RegistryClient {
	return &registryClient{
		baseUrl:               baseUrl,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// Register registers a service instance
func (rc *registryClient) Register(ctx context.Context, req requests.AddRegistrationRequest) errors.IIOT {
	err := utils.PostRequestWithRawData(ctx, &emptyResponse, rc.baseUrl, common.ApiRegisterRoute, nil, req, rc.authInjector)
	if err != nil {
		return errors.NewCommonIIOTWrapper(err)
	}
	return nil
}

// UpdateRegister updates the registration data of the service
func (rc *registryClient) UpdateRegister(ctx context.Context, req requests.AddRegistrationRequest) errors.IIOT {
	err := utils.PutRequest(ctx, &emptyResponse, rc.baseUrl, common.ApiRegisterRoute, nil, req, rc.authInjector)
	if err != nil {
		return errors.NewCommonIIOTWrapper(err)
	}
	return nil
}

// RegistrationByServiceId returns the registration data by service id
func (rc *registryClient) RegistrationByServiceId(ctx context.Context, serviceId string) (responses.RegistrationResponse, errors.IIOT) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(rc.enableNameFieldEscape).
		SetPath(common.ApiRegisterRoute).SetPath(common.ServiceId).SetNameFieldPath(serviceId).BuildPath()
	res := responses.RegistrationResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, nil, rc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// AllRegistry returns the registration data of all registered service
func (rc *registryClient) AllRegistry(ctx context.Context, deregistered bool) (responses.MultiRegistrationsResponse, errors.IIOT) {
	requestParams := url.Values{}
	requestParams.Set(common.Deregistered, strconv.FormatBool(deregistered))

	res := responses.MultiRegistrationsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, common.ApiAllRegistrationsRoute, requestParams, rc.authInjector)
	if err != nil {
		return res, errors.NewCommonIIOTWrapper(err)
	}
	return res, nil
}

// Deregister deregisters a service by service id
func (rc *registryClient) Deregister(ctx context.Context, serviceId string) errors.IIOT {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(rc.enableNameFieldEscape).
		SetPath(common.ApiRegisterRoute).SetPath(common.ServiceId).SetNameFieldPath(serviceId).BuildPath()
	err := utils.DeleteRequest(ctx, &emptyResponse, rc.baseUrl, requestPath, rc.authInjector)
	if err != nil {
		return errors.NewCommonIIOTWrapper(err)
	}
	return nil
}
