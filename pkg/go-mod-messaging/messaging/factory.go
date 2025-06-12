//go:build !no_messagebus
// +build !no_messagebus

//
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

package messaging

import (
	"fmt"
	"strings"

	"iiot-backend/pkg/go-mod-messaging/internal/pkg/mqtt"
	"iiot-backend/pkg/go-mod-messaging/internal/pkg/nats"
	"iiot-backend/pkg/go-mod-messaging/internal/pkg/nats/jetstream"
	"iiot-backend/pkg/go-mod-messaging/pkg/types"
)

const (
	// MQTT messaging implementation
	MQTT = "mqtt"

	// NatsCore implementation
	NatsCore = "nats-core"

	// NatsJetStream implementation
	NatsJetStream = "nats-jetstream"
)

// NewMessageClient is a factory function to instantiate different message client depending on
// the "Type" from the configuration
func NewMessageClient(msgConfig types.MessageBusConfig) (MessageClient, error) {

	if msgConfig.Broker.IsHostInfoEmpty() {
		return nil, fmt.Errorf("unable to create messageClient: Broker info not set")
	}

	switch lowerMsgType := strings.ToLower(msgConfig.Type); lowerMsgType {
	case MQTT:
		return mqtt.NewMQTTClient(msgConfig)
	case NatsCore:
		return nats.NewClient(msgConfig)
	case NatsJetStream:
		return jetstream.NewClient(msgConfig)
	default:
		return nil, fmt.Errorf("unknown message type '%s' requested", msgConfig.Type)
	}
}
