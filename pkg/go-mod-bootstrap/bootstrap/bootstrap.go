//
// Copyright (C) 2021-2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package bootstrap

import (
	"context"
	"sync"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/interfaces"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/startup"
	"iiot-backend/pkg/go-mod-bootstrap/config"
	"iiot-backend/pkg/go-mod-bootstrap/di"
)

// Bootstrap contains references to dependencies required by the bootstrap implementation.
type Bootstrap struct {
	serviceName    string
	serviceVersion string
	configDir      string
	configFile     string
	configProfile  string
	serviceType    config.ServiceType
}

// NewBootstrap is a factory method that returns an initialized Bootstrap receiver struct.
func NewBootstrap(serviceName, serviceVersion, configDir, configFile, configProfile string, serviceType config.ServiceType) *Bootstrap {
	return &Bootstrap{
		serviceName:    serviceName,
		serviceVersion: serviceVersion,
		configDir:      configDir,
		configFile:     configFile,
		configProfile:  configProfile,
		serviceType:    serviceType,
	}
}

// RunAndReturnWaitGroup bootstraps the service and returns a WaitGroup reference
func (b *Bootstrap) RunAndReturnWaitGroup(
	ctx context.Context,
	cancel context.CancelFunc,
	serviceKey string,
	handlers []interfaces.BootstrapHandler) (*sync.WaitGroup, bool, error) {

	startupTimer := startup.NewStartUpTimer(serviceKey)

	dic := di.NewContainer(di.ServiceConstructorMap{
		"serviceName":    func(get di.Get) interface{} { return b.serviceName },
		"serviceVersion": func(get di.Get) interface{} { return b.serviceVersion },
	})

	wg := &sync.WaitGroup{}
	
	// Execute all bootstrap handlers
	for _, handler := range handlers {
		if !handler.BootstrapHandler(ctx, wg, startupTimer, dic) {
			return nil, false, nil
		}
	}

	startupTimer.SeedRequired()

	return wg, true, nil
}

// Run bootstraps the service and waits until shutdown
func (b *Bootstrap) Run(
	ctx context.Context,
	cancel context.CancelFunc,
	serviceKey string,
	handlers []interfaces.BootstrapHandler) error {

	wg, success, err := b.RunAndReturnWaitGroup(ctx, cancel, serviceKey, handlers)
	if err != nil {
		return err
	}
	
	if !success {
		return nil
	}

	// Wait for all goroutines to complete
	wg.Wait()
	
	return nil
}