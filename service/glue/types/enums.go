// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CatalogEncryptionMode string

// Enum values for CatalogEncryptionMode
const (
	CatalogEncryptionModeDisabled CatalogEncryptionMode = "DISABLED"
	CatalogEncryptionModeSsekms   CatalogEncryptionMode = "SSE-KMS"
)

type CloudWatchEncryptionMode string

// Enum values for CloudWatchEncryptionMode
const (
	CloudWatchEncryptionModeDisabled CloudWatchEncryptionMode = "DISABLED"
	CloudWatchEncryptionModeSsekms   CloudWatchEncryptionMode = "SSE-KMS"
)

type ColumnStatisticsType string

// Enum values for ColumnStatisticsType
const (
	ColumnStatisticsTypeBoolean ColumnStatisticsType = "BOOLEAN"
	ColumnStatisticsTypeDate    ColumnStatisticsType = "DATE"
	ColumnStatisticsTypeDecimal ColumnStatisticsType = "DECIMAL"
	ColumnStatisticsTypeDouble  ColumnStatisticsType = "DOUBLE"
	ColumnStatisticsTypeLong    ColumnStatisticsType = "LONG"
	ColumnStatisticsTypeString  ColumnStatisticsType = "STRING"
	ColumnStatisticsTypeBinary  ColumnStatisticsType = "BINARY"
)

type Comparator string

// Enum values for Comparator
const (
	ComparatorEquals              Comparator = "EQUALS"
	ComparatorGreater_than        Comparator = "GREATER_THAN"
	ComparatorLess_than           Comparator = "LESS_THAN"
	ComparatorGreater_than_equals Comparator = "GREATER_THAN_EQUALS"
	ComparatorLess_than_equals    Comparator = "LESS_THAN_EQUALS"
)

type ConnectionPropertyKey string

// Enum values for ConnectionPropertyKey
const (
	ConnectionPropertyKeyHost                              ConnectionPropertyKey = "HOST"
	ConnectionPropertyKeyPort                              ConnectionPropertyKey = "PORT"
	ConnectionPropertyKeyUser_name                         ConnectionPropertyKey = "USERNAME"
	ConnectionPropertyKeyPassword                          ConnectionPropertyKey = "PASSWORD"
	ConnectionPropertyKeyEncrypted_password                ConnectionPropertyKey = "ENCRYPTED_PASSWORD"
	ConnectionPropertyKeyJdbc_driver_jar_uri               ConnectionPropertyKey = "JDBC_DRIVER_JAR_URI"
	ConnectionPropertyKeyJdbc_driver_class_name            ConnectionPropertyKey = "JDBC_DRIVER_CLASS_NAME"
	ConnectionPropertyKeyJdbc_engine                       ConnectionPropertyKey = "JDBC_ENGINE"
	ConnectionPropertyKeyJdbc_engine_version               ConnectionPropertyKey = "JDBC_ENGINE_VERSION"
	ConnectionPropertyKeyConfig_files                      ConnectionPropertyKey = "CONFIG_FILES"
	ConnectionPropertyKeyInstance_id                       ConnectionPropertyKey = "INSTANCE_ID"
	ConnectionPropertyKeyJdbc_connection_url               ConnectionPropertyKey = "JDBC_CONNECTION_URL"
	ConnectionPropertyKeyJdbc_enforce_ssl                  ConnectionPropertyKey = "JDBC_ENFORCE_SSL"
	ConnectionPropertyKeyCustom_jdbc_cert                  ConnectionPropertyKey = "CUSTOM_JDBC_CERT"
	ConnectionPropertyKeySkip_custom_jdbc_cert_validation  ConnectionPropertyKey = "SKIP_CUSTOM_JDBC_CERT_VALIDATION"
	ConnectionPropertyKeyCustom_jdbc_cert_string           ConnectionPropertyKey = "CUSTOM_JDBC_CERT_STRING"
	ConnectionPropertyKeyConnection_url                    ConnectionPropertyKey = "CONNECTION_URL"
	ConnectionPropertyKeyKafka_bootstrap_servers           ConnectionPropertyKey = "KAFKA_BOOTSTRAP_SERVERS"
	ConnectionPropertyKeyKafka_ssl_enabled                 ConnectionPropertyKey = "KAFKA_SSL_ENABLED"
	ConnectionPropertyKeyKafka_custom_cert                 ConnectionPropertyKey = "KAFKA_CUSTOM_CERT"
	ConnectionPropertyKeyKafka_skip_custom_cert_validation ConnectionPropertyKey = "KAFKA_SKIP_CUSTOM_CERT_VALIDATION"
)

type ConnectionType string

// Enum values for ConnectionType
const (
	ConnectionTypeJdbc    ConnectionType = "JDBC"
	ConnectionTypeSftp    ConnectionType = "SFTP"
	ConnectionTypeMongodb ConnectionType = "MONGODB"
	ConnectionTypeKafka   ConnectionType = "KAFKA"
)

type CrawlerState string

// Enum values for CrawlerState
const (
	CrawlerStateReady    CrawlerState = "READY"
	CrawlerStateRunning  CrawlerState = "RUNNING"
	CrawlerStateStopping CrawlerState = "STOPPING"
)

type CrawlState string

// Enum values for CrawlState
const (
	CrawlStateRunning    CrawlState = "RUNNING"
	CrawlStateCancelling CrawlState = "CANCELLING"
	CrawlStateCancelled  CrawlState = "CANCELLED"
	CrawlStateSucceeded  CrawlState = "SUCCEEDED"
	CrawlStateFailed     CrawlState = "FAILED"
)

type CsvHeaderOption string

// Enum values for CsvHeaderOption
const (
	CsvHeaderOptionUnknown CsvHeaderOption = "UNKNOWN"
	CsvHeaderOptionPresent CsvHeaderOption = "PRESENT"
	CsvHeaderOptionAbsent  CsvHeaderOption = "ABSENT"
)

type DeleteBehavior string

// Enum values for DeleteBehavior
const (
	DeleteBehaviorLog                   DeleteBehavior = "LOG"
	DeleteBehaviorDelete_from_database  DeleteBehavior = "DELETE_FROM_DATABASE"
	DeleteBehaviorDeprecate_in_database DeleteBehavior = "DEPRECATE_IN_DATABASE"
)

type EnableHybridValues string

// Enum values for EnableHybridValues
const (
	EnableHybridValuesTrue  EnableHybridValues = "TRUE"
	EnableHybridValuesFalse EnableHybridValues = "FALSE"
)

type ExistCondition string

// Enum values for ExistCondition
const (
	ExistConditionMust_exist ExistCondition = "MUST_EXIST"
	ExistConditionNot_exist  ExistCondition = "NOT_EXIST"
	ExistConditionNone       ExistCondition = "NONE"
)

type JobBookmarksEncryptionMode string

// Enum values for JobBookmarksEncryptionMode
const (
	JobBookmarksEncryptionModeDisabled JobBookmarksEncryptionMode = "DISABLED"
	JobBookmarksEncryptionModeCsekms   JobBookmarksEncryptionMode = "CSE-KMS"
)

type JobRunState string

// Enum values for JobRunState
const (
	JobRunStateStarting  JobRunState = "STARTING"
	JobRunStateRunning   JobRunState = "RUNNING"
	JobRunStateStopping  JobRunState = "STOPPING"
	JobRunStateStopped   JobRunState = "STOPPED"
	JobRunStateSucceeded JobRunState = "SUCCEEDED"
	JobRunStateFailed    JobRunState = "FAILED"
	JobRunStateTimeout   JobRunState = "TIMEOUT"
)

type Language string

// Enum values for Language
const (
	LanguagePython Language = "PYTHON"
	LanguageScala  Language = "SCALA"
)

type LastCrawlStatus string

// Enum values for LastCrawlStatus
const (
	LastCrawlStatusSucceeded LastCrawlStatus = "SUCCEEDED"
	LastCrawlStatusCancelled LastCrawlStatus = "CANCELLED"
	LastCrawlStatusFailed    LastCrawlStatus = "FAILED"
)

type Logical string

// Enum values for Logical
const (
	LogicalAnd Logical = "AND"
	LogicalAny Logical = "ANY"
)

type LogicalOperator string

// Enum values for LogicalOperator
const (
	LogicalOperatorEquals LogicalOperator = "EQUALS"
)

type NodeType string

// Enum values for NodeType
const (
	NodeTypeCrawler NodeType = "CRAWLER"
	NodeTypeJob     NodeType = "JOB"
	NodeTypeTrigger NodeType = "TRIGGER"
)

type Permission string

// Enum values for Permission
const (
	PermissionAll                  Permission = "ALL"
	PermissionSelect               Permission = "SELECT"
	PermissionAlter                Permission = "ALTER"
	PermissionDrop                 Permission = "DROP"
	PermissionDelete               Permission = "DELETE"
	PermissionInsert               Permission = "INSERT"
	PermissionCreate_database      Permission = "CREATE_DATABASE"
	PermissionCreate_table         Permission = "CREATE_TABLE"
	PermissionData_location_access Permission = "DATA_LOCATION_ACCESS"
)

type PrincipalType string

// Enum values for PrincipalType
const (
	PrincipalTypeUser  PrincipalType = "USER"
	PrincipalTypeRole  PrincipalType = "ROLE"
	PrincipalTypeGroup PrincipalType = "GROUP"
)

type ResourceShareType string

// Enum values for ResourceShareType
const (
	ResourceShareTypeForeign ResourceShareType = "FOREIGN"
	ResourceShareTypeAll     ResourceShareType = "ALL"
)

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeJar     ResourceType = "JAR"
	ResourceTypeFile    ResourceType = "FILE"
	ResourceTypeArchive ResourceType = "ARCHIVE"
)

type S3EncryptionMode string

// Enum values for S3EncryptionMode
const (
	S3EncryptionModeDisabled S3EncryptionMode = "DISABLED"
	S3EncryptionModeSsekms   S3EncryptionMode = "SSE-KMS"
	S3EncryptionModeSses3    S3EncryptionMode = "SSE-S3"
)

type ScheduleState string

// Enum values for ScheduleState
const (
	ScheduleStateScheduled     ScheduleState = "SCHEDULED"
	ScheduleStateNot_scheduled ScheduleState = "NOT_SCHEDULED"
	ScheduleStateTransitioning ScheduleState = "TRANSITIONING"
)

type Sort string

// Enum values for Sort
const (
	SortAscending  Sort = "ASC"
	SortDescending Sort = "DESC"
)

type SortDirectionType string

// Enum values for SortDirectionType
const (
	SortDirectionTypeDescending SortDirectionType = "DESCENDING"
	SortDirectionTypeAscending  SortDirectionType = "ASCENDING"
)

type TaskRunSortColumnType string

// Enum values for TaskRunSortColumnType
const (
	TaskRunSortColumnTypeTask_run_type TaskRunSortColumnType = "TASK_RUN_TYPE"
	TaskRunSortColumnTypeStatus        TaskRunSortColumnType = "STATUS"
	TaskRunSortColumnTypeStarted       TaskRunSortColumnType = "STARTED"
)

type TaskStatusType string

// Enum values for TaskStatusType
const (
	TaskStatusTypeStarting  TaskStatusType = "STARTING"
	TaskStatusTypeRunning   TaskStatusType = "RUNNING"
	TaskStatusTypeStopping  TaskStatusType = "STOPPING"
	TaskStatusTypeStopped   TaskStatusType = "STOPPED"
	TaskStatusTypeSucceeded TaskStatusType = "SUCCEEDED"
	TaskStatusTypeFailed    TaskStatusType = "FAILED"
	TaskStatusTypeTimeout   TaskStatusType = "TIMEOUT"
)

type TaskType string

// Enum values for TaskType
const (
	TaskTypeEvaluation              TaskType = "EVALUATION"
	TaskTypeLabeling_set_generation TaskType = "LABELING_SET_GENERATION"
	TaskTypeImport_labels           TaskType = "IMPORT_LABELS"
	TaskTypeExport_labels           TaskType = "EXPORT_LABELS"
	TaskTypeFind_matches            TaskType = "FIND_MATCHES"
)

type TransformSortColumnType string

// Enum values for TransformSortColumnType
const (
	TransformSortColumnTypeName           TransformSortColumnType = "NAME"
	TransformSortColumnTypeTransform_type TransformSortColumnType = "TRANSFORM_TYPE"
	TransformSortColumnTypeStatus         TransformSortColumnType = "STATUS"
	TransformSortColumnTypeCreated        TransformSortColumnType = "CREATED"
	TransformSortColumnTypeLast_modified  TransformSortColumnType = "LAST_MODIFIED"
)

type TransformStatusType string

// Enum values for TransformStatusType
const (
	TransformStatusTypeNot_ready TransformStatusType = "NOT_READY"
	TransformStatusTypeReady     TransformStatusType = "READY"
	TransformStatusTypeDeleting  TransformStatusType = "DELETING"
)

type TransformType string

// Enum values for TransformType
const (
	TransformTypeFind_matches TransformType = "FIND_MATCHES"
)

type TriggerState string

// Enum values for TriggerState
const (
	TriggerStateCreating     TriggerState = "CREATING"
	TriggerStateCreated      TriggerState = "CREATED"
	TriggerStateActivating   TriggerState = "ACTIVATING"
	TriggerStateActivated    TriggerState = "ACTIVATED"
	TriggerStateDeactivating TriggerState = "DEACTIVATING"
	TriggerStateDeactivated  TriggerState = "DEACTIVATED"
	TriggerStateDeleting     TriggerState = "DELETING"
	TriggerStateUpdating     TriggerState = "UPDATING"
)

type TriggerType string

// Enum values for TriggerType
const (
	TriggerTypeScheduled   TriggerType = "SCHEDULED"
	TriggerTypeConditional TriggerType = "CONDITIONAL"
	TriggerTypeOn_demand   TriggerType = "ON_DEMAND"
)

type UpdateBehavior string

// Enum values for UpdateBehavior
const (
	UpdateBehaviorLog                UpdateBehavior = "LOG"
	UpdateBehaviorUpdate_in_database UpdateBehavior = "UPDATE_IN_DATABASE"
)

type WorkerType string

// Enum values for WorkerType
const (
	WorkerTypeStandard WorkerType = "Standard"
	WorkerTypeG1x      WorkerType = "G.1X"
	WorkerTypeG2x      WorkerType = "G.2X"
)

type WorkflowRunStatus string

// Enum values for WorkflowRunStatus
const (
	WorkflowRunStatusRunning   WorkflowRunStatus = "RUNNING"
	WorkflowRunStatusCompleted WorkflowRunStatus = "COMPLETED"
	WorkflowRunStatusStopping  WorkflowRunStatus = "STOPPING"
	WorkflowRunStatusStopped   WorkflowRunStatus = "STOPPED"
)
