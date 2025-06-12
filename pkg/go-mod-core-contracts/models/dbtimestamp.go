//
//
// SPDX-License-Identifier: Apache-2.0

package models

type DBTimestamp struct {
	Created  int64 // Created is a timestamp indicating when the entity was created.
	Modified int64 // Modified is a timestamp indicating when the entity was last modified.
}
