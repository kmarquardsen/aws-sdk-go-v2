// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type BehaviorOnMxFailure string

// Enum values for BehaviorOnMxFailure
const (
	BehaviorOnMxFailureUse_default_value BehaviorOnMxFailure = "USE_DEFAULT_VALUE"
	BehaviorOnMxFailureReject_message    BehaviorOnMxFailure = "REJECT_MESSAGE"
)

type BulkEmailStatus string

// Enum values for BulkEmailStatus
const (
	BulkEmailStatusSuccess                          BulkEmailStatus = "SUCCESS"
	BulkEmailStatusMessage_rejected                 BulkEmailStatus = "MESSAGE_REJECTED"
	BulkEmailStatusMail_from_domain_not_verified    BulkEmailStatus = "MAIL_FROM_DOMAIN_NOT_VERIFIED"
	BulkEmailStatusConfiguration_set_not_found      BulkEmailStatus = "CONFIGURATION_SET_NOT_FOUND"
	BulkEmailStatusTemplate_not_found               BulkEmailStatus = "TEMPLATE_NOT_FOUND"
	BulkEmailStatusAccount_suspended                BulkEmailStatus = "ACCOUNT_SUSPENDED"
	BulkEmailStatusAccount_throttled                BulkEmailStatus = "ACCOUNT_THROTTLED"
	BulkEmailStatusAccount_daily_quota_exceeded     BulkEmailStatus = "ACCOUNT_DAILY_QUOTA_EXCEEDED"
	BulkEmailStatusInvalid_sending_pool_name        BulkEmailStatus = "INVALID_SENDING_POOL_NAME"
	BulkEmailStatusAccount_sending_paused           BulkEmailStatus = "ACCOUNT_SENDING_PAUSED"
	BulkEmailStatusConfiguration_set_sending_paused BulkEmailStatus = "CONFIGURATION_SET_SENDING_PAUSED"
	BulkEmailStatusInvalid_parameter                BulkEmailStatus = "INVALID_PARAMETER"
	BulkEmailStatusTransient_failure                BulkEmailStatus = "TRANSIENT_FAILURE"
	BulkEmailStatusFailed                           BulkEmailStatus = "FAILED"
)

type ContactLanguage string

// Enum values for ContactLanguage
const (
	ContactLanguageEn ContactLanguage = "EN"
	ContactLanguageJa ContactLanguage = "JA"
)

type DeliverabilityDashboardAccountStatus string

// Enum values for DeliverabilityDashboardAccountStatus
const (
	DeliverabilityDashboardAccountStatusActive             DeliverabilityDashboardAccountStatus = "ACTIVE"
	DeliverabilityDashboardAccountStatusPending_expiration DeliverabilityDashboardAccountStatus = "PENDING_EXPIRATION"
	DeliverabilityDashboardAccountStatusDisabled           DeliverabilityDashboardAccountStatus = "DISABLED"
)

type DeliverabilityTestStatus string

// Enum values for DeliverabilityTestStatus
const (
	DeliverabilityTestStatusIn_progress DeliverabilityTestStatus = "IN_PROGRESS"
	DeliverabilityTestStatusCompleted   DeliverabilityTestStatus = "COMPLETED"
)

type DimensionValueSource string

// Enum values for DimensionValueSource
const (
	DimensionValueSourceMessage_tag  DimensionValueSource = "MESSAGE_TAG"
	DimensionValueSourceEmail_header DimensionValueSource = "EMAIL_HEADER"
	DimensionValueSourceLink_tag     DimensionValueSource = "LINK_TAG"
)

type DkimSigningAttributesOrigin string

// Enum values for DkimSigningAttributesOrigin
const (
	DkimSigningAttributesOriginAws_ses  DkimSigningAttributesOrigin = "AWS_SES"
	DkimSigningAttributesOriginExternal DkimSigningAttributesOrigin = "EXTERNAL"
)

type DkimStatus string

// Enum values for DkimStatus
const (
	DkimStatusPending           DkimStatus = "PENDING"
	DkimStatusSuccess           DkimStatus = "SUCCESS"
	DkimStatusFailed            DkimStatus = "FAILED"
	DkimStatusTemporary_failure DkimStatus = "TEMPORARY_FAILURE"
	DkimStatusNot_started       DkimStatus = "NOT_STARTED"
)

type EventType string

// Enum values for EventType
const (
	EventTypeSend              EventType = "SEND"
	EventTypeReject            EventType = "REJECT"
	EventTypeBounce            EventType = "BOUNCE"
	EventTypeComplaint         EventType = "COMPLAINT"
	EventTypeDelivery          EventType = "DELIVERY"
	EventTypeOpen              EventType = "OPEN"
	EventTypeClick             EventType = "CLICK"
	EventTypeRendering_failure EventType = "RENDERING_FAILURE"
	EventTypeDelivery_delay    EventType = "DELIVERY_DELAY"
)

type IdentityType string

// Enum values for IdentityType
const (
	IdentityTypeEmail_address  IdentityType = "EMAIL_ADDRESS"
	IdentityTypeDomain         IdentityType = "DOMAIN"
	IdentityTypeManaged_domain IdentityType = "MANAGED_DOMAIN"
)

type MailFromDomainStatus string

// Enum values for MailFromDomainStatus
const (
	MailFromDomainStatusPending           MailFromDomainStatus = "PENDING"
	MailFromDomainStatusSuccess           MailFromDomainStatus = "SUCCESS"
	MailFromDomainStatusFailed            MailFromDomainStatus = "FAILED"
	MailFromDomainStatusTemporary_failure MailFromDomainStatus = "TEMPORARY_FAILURE"
)

type MailType string

// Enum values for MailType
const (
	MailTypeMarketing     MailType = "MARKETING"
	MailTypeTransactional MailType = "TRANSACTIONAL"
)

type ReviewStatus string

// Enum values for ReviewStatus
const (
	ReviewStatusPending ReviewStatus = "PENDING"
	ReviewStatusFailed  ReviewStatus = "FAILED"
	ReviewStatusGranted ReviewStatus = "GRANTED"
	ReviewStatusDenied  ReviewStatus = "DENIED"
)

type SuppressionListReason string

// Enum values for SuppressionListReason
const (
	SuppressionListReasonBounce    SuppressionListReason = "BOUNCE"
	SuppressionListReasonComplaint SuppressionListReason = "COMPLAINT"
)

type TlsPolicy string

// Enum values for TlsPolicy
const (
	TlsPolicyRequire  TlsPolicy = "REQUIRE"
	TlsPolicyOptional TlsPolicy = "OPTIONAL"
)

type WarmupStatus string

// Enum values for WarmupStatus
const (
	WarmupStatusIn_progress WarmupStatus = "IN_PROGRESS"
	WarmupStatusDone        WarmupStatus = "DONE"
)
