package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
	"iiot-backend/pkg/go-mod-core-contracts/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// DeviceHandlerClient implements the DeviceHandlerClient interface
type DeviceHandlerClient struct {
	baseUrl    string
	httpClient *http.Client
}

// NewDeviceHandlerClient creates a new device service HTTP client
func NewDeviceHandlerClient(baseUrl string) interfaces.DeviceHandlerClient {
	return &DeviceHandlerClient{
		baseUrl:    baseUrl,
		httpClient: &http.Client{},
	}
}

// Add adds new device services
func (dsc *DeviceHandlerClient) Add(ctx context.Context, reqs []dtos.DeviceHandlerRequest) ([]dtos.BaseWithIdResponse, errors.IIOTError) {
	var responses []dtos.BaseWithIdResponse
	
	for _, req := range reqs {
		jsonData, err := json.Marshal(req)
		if err != nil {
			return nil, errors.NewCommonIIOT(errors.ErrKindContractInvalid, "failed to marshal request", err)
		}

		httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, dsc.baseUrl+common.ApiDeviceHandlerRoute, bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, errors.NewCommonIIOT(errors.ErrKindUnexpectedError, "failed to create HTTP request", err)
		}
		httpReq.Header.Set("Content-Type", common.ContentTypeJSON)

		resp, err := dsc.httpClient.Do(httpReq)
		if err != nil {
			return nil, errors.NewCommonIIOT(errors.ErrKindServiceUnavailable, "HTTP request failed", err)
		}
		defer resp.Body.Close()

		var response dtos.BaseWithIdResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, errors.NewCommonIIOT(errors.ErrKindContractInvalid, "failed to decode response", err)
		}

		responses = append(responses, response)
	}

	return responses, errors.IIOTError{}
}

// Update updates device services
func (dsc *DeviceHandlerClient) Update(ctx context.Context, reqs []dtos.UpdateDeviceHandlerRequest) ([]dtos.BaseResponse, errors.IIOTError) {
	var responses []dtos.BaseResponse
	
	for _, req := range reqs {
		jsonData, err := json.Marshal(req)
		if err != nil {
			return nil, errors.NewCommonIIOT(errors.ErrKindContractInvalid, "failed to marshal request", err)
		}

		httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, dsc.baseUrl+common.ApiDeviceHandlerRoute, bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, errors.NewCommonIIOT(errors.ErrKindUnexpectedError, "failed to create HTTP request", err)
		}
		httpReq.Header.Set("Content-Type", common.ContentTypeJSON)

		resp, err := dsc.httpClient.Do(httpReq)
		if err != nil {
			return nil, errors.NewCommonIIOT(errors.ErrKindServiceUnavailable, "HTTP request failed", err)
		}
		defer resp.Body.Close()

		var response dtos.BaseResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, errors.NewCommonIIOT(errors.ErrKindContractInvalid, "failed to decode response", err)
		}

		responses = append(responses, response)
	}

	return responses, errors.IIOTError{}
}

// AllDeviceHandlers retrieves all device services
func (dsc *DeviceHandlerClient) AllDeviceHandlers(ctx context.Context, labels []string, offset int, limit int) (dtos.MultiDeviceHandlersResponse, errors.IIOTError) {
	requestParams := url.Values{}
	requestParams.Set("offset", strconv.Itoa(offset))
	requestParams.Set("limit", strconv.Itoa(limit))
	for _, label := range labels {
		requestParams.Add("labels", label)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, dsc.baseUrl+common.ApiDeviceHandlerRoute+"?"+requestParams.Encode(), nil)
	if err != nil {
		return dtos.MultiDeviceHandlersResponse{}, errors.NewCommonIIOT(errors.ErrKindUnexpectedError, "failed to create HTTP request", err)
	}

	resp, err := dsc.httpClient.Do(httpReq)
	if err != nil {
		return dtos.MultiDeviceHandlersResponse{}, errors.NewCommonIIOT(errors.ErrKindServiceUnavailable, "HTTP request failed", err)
	}
	defer resp.Body.Close()

	var response dtos.MultiDeviceHandlersResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return dtos.MultiDeviceHandlersResponse{}, errors.NewCommonIIOT(errors.ErrKindContractInvalid, "failed to decode response", err)
	}

	return response, errors.IIOTError{}
}

// DeviceHandlerByName retrieves device service by name
func (dsc *DeviceHandlerClient) DeviceHandlerByName(ctx context.Context, name string) (dtos.DeviceHandlerResponse, errors.IIOTError) {
	path := fmt.Sprintf("%s/name/%s", common.ApiDeviceHandlerRoute, url.PathEscape(name))
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, dsc.baseUrl+path, nil)
	if err != nil {
		return dtos.DeviceHandlerResponse{}, errors.NewCommonIIOT(errors.ErrKindUnexpectedError, "failed to create HTTP request", err)
	}

	resp, err := dsc.httpClient.Do(httpReq)
	if err != nil {
		return dtos.DeviceHandlerResponse{}, errors.NewCommonIIOT(errors.ErrKindServiceUnavailable, "HTTP request failed", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return dtos.DeviceHandlerResponse{}, errors.NewEntityDoesNotExistError("DeviceHandler", name)
	}

	var response dtos.DeviceHandlerResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return dtos.DeviceHandlerResponse{}, errors.NewCommonIIOT(errors.ErrKindContractInvalid, "failed to decode response", err)
	}

	return response, errors.IIOTError{}
}

// DeleteDeviceHandlerByName deletes device service by name
func (dsc *DeviceHandlerClient) DeleteDeviceHandlerByName(ctx context.Context, name string) (dtos.BaseResponse, errors.IIOTError) {
	path := fmt.Sprintf("%s/name/%s", common.ApiDeviceHandlerRoute, url.PathEscape(name))
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodDelete, dsc.baseUrl+path, nil)
	if err != nil {
		return dtos.BaseResponse{}, errors.NewCommonIIOT(errors.ErrKindUnexpectedError, "failed to create HTTP request", err)
	}

	resp, err := dsc.httpClient.Do(httpReq)
	if err != nil {
		return dtos.BaseResponse{}, errors.NewCommonIIOT(errors.ErrKindServiceUnavailable, "HTTP request failed", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return dtos.BaseResponse{}, errors.NewEntityDoesNotExistError("DeviceHandler", name)
	}

	var response dtos.BaseResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return dtos.BaseResponse{}, errors.NewCommonIIOT(errors.ErrKindContractInvalid, "failed to decode response", err)
	}

	return response, errors.IIOTError{}
}