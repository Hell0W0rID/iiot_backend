//
//
//
//
// Unless required by applicable law or agreed to in writing, software

package container

import (
	"iiot-backend/pkg/go-mod-bootstrap/di"
	"iiot-backend/pkg/go-mod-messaging/messaging"
)

// MessagingClientName contains the name of the messaging client instance in the DIC.
var MessagingClientName = di.TypeInstanceToName((*messaging.MessageClient)(nil))

// MessagingClientFrom helper function queries the DIC and returns the messaging client.
func MessagingClientFrom(get di.Get) messaging.MessageClient {
	client, ok := get(MessagingClientName).(messaging.MessageClient)
	if !ok {
		return nil
	}

	return client
}
