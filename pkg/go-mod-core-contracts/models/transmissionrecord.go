//
//
// SPDX-License-Identifier: Apache-2.0

package models

type DeliveryRecord struct {
	Status   DeliveryStatus
	Response string
	Sent     int64
}
