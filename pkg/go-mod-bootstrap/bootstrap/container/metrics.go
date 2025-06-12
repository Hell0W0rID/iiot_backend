//
//
//
//
// Unless required by applicable law or agreed to in writing, software

package container

import (
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/interfaces"
	"iiot-backend/pkg/go-mod-bootstrap/di"
)

// MetricsManagerInterfaceName contains the name of the metrics.Manager implementation in the DIC.
var MetricsManagerInterfaceName = di.TypeInstanceToName((*interfaces.MetricsManager)(nil))

// MetricsManagerFrom helper function queries the DIC and returns the metrics.Manager implementation.
func MetricsManagerFrom(get di.Get) interfaces.MetricsManager {
	manager, ok := get(MetricsManagerInterfaceName).(interfaces.MetricsManager)
	if !ok {
		return nil
	}

	return manager
}
