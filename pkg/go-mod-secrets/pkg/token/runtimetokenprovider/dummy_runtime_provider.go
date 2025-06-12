//go:build non_delayedstart
// +build non_delayedstart

//
//
//
//
//
// SPDX-License-Identifier: Apache-2.0'
//

package runtimetokenprovider

import (
	"context"
	"fmt"

	"iiot-backend/pkg/go-mod-core-contracts/clients/logger"
	"iiot-backend/pkg/go-mod-secrets/pkg/types"
)

type runtimetokenprovider struct{}

func NewRuntimeTokenProvider(_ context.Context, _ logger.LoggerClient,
	_ types.RuntimeTokenProviderInfo) RuntimeTokenProvider {
	return &runtimetokenprovider{}
}

func (p *runtimetokenprovider) GetRawToken(serviceKey string) (string, error) {
	return "", fmt.Errorf("wrong build: RuntimeTokenProvider is not available. " +
		"Build without \"-tags non_delayedstart\" on the go build command line to enable runtime support for this feature.")
}
