//
// Copyright (c) 2022 One Track Consulting
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

//go:build include_nats_messaging

package nats

import (
	"iiot-backend/pkg/go-mod-messaging/internal/pkg/nats/interfaces"
)

// ConnectNats is a function that can be provided to determine the underlying connection for clients to use.
type ConnectNats func(ClientConfig) (interfaces.Connection, error)
