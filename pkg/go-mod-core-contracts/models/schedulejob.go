//
//
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"encoding/json"
	"github.com/google/uuid"

	"iiot-backend/pkg/go-mod-core-contracts/common"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

type ScheduleJob struct {
	DBTimestamp
	Id                       string
	Name                     string
	Definition               ScheduleDef
	AutoTriggerMissedRecords bool
	Actions                  []ScheduleAction
	ServiceState               ServiceState
	Labels                   []string
	Properties               map[string]any
}

func (scheduleJob *ScheduleJob) UnmarshalJSON(b []byte) error {
	var alias struct {
		DBTimestamp
		Id                       string
		Name                     string
		Definition               any
		AutoTriggerMissedRecords bool
		Actions                  []any
		ServiceState               ServiceState
		Labels                   []string
		Properties               map[string]any
	}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal ScheduleJob.", err)
	}

	def, err := instantiateScheduleDef(alias.Definition)
	if err != nil {
		return errors.NewCommonIIOTWrapper(err)
	}

	actions := make([]ScheduleAction, len(alias.Actions))
	for i, a := range alias.Actions {
		action, err := instantiateScheduleAction(a)
		if err != nil {
			return errors.NewCommonIIOTWrapper(err)
		}
		actions[i] = action
	}

	*scheduleJob = ScheduleJob{
		DBTimestamp:              alias.DBTimestamp,
		Id:                       alias.Id,
		Name:                     alias.Name,
		Definition:               def,
		AutoTriggerMissedRecords: alias.AutoTriggerMissedRecords,
		Actions:                  actions,
		ServiceState:               alias.ServiceState,
		Labels:                   alias.Labels,
		Properties:               alias.Properties,
	}
	return nil
}

type ScheduleDef interface {
	GetBaseScheduleDef() BaseScheduleDef
}

// instantiateScheduleDef instantiate the interface to the corresponding schedule definition type
func instantiateScheduleDef(i any) (def ScheduleDef, err error) {
	d, err := json.Marshal(i)
	if err != nil {
		return def, errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to marshal ScheduleDef.", err)
	}
	return unmarshalScheduleDef(d)
}

func unmarshalScheduleDef(b []byte) (def ScheduleDef, err error) {
	var alias struct {
		Type string
	}
	if err = json.Unmarshal(b, &alias); err != nil {
		return def, errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal ScheduleDef.", err)
	}
	switch alias.Type {
	case common.DefInterval:
		var intervalDef IntervalScheduleDef
		if err = json.Unmarshal(b, &intervalDef); err != nil {
			return def, errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal INTERVAL ScheduleDef.", err)
		}
		def = intervalDef
	case common.DefCron:
		var cronDef CronScheduleDef
		if err = json.Unmarshal(b, &cronDef); err != nil {
			return def, errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal CRON ScheduleDef.", err)
		}
		def = cronDef
	default:
		return def, errors.NewCommonIIOT(errors.KindContractInvalid, "Unsupported schedule definition type", err)
	}
	return def, nil
}

type BaseScheduleDef struct {
	Type           ScheduleDefType
	StartTimestamp int64
	EndTimestamp   int64
}

type IntervalScheduleDef struct {
	BaseScheduleDef
	// Interval specifies the time interval between two consecutive executions
	Interval string
}

func (d IntervalScheduleDef) GetBaseScheduleDef() BaseScheduleDef {
	return d.BaseScheduleDef
}

type CronScheduleDef struct {
	BaseScheduleDef
	// Crontab is the cron expression
	Crontab string
}

func (c CronScheduleDef) GetBaseScheduleDef() BaseScheduleDef {
	return c.BaseScheduleDef
}

type ScheduleAction interface {
	GetBaseScheduleAction() BaseScheduleAction
	// WithEmptyPayloadAndId returns a copy of the ScheduleAction with empty payload and Id, which is used by ScheduleActionRecord to remove the payload and id before storing the record into database
	WithEmptyPayloadAndId() ScheduleAction
	// WithId returns a copy of the ScheduleAction with ID or generates a new ID if the ID is empty, which is used to identify the action and record in the database
	WithId(id string) ScheduleAction
}

// instantiateScheduleAction instantiate the interface to the corresponding schedule action type
func instantiateScheduleAction(i any) (action ScheduleAction, err error) {
	a, err := json.Marshal(i)
	if err != nil {
		return action, errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to marshal ScheduleAction.", err)
	}
	return UnmarshalScheduleAction(a)
}

func UnmarshalScheduleAction(b []byte) (action ScheduleAction, err error) {
	var alias struct {
		Type string
	}
	if err = json.Unmarshal(b, &alias); err != nil {
		return action, errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal ScheduleAction.", err)
	}
	switch alias.Type {
	case common.ActionIIOTMessageBus:
		var edgeXMessageBusAction IIOTMessageBusAction
		if err = json.Unmarshal(b, &edgeXMessageBusAction); err != nil {
			return action, errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal IIOTMESSAGEBUS ScheduleAction.", err)
		}
		action = edgeXMessageBusAction
	case common.ActionREST:
		var restAction RESTAction
		if err = json.Unmarshal(b, &restAction); err != nil {
			return action, errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal REST ScheduleAction.", err)
		}
		action = restAction
	case common.ActionDeviceControl:
		var deviceControlAction DeviceControlAction
		if err = json.Unmarshal(b, &deviceControlAction); err != nil {
			return action, errors.NewCommonIIOT(errors.KindContractInvalid, "Failed to unmarshal DEVICECONTROL ScheduleAction.", err)
		}
		action = deviceControlAction
	default:
		return action, errors.NewCommonIIOT(errors.KindContractInvalid, "Unsupported schedule action type", err)
	}
	return action, nil
}

type BaseScheduleAction struct {
	// Id is the identifier of the action, no need to be in the DTO
	Id          string
	Type        ScheduleActionType
	ContentType string
	Payload     []byte
}

type IIOTMessageBusAction struct {
	BaseScheduleAction
	Topic string
}

func (m IIOTMessageBusAction) GetBaseScheduleAction() BaseScheduleAction {
	return m.BaseScheduleAction
}
func (m IIOTMessageBusAction) WithEmptyPayloadAndId() ScheduleAction {
	m.Id = ""
	m.Payload = nil
	return m
}
func (m IIOTMessageBusAction) WithId(id string) ScheduleAction {
	if len(m.Id) == 0 {
		if id != "" {
			m.Id = id
		} else {
			m.Id = uuid.New().String()
		}
	}
	return m
}

type RESTAction struct {
	BaseScheduleAction
	Address         string
	Method          string
	InjectIIOTAuth bool
}

func (r RESTAction) GetBaseScheduleAction() BaseScheduleAction {
	return r.BaseScheduleAction
}
func (r RESTAction) WithEmptyPayloadAndId() ScheduleAction {
	r.Id = ""
	r.Payload = nil
	return r
}
func (r RESTAction) WithId(id string) ScheduleAction {
	if len(r.Id) == 0 {
		if id != "" {
			r.Id = id
		} else {
			r.Id = uuid.New().String()
		}
	}
	return r
}

type DeviceControlAction struct {
	BaseScheduleAction
	DeviceName string
	SourceName string
}

func (d DeviceControlAction) GetBaseScheduleAction() BaseScheduleAction {
	return d.BaseScheduleAction
}
func (d DeviceControlAction) WithEmptyPayloadAndId() ScheduleAction {
	d.Id = ""
	d.Payload = nil
	return d
}
func (d DeviceControlAction) WithId(id string) ScheduleAction {
	if len(d.Id) == 0 {
		if id != "" {
			d.Id = id
		} else {
			d.Id = uuid.New().String()
		}
	}
	return d
}

// ScheduleDefType is used to identify the schedule definition type, i.e., INTERVAL or CRON
type ScheduleDefType string

// ScheduleActionType is used to identify the schedule action type, i.e., IIOTMESSAGEBUS, REST, or DEVICECONTROL
type ScheduleActionType string
