//
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/requests"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// ScheduleJobClient defines the interface for interactions with the ScheduleJob endpoint on the IIOT Foundry support-scheduler service.
type ScheduleJobClient interface {
	// Add adds new schedule jobs.
	Add(ctx context.Context, reqs []requests.AddScheduleJobRequest) ([]common.BaseWithIdResponse, errors.IIOT)
	// Update updates schedule jobs.
	Update(ctx context.Context, reqs []requests.UpdateScheduleJobRequest) ([]common.BaseResponse, errors.IIOT)
	// AllScheduleJobs returns all schedule jobs.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	AllScheduleJobs(ctx context.Context, labels []string, offset int, limit int) (responses.MultiScheduleJobsResponse, errors.IIOT)
	// ScheduleJobByName returns a schedule job by name.
	ScheduleJobByName(ctx context.Context, name string) (responses.ScheduleJobResponse, errors.IIOT)
	// DeleteScheduleJobByName deletes a schedule job by name.
	DeleteScheduleJobByName(ctx context.Context, name string) (common.BaseResponse, errors.IIOT)
	// TriggerScheduleJobByName triggers a schedule job by name.
	TriggerScheduleJobByName(ctx context.Context, name string) (common.BaseResponse, errors.IIOT)
}
