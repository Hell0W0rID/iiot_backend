/*******************************************************************************
 * Copyright 2021-2025 IOTech Ltd
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package interfaces

import (
	"context"

	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
)

// CommonClient defines the interface for interactions with the common endpoints.
type CommonClient interface {
	// Configuration obtains configuration information from the target service.
	Configuration(ctx context.Context) (common.ConfigResponse, error)
	// Metrics obtains metrics information from the target service.
	Metrics(ctx context.Context) (common.MetricsResponse, error)
	// Ping tests whether the service is working
	Ping(ctx context.Context) (common.PingResponse, error)
	// Version obtains version information from the target service.
	Version(ctx context.Context) (common.VersionResponse, error)
}