//
//
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"encoding/json"

	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

type ScheduleActionRecord struct {
	Id          string
	JobName     string
	Action      ScheduleAction
	Status      ScheduleActionRecordStatus
	ScheduledAt int64
	Created     int64
}

// ScheduleActionRecordStatus indicates the most recent success/failure of a given schedule action attempt or a missed record.
type ScheduleActionRecordStatus string

func (scheduleActionRecord *ScheduleActionRecord) UnmarshalJSON(b []byte) error {
	var alias struct {
		Id          string
		JobName     string
		Action      any
		Status      ScheduleActionRecordStatus
		ScheduledAt int64
		Created     int64
	}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal ScheduleActionRecord.", err)
	}

	action, err := instantiateScheduleAction(alias.Action)
	if err != nil {
		return errors.NewCommonIIOTWrapper(err)
	}

	*scheduleActionRecord = ScheduleActionRecord{
		Id:          alias.Id,
		JobName:     alias.JobName,
		Action:      action,
		Status:      alias.Status,
		ScheduledAt: alias.ScheduledAt,
		Created:     alias.Created,
	}
	return nil
}
