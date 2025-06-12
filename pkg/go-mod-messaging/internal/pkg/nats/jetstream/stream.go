//
// Copyright (c) 2022 One Track Consulting
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

//go:build include_nats_messaging

package jetstream

import (
	"strings"

	localnats "iiot-backend/pkg/go-mod-messaging/internal/pkg/nats"
)

const (
	streamSeparator = "_"
	streamWildcard  = ""
)

var streamReplacer = strings.NewReplacer(localnats.Separator, streamSeparator, localnats.SingleLevelWildcard, streamWildcard, localnats.MultiLevelWildcard, streamWildcard)

func subjectToStreamName(subject string) string {
	return strings.TrimSuffix(streamReplacer.Replace(subject), streamSeparator)
}
