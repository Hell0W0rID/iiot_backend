/*******************************************************************************
 *******************************************************************************/

package handlers

import (
	"context"
	"fmt"
	"sync"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/container"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/startup"
	"iiot-backend/pkg/go-mod-bootstrap/di"
)

// StartMessage contains references to dependencies required by the start message handler.
type StartMessage struct {
	serviceKey string
	version    string
}

// NewStartMessage is a factory method that returns an initialized StartMessage receiver struct.
func NewStartMessage(serviceKey, version string) *StartMessage {
	return &StartMessage{
		serviceKey: serviceKey,
		version:    version,
	}
}

// BootstrapHandler fulfills the BootstrapHandler contract.  It creates no go routines.  It logs a "standard" set of
// messages when the service first starts up successfully.
func (h StartMessage) BootstrapHandler(
	_ context.Context,
	_ *sync.WaitGroup,
	startupTimer startup.Timer,
	dic *di.Container) bool {

	lc := container.LoggerClientFrom(dic.Get)
	lc.Info("Service dependencies resolved...")
	lc.Info(fmt.Sprintf("Starting %s %s ", h.serviceKey, h.version))

	bootstrapConfig := container.ConfigurationFrom(dic.Get).GetBootstrap()
	if len(bootstrapConfig.Service.StartupMsg) > 0 {
		lc.Info(bootstrapConfig.Service.StartupMsg)
	}

	lc.Info("Service started in: " + startupTimer.SinceAsString())

	return true
}
