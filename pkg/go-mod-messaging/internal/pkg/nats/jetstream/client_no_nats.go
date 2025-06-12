// Copyright (c) 2022 One Track Consulting
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

//go:build !include_nats_messaging

package jetstream

import (
	"fmt"

	"iiot-backend/pkg/go-mod-messaging/internal/pkg"
	"iiot-backend/pkg/go-mod-messaging/pkg/types"
)

// NewClient initializes creates a new client using NATS core
func NewClient(_ types.MessageBusConfig) (*pkg.NoopClient, error) {
	return nil, fmt.Errorf("to enable NATS message bus options please build using the flag include_nats_messaging")
}
