/********************************************************************************
 *******************************************************************************/

package mqtt

import (
	"fmt"
)

const (
	// Different Client operations.
	PublishOperation   = "Publish"
	SubscribeOperation = "Subscribe"
	ConnectOperation   = "Connect"
)

// TimeoutErr defines an error representing operations which have not completed and surpassed the allowed wait time.
type TimeoutErr struct {
	operation string
	message   string
}

func (te TimeoutErr) Error() string {
	return fmt.Sprintf("Timeout occured while performing a '%s' operation: %s", te.operation, te.message)
}

// NewTimeoutError creates a new TimeoutErr.
func NewTimeoutError(operation string, message string) TimeoutErr {
	return TimeoutErr{
		operation: operation,
		message:   message,
	}
}

// OperationErr defines an error representing operations which have failed.
type OperationErr struct {
	operation string
	message   string
}

func (oe OperationErr) Error() string {
	return fmt.Sprintf("An error occured while performing a '%s' operation: %s", oe.operation, oe.message)
}

// NewOperationErr creates a new OperationErr
func NewOperationErr(operation string, message string) OperationErr {
	return OperationErr{
		operation: operation,
		message:   message,
	}
}
