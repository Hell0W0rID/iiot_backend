//
// Copyright (c) 2022 One Track Consulting
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

package pkg

import (
	"fmt"
	"time"

	"iiot-backend/pkg/go-mod-messaging/pkg/types"
)

type NoopClient struct{}

func (n NoopClient) Request(message types.MessageEnvelope, targetServiceName string, requestTopic string, timeout time.Duration) (*types.MessageEnvelope, error) {
	panic("implement me")
}

func (n NoopClient) Unsubscribe(topics ...string) error {
	panic("implement me")
}

func (n NoopClient) Connect() error {
	panic("implement me")
}

func (n NoopClient) Publish(message types.MessageEnvelope, topic string) error {
	panic("implement me")
}

func (n NoopClient) PublishWithSizeLimit(message types.MessageEnvelope, topic string, limit int64) error {
	panic("implement me")
}

func (n NoopClient) Subscribe(topics []types.TopicChannel, messageErrors chan error) error {
	panic("implement me")
}

func (n NoopClient) Disconnect() error {
	panic("implement me")
}

func (n NoopClient) PublishBinaryData(data []byte, topic string) error {
	return fmt.Errorf("not supported PublishBinaryData func")
}
func (n NoopClient) SubscribeBinaryData(topics []types.TopicChannel, messageErrors chan error) error {
	return fmt.Errorf("not supported SubscribeBinaryData func")
}
