/*******************************************************************************
 *******************************************************************************/

package handlers

import (
	"context"
	"sync"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/startup"
	"iiot-backend/pkg/go-mod-bootstrap/di"
)

// httpServer defines the contract used to determine whether or not the http httpServer is running.
type httpServer interface {
	IsRunning() bool
}

// Ready contains references to dependencies required by the testing implementation.
type Ready struct {
	httpServer httpServer
	stream     chan<- bool
}

// NewReady is a factory method that returns an initialized Ready receiver struct.
func NewReady(httpServer httpServer, stream chan<- bool) *Ready {
	return &Ready{
		httpServer: httpServer,
		stream:     stream,
	}
}

// BootstrapHandler fulfills the BootstrapHandler contract.  During normal production execution, a nil stream
// will be supplied.  A non-nil stream indicates we're running within the test runner context and that we should
// wait for the httpServer to start running before sending confirmation over the stream.  If the httpServer doesn't
// start running within the defined startup time, no confirmation is sent over the stream and the application
// bootstrapping is aborted.
func (r *Ready) BootstrapHandler(
	_ context.Context,
	_ *sync.WaitGroup,
	startupTimer startup.Timer,
	_ *di.Container) bool {

	if r.stream != nil {
		for startupTimer.HasNotElapsed() {
			if r.httpServer.IsRunning() {
				r.stream <- true
				return true
			}
			startupTimer.SleepForInterval()
		}
		return false
	}
	return true
}
