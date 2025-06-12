/*******************************************************************************
 *******************************************************************************/

package container

import (
	"context"

	"iiot-backend/pkg/go-mod-bootstrap/di"
)

// CancelFuncName contains the name of the context.CancelFunc in the DIC.
var CancelFuncName = di.TypeInstanceToName((*context.CancelFunc)(nil))

// CancelFuncFrom helper function queries the DIC and returns the context.CancelFunc.
func CancelFuncFrom(get di.Get) context.CancelFunc {
	cancelFunc, ok := get(CancelFuncName).(context.CancelFunc)
	if !ok {
		return nil
	}

	return cancelFunc
}
