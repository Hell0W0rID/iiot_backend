//
// Copyright (c) 2022 One Track Consulting
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

//go:build include_nats_messaging

package nats

import (
	"strings"
)

const (
	StandardSeparator           = "/"
	StandardSingleLevelWildcard = "+"
	StandardMultiLevelWildcard  = "#"
	Separator                   = "."
	SingleLevelWildcard         = "*"
	MultiLevelWildcard          = ">"
)

var subjectReplacer = strings.NewReplacer(StandardSeparator, Separator, StandardSingleLevelWildcard, SingleLevelWildcard, StandardMultiLevelWildcard, MultiLevelWildcard)

// TopicToSubject formats an IIOT topic into a NATS subject
func TopicToSubject(topic string) string {
	return subjectReplacer.Replace(topic)
}

var topicReplacer = strings.NewReplacer(Separator, StandardSeparator, SingleLevelWildcard, StandardSingleLevelWildcard, MultiLevelWildcard, StandardMultiLevelWildcard)

// subjectToTopic formats a NATS subject into an IIOT topic
func subjectToTopic(subject string) string {
	return topicReplacer.Replace(subject)
}
