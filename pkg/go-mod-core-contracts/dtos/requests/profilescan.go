//
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

	"iiot-backend/pkg/go-mod-core-contracts/common"
	dtoCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// ProfileScanRequest is the struct for requesting a profile for a specified device.
type ProfileScanRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	DeviceName            string `json:"deviceName" validate:"required"`
	ProfileName           string `json:"profileName,omitempty"`
	Options               any    `json:"options,omitempty"`
}

// Validate satisfies the Validator interface
func (request ProfileScanRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddDeviceActionRequest type
func (psr *ProfileScanRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		DeviceName  string
		ProfileName string
		Options     any
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*psr = ProfileScanRequest(alias)

	if err := psr.Validate(); err != nil {
		return err
	}

	return nil
}
