//
// Copyright (C) 2021-2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"
	"sync"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/startup"
	"iiot-backend/pkg/go-mod-bootstrap/di"
)

// BootstrapHandler defines the contract for bootstrap handlers
type BootstrapHandler func(ctx context.Context, wg *sync.WaitGroup, startupTimer startup.Timer, dic *di.Container) bool