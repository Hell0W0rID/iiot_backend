//
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
)

// ScheduleJobResponse defines the Response Content for GET ScheduleJob DTO.
type ScheduleJobResponse struct {
	common.BaseResponse `json:",inline"`
	ScheduleJob         dtos.ScheduleJob `json:"scheduleJob"`
}

func NewScheduleJobResponse(requestId string, message string, statusCode int, scheduleJob dtos.ScheduleJob) ScheduleJobResponse {
	return ScheduleJobResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		ScheduleJob:  scheduleJob,
	}
}

// MultiScheduleJobsResponse defines the Response Content for GET multiple ScheduleJob DTOs.
type MultiScheduleJobsResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	ScheduleJobs                      []dtos.ScheduleJob `json:"scheduleJobs"`
}

func NewMultiScheduleJobsResponse(requestId string, message string, statusCode int, totalCount uint32, scheduleJobs []dtos.ScheduleJob) MultiScheduleJobsResponse {
	return MultiScheduleJobsResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		ScheduleJobs:               scheduleJobs,
	}
}
