// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type StartSelectorType string

// Enum values for StartSelectorType
const (
	StartSelectorTypeFragment_number    StartSelectorType = "FRAGMENT_NUMBER"
	StartSelectorTypeServer_timestamp   StartSelectorType = "SERVER_TIMESTAMP"
	StartSelectorTypeProducer_timestamp StartSelectorType = "PRODUCER_TIMESTAMP"
	StartSelectorTypeNow                StartSelectorType = "NOW"
	StartSelectorTypeEarliest           StartSelectorType = "EARLIEST"
	StartSelectorTypeContinuation_token StartSelectorType = "CONTINUATION_TOKEN"
)
