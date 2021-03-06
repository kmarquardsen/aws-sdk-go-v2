// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ChannelStatus string

// Enum values for ChannelStatus
const (
	ChannelStatusCreating ChannelStatus = "CREATING"
	ChannelStatusActive   ChannelStatus = "ACTIVE"
	ChannelStatusDeleting ChannelStatus = "DELETING"
)

type ComputeType string

// Enum values for ComputeType
const (
	ComputeTypeAcu_1 ComputeType = "ACU_1"
	ComputeTypeAcu_2 ComputeType = "ACU_2"
)

type DatasetActionType string

// Enum values for DatasetActionType
const (
	DatasetActionTypeQuery     DatasetActionType = "QUERY"
	DatasetActionTypeContainer DatasetActionType = "CONTAINER"
)

type DatasetContentState string

// Enum values for DatasetContentState
const (
	DatasetContentStateCreating  DatasetContentState = "CREATING"
	DatasetContentStateSucceeded DatasetContentState = "SUCCEEDED"
	DatasetContentStateFailed    DatasetContentState = "FAILED"
)

type DatasetStatus string

// Enum values for DatasetStatus
const (
	DatasetStatusCreating DatasetStatus = "CREATING"
	DatasetStatusActive   DatasetStatus = "ACTIVE"
	DatasetStatusDeleting DatasetStatus = "DELETING"
)

type DatastoreStatus string

// Enum values for DatastoreStatus
const (
	DatastoreStatusCreating DatastoreStatus = "CREATING"
	DatastoreStatusActive   DatastoreStatus = "ACTIVE"
	DatastoreStatusDeleting DatastoreStatus = "DELETING"
)

type LoggingLevel string

// Enum values for LoggingLevel
const (
	LoggingLevelError LoggingLevel = "ERROR"
)

type ReprocessingStatus string

// Enum values for ReprocessingStatus
const (
	ReprocessingStatusRunning   ReprocessingStatus = "RUNNING"
	ReprocessingStatusSucceeded ReprocessingStatus = "SUCCEEDED"
	ReprocessingStatusCancelled ReprocessingStatus = "CANCELLED"
	ReprocessingStatusFailed    ReprocessingStatus = "FAILED"
)
