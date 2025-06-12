//
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"encoding/json"
	"os"

	"github.com/fxamacker/cbor/v2"

	"iiot-backend/pkg/go-mod-core-contracts/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// DataEventResponse defines the Response Content for GET event DTOs.
type DataEventResponse struct {
	dtoCommon.BaseResponse `json:",inline"`
	DataEvent                  dtos.DataEvent `json:"event"`
}

// MultiDataEventsResponse defines the Response Content for GET multiple event DTOs.
type MultiDataEventsResponse struct {
	dtoCommon.BaseWithTotalCountResponse `json:",inline"`
	DataEvents                               []dtos.DataEvent `json:"events"`
}

func NewDataEventResponse(requestId string, message string, statusCode int, event dtos.DataEvent) DataEventResponse {
	return DataEventResponse{
		BaseResponse: dtoCommon.NewBaseResponse(requestId, message, statusCode),
		DataEvent:        event,
	}
}

func NewMultiDataEventsResponse(requestId string, message string, statusCode int, totalCount uint32, events []dtos.DataEvent) MultiDataEventsResponse {
	return MultiDataEventsResponse{
		BaseWithTotalCountResponse: dtoCommon.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		DataEvents:                     events,
	}
}

func (e *DataEventResponse) Encode() ([]byte, string, error) {
	var encoding = e.GetEncodingContentType()
	var err error
	var encodedData []byte

	switch encoding {
	case common.ContentTypeCBOR:
		encodedData, err = cbor.Marshal(e)
		if err != nil {
			return nil, "", errors.NewCommonIIOT(errors.KindContractInvalid, "failed to encode DataEventResponse to CBOR", err)
		}
	case common.ContentTypeJSON:
		encodedData, err = json.Marshal(e)
		if err != nil {
			return nil, "", errors.NewCommonIIOT(errors.KindContractInvalid, "failed to encode DataEventResponse to JSON", err)
		}
	}

	return encodedData, encoding, nil
}

// GetEncodingContentType determines which content type should be used to encode and decode this object
func (e *DataEventResponse) GetEncodingContentType() string {
	if v := os.Getenv(common.EnvEncodeAllDataEvents); v == common.ValueTrue {
		return common.ContentTypeCBOR
	}
	for _, r := range e.DataEvent.Measurements {
		if r.ValueType == common.ValueTypeBinary {
			return common.ContentTypeCBOR
		}
	}

	return common.ContentTypeJSON
}
