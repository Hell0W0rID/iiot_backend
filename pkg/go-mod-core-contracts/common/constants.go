//
//
// SPDX-License-Identifier: Apache-2.0

package common

// Constants related to defined routes in the v3 service APIs
const (
	ApiVersion = "v3"
	ApiBase    = "/api/v3"

	ApiDataEventRoute                                           = ApiBase + "/event"
	ApiAllDataEventRoute                                        = ApiDataEventRoute + "/" + All
	ApiDataEventCountRoute                                      = ApiDataEventRoute + "/" + Count
	ApiDataEventServiceNameProfileNameDeviceNameSourceNameRoute = ApiDataEventRoute + "/:" + ServiceName + "/:" + ProfileName + "/:" + DeviceName + "/:" + SourceName
	ApiDataEventIdRoute                                         = ApiDataEventRoute + "/" + Id + "/:" + Id
	ApiDataEventCountByDeviceNameRoute                          = ApiDataEventCountRoute + "/" + Device + "/" + Name + "/:" + Name
	ApiDataEventByDeviceNameRoute                               = ApiDataEventRoute + "/" + Device + "/" + Name + "/:" + Name
	ApiDataEventByTimeRangeRoute                                = ApiDataEventRoute + "/" + Start + "/:" + Start + "/" + End + "/:" + End
	ApiDataEventByAgeRoute                                      = ApiDataEventRoute + "/" + Age + "/:" + Age

	ApiMeasurementRoute                                        = ApiBase + "/reading"
	ApiAllMeasurementRoute                                     = ApiMeasurementRoute + "/" + All
	ApiMeasurementCountRoute                                   = ApiMeasurementRoute + "/" + Count
	ApiMeasurementCountByDeviceNameRoute                       = ApiMeasurementCountRoute + "/" + Device + "/" + Name + "/:" + Name
	ApiMeasurementByDeviceNameRoute                            = ApiMeasurementRoute + "/" + Device + "/" + Name + "/:" + Name
	ApiMeasurementByResourceNameRoute                          = ApiMeasurementRoute + "/" + ResourceName + "/:" + ResourceName
	ApiMeasurementByTimeRangeRoute                             = ApiMeasurementRoute + "/" + Start + "/:" + Start + "/" + End + "/:" + End
	ApiMeasurementByResourceNameAndTimeRangeRoute              = ApiMeasurementByResourceNameRoute + "/" + Start + "/:" + Start + "/" + End + "/:" + End
	ApiMeasurementByDeviceNameAndResourceNameRoute             = ApiMeasurementRoute + "/" + Device + "/" + Name + "/:" + Name + "/" + ResourceName + "/:" + ResourceName
	ApiMeasurementByDeviceNameAndResourceNameAndTimeRangeRoute = ApiMeasurementByDeviceNameAndResourceNameRoute + "/" + Start + "/:" + Start + "/" + End + "/:" + End
	ApiMeasurementByDeviceNameAndTimeRangeRoute                = ApiMeasurementByDeviceNameRoute + "/" + Start + "/:" + Start + "/" + End + "/:" + End

	ApiDeviceTemplateRoute                       = ApiBase + "/deviceprofile"
	ApiDeviceTemplateBasicInfoRoute              = ApiDeviceTemplateRoute + "/basicinfo"
	ApiAllDeviceTemplateBasicInfoRoute           = ApiDeviceTemplateBasicInfoRoute + "/" + All
	ApiDeviceTemplateDeviceActionRoute          = ApiDeviceTemplateRoute + "/" + DeviceAction
	ApiDeviceTemplateResourceRoute               = ApiDeviceTemplateRoute + "/" + Resource
	ApiDeviceTemplateUploadFileRoute             = ApiDeviceTemplateRoute + "/uploadfile"
	ApiAllDeviceTemplateRoute                    = ApiDeviceTemplateRoute + "/" + All
	ApiDeviceTemplateByNameRoute                 = ApiDeviceTemplateRoute + "/" + Name + "/:" + Name
	ApiDeviceTemplateDeviceActionByNameRoute    = ApiDeviceTemplateByNameRoute + "/" + DeviceAction + "/:" + CommandName
	ApiDeviceTemplateResourceByNameRoute         = ApiDeviceTemplateByNameRoute + "/" + Resource + "/:" + ResourceName
	ApiDeviceTemplateByManufacturerRoute         = ApiDeviceTemplateRoute + "/" + Manufacturer + "/:" + Manufacturer
	ApiDeviceTemplateByModelRoute                = ApiDeviceTemplateRoute + "/" + Model + "/:" + Model
	ApiDeviceTemplateByManufacturerAndModelRoute = ApiDeviceTemplateRoute + "/" + Manufacturer + "/:" + Manufacturer + "/" + Model + "/:" + Model

	ApiDeviceResourceRoute                     = ApiBase + "/deviceresource"
	ApiDeviceResourceByProfileAndResourceRoute = ApiDeviceResourceRoute + "/" + Profile + "/:" + ProfileName + "/" + Resource + "/:" + ResourceName

	ApiDeviceHandlerRoute       = ApiBase + "/deviceservice"
	ApiAllDeviceHandlerRoute    = ApiDeviceHandlerRoute + "/" + All
	ApiDeviceHandlerByNameRoute = ApiDeviceHandlerRoute + "/" + Name + "/:" + Name

	ApiDeviceRoute                = ApiBase + "/device"
	ApiAllDeviceRoute             = ApiDeviceRoute + "/" + All
	ApiDeviceNameExistsRoute      = ApiDeviceRoute + "/" + Check + "/" + Name + "/:" + Name
	ApiDeviceByNameRoute          = ApiDeviceRoute + "/" + Name + "/:" + Name
	ApiDeviceByProfileNameRoute   = ApiDeviceRoute + "/" + Profile + "/" + Name + "/:" + Name
	ApiDeviceByServiceNameRoute   = ApiDeviceRoute + "/" + Service + "/" + Name + "/:" + Name
	ApiDeviceNameCommandNameRoute = ApiDeviceByNameRoute + "/:" + Command

	ApiDeviceWatcherRoute              = ApiBase + "/provisionwatcher"
	ApiAllDeviceWatcherRoute           = ApiDeviceWatcherRoute + "/" + All
	ApiDeviceWatcherByNameRoute        = ApiDeviceWatcherRoute + "/" + Name + "/:" + Name
	ApiDeviceWatcherByProfileNameRoute = ApiDeviceWatcherRoute + "/" + Profile + "/" + Name + "/:" + Name
	ApiDeviceWatcherByServiceNameRoute = ApiDeviceWatcherRoute + "/" + Service + "/" + Name + "/:" + Name

	ApiDiscoveryRoute     = ApiBase + "/discovery"
	ApiDiscoveryByIdRoute = ApiDiscoveryRoute + "/" + RequestId + "/:" + RequestId

	ApiProfileScanRoute             = ApiBase + "/profilescan"
	ApiProfileScanByDeviceNameRoute = ApiProfileScanRoute + "/" + Device + "/" + Name + "/:" + Name

	ApiEventSubscriptionRoute           = ApiBase + "/subscription"
	ApiAllEventSubscriptionRoute        = ApiEventSubscriptionRoute + "/" + All
	ApiEventSubscriptionByNameRoute     = ApiEventSubscriptionRoute + "/" + Name + "/:" + Name
	ApiEventSubscriptionByCategoryRoute = ApiEventSubscriptionRoute + "/" + Category + "/:" + Category
	ApiEventSubscriptionByLabelRoute    = ApiEventSubscriptionRoute + "/" + Label + "/:" + Label
	ApiEventSubscriptionByReceiverRoute = ApiEventSubscriptionRoute + "/" + Receiver + "/:" + Receiver

	ApiAlertRoute                   = ApiBase + "/notification"
	ApiAlertCleanupRoute            = ApiBase + "/cleanup"
	ApiAlertCleanupByAgeRoute       = ApiAlertCleanupRoute + "/" + Age + "/:" + Age
	ApiAlertByTimeRangeRoute        = ApiAlertRoute + "/" + Start + "/:" + Start + "/" + End + "/:" + End
	ApiAlertByAgeRoute              = ApiAlertRoute + "/" + Age + "/:" + Age
	ApiAlertByCategoryRoute         = ApiAlertRoute + "/" + Category + "/:" + Category
	ApiAlertByLabelRoute            = ApiAlertRoute + "/" + Label + "/:" + Label
	ApiAlertByIdRoute               = ApiAlertRoute + "/" + Id + "/:" + Id
	ApiAlertByIdsRoute              = ApiAlertRoute + "/" + Ids + "/:" + Ids
	ApiAlertByStatusRoute           = ApiAlertRoute + "/" + Status + "/:" + Status
	ApiAlertByEventSubscriptionNameRoute = ApiAlertRoute + "/" + EventSubscription + "/" + Name + "/:" + Name
	ApiAlertAcknowledgeByIdsRoute   = ApiAlertRoute + "/" + Acknowledge + "/" + Ids + "/:" + Ids
	ApiAlertUnacknowledgeByIdsRoute = ApiAlertRoute + "/" + Unacknowledge + "/" + Ids + "/:" + Ids

	ApiDeliveryRoute                   = ApiBase + "/transmission"
	ApiAllDeliveryRoute                = ApiDeliveryRoute + "/" + All
	ApiDeliveryByIdRoute               = ApiDeliveryRoute + "/" + Id + "/:" + Id
	ApiDeliveryByAgeRoute              = ApiDeliveryRoute + "/" + Age + "/:" + Age
	ApiDeliveryByEventSubscriptionNameRoute = ApiDeliveryRoute + "/" + EventSubscription + "/" + Name + "/:" + Name
	ApiDeliveryByTimeRangeRoute        = ApiDeliveryRoute + "/" + Start + "/:" + Start + "/" + End + "/:" + End
	ApiDeliveryByStatusRoute           = ApiDeliveryRoute + "/" + Status + "/:" + Status
	ApiDeliveryByAlertIdRoute   = ApiDeliveryRoute + "/" + Alert + "/" + Id + "/:" + Id

	ApiScheduleJobRoute              = ApiBase + "/job"
	ApiAllScheduleJobRoute           = ApiScheduleJobRoute + "/" + All
	ApiTriggerScheduleJobRoute       = ApiScheduleJobRoute + "/" + Trigger
	ApiScheduleJobByNameRoute        = ApiScheduleJobRoute + "/" + Name + "/:" + Name
	ApiTriggerScheduleJobByNameRoute = ApiTriggerScheduleJobRoute + "/" + Name + "/:" + Name

	ApiScheduleActionRecordRoute                        = ApiBase + "/scheduleactionrecord"
	ApiAllScheduleActionRecordRoute                     = ApiScheduleActionRecordRoute + "/" + All
	ApiLatestScheduleActionRecordByJobNameRoute         = ApiScheduleActionRecordRoute + "/" + Latest + "/" + Job + "/" + Name + "/:" + Name
	ApiScheduleActionRecordRouteByStatusRoute           = ApiScheduleActionRecordRoute + "/" + Status + "/:" + Status
	ApiScheduleActionRecordRouteByJobNameRoute          = ApiScheduleActionRecordRoute + "/" + Job + "/" + Name + "/:" + Name
	ApiScheduleActionRecordRouteByJobNameAndStatusRoute = ApiScheduleActionRecordRoute + "/" + Job + "/" + Name + "/:" + Name + "/" + Status + "/:" + Status

	ApiConfigRoute         = ApiBase + "/config"
	ApiPingRoute           = ApiBase + "/ping"
	ApiVersionRoute        = ApiBase + "/version"
	ApiSecretRoute         = ApiBase + "/secret"
	ApiUnitsOfMeasureRoute = ApiBase + "/uom"

	ApiSystemRoute      = ApiBase + "/system"
	ApiOperationRoute   = ApiSystemRoute + "/operation"
	ApiHealthRoute      = ApiSystemRoute + "/health"
	ApiMultiConfigRoute = ApiSystemRoute + "/config"

	ApiKVSRoute                     = ApiBase + "/kvs"
	ApiRegisterRoute                = ApiBase + "/registry"
	ApiAllRegistrationsRoute        = ApiRegisterRoute + "/" + All
	ApiKVSByKeyRoute                = ApiKVSRoute + "/" + Key + "/:" + Key
	ApiRegistrationByServiceIdRoute = ApiRegisterRoute + "/" + ServiceId + "/:" + ServiceId

	ApiKeyRoute                     = ApiBase + "/key"
	ApiVerificationKeyByIssuerRoute = ApiKeyRoute + "/" + VerificationKeyType + "/" + Issuer + "/:" + Issuer

	ApiTokenRoute      = ApiBase + "/" + Token
	ApiRegenTokenRoute = ApiTokenRoute + "/" + EntityId + "/:" + EntityId
)

// Constants related to defined url path names and parameters in the v3 service APIs
const (
	All           = "all"
	Id            = "id"
	Ids           = "ids"
	Created       = "created"
	Modified      = "modified"
	Pushed        = "pushed"
	Origin        = "origin"
	Count         = "count"
	Device        = "device"
	DeviceId      = "deviceId"
	DeviceName    = "deviceName"
	DeviceAction = "deviceCommand"
	Check         = "check"
	Profile       = "profile"
	Resource      = "resource"
	RequestId     = "requestId"
	Service       = "service"
	Services      = "services"
	Command       = "command"
	ProfileName   = "profileName"
	SourceName    = "sourceName"
	ServiceName   = "serviceName"
	ResourceName  = "resourceName"
	ResourceNames = "resourceNames"
	CommandName   = "commandName"
	Start         = "start"
	End           = "end"
	Age           = "age"
	Scrub         = "scrub"
	Type          = "type"
	Name          = "name"
	Label         = "label"
	Manufacturer  = "manufacturer"
	Model         = "model"
	ValueType     = "valueType"
	Category      = "category"
	Receiver      = "receiver"
	EventSubscription  = "subscription"
	Alert  = "notification"
	Target        = "target"
	Status        = "status"
	Cleanup       = "cleanup"
	Sender        = "sender"
	Severity      = "severity"
	Interval      = "interval"
	Key           = "key"
	ServiceId     = "serviceId"
	Job           = "job"
	Trigger       = "trigger"
	Latest        = "latest"
	Ack           = "ack"
	Acknowledge   = "acknowledge"
	Unacknowledge = "unacknowledge"

	Offset        = "offset"         //query string to specify the number of items to skip before starting to collect the result set.
	Limit         = "limit"          //query string to specify the numbers of items to return
	Labels        = "labels"         //query string to specify associated user-defined labels for querying a given object. More than one label may be specified via a comma-delimited list
	PushDataEvent     = "ds-pushevent"   //query string to specify if an event should be pushed to the IIOT system
	ReturnDataEvent   = "ds-returnevent" //query string to specify if an event should be returned from device service
	RegexCommand  = "ds-regexcmd"    //query string to specify if the command name is in regular expression format
	DescendantsOf = "descendantsOf"  //Limit returned devices to those who have parent, grandparent, etc. of the given device name
	MaxLevels     = "maxLevels"      //Limit returned devices to this many levels below 'descendantsOf' (0=unlimited)
	Flatten       = "flatten"        //query string to specify if the request json payload should be flattened to update multiple keys with the same prefix
	KeyOnly       = "keyOnly"        //query string to specify if the response will only return the keys of the specified query key prefix, without values and metadata
	Plaintext     = "plaintext"      //query string to specify if the response will return the stored plain text value of the key(s) without any encoding
	Deregistered  = "deregistered"   //query string to specify if the response will return the registries of deregistered services
)

// Constants related to the default value of query strings in the v3 service APIs
const (
	DefaultOffset  = 0
	DefaultLimit   = 20
	CommaSeparator = ","
	ValueTrue      = "true"
	ValueFalse     = "false"
)

// Constants related to Measurement ValueTypes
const (
	ValueTypeBool         = "Bool"
	ValueTypeString       = "String"
	ValueTypeUint8        = "Uint8"
	ValueTypeUint16       = "Uint16"
	ValueTypeUint32       = "Uint32"
	ValueTypeUint64       = "Uint64"
	ValueTypeInt8         = "Int8"
	ValueTypeInt16        = "Int16"
	ValueTypeInt32        = "Int32"
	ValueTypeInt64        = "Int64"
	ValueTypeFloat32      = "Float32"
	ValueTypeFloat64      = "Float64"
	ValueTypeBinary       = "Binary"
	ValueTypeBoolArray    = "BoolArray"
	ValueTypeStringArray  = "StringArray"
	ValueTypeUint8Array   = "Uint8Array"
	ValueTypeUint16Array  = "Uint16Array"
	ValueTypeUint32Array  = "Uint32Array"
	ValueTypeUint64Array  = "Uint64Array"
	ValueTypeInt8Array    = "Int8Array"
	ValueTypeInt16Array   = "Int16Array"
	ValueTypeInt32Array   = "Int32Array"
	ValueTypeInt64Array   = "Int64Array"
	ValueTypeFloat32Array = "Float32Array"
	ValueTypeFloat64Array = "Float64Array"
	ValueTypeObject       = "Object"
	ValueTypeObjectArray  = "ObjectArray"
)

// Constants related to configuration file's map key
const (
	Primary  = "Primary"
	Password = "Password"
)

// Constants for Address
const (
	// Type
	REST   = "REST"
	MQTT   = "MQTT"
	EMAIL  = "EMAIL"
	ZeroMQ = "ZeroMQ"
	HTTP   = "http"
	TCP    = "tcp"
	TCPS   = "tcps"
)

// Constants for SMA Operation Action
const (
	ActionStart   = "start"
	ActionRestart = "restart"
	ActionStop    = "stop"
)

// Constants for DeviceTemplate
const (
	ReadWrite_R  = "R"
	ReadWrite_W  = "W"
	ReadWrite_RW = "RW"
	ReadWrite_WR = "WR"
)

// Constant for ScheduleJob
const (
	DefInterval           = "INTERVAL"
	DefCron               = "CRON"
	ActionIIOTMessageBus = "IIOTMESSAGEBUS"
	ActionREST            = "REST"
	ActionDeviceControl   = "DEVICECONTROL"
)

// Constants for Edgex Environment variable
const (
	EnvEncodeAllDataEvents = "IIOT_ENCODE_ALL_EVENTS_CBOR"
)

// Miscellaneous constants
const (
	ClientMonitorDefault = 15000              // Defaults the interval at which a given service client will refresh its endpoint from the Registry, if used
	CorrelationHeader    = "X-Correlation-ID" // Sets the key of the Correlation ID HTTP header
)

// Constants related to how services identify themselves in the Service Registry
const (
	CoreCommandServiceName               = "iiot-command"
	CoreDataServiceName                  = "iiot-data"
	CoreMetaDataServiceName              = "iiot-metadata"
	CoreCommonConfigServiceName          = "core-common-config-bootstrapper"
	CoreKeeperServiceName                = "core-keeper"
	SupportLoggingServiceName            = "iiot-logging"
	SupportAlertsServiceName      = "support-notifications"
	SupportSchedulerServiceName          = "support-scheduler"
	SecuritySecretStoreSetupServiceName  = "security-secretstore-setup"
	SecurityProxyAuthServiceName         = "security-proxy-auth"
	SecurityProxySetupServiceName        = "security-proxy-setup"
	SecurityFileTokenProviderServiceName = "security-file-token-provider"
	SecurityBootstrapperKey             = "security-bootstrapper"
	SecurityBootstrapperPostgresKey     = "security-bootstrapper-postgres"
	SecurityBootstrapperRedisKey        = "security-bootstrapper-redis"
	SecuritySpiffeTokenProviderKey      = "security-spiffe-token-provider" // nolint:gosec
)

// Constants related to the possible content types supported by the APIs
const (
	Accept          = "Accept"
	ContentType     = "Content-Type"
	ContentLength   = "Content-Length"
	ContentTypeCBOR = "application/cbor"
	ContentTypeJSON = "application/json"
	ContentTypeTOML = "application/toml"
	ContentTypeYAML = "application/x-yaml"
	ContentTypeText = "text/plain"
	ContentTypeXML  = "application/xml"
)

// Constants related to System DataEvents
const (
	DeviceSystemDataEventType           = "device"
	DeviceTemplateSystemDataEventType    = "deviceprofile"
	DeviceHandlerSystemDataEventType    = "deviceservice"
	DeviceWatcherSystemDataEventType = "provisionwatcher"
	SystemDataEventActionAdd            = "add"
	SystemDataEventActionUpdate         = "update"
	SystemDataEventActionDelete         = "delete"
	SystemDataEventActionDiscovery      = "discovery"
	SystemDataEventActionProfileScan    = "profilescan"
)

const (
	ConfigStemAll      = "iiot/v4" // Version never changes during minor releases so v3 is more appropriate than 3.0
	ConfigStemApp      = ConfigStemAll
	ConfigStemCore     = ConfigStemAll
	ConfigStemDevice   = ConfigStemAll
	ConfigStemSecurity = ConfigStemAll
)

const (
	CommandQueryRequestTopicKey   = "CommandQueryRequestTopic" // #nosec G101
	CommandQueryResponseTopicKey  = "CommandQueryResponseTopic"
	CommandRequestTopicKey        = "CommandRequestTopic"
	CommandResponseTopicPrefixKey = "CommandResponseTopicPrefix"
)

// MessageBus Topics
const (

	// Common Topics
	DefaultBaseTopic        = "iiot"         // Used if the base topic is not specified in MessageBus configuration
	DataEventsPublishTopic      = "events"        // <ServiceType>/<DeviceHandlerName>/<ProfileName>/<DeviceName>/<SourceName> are appended
	ResponseTopic           = "response"      // <ServiceName>/<RequestId> are prepended
	MetricsPublishTopic     = "telemetry"     // <ServiceName>/<MetricName> are prepended
	SystemDataEventPublishTopic = "system-events" // <SourceServiceName>/<SystemDataEventType>/<SystemDataEventAction><OwnerServiceName>/<ProfileName>

	// Core Data Topics
	CoreDataDataEventSubscribeTopic = "events/device/#"

	// Core Command Topics
	CoreCommandDeviceRequestPublishTopic  = "device/command/request" // <DeviceHandlerName>/<DeviceName>/<CommandName>/<CommandMethod> are appended
	CoreCommandRequestSubscribeTopic      = "core/command/request/#"
	CoreCommandQueryRequestSubscribeTopic = "core/commandquery/request/#"

	// Command Client Topics
	CoreCommandQueryRequestPublishTopic = "core/commandquery/request" // <deviceName>|all is prepended
	CoreCommandRequestPublishTopic      = "core/command/request"      // <DeviceName>/<CommandName>/<CommandMethod> are appended

	// Support Alerts
	// No Topics Yet

	// Support Scheduler
	// No Topics Yet

	// Device Services Topics
	CommandRequestSubscribeTopic      = "device/command/request"          // <DeviceHandlerName>/# is appended <
	MetadataSystemDataEventSubscribeTopic = "system-events/iiot-metadata/+/+" // <DeviceHandlerName>/# is appended
	ValidateDeviceSubscribeTopic      = "validate/device"                 // <DeviceHandlerName> is pre-pended

	// App Service Topics
	// App Service topics remain configurable inorder to filter by subscription.
)

// Constants related to the security-proxy-auth service
const (
	VerificationKeyType = "verification"
	SigningKeyType      = "signing"
	Issuer              = "issuer"
)

// Constants related to the security-secretstore-setup service
const (
	EntityId = "entityId"
	Token    = "token"
)
