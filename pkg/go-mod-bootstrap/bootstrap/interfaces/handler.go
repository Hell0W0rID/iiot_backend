/*******************************************************************************
 *******************************************************************************/

package interfaces

import (
	"context"
	"sync"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/startup"
	"iiot-backend/pkg/go-mod-bootstrap/di"
)

// BootstrapHandler defines the contract each bootstrap handler must fulfill.  Implementation returns true if the
// handler completed successfully, false if it did not.
type BootstrapHandler func(
	ctx context.Context,
	wg *sync.WaitGroup,
	startupTimer startup.Timer,
	dic *di.Container) (success bool)
