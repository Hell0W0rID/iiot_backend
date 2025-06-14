//
// Copyright (C) 2022-2025 IOTech Ltd
// Copyright (C) 2023 Intel Inc.
//
// SPDX-License-Identifier: Apache-2.0

package messaging

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	bootstrapContainer "iiot-backend/pkg/go-mod-bootstrap/bootstrap/container"
	"iiot-backend/pkg/go-mod-bootstrap/di"
	"iiot-backend/pkg/go-mod-core-contracts/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"

	"iiot-backend/pkg/go-mod-messaging/pkg/types"

	"iiot-backend/services/core/command/application"
)

// retrieveServiceNameByDevice validates the existence of device and device service,
// returns the service name to which the command request will be sent.
func retrieveServiceNameByDevice(deviceName string, dic *di.Container) (string, error) {
	// retrieve device information through Metadata DeviceClient
	dc := bootstrapContainer.DeviceClientFrom(dic.Get)
	if dc == nil {
		return "", errors.New("nil Device Client")
	}
	deviceResponse, err := dc.DeviceByName(context.Background(), deviceName)
	if err != nil {
		return "", fmt.Errorf("failed to get Device by name %s: %v", deviceName, err)
	}

	// retrieve device service information through Metadata DeviceClient
	dsc := bootstrapContainer.DeviceHandlerClientFrom(dic.Get)
	if dsc == nil {
		return "", errors.New("nil DeviceService Client")
	}
	deviceServiceResponse, err := dsc.DeviceServiceByName(context.Background(), deviceResponse.Device.ServiceName)
	if err != nil {
		return "", fmt.Errorf("failed to get DeviceService by name %s: %v", deviceResponse.Device.ServiceName, err)
	}
	return deviceServiceResponse.Service.Name, nil
}

// validateGetCommandQueryParameters validates the value is valid for device service's reserved query parameters
func validateGetCommandQueryParameters(queryParams map[string]string) error {
	if dsReturnEvent, ok := queryParams[common.ReturnDataEvent]; ok {
		if dsReturnEvent != common.ValueTrue && dsReturnEvent != common.ValueFalse {
			return fmt.Errorf("invalid query parameter, %s has to be '%s' or '%s'", common.ReturnDataEvent, common.ValueTrue, common.ValueFalse)
		}
	}
	if dsPushEvent, ok := queryParams[common.PushDataEvent]; ok {
		if dsPushEvent != common.ValueTrue && dsPushEvent != common.ValueFalse {
			return fmt.Errorf("invalid query parameter, %s has to be '%s' or '%s'", common.PushDataEvent, common.ValueTrue, common.ValueFalse)
		}
	}

	return nil
}

// getCommandQueryResponseEnvelope returns the MessageEnvelope containing the DeviceCoreCommand payload bytes
func getCommandQueryResponseEnvelope(requestEnvelope types.MessageEnvelope, deviceName string, dic *di.Container) (types.MessageEnvelope, error) {
	var commandsResponse any
	var err error

	switch deviceName {
	case common.All:
		offset, limit := common.DefaultOffset, common.DefaultLimit
		if requestEnvelope.QueryParams != nil {
			if offsetRaw, ok := requestEnvelope.QueryParams[common.Offset]; ok {
				offset, err = strconv.Atoi(offsetRaw)
				if err != nil {
					return types.MessageEnvelope{}, fmt.Errorf("failed to convert 'offset' query parameter to intger: %s", err.Error())
				}
			}
			if limitRaw, ok := requestEnvelope.QueryParams[common.Limit]; ok {
				limit, err = strconv.Atoi(limitRaw)
				if err != nil {
					return types.MessageEnvelope{}, fmt.Errorf("failed to convert 'limit' query parameter to integer: %s", err.Error())
				}
			}
		}

		commands, totalCounts, iiotError := application.AllCommands(offset, limit, dic)
		if iiotError != nil {
			return types.MessageEnvelope{}, fmt.Errorf("failed to get all commands: %s", iiotError.Error())
		}

		commandsResponse = responses.NewMultiDeviceCoreCommandsResponse(requestEnvelope.RequestID, "", http.StatusOK, totalCounts, commands)
	default:
		commands, iiotError := application.CommandsByDeviceName(deviceName, dic)
		if iiotError != nil {
			return types.MessageEnvelope{}, fmt.Errorf("failed to get commands by device name '%s': %s", deviceName, iiotError.Error())
		}

		commandsResponse = responses.NewDeviceCoreCommandResponse(requestEnvelope.RequestID, "", http.StatusOK, commands)
	}

	responseEnvelope, err := types.NewMessageEnvelopeForResponse(commandsResponse, requestEnvelope.RequestID, requestEnvelope.CorrelationID, common.ContentTypeJSON)
	if err != nil {
		return types.MessageEnvelope{}, fmt.Errorf("failed to create response MessageEnvelope: %s", err.Error())
	}

	return responseEnvelope, nil
}
