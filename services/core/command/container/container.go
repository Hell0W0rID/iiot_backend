/*******************************************************************************
 * Copyright 2017 Dell Inc.
 * Copyright (c) 2019-2023 Intel Corporation
 * Copyright (C) 2021 IOTech Ltd
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

package container

import (
	"iiot-backend/pkg/go-mod-bootstrap/di"
	"iiot-backend/services/core/command/config"
)

// ConfigurationName contains the name of command service's config.ConfigurationStruct implementation in the DIC.
var ConfigurationName = di.TypeInstanceToName((*config.ConfigurationStruct)(nil))

// ConfigurationFrom helper function queries the DIC and returns command service's config.ConfigurationStruct implementation.
func ConfigurationFrom(get di.Get) *config.ConfigurationStruct {
	return get(ConfigurationName).(*config.ConfigurationStruct)
}