//
//
// SPDX-License-Identifier: Apache-2.0

package container

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"

	"iiot-backend/pkg/go-mod-bootstrap/di"
)

// ExternalMQTTMessagingClientName contains the name of the external messaging client instance in the DIC.
var ExternalMQTTMessagingClientName = di.TypeInstanceToName((*mqtt.Client)(nil))

// ExternalMQTTMessagingClientFrom helper function queries the DIC and returns the external messaging client.
func ExternalMQTTMessagingClientFrom(get di.Get) mqtt.Client {
	client, ok := get(ExternalMQTTMessagingClientName).(mqtt.Client)
	if !ok {
		return nil
	}

	return client
}
