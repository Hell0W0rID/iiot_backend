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
        "iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
)

// DeviceClient defines the interface for interactions with the Device endpoint on core-metadata service.
type DeviceClient interface {
        // AllDevices returns all devices. 
        AllDevices(ctx context.Context, labels []string, offset int, limit int) (responses.MultiDevicesResponse, error)
        // DeviceByName returns a device by name
        DeviceByName(ctx context.Context, name string) (responses.DeviceResponse, error)
        // DevicesByProfileName returns devices associated with the specified device profile name.
        DevicesByProfileName(ctx context.Context, name string, offset int, limit int) (responses.MultiDevicesResponse, error)
        // DevicesByServiceName returns devices associated with the specified device service name.
        DevicesByServiceName(ctx context.Context, name string, offset int, limit int) (responses.MultiDevicesResponse, error)
}

// DataEventClient defines the interface for interactions with the DataEvent endpoint on core-data service.
type DataEventClient interface {
        // EventCount returns the count of Events currently stored in the database.
        EventCount(ctx context.Context) (common.CountResponse, error)
        // EventCountByDeviceName returns the count of Events associated with the specified device name.
        EventCountByDeviceName(ctx context.Context, name string) (common.CountResponse, error)
}

// MeasurementClient defines the interface for interactions with the Measurement endpoint.
type MeasurementClient interface {
        // MeasurementsByDeviceName returns measurements associated with the specified device name.
        MeasurementsByDeviceName(ctx context.Context, name string, offset int, limit int) (responses.MultiMeasurementsResponse, error)
}

// CommandClient defines the interface for interactions with the Command endpoint.
type CommandClient interface {
        // GetCommand issues a get command to the specified device.
        GetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams map[string]string) (responses.EventResponse, error)
        // SetCommand issues a set command to the specified device.
        SetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams map[string]string, settings map[string]interface{}) (common.BaseResponse, error)
}

// AlertClient defines the interface for interactions with the Alert endpoint.
type AlertClient interface {
        // AlertsByDeviceName returns alerts associated with the specified device name.
        AlertsByDeviceName(ctx context.Context, name string, offset int, limit int) (responses.MultiAlertsResponse, error)
}

// EventSubscriptionClient defines the interface for interactions with the EventSubscription endpoint.
type EventSubscriptionClient interface {
        // EventSubscriptionsByDeviceName returns event subscriptions associated with the specified device name.
        EventSubscriptionsByDeviceName(ctx context.Context, name string, offset int, limit int) (responses.MultiEventSubscriptionsResponse, error)
}

// DeviceHandlerClient defines the interface for interactions with the DeviceHandler endpoint on core-metadata service.
type DeviceHandlerClient interface {
        // DeviceHandlerByName returns a device handler by name
        DeviceHandlerByName(ctx context.Context, name string) (responses.DeviceServiceResponse, error)
}

// DeviceTemplateClient defines the interface for interactions with the DeviceTemplate endpoint on core-metadata service.
type DeviceTemplateClient interface {
        // DeviceTemplateByName returns a device template by name
        DeviceTemplateByName(ctx context.Context, name string) (responses.DeviceProfileResponse, error)
}

// DeviceWatcherClient defines the interface for interactions with the DeviceWatcher endpoint on core-metadata service.
type DeviceWatcherClient interface {
        // DeviceWatcherByName returns a device watcher by name
        DeviceWatcherByName(ctx context.Context, name string) (responses.DeviceWatcherResponse, error)
}

// DeviceHandlerCommandClient defines the interface for command interactions with device handlers.
type DeviceHandlerCommandClient interface {
        // GetCommand issues a get command to the specified device through its handler.
        GetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams map[string]string) (responses.EventResponse, error)
        // SetCommand issues a set command to the specified device through its handler.
        SetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams map[string]string, settings map[string]interface{}) (common.BaseResponse, error)
        // SetCommandWithObject issues a set command with object payload to the specified device through its handler.
        SetCommandWithObject(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams map[string]string, settings map[string]interface{}) (common.BaseResponse, error)
}