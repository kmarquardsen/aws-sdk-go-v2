// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DetailType string

// Enum values for DetailType
const (
	DetailTypeBasic DetailType = "BASIC"
	DetailTypeFull  DetailType = "FULL"
)

type ListEventTypesFilterName string

// Enum values for ListEventTypesFilterName
const (
	ListEventTypesFilterNameResource_type ListEventTypesFilterName = "RESOURCE_TYPE"
	ListEventTypesFilterNameService_name  ListEventTypesFilterName = "SERVICE_NAME"
)

type ListNotificationRulesFilterName string

// Enum values for ListNotificationRulesFilterName
const (
	ListNotificationRulesFilterNameEvent_type_id  ListNotificationRulesFilterName = "EVENT_TYPE_ID"
	ListNotificationRulesFilterNameCreated_by     ListNotificationRulesFilterName = "CREATED_BY"
	ListNotificationRulesFilterNameResource       ListNotificationRulesFilterName = "RESOURCE"
	ListNotificationRulesFilterNameTarget_address ListNotificationRulesFilterName = "TARGET_ADDRESS"
)

type ListTargetsFilterName string

// Enum values for ListTargetsFilterName
const (
	ListTargetsFilterNameTarget_type    ListTargetsFilterName = "TARGET_TYPE"
	ListTargetsFilterNameTarget_address ListTargetsFilterName = "TARGET_ADDRESS"
	ListTargetsFilterNameTarget_status  ListTargetsFilterName = "TARGET_STATUS"
)

type NotificationRuleStatus string

// Enum values for NotificationRuleStatus
const (
	NotificationRuleStatusEnabled  NotificationRuleStatus = "ENABLED"
	NotificationRuleStatusDisabled NotificationRuleStatus = "DISABLED"
)

type TargetStatus string

// Enum values for TargetStatus
const (
	TargetStatusPending     TargetStatus = "PENDING"
	TargetStatusActive      TargetStatus = "ACTIVE"
	TargetStatusUnreachable TargetStatus = "UNREACHABLE"
	TargetStatusInactive    TargetStatus = "INACTIVE"
	TargetStatusDeactivated TargetStatus = "DEACTIVATED"
)
