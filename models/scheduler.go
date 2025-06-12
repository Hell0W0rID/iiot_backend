package models

import (
	"time"
)

// Interval represents a scheduled interval
type Interval struct {
	ID       string    `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Start    string    `json:"start" db:"start_time"`
	End      string    `json:"end" db:"end_time"`
	Interval string    `json:"interval" db:"interval_time"`
	RunOnce  bool      `json:"runOnce" db:"run_once"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
}

// IntervalAction represents an action to be executed on an interval
type IntervalAction struct {
	ID           string    `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	IntervalName string    `json:"intervalName" db:"interval_name"`
	Protocol     string    `json:"protocol" db:"protocol"`
	Host         string    `json:"host" db:"host"`
	Port         int       `json:"port" db:"port"`
	Path         string    `json:"path" db:"path"`
	Parameters   string    `json:"parameters" db:"parameters"`
	HTTPMethod   string    `json:"httpMethod" db:"http_method"`
	Address      string    `json:"address" db:"address"`
	Publisher    string    `json:"publisher" db:"publisher"`
	Target       string    `json:"target" db:"target"`
	User         string    `json:"user" db:"user"`
	Password     string    `json:"password" db:"password"`
	Topic        string    `json:"topic" db:"topic"`
	Created      time.Time `json:"created" db:"created"`
	Modified     time.Time `json:"modified" db:"modified"`
}

// IntervalRequest represents a request to create/update an interval
type IntervalRequest struct {
	Name     string `json:"name" validate:"required"`
	Start    string `json:"start" validate:"required"`
	End      string `json:"end"`
	Interval string `json:"interval" validate:"required"`
	RunOnce  bool   `json:"runOnce"`
}

// IntervalActionRequest represents a request to create/update an interval action
type IntervalActionRequest struct {
	Name         string `json:"name" validate:"required"`
	IntervalName string `json:"intervalName" validate:"required"`
	Protocol     string `json:"protocol" validate:"required"`
	Host         string `json:"host" validate:"required"`
	Port         int    `json:"port" validate:"required"`
	Path         string `json:"path"`
	Parameters   string `json:"parameters"`
	HTTPMethod   string `json:"httpMethod"`
	Address      string `json:"address"`
	Publisher    string `json:"publisher"`
	Target       string `json:"target"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Topic        string `json:"topic"`
}

// ScheduleStatus represents the status of a scheduled task
type ScheduleStatus struct {
	IntervalName   string    `json:"intervalName"`
	LastExecution  time.Time `json:"lastExecution"`
	NextExecution  time.Time `json:"nextExecution"`
	ExecutionCount int64     `json:"executionCount"`
	Status         string    `json:"status"`
	Message        string    `json:"message"`
}
