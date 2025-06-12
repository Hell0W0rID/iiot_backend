//
//
//
//
// Unless required by applicable law or agreed to in writing, software

package container

import (
        "iiot-backend/pkg/go-mod-bootstrap/di"
        "iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
)

// DataEventClientName contains the name of the DataEventClient's implementation in the DIC.
var DataEventClientName = di.TypeInstanceToName((*interfaces.DataEventClient)(nil))

// DataEventClientFrom helper function queries the DIC and returns the DataEventClient's implementation.
func DataEventClientFrom(get di.Get) interfaces.DataEventClient {
        if get(DataEventClientName) == nil {
                return nil
        }

        return get(DataEventClientName).(interfaces.DataEventClient)
}

// MeasurementClientName contains the name of the MeasurementClient instance in the DIC.
var MeasurementClientName = di.TypeInstanceToName((*interfaces.MeasurementClient)(nil))

// MeasurementClientFrom helper function queries the DIC and returns the MeasurementClient instance.
func MeasurementClientFrom(get di.Get) interfaces.MeasurementClient {
        client, ok := get(MeasurementClientName).(interfaces.MeasurementClient)
        if !ok {
                return nil
        }

        return client
}

// CommandClientName contains the name of the CommandClient's implementation in the DIC.
var CommandClientName = di.TypeInstanceToName((*interfaces.CommandClient)(nil))

// CommandClientFrom helper function queries the DIC and returns the CommandClient's implementation.
func CommandClientFrom(get di.Get) interfaces.CommandClient {
        if get(CommandClientName) == nil {
                return nil
        }

        return get(CommandClientName).(interfaces.CommandClient)
}

// AlertClientName contains the name of the AlertClient's implementation in the DIC.
var AlertClientName = di.TypeInstanceToName((*interfaces.AlertClient)(nil))

// AlertClientFrom helper function queries the DIC and returns the AlertClient's implementation.
func AlertClientFrom(get di.Get) interfaces.AlertClient {
        if get(AlertClientName) == nil {
                return nil
        }

        return get(AlertClientName).(interfaces.AlertClient)
}

// EventSubscriptionClientName contains the name of the EventSubscriptionClient's implementation in the DIC.
var EventSubscriptionClientName = di.TypeInstanceToName((*interfaces.EventSubscriptionClient)(nil))

// EventSubscriptionClientFrom helper function queries the DIC and returns the EventSubscriptionClient's implementation.
func EventSubscriptionClientFrom(get di.Get) interfaces.EventSubscriptionClient {
        if get(EventSubscriptionClientName) == nil {
                return nil
        }

        return get(EventSubscriptionClientName).(interfaces.EventSubscriptionClient)
}



// DeviceClientName contains the name of the DeviceClient's implementation in the DIC.
var DeviceClientName = di.TypeInstanceToName((*interfaces.DeviceClient)(nil))

// DeviceClientFrom helper function queries the DIC and returns the DeviceClient's implementation.
func DeviceClientFrom(get di.Get) interfaces.DeviceClient {
        if get(DeviceClientName) == nil {
                return nil
        }

        return get(DeviceClientName).(interfaces.DeviceClient)
}

// DeviceWatcherClientName contains the name of the DeviceWatcherClient's implementation in the DIC.
var DeviceWatcherClientName = di.TypeInstanceToName((*interfaces.DeviceWatcherClient)(nil))

// DeviceWatcherClientFrom helper function queries the DIC and returns the DeviceWatcherClient's implementation.
func DeviceWatcherClientFrom(get di.Get) interfaces.DeviceWatcherClient {
        if get(DeviceWatcherClientName) == nil {
                return nil
        }

        return get(DeviceWatcherClientName).(interfaces.DeviceWatcherClient)
}

// DeviceHandlerCommandClientName contains the name of the DeviceHandlerCommandClient instance in the DIC.
var DeviceHandlerCommandClientName = di.TypeInstanceToName((*interfaces.DeviceHandlerCommandClient)(nil))

// DeviceHandlerCommandClientFrom helper function queries the DIC and returns the DeviceHandlerCommandClient instance.
func DeviceHandlerCommandClientFrom(get di.Get) interfaces.DeviceHandlerCommandClient {
        client, ok := get(DeviceHandlerCommandClientName).(interfaces.DeviceHandlerCommandClient)
        if !ok {
                return nil
        }

        return client
}

// ScheduleJobClientName contains the name of the ScheduleJobClient's implementation in the DIC.
var ScheduleJobClientName = di.TypeInstanceToName((*interfaces.ScheduleJobClient)(nil))

// ScheduleJobClientFrom helper function queries the DIC and returns the ScheduleJobClient's implementation.
func ScheduleJobClientFrom(get di.Get) interfaces.ScheduleJobClient {
        if get(ScheduleJobClientName) == nil {
                return nil
        }

        return get(ScheduleJobClientName).(interfaces.ScheduleJobClient)
}

// ScheduleActionRecordClientName contains the name of the ScheduleActionRecordClient's implementation in the DIC.
var ScheduleActionRecordClientName = di.TypeInstanceToName((*interfaces.ScheduleActionRecordClient)(nil))

// ScheduleActionRecordClientFrom helper function queries the DIC and returns the ScheduleActionRecordClient's implementation.
func ScheduleActionRecordClientFrom(get di.Get) interfaces.ScheduleActionRecordClient {
        if get(ScheduleActionRecordClientName) == nil {
                return nil
        }

        return get(ScheduleActionRecordClientName).(interfaces.ScheduleActionRecordClient)
}

// SecurityProxyAuthClientName contains the name of the AuthClient's implementation in the DIC.
var SecurityProxyAuthClientName = di.TypeInstanceToName((*interfaces.AuthClient)(nil))

// SecurityProxyAuthClientFrom helper function queries the DIC and returns the AuthClient's implementation.
func SecurityProxyAuthClientFrom(get di.Get) interfaces.AuthClient {
        if get(SecurityProxyAuthClientName) == nil {
                return nil
        }

        return get(SecurityProxyAuthClientName).(interfaces.AuthClient)
}

// DeviceHandlerClientName contains the name of the DeviceServiceClient's implementation in the DIC.
var DeviceHandlerClientName = di.TypeInstanceToName((*interfaces.DeviceServiceClient)(nil))

// DeviceHandlerClientFrom helper function queries the DIC and returns the DeviceServiceClient's implementation.
func DeviceHandlerClientFrom(get di.Get) interfaces.DeviceServiceClient {
        if get(DeviceHandlerClientName) == nil {
                return nil
        }

        return get(DeviceHandlerClientName).(interfaces.DeviceServiceClient)
}

// DeviceTemplateClientName contains the name of the DeviceProfileClient's implementation in the DIC.
var DeviceTemplateClientName = di.TypeInstanceToName((*interfaces.DeviceProfileClient)(nil))

// DeviceTemplateClientFrom helper function queries the DIC and returns the DeviceProfileClient's implementation.
func DeviceTemplateClientFrom(get di.Get) interfaces.DeviceProfileClient {
        if get(DeviceTemplateClientName) == nil {
                return nil
        }

        return get(DeviceTemplateClientName).(interfaces.DeviceProfileClient)
}




