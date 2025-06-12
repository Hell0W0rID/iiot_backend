//
// Copyright (c) 2022 One Track Consulting
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

//go:build include_nats_messaging

package interfaces

import "github.com/nats-io/nats.go"

// Connection provides an interface over basic *nats.Conn methods that we need to interact with the broker
type Connection interface {
	// QueueSubscribe subscribes to a NATS subject, equivalent to default Subscribe if queuegroup not supplied.
	QueueSubscribe(string, string, nats.MsgHandler) (*nats.EventSubscription, error)
	// PublishMsg sends the provided NATS message to the broker.
	PublishMsg(*nats.Msg) error
	// Drain will end all active subscription interest and attempt to wait for in-flight messages to process before closing.
	Drain() error
}
