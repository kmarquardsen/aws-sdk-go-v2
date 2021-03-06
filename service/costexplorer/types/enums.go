// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccountScope string

// Enum values for AccountScope
const (
	AccountScopePayer  AccountScope = "PAYER"
	AccountScopeLinked AccountScope = "LINKED"
)

type Context string

// Enum values for Context
const (
	ContextCost_and_usage Context = "COST_AND_USAGE"
	ContextReservations   Context = "RESERVATIONS"
	ContextSavings_plans  Context = "SAVINGS_PLANS"
)

type CostCategoryRuleVersion string

// Enum values for CostCategoryRuleVersion
const (
	CostCategoryRuleVersionCostcategoryexpressionv1 CostCategoryRuleVersion = "CostCategoryExpression.v1"
)

type Dimension string

// Enum values for Dimension
const (
	DimensionAz                   Dimension = "AZ"
	DimensionInstance_type        Dimension = "INSTANCE_TYPE"
	DimensionLinked_account       Dimension = "LINKED_ACCOUNT"
	DimensionLinked_account_name  Dimension = "LINKED_ACCOUNT_NAME"
	DimensionOperation            Dimension = "OPERATION"
	DimensionPurchase_type        Dimension = "PURCHASE_TYPE"
	DimensionRegion               Dimension = "REGION"
	DimensionService              Dimension = "SERVICE"
	DimensionService_code         Dimension = "SERVICE_CODE"
	DimensionUsage_type           Dimension = "USAGE_TYPE"
	DimensionUsage_type_group     Dimension = "USAGE_TYPE_GROUP"
	DimensionRecord_type          Dimension = "RECORD_TYPE"
	DimensionOperating_system     Dimension = "OPERATING_SYSTEM"
	DimensionTenancy              Dimension = "TENANCY"
	DimensionScope                Dimension = "SCOPE"
	DimensionPlatform             Dimension = "PLATFORM"
	DimensionSubscription_id      Dimension = "SUBSCRIPTION_ID"
	DimensionLegal_entity_name    Dimension = "LEGAL_ENTITY_NAME"
	DimensionDeployment_option    Dimension = "DEPLOYMENT_OPTION"
	DimensionDatabase_engine      Dimension = "DATABASE_ENGINE"
	DimensionCache_engine         Dimension = "CACHE_ENGINE"
	DimensionInstance_type_family Dimension = "INSTANCE_TYPE_FAMILY"
	DimensionBilling_entity       Dimension = "BILLING_ENTITY"
	DimensionReservation_id       Dimension = "RESERVATION_ID"
	DimensionResource_id          Dimension = "RESOURCE_ID"
	DimensionRightsizing_type     Dimension = "RIGHTSIZING_TYPE"
	DimensionSavings_plans_type   Dimension = "SAVINGS_PLANS_TYPE"
	DimensionSavings_plan_arn     Dimension = "SAVINGS_PLAN_ARN"
	DimensionPayment_option       Dimension = "PAYMENT_OPTION"
)

type Granularity string

// Enum values for Granularity
const (
	GranularityDaily   Granularity = "DAILY"
	GranularityMonthly Granularity = "MONTHLY"
	GranularityHourly  Granularity = "HOURLY"
)

type GroupDefinitionType string

// Enum values for GroupDefinitionType
const (
	GroupDefinitionTypeDimension     GroupDefinitionType = "DIMENSION"
	GroupDefinitionTypeTag           GroupDefinitionType = "TAG"
	GroupDefinitionTypeCost_category GroupDefinitionType = "COST_CATEGORY"
)

type LookbackPeriodInDays string

// Enum values for LookbackPeriodInDays
const (
	LookbackPeriodInDaysSeven_days  LookbackPeriodInDays = "SEVEN_DAYS"
	LookbackPeriodInDaysThirty_days LookbackPeriodInDays = "THIRTY_DAYS"
	LookbackPeriodInDaysSixty_days  LookbackPeriodInDays = "SIXTY_DAYS"
)

type MatchOption string

// Enum values for MatchOption
const (
	MatchOptionEquals           MatchOption = "EQUALS"
	MatchOptionStarts_with      MatchOption = "STARTS_WITH"
	MatchOptionEnds_with        MatchOption = "ENDS_WITH"
	MatchOptionContains         MatchOption = "CONTAINS"
	MatchOptionCase_sensitive   MatchOption = "CASE_SENSITIVE"
	MatchOptionCase_insensitive MatchOption = "CASE_INSENSITIVE"
)

type Metric string

// Enum values for Metric
const (
	MetricBlended_cost            Metric = "BLENDED_COST"
	MetricUnblended_cost          Metric = "UNBLENDED_COST"
	MetricAmortized_cost          Metric = "AMORTIZED_COST"
	MetricNet_unblended_cost      Metric = "NET_UNBLENDED_COST"
	MetricNet_amortized_cost      Metric = "NET_AMORTIZED_COST"
	MetricUsage_quantity          Metric = "USAGE_QUANTITY"
	MetricNormalized_usage_amount Metric = "NORMALIZED_USAGE_AMOUNT"
)

type OfferingClass string

// Enum values for OfferingClass
const (
	OfferingClassStandard    OfferingClass = "STANDARD"
	OfferingClassConvertible OfferingClass = "CONVERTIBLE"
)

type PaymentOption string

// Enum values for PaymentOption
const (
	PaymentOptionNo_upfront         PaymentOption = "NO_UPFRONT"
	PaymentOptionPartial_upfront    PaymentOption = "PARTIAL_UPFRONT"
	PaymentOptionAll_upfront        PaymentOption = "ALL_UPFRONT"
	PaymentOptionLight_utilization  PaymentOption = "LIGHT_UTILIZATION"
	PaymentOptionMedium_utilization PaymentOption = "MEDIUM_UTILIZATION"
	PaymentOptionHeavy_utilization  PaymentOption = "HEAVY_UTILIZATION"
)

type RecommendationTarget string

// Enum values for RecommendationTarget
const (
	RecommendationTargetSame_instance_family  RecommendationTarget = "SAME_INSTANCE_FAMILY"
	RecommendationTargetCross_instance_family RecommendationTarget = "CROSS_INSTANCE_FAMILY"
)

type RightsizingType string

// Enum values for RightsizingType
const (
	RightsizingTypeTerminate RightsizingType = "TERMINATE"
	RightsizingTypeModify    RightsizingType = "MODIFY"
)

type SupportedSavingsPlansType string

// Enum values for SupportedSavingsPlansType
const (
	SupportedSavingsPlansTypeCompute_sp      SupportedSavingsPlansType = "COMPUTE_SP"
	SupportedSavingsPlansTypeEc2_instance_sp SupportedSavingsPlansType = "EC2_INSTANCE_SP"
)

type TermInYears string

// Enum values for TermInYears
const (
	TermInYearsOne_year    TermInYears = "ONE_YEAR"
	TermInYearsThree_years TermInYears = "THREE_YEARS"
)
