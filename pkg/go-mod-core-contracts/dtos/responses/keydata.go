//
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
)

// KeyDataResponse defines the Response Content for GET KeyData DTOs.
type KeyDataResponse struct {
	dtoCommon.BaseResponse `json:",inline"`
	KeyData                dtos.KeyData `json:"keyData"`
}

func NewKeyDataResponse(requestId string, message string, statusCode int, keyData dtos.KeyData) KeyDataResponse {
	return KeyDataResponse{
		BaseResponse: dtoCommon.NewBaseResponse(requestId, message, statusCode),
		KeyData:      keyData,
	}
}
