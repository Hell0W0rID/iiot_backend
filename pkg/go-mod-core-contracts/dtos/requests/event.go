//
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"os"

	"iiot-backend/pkg/go-mod-core-contracts/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
	"iiot-backend/pkg/go-mod-core-contracts/models"

	"github.com/fxamacker/cbor/v2"
)

// AddDataEventRequest defines the Request Content for POST event DTO.
type AddDataEventRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	DataEvent                 dtos.DataEvent `json:"event" validate:"required"`
}

// NewAddDataEventRequest creates, initializes and returns an AddDataEventRequests
func NewAddDataEventRequest(event dtos.DataEvent) AddDataEventRequest {
	return AddDataEventRequest{
		BaseRequest: dtoCommon.NewBaseRequest(),
		DataEvent:       event,
	}
}

// Validate satisfies the Validator interface
func (a AddDataEventRequest) Validate() error {
	if err := common.Validate(a); err != nil {
		return err
	}

	// BaseMeasurement has the skip("-") validation annotation for BinaryMeasurement and SimpleMeasurement
	// Otherwise error will occur as only one of them exists
	// Therefore, need to validate the nested BinaryMeasurement and SimpleMeasurement struct here
	for _, r := range a.DataEvent.Measurements {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type unmarshal func([]byte, interface{}) error

func (a *AddDataEventRequest) UnmarshalJSON(b []byte) error {
	return a.Unmarshal(b, json.Unmarshal)
}

func (a *AddDataEventRequest) UnmarshalCBOR(b []byte) error {
	return a.Unmarshal(b, cbor.Unmarshal)
}

func (a *AddDataEventRequest) Unmarshal(b []byte, f unmarshal) error {
	// To avoid recursively invoke unmarshaler interface, intentionally create a struct to represent AddDataEventRequest DTO
	var addDataEvent struct {
		dtoCommon.BaseRequest
		DataEvent dtos.DataEvent
	}
	if err := f(b, &addDataEvent); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal the byte array.", err)
	}

	*a = AddDataEventRequest(addDataEvent)

	// validate AddDataEventRequest DTO
	if err := a.Validate(); err != nil {
		return err
	}

	// Normalize reading's value type
	for i, r := range a.DataEvent.Measurements {
		valueType, err := common.NormalizeValueType(r.ValueType)
		if err != nil {
			return errors.NewCommonIIOTWrapper(err)
		}
		a.DataEvent.Measurements[i].ValueType = valueType
	}
	return nil
}

func (a *AddDataEventRequest) Encode() ([]byte, string, error) {
	var encoding = a.GetEncodingContentType()
	var err error
	var encodedData []byte

	switch encoding {
	case common.ContentTypeCBOR:
		encodedData, err = cbor.Marshal(a)
		if err != nil {
			return nil, "", errors.NewCommonIIOT(errors.KindContractInvalid, "failed to encode AddDataEventRequest to CBOR", err)
		}
	case common.ContentTypeJSON:
		encodedData, err = json.Marshal(a)
		if err != nil {
			return nil, "", errors.NewCommonIIOT(errors.KindContractInvalid, "failed to encode AddDataEventRequest to JSON", err)
		}
	}

	return encodedData, encoding, nil
}

// GetEncodingContentType determines which content type should be used to encode and decode this object
func (a *AddDataEventRequest) GetEncodingContentType() string {
	if v := os.Getenv(common.EnvEncodeAllDataEvents); v == common.ValueTrue {
		return common.ContentTypeCBOR
	}
	for _, r := range a.DataEvent.Measurements {
		if r.ValueType == common.ValueTypeBinary {
			return common.ContentTypeCBOR
		}
	}

	return common.ContentTypeJSON
}

// AddDataEventReqToDataEventModel transforms the AddDataEventRequest DTO to the DataEvent model
func AddDataEventReqToDataEventModel(addDataEventReq AddDataEventRequest) (event models.DataEvent) {
	readings := make([]models.Measurement, len(addDataEventReq.DataEvent.Measurements))
	for i, r := range addDataEventReq.DataEvent.Measurements {
		readings[i] = dtos.ToMeasurementModel(r)
	}

	tags := make(map[string]interface{})
	for tag, value := range addDataEventReq.DataEvent.Tags {
		tags[tag] = value
	}

	return models.DataEvent{
		Id:          addDataEventReq.DataEvent.Id,
		DeviceName:  addDataEventReq.DataEvent.DeviceName,
		ProfileName: addDataEventReq.DataEvent.ProfileName,
		SourceName:  addDataEventReq.DataEvent.SourceName,
		Origin:      addDataEventReq.DataEvent.Origin,
		Measurements:    readings,
		Tags:        tags,
	}
}
