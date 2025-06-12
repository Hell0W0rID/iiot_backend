package models

import (
	"time"
)

// Rule represents a rule in the rules engine
type Rule struct {
	ID          string                 `json:"id" db:"id"`
	Name        string                 `json:"name" db:"name"`
	Description string                 `json:"description" db:"description"`
	Enabled     bool                   `json:"enabled" db:"enabled"`
	Priority    int                    `json:"priority" db:"priority"`
	Conditions  []RuleCondition        `json:"conditions" db:"conditions"`
	Actions     []RuleAction           `json:"actions" db:"actions"`
	Tags        map[string]string      `json:"tags" db:"tags"`
	Created     time.Time              `json:"created" db:"created"`
	Modified    time.Time              `json:"modified" db:"modified"`
}

// RuleCondition represents a condition in a rule
type RuleCondition struct {
	Type       string                 `json:"type"`
	Device     string                 `json:"device"`
	Resource   string                 `json:"resource"`
	Operator   string                 `json:"operator"`
	Value      interface{}            `json:"value"`
	LogicalOp  string                 `json:"logicalOp"`
	Expression string                 `json:"expression"`
	Metadata   map[string]interface{} `json:"metadata"`
}

// RuleAction represents an action in a rule
type RuleAction struct {
	Type       string                 `json:"type"`
	Target     string                 `json:"target"`
	Command    string                 `json:"command"`
	Parameters map[string]interface{} `json:"parameters"`
	Template   string                 `json:"template"`
	Enabled    bool                   `json:"enabled"`
}

// Pipeline represents a data processing pipeline
type Pipeline struct {
	ID          string                 `json:"id" db:"id"`
	Name        string                 `json:"name" db:"name"`
	Description string                 `json:"description" db:"description"`
	Enabled     bool                   `json:"enabled" db:"enabled"`
	Functions   []PipelineFunction     `json:"functions" db:"functions"`
	Triggers    []PipelineTrigger      `json:"triggers" db:"triggers"`
	Targets     []PipelineTarget       `json:"targets" db:"targets"`
	Tags        map[string]string      `json:"tags" db:"tags"`
	Created     time.Time              `json:"created" db:"created"`
	Modified    time.Time              `json:"modified" db:"modified"`
}

// PipelineFunction represents a function in a pipeline
type PipelineFunction struct {
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	Config     map[string]interface{} `json:"config"`
	Order      int                    `json:"order"`
	Enabled    bool                   `json:"enabled"`
}

// PipelineTrigger represents a trigger for a pipeline
type PipelineTrigger struct {
	Type       string                 `json:"type"`
	Source     string                 `json:"source"`
	Filter     map[string]interface{} `json:"filter"`
	Config     map[string]interface{} `json:"config"`
}

// PipelineTarget represents a target for a pipeline
type PipelineTarget struct {
	Type       string                 `json:"type"`
	Name       string                 `json:"name"`
	Config     map[string]interface{} `json:"config"`
	Enabled    bool                   `json:"enabled"`
}

// RuleRequest represents a request to create/update a rule
type RuleRequest struct {
	Name        string                 `json:"name" validate:"required"`
	Description string                 `json:"description"`
	Enabled     bool                   `json:"enabled"`
	Priority    int                    `json:"priority"`
	Conditions  []RuleCondition        `json:"conditions" validate:"required"`
	Actions     []RuleAction           `json:"actions" validate:"required"`
	Tags        map[string]string      `json:"tags"`
}

// PipelineRequest represents a request to create/update a pipeline
type PipelineRequest struct {
	Name        string                 `json:"name" validate:"required"`
	Description string                 `json:"description"`
	Enabled     bool                   `json:"enabled"`
	Functions   []PipelineFunction     `json:"functions"`
	Triggers    []PipelineTrigger      `json:"triggers" validate:"required"`
	Targets     []PipelineTarget       `json:"targets" validate:"required"`
	Tags        map[string]string      `json:"tags"`
}

// RuleExecution represents the execution of a rule
type RuleExecution struct {
	RuleID      string    `json:"ruleId"`
	RuleName    string    `json:"ruleName"`
	ExecutionID string    `json:"executionId"`
	Status      string    `json:"status"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Duration    int64     `json:"duration"`
	Result      string    `json:"result"`
	Error       string    `json:"error,omitempty"`
}
