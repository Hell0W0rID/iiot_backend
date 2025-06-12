//
// Copyright (c) 2022 One Track Consulting
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

//go:build include_nats_messaging

package interfaces

import (
	"iiot-backend/pkg/go-mod-messaging/pkg/types"
	"github.com/nats-io/nats.go"
)

// MarshallerUnmarshaller provides translation between NATS and IIOT formats
type MarshallerUnmarshaller interface {
	Marshal(v types.MessageEnvelope, publishTopic string) (*nats.Msg, error)
	Unmarshal(msg *nats.Msg, v *types.MessageEnvelope) error
}
