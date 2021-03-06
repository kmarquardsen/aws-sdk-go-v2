// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AlarmType string

// Enum values for AlarmType
const (
	AlarmTypeCompositealarm AlarmType = "CompositeAlarm"
	AlarmTypeMetricalarm    AlarmType = "MetricAlarm"
)

type AnomalyDetectorStateValue string

// Enum values for AnomalyDetectorStateValue
const (
	AnomalyDetectorStateValuePending_training          AnomalyDetectorStateValue = "PENDING_TRAINING"
	AnomalyDetectorStateValueTrained_insufficient_data AnomalyDetectorStateValue = "TRAINED_INSUFFICIENT_DATA"
	AnomalyDetectorStateValueTrained                   AnomalyDetectorStateValue = "TRAINED"
)

type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	ComparisonOperatorGreaterthanorequaltothreshold            ComparisonOperator = "GreaterThanOrEqualToThreshold"
	ComparisonOperatorGreaterthanthreshold                     ComparisonOperator = "GreaterThanThreshold"
	ComparisonOperatorLessthanthreshold                        ComparisonOperator = "LessThanThreshold"
	ComparisonOperatorLessthanorequaltothreshold               ComparisonOperator = "LessThanOrEqualToThreshold"
	ComparisonOperatorLessthanlowerorgreaterthanupperthreshold ComparisonOperator = "LessThanLowerOrGreaterThanUpperThreshold"
	ComparisonOperatorLessthanlowerthreshold                   ComparisonOperator = "LessThanLowerThreshold"
	ComparisonOperatorGreaterthanupperthreshold                ComparisonOperator = "GreaterThanUpperThreshold"
)

type HistoryItemType string

// Enum values for HistoryItemType
const (
	HistoryItemTypeConfigurationupdate HistoryItemType = "ConfigurationUpdate"
	HistoryItemTypeStateupdate         HistoryItemType = "StateUpdate"
	HistoryItemTypeAction              HistoryItemType = "Action"
)

type RecentlyActive string

// Enum values for RecentlyActive
const (
	RecentlyActivePt3h RecentlyActive = "PT3H"
)

type ScanBy string

// Enum values for ScanBy
const (
	ScanByTimestamp_descending ScanBy = "TimestampDescending"
	ScanByTimestamp_ascending  ScanBy = "TimestampAscending"
)

type StandardUnit string

// Enum values for StandardUnit
const (
	StandardUnitSeconds         StandardUnit = "Seconds"
	StandardUnitMicroseconds    StandardUnit = "Microseconds"
	StandardUnitMilliseconds    StandardUnit = "Milliseconds"
	StandardUnitBytes           StandardUnit = "Bytes"
	StandardUnitKilobytes       StandardUnit = "Kilobytes"
	StandardUnitMegabytes       StandardUnit = "Megabytes"
	StandardUnitGigabytes       StandardUnit = "Gigabytes"
	StandardUnitTerabytes       StandardUnit = "Terabytes"
	StandardUnitBits            StandardUnit = "Bits"
	StandardUnitKilobits        StandardUnit = "Kilobits"
	StandardUnitMegabits        StandardUnit = "Megabits"
	StandardUnitGigabits        StandardUnit = "Gigabits"
	StandardUnitTerabits        StandardUnit = "Terabits"
	StandardUnitPercent         StandardUnit = "Percent"
	StandardUnitCount           StandardUnit = "Count"
	StandardUnitBytesSecond     StandardUnit = "Bytes/Second"
	StandardUnitKilobytesSecond StandardUnit = "Kilobytes/Second"
	StandardUnitMegabytesSecond StandardUnit = "Megabytes/Second"
	StandardUnitGigabytesSecond StandardUnit = "Gigabytes/Second"
	StandardUnitTerabytesSecond StandardUnit = "Terabytes/Second"
	StandardUnitBitsSecond      StandardUnit = "Bits/Second"
	StandardUnitKilobitsSecond  StandardUnit = "Kilobits/Second"
	StandardUnitMegabitsSecond  StandardUnit = "Megabits/Second"
	StandardUnitGigabitsSecond  StandardUnit = "Gigabits/Second"
	StandardUnitTerabitsSecond  StandardUnit = "Terabits/Second"
	StandardUnitCountSecond     StandardUnit = "Count/Second"
	StandardUnitNone            StandardUnit = "None"
)

type StateValue string

// Enum values for StateValue
const (
	StateValueOk                StateValue = "OK"
	StateValueAlarm             StateValue = "ALARM"
	StateValueInsufficient_data StateValue = "INSUFFICIENT_DATA"
)

type Statistic string

// Enum values for Statistic
const (
	StatisticSamplecount Statistic = "SampleCount"
	StatisticAverage     Statistic = "Average"
	StatisticSum         Statistic = "Sum"
	StatisticMinimum     Statistic = "Minimum"
	StatisticMaximum     Statistic = "Maximum"
)

type StatusCode string

// Enum values for StatusCode
const (
	StatusCodeComplete       StatusCode = "Complete"
	StatusCodeInternal_error StatusCode = "InternalError"
	StatusCodePartial_data   StatusCode = "PartialData"
)
