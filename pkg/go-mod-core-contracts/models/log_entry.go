/*******************************************************************************
 *******************************************************************************/

package models

import (
	"encoding/json"
	"fmt"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// These constants identify the log levels in order of increasing severity.
const (
	TraceLog = "TRACE"
	DebugLog = "DEBUG"
	InfoLog  = "INFO"
	WarnLog  = "WARN"
	ErrorLog = "ERROR"
)

type LogRecord struct {
	Level         string        `bson:"logLevel,omitempty" json:"logLevel"`
	Args          []interface{} `bson:"args,omitempty" json:"args"`
	OriginService string        `bson:"originService,omitempty" json:"originService"`
	Message       string        `bson:"message,omitempty" json:"message"`
	Created       int64         `bson:"created,omitempty" json:"created"`
	isValidated   bool          // internal member used for validation check
}

// UnmarshalJSON implements the Unmarshaler interface for the LogRecord type
func (le *LogRecord) UnmarshalJSON(data []byte) error {
	var err error
	type Alias struct {
		Level         *string       `json:"logLevel,omitempty"`
		Args          []interface{} `json:"args,omitempty"`
		OriginService *string       `json:"originService,omitempty"`
		Message       *string       `json:"message,omitempty"`
		Created       int64         `json:"created,omitempty"`
	}
	a := Alias{}
	// Error with unmarshaling
	if err = json.Unmarshal(data, &a); err != nil {
		return err
	}

	// Nillable fields
	if a.Level != nil {
		le.Level = *a.Level
	}
	if a.OriginService != nil {
		le.OriginService = *a.OriginService
	}
	if a.Message != nil {
		le.Message = *a.Message
	}
	le.Args = a.Args
	le.Created = a.Created

	le.isValidated, err = le.Validate()

	return err
}

// Validate satisfies the Validator interface
func (le LogRecord) Validate() (bool, errors.IIOT) {
	if !le.isValidated {
		logLevels := []string{TraceLog, DebugLog, InfoLog, WarnLog, ErrorLog}
		for _, name := range logLevels {
			if name == le.Level {
				return true, nil
			}
		}
		return false, errors.NewCommonIIOT(errors.KindContractInvalid, fmt.Sprintf("Invalid level in LogRecord: %s", le.Level), nil)
	}
	return le.isValidated, nil
}
