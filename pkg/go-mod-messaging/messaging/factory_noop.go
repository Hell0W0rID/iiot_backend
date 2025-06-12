//go:build no_messagebus
// +build no_messagebus

//
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

package messaging

import (
	"errors"

	"iiot-backend/pkg/go-mod-messaging/pkg/types"
)

// NewMessageClient is noop implementation when service doesn't need the message bus.
// This is need when this module is included in the common go-mod-bootstrap, but some service
// such as security service have no need for messaging.
func NewMessageClient(msgConfig types.MessageBusConfig) (MessageClient, error) {
	return nil, errors.New("messaging was disabled during build with the no_messagebus build flag")
}
