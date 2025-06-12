//
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"fmt"

	"iiot-backend/pkg/go-mod-core-contracts/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
	"iiot-backend/pkg/go-mod-core-contracts/models"
)

var supportedChannelTypes = []string{common.EMAIL, common.REST, common.MQTT, common.ZeroMQ}

// AddEventSubscriptionRequest defines the Request Content for POST EventSubscription DTO.
type AddEventSubscriptionRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	EventSubscription          dtos.EventSubscription `json:"subscription"`
}

// Validate satisfies the Validator interface
func (request AddEventSubscriptionRequest) Validate() error {
	err := common.Validate(request)
	if err != nil {
		return errors.NewCommonIIOTWrapper(err)
	}
	for _, c := range request.EventSubscription.Channels {
		err = c.Validate()
		if err != nil {
			return errors.NewCommonIIOTWrapper(err)
		} else if !contains(supportedChannelTypes, c.Type) {
			return errors.NewCommonIIOT(errors.KindContractInvalid, fmt.Sprintf("%s is not valid type for Channel", c.Type), nil)
		}
	}
	return nil
}

// UnmarshalJSON implements the Unmarshaler interface for the AddEventSubscriptionRequest type
func (request *AddEventSubscriptionRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		EventSubscription dtos.EventSubscription
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*request = AddEventSubscriptionRequest(alias)

	// validate AddEventSubscriptionRequest DTO
	if err := request.Validate(); err != nil {
		return errors.NewCommonIIOTWrapper(err)
	}
	return nil
}

// AddEventSubscriptionReqToEventSubscriptionModels transforms the AddEventSubscriptionRequest DTO array to the AddEventSubscriptionRequest model array
func AddEventSubscriptionReqToEventSubscriptionModels(reqs []AddEventSubscriptionRequest) (s []models.EventSubscription) {
	for _, req := range reqs {
		d := dtos.ToEventSubscriptionModel(req.EventSubscription)
		s = append(s, d)
	}
	return s
}

// UpdateEventSubscriptionRequest defines the Request Content for PUT event as pushed DTO.
type UpdateEventSubscriptionRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	EventSubscription          dtos.UpdateEventSubscription `json:"subscription"`
}

// Validate satisfies the Validator interface
func (request UpdateEventSubscriptionRequest) Validate() error {
	err := common.Validate(request)
	if err != nil {
		return errors.NewCommonIIOTWrapper(err)
	}
	for _, c := range request.EventSubscription.Channels {
		err = c.Validate()
		if err != nil {
			return errors.NewCommonIIOTWrapper(err)
		} else if !contains(supportedChannelTypes, c.Type) {
			return errors.NewCommonIIOT(errors.KindContractInvalid, fmt.Sprintf("%s is not valid type for Channel", c.Type), nil)
		}
	}
	if request.EventSubscription.Categories != nil && request.EventSubscription.Labels != nil &&
		len(request.EventSubscription.Categories) == 0 && len(request.EventSubscription.Labels) == 0 {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "categories and labels can not be both empty", nil)
	}
	return nil
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateEventSubscriptionRequest type
func (request *UpdateEventSubscriptionRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		EventSubscription dtos.UpdateEventSubscription
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*request = UpdateEventSubscriptionRequest(alias)

	// validate UpdateEventSubscriptionRequest DTO
	if err := request.Validate(); err != nil {
		return errors.NewCommonIIOTWrapper(err)
	}
	return nil
}

// ReplaceEventSubscriptionModelFieldsWithDTO replace existing EventSubscription's fields with DTO patch
func ReplaceEventSubscriptionModelFieldsWithDTO(s *models.EventSubscription, patch dtos.UpdateEventSubscription) {
	if patch.Channels != nil {
		s.Channels = dtos.ToAddressModels(patch.Channels)
	}
	if patch.Categories != nil {
		s.Categories = patch.Categories
	}
	if patch.Labels != nil {
		s.Labels = patch.Labels
	}
	if patch.Description != nil {
		s.Description = *patch.Description
	}
	if patch.Receiver != nil {
		s.Receiver = *patch.Receiver
	}
	if patch.ResendLimit != nil {
		s.ResendLimit = *patch.ResendLimit
	}
	if patch.ResendInterval != nil {
		s.ResendInterval = *patch.ResendInterval
	}
	if patch.ServiceState != nil {
		s.ServiceState = models.ServiceState(*patch.ServiceState)
	}
}

func NewAddEventSubscriptionRequest(dto dtos.EventSubscription) AddEventSubscriptionRequest {
	return AddEventSubscriptionRequest{
		BaseRequest:  dtoCommon.NewBaseRequest(),
		EventSubscription: dto,
	}
}

func NewUpdateEventSubscriptionRequest(dto dtos.UpdateEventSubscription) UpdateEventSubscriptionRequest {
	return UpdateEventSubscriptionRequest{
		BaseRequest:  dtoCommon.NewBaseRequest(),
		EventSubscription: dto,
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
