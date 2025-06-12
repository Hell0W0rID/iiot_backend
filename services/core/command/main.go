/*******************************************************************************
 * Copyright 2020 Dell Inc.
 * Copyright 2022-2023 IOTech Ltd.
 * Copyright 2023 Intel Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package command

import (
	"context"
	"sync"
	"time"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap"
	bootstrapContainer "iiot-backend/pkg/go-mod-bootstrap/bootstrap/container"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/flags"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/handlers"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/interfaces"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/startup"
	bootstrapConfig "iiot-backend/pkg/go-mod-bootstrap/config"
	"iiot-backend/pkg/go-mod-bootstrap/di"

	"iiot-backend/pkg/go-mod-core-contracts/version"
	"iiot-backend/services/core/command/config"
	"iiot-backend/services/core/command/container"
	"iiot-backend/services/core/command/controller/messaging"

	"iiot-backend/pkg/go-mod-core-contracts/common"

	"github.com/labstack/echo/v4"
)

func Main(ctx context.Context, cancel context.CancelFunc, router *echo.Echo, args []string) {
	startupTimer := startup.NewStartUpTimer(common.CoreCommandServiceName)

	// All common command-line flags have been moved to DefaultCommonFlags. Service specific flags can be added here,
	// by inserting service specific flag prior to call to commonFlags.Parse().
	// Example:
	// 		flags.FlagSet.StringVar(&myvar, "m", "", "Specify a ....")
	//      ....
	//      flags.Parse(args)
	//
	f := flags.New()
	f.Parse(args)

	configuration := &config.ConfigurationStruct{}
	dic := di.NewContainer(di.ServiceConstructorMap{
		container.ConfigurationName: func(get di.Get) interface{} {
			return configuration
		},
	})

	httpServer := handlers.NewHttpServer(router, true, common.CoreCommandServiceName)

	bootstrap.Run(
		ctx,
		cancel,
		f,
		common.CoreCommandServiceName,
		common.ConfigStemCore,
		configuration,
		startupTimer,
		dic,
		true,
		bootstrapConfig.ServiceTypeOther,
		[]interfaces.BootstrapHandler{
			handlers.NewClientsBootstrap().BootstrapHandler,
			MessagingBootstrapHandler,
			handlers.NewServiceMetrics(common.CoreCommandServiceName).BootstrapHandler, // Must be after Messaging
			NewBootstrap(router, common.CoreCommandServiceName).BootstrapHandler,
			httpServer.BootstrapHandler,
			handlers.NewStartMessage(common.CoreCommandServiceName, version.CoreCommandVersion).BootstrapHandler,
		})

	// code here!
}

// MessagingBootstrapHandler sets up the MessageBus and External MQTT connections as well as subscriptions
func MessagingBootstrapHandler(ctx context.Context, wg *sync.WaitGroup, startupTimer startup.Timer, dic *di.Container) bool {
	lc := bootstrapContainer.LoggingClientFrom(dic.Get)
	configuration := container.ConfigurationFrom(dic.Get)

	if len(configuration.Service.RequestTimeout) == 0 {
		lc.Error("Service.RequestTimeout found empty in service's configuration, missing common config? Use -cp or -cc flags for common config")
		return false
	}

	requestTimeout, err := time.ParseDuration(configuration.Service.RequestTimeout)
	if err != nil {
		lc.Errorf("Failed to parse Service.RequestTimeout configuration value: %v", err)
		return false
	}

	if configuration.ExternalMQTT.Enabled {
		if !handlers.NewExternalMQTT(messaging.OnConnectHandler(requestTimeout, dic)).BootstrapHandler(ctx, wg, startupTimer, dic) {
			return false
		}
	}

	if !handlers.MessagingBootstrapHandler(ctx, wg, startupTimer, dic) {
		return false
	}
	if err := messaging.SubscribeCommandRequests(ctx, requestTimeout, dic); err != nil {
		lc.Errorf("Failed to subscribe commands request from internal message bus, %v", err)
		return false
	}

	if err := messaging.SubscribeCommandQueryRequests(ctx, dic); err != nil {
		lc.Errorf("Failed to subscribe command query request from internal message bus, %v", err)
		return false
	}

	return true
}