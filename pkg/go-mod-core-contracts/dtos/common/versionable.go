//
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"iiot-backend/pkg/go-mod-core-contracts/common"
)

// Versionable shows the API version in DTOs
type Versionable struct {
	ApiVersion string `json:"apiVersion,omitempty"`
}

// NewVersionable creates versionable with the latest API version
func NewVersionable() Versionable {
	return Versionable{ApiVersion: common.ApiVersion}
}