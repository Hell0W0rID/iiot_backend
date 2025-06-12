//
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
	"iiot-backend/pkg/go-mod-core-contracts/common"
	commonDTO "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	iiotErr "iiot-backend/pkg/go-mod-core-contracts/errors"

	"iiot-backend/pkg/go-mod-messaging/messaging"
	"iiot-backend/pkg/go-mod-messaging/pkg/types"
)

type CommandClient struct {
	messageBus            messaging.MessageClient
	baseTopic             string
	responseTopicPrefix   string
	timeout               time.Duration
	enableNameFieldEscape bool
}

// NewCommandClient returns the command client with the disabled NameFieldEscape
func NewCommandClient(messageBus messaging.MessageClient, baseTopic string, timeout time.Duration) interfaces.CommandClient {
	client := &CommandClient{
		messageBus:          messageBus,
		baseTopic:           baseTopic,
		responseTopicPrefix: common.BuildTopic(baseTopic, common.ResponseTopic, common.CoreCommandServiceName),
		timeout:             timeout,
	}

	return client
}

// NewCommandClientWithNameFieldEscape returns the command client with the enabled NameFieldEscape
func NewCommandClientWithNameFieldEscape(messageBus messaging.MessageClient, baseTopic string, timeout time.Duration) interfaces.CommandClient {
	client := &CommandClient{
		messageBus:            messageBus,
		baseTopic:             baseTopic,
		responseTopicPrefix:   common.BuildTopic(baseTopic, common.ResponseTopic, common.CoreCommandServiceName),
		timeout:               timeout,
		enableNameFieldEscape: true,
	}

	return client
}

func (c *CommandClient) AllDeviceCoreCommands(_ context.Context, offset int, limit int) (responses.MultiDeviceCoreCommandsResponse, iiotErr.IIOT) {
	queryParams := map[string]string{common.Offset: strconv.Itoa(offset), common.Limit: strconv.Itoa(limit)}
	requestEnvelope := types.NewMessageEnvelopeForRequest(nil, queryParams)

	requestTopic := common.BuildTopic(c.baseTopic, common.CoreCommandQueryRequestPublishTopic, common.All)
	responseEnvelope, err := c.messageBus.Request(requestEnvelope, common.CoreCommandServiceName, requestTopic, c.timeout)
	if err != nil {
		return responses.MultiDeviceCoreCommandsResponse{}, iiotErr.NewCommonIIOTWrapper(err)
	}

	if responseEnvelope.ErrorCode == 1 {
		return responses.MultiDeviceCoreCommandsResponse{}, iiotErr.NewCommonIIOTWrapper(fmt.Errorf("%v", responseEnvelope.Payload))
	}

	var res responses.MultiDeviceCoreCommandsResponse
	res, err = types.GetMsgPayload[responses.MultiDeviceCoreCommandsResponse](*responseEnvelope)
	if err != nil {
		return responses.MultiDeviceCoreCommandsResponse{}, iiotErr.NewCommonIIOTWrapper(err)
	}

	return res, nil
}

func (c *CommandClient) DeviceCoreCommandsByDeviceName(_ context.Context, deviceName string) (responses.DeviceCoreCommandResponse, iiotErr.IIOT) {
	requestEnvelope := types.NewMessageEnvelopeForRequest(nil, nil)
	requestTopic := common.NewPathBuilder().EnableNameFieldEscape(c.enableNameFieldEscape).
		SetPath(c.baseTopic).SetPath(common.CoreCommandQueryRequestPublishTopic).SetNameFieldPath(deviceName).BuildPath()
	responseEnvelope, err := c.messageBus.Request(requestEnvelope, requestTopic, c.responseTopicPrefix, c.timeout)
	if err != nil {
		return responses.DeviceCoreCommandResponse{}, iiotErr.NewCommonIIOTWrapper(err)
	}

	if responseEnvelope.ErrorCode == 1 {
		return responses.DeviceCoreCommandResponse{}, iiotErr.NewCommonIIOTWrapper(fmt.Errorf("%v", responseEnvelope.Payload))
	}

	var res responses.DeviceCoreCommandResponse
	res, err = types.GetMsgPayload[responses.DeviceCoreCommandResponse](*responseEnvelope)
	if err != nil {
		return responses.DeviceCoreCommandResponse{}, iiotErr.NewCommonIIOTWrapper(err)
	}

	return res, nil
}

func (c *CommandClient) IssueGetCommandByName(ctx context.Context, deviceName string, commandName string, dsPushDataEvent bool, dsReturnDataEvent bool) (*responses.DataEventResponse, iiotErr.IIOT) {
	queryParams := map[string]string{common.PushDataEvent: strconv.FormatBool(dsPushDataEvent), common.ReturnDataEvent: strconv.FormatBool(dsReturnDataEvent)}
	return c.IssueGetCommandByNameWithQueryParams(ctx, deviceName, commandName, queryParams)
}

func (c *CommandClient) IssueGetCommandByNameWithQueryParams(_ context.Context, deviceName string, commandName string, queryParams map[string]string) (*responses.DataEventResponse, iiotErr.IIOT) {
	requestEnvelope := types.NewMessageEnvelopeForRequest(nil, queryParams)
	requestTopic := common.NewPathBuilder().EnableNameFieldEscape(c.enableNameFieldEscape).
		SetPath(c.baseTopic).SetPath(common.CoreCommandRequestPublishTopic).SetNameFieldPath(deviceName).SetNameFieldPath(commandName).SetPath("get").BuildPath()
	responseEnvelope, err := c.messageBus.Request(requestEnvelope, requestTopic, c.responseTopicPrefix, c.timeout)
	if err != nil {
		return nil, iiotErr.NewCommonIIOTWrapper(err)
	}

	if responseEnvelope.ErrorCode == 1 {
		return nil, iiotErr.NewCommonIIOTWrapper(fmt.Errorf("%v", responseEnvelope.Payload))
	}

	var res responses.DataEventResponse
	returnDataEvent, ok := queryParams[common.ReturnDataEvent]
	if ok && returnDataEvent == common.ValueFalse {
		res.ApiVersion = common.ApiVersion
		res.RequestId = responseEnvelope.RequestID
		res.StatusCode = http.StatusOK
	} else {
		res, err = types.GetMsgPayload[responses.DataEventResponse](*responseEnvelope)
		if err != nil {
			return nil, iiotErr.NewCommonIIOTWrapper(err)
		}
	}

	return &res, nil
}

func (c *CommandClient) IssueSetCommandByName(_ context.Context, deviceName string, commandName string, settings map[string]any) (commonDTO.BaseResponse, iiotErr.IIOT) {
	requestEnvelope := types.NewMessageEnvelopeForRequest(settings, nil)
	requestTopic := common.NewPathBuilder().EnableNameFieldEscape(c.enableNameFieldEscape).
		SetPath(c.baseTopic).SetPath(common.CoreCommandRequestPublishTopic).SetNameFieldPath(deviceName).SetNameFieldPath(commandName).SetPath("set").BuildPath()
	responseEnvelope, err := c.messageBus.Request(requestEnvelope, requestTopic, c.responseTopicPrefix, c.timeout)
	if err != nil {
		return commonDTO.BaseResponse{}, iiotErr.NewCommonIIOTWrapper(err)
	}

	if responseEnvelope.ErrorCode == 1 {
		return commonDTO.BaseResponse{}, iiotErr.NewCommonIIOTWrapper(fmt.Errorf("%v", responseEnvelope.Payload))
	}

	res := commonDTO.NewBaseResponse(responseEnvelope.RequestID, "", http.StatusOK)
	return res, nil
}
