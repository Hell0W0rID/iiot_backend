/*******************************************************************************
 *******************************************************************************/

package dtos

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
)

// SystemDataEvent defines the data for a system event
type SystemDataEvent struct {
	common.Versionable `json:",inline"`
	Type               string            `json:"type"`
	Action             string            `json:"action"`
	Source             string            `json:"source"`
	Owner              string            `json:"owner"`
	Tags               map[string]string `json:"tags"`
	Details            any               `json:"details"`
	Timestamp          int64             `json:"timestamp"`
}

// NewSystemDataEvent creates a new SystemDataEvent for the specified data
func NewSystemDataEvent(eventType, action, source, owner string, tags map[string]string, details any) SystemDataEvent {
	return SystemDataEvent{
		Versionable: common.NewVersionable(),
		Type:        eventType,
		Action:      action,
		Source:      source,
		Owner:       owner,
		Tags:        tags,
		Details:     details,
		Timestamp:   time.Now().UnixNano(),
	}
}

// DecodeDetails decodes the details (any type) into the passed in object
func (s *SystemDataEvent) DecodeDetails(details any) error {
	if s.Details == nil {
		return errors.New("unable to decode System DataEvent details: Details are nil")
	}

	// Must encode the details to JSON since if the target SystemDataEvent was decoded from JSON the details are
	// captured in a map[string]interface{}.
	data, err := json.Marshal(s.Details)
	if err != nil {
		return fmt.Errorf("unable to encode System DataEvent details to JSON: %s", err.Error())
	}

	err = json.Unmarshal(data, details)
	if err != nil {
		return fmt.Errorf("unable to decode System DataEvent details from JSON: %s", err.Error())
	}

	return nil
}
