// Copyright (C) 2021-2025 IOTech Ltd
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package command

import (
	"iiot-backend/pkg/go-mod-core-contracts/version"
	commandController "iiot-backend/services/core/command/controller/http"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/controller"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/handlers"
	"iiot-backend/pkg/go-mod-bootstrap/di"
	"iiot-backend/pkg/go-mod-core-contracts/common"

	"github.com/labstack/echo/v4"
)

func LoadRestRoutes(r *echo.Echo, dic *di.Container, serviceName string) {
	authenticationHook := handlers.AutoConfigAuthenticationFunc(dic)

	// Common
	_ = controller.NewCommonController(dic, r, serviceName, version.CoreCommandVersion)

	// Command
	cmd := commandController.NewCommandController(dic)
	r.GET(common.ApiAllDeviceRoute, cmd.AllCommands, authenticationHook)
	r.GET(common.ApiDeviceByNameRoute, cmd.CommandsByDeviceName, authenticationHook)
	r.GET(common.ApiDeviceNameCommandNameRoute, cmd.IssueGetCommandByName, authenticationHook)
	r.PUT(common.ApiDeviceNameCommandNameRoute, cmd.IssueSetCommandByName, authenticationHook)
}