package models

// Constants for ServiceState
const (
	// Locked : device is locked
	// Unlocked : device is unlocked
	Locked   = "LOCKED"
	Unlocked = "UNLOCKED"
)

// Constants for ChannelType
const (
	Rest  = "REST"
	Email = "EMAIL"
)

// Constants for AlertSeverity
const (
	Minor    = "MINOR"
	Critical = "CRITICAL"
	Normal   = "NORMAL"
)

// Constants for AlertStatus
const (
	New       = "NEW"
	Processed = "PROCESSED"

	EscalationEventSubscriptionName = "ESCALATION"
	EscalationPrefix           = "escalated-"
	EscalatedContentNotice     = "This notification is escalated by the transmission"
)

// Constants for DeliveryStatus and ScheduleActionRecordStatus
const (
	Failed       = "FAILED"
	Sent         = "SENT"
	Acknowledged = "ACKNOWLEDGED"
	RESENDING    = "RESENDING"

	// Constants for ScheduleActionRecordStatus only
	Succeeded = "SUCCEEDED"
	Missed    = "MISSED"
)

// Constants for both AlertStatus and DeliveryStatus
const (
	Escalated = "ESCALATED"
)

// Constants for OperatingState
const (
	Up      = "UP"
	Down    = "DOWN"
	Unknown = "UNKNOWN"
)

// Constant for Keeper health status
const Halt = "HALT"
