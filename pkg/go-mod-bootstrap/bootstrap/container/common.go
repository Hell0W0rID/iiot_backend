//
//
// SPDX-License-Identifier: Apache-2.0

package container

import (
	"iiot-backend/pkg/go-mod-bootstrap/di"
	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
)

// CommonClientName contains the name of the CommonClient instance in the DIC.
var CommonClientName = di.TypeInstanceToName((*interfaces.CommonClient)(nil))

// CommonClientFrom helper function queries the DIC and returns the CommonClient instance.
func CommonClientFrom(get di.Get) interfaces.CommonClient {
	client, ok := get(CommonClientName).(interfaces.CommonClient)
	if !ok {
		return nil
	}

	return client
}
