//
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

package pkg

import (
	"fmt"
	"strings"
	"time"

	"iiot-backend/pkg/go-mod-messaging/pkg/types"
	"github.com/google/uuid"
)

// DoRequest publishes a request containing a RequestID to the specified topic,
// then subscribes to a response topic which contains the RequestID. Once the response is received, the
// response topic is unsubscribed and the response data is returned. If no response is received within
// the timeout period, a timed out  error returned.
func DoRequest(
	subscribe func(topics []types.TopicChannel, messageErrors chan error) error,
	unsubscribe func(topics ...string) error,
	publish func(message types.MessageEnvelope, topic string) error,
	requestMessage types.MessageEnvelope,
	requestTopic string,
	responseTopicPrefix string,
	requestTimeout time.Duration) (*types.MessageEnvelope, error) {
	if len(strings.TrimSpace(requestMessage.RequestID)) == 0 {
		requestMessage.RequestID = uuid.NewString()
	}

	// Format of response topic is <prefix>/<request-id>
	responseTopic := strings.Join([]string{responseTopicPrefix, requestMessage.RequestID}, "/")

	errs := make(chan error, 1)
	messages := make(chan types.MessageEnvelope, 1)
	responseTopicChan := types.TopicChannel{
		Topic:    responseTopic,
		Messages: messages,
	}

	// Must create the subscription first so that it is in place when the request is handled and response published back
	err := subscribe([]types.TopicChannel{responseTopicChan}, errs)
	if err != nil {
		return nil, fmt.Errorf("unable to create response subscription: %w", err)
	}

	defer func() {
		_ = unsubscribe(responseTopicChan.Topic)
		close(errs)
		close(messages)
	}()

	err = publish(requestMessage, requestTopic)
	if err != nil {
		return nil, fmt.Errorf("unable to create publish request to %s: %w", requestTopic, err)
	}

	timer := time.NewTimer(requestTimeout)
	defer timer.Stop()

	select {
	case <-timer.C:
		return nil, fmt.Errorf("timed out waiting for response on %s topic", responseTopicChan.Topic)

	case err = <-errs:
		return nil, fmt.Errorf("encountered error waiting for response to %s: %w", requestTopic, err)

	case responseMessage := <-messages:
		return &responseMessage, nil
	}
}
