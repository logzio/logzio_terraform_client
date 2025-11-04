package unified_alerts

import (
	"fmt"

	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	unifiedAlertsServiceEndpoint = "%s/poc/unified-alerts"
)

const (
	// Alert types
	TypeLogAlert    string = "LOG_ALERT"
	TypeMetricAlert string = "METRIC_ALERT"

	// URL type parameters
	UrlTypeLogs    string = "logs"
	UrlTypeMetrics string = "metrics"

	// Aggregation types
	AggregationTypeSum         string = "SUM"
	AggregationTypeMin         string = "MIN"
	AggregationTypeMax         string = "MAX"
	AggregationTypeAvg         string = "AVG"
	AggregationTypeCount       string = "COUNT"
	AggregationTypeUniqueCount string = "UNIQUE_COUNT"
	AggregationTypeNone        string = "NONE"

	// Operators
	OperatorLessThan            string = "LESS_THAN"
	OperatorGreaterThan         string = "GREATER_THAN"
	OperatorLessThanOrEquals    string = "LESS_THAN_OR_EQUALS"
	OperatorGreaterThanOrEquals string = "GREATER_THAN_OR_EQUALS"
	OperatorEquals              string = "EQUALS"
	OperatorNotEquals           string = "NOT_EQUALS"

	// Severity levels
	SeverityInfo   string = "INFO"
	SeverityLow    string = "LOW"
	SeverityMedium string = "MEDIUM"
	SeverityHigh   string = "HIGH"
	SeveritySevere string = "SEVERE"

	// Trigger types
	TriggerTypeThreshold      string = "THRESHOLD"
	TriggerTypeMathExpression string = "MATH_EXPRESSION"

	// Metric operators
	MetricOperatorAbove        string = "ABOVE"
	MetricOperatorBelow        string = "BELOW"
	MetricOperatorWithinRange  string = "WITHIN_RANGE"
	MetricOperatorOutsideRange string = "OUTSIDE_RANGE"

	// Sort directions
	SortDesc string = "DESC"
	SortAsc  string = "ASC"

	// Output types
	OutputTypeJson  string = "JSON"
	OutputTypeTable string = "TABLE"

	// Operation names for logging
	createUnifiedAlertOperation = "CreateUnifiedAlert"
	getUnifiedAlertOperation    = "GetUnifiedAlert"
	updateUnifiedAlertOperation = "UpdateUnifiedAlert"
	deleteUnifiedAlertOperation = "DeleteUnifiedAlert"

	unifiedAlertResourceName = "unified_alert"
)

type UnifiedAlertsClient struct {
	*client.Client
}

// CreateUnifiedAlert represents the request payload for creating a unified alert
type CreateUnifiedAlert struct {
	Title                               string             `json:"title,omitempty"`
	Type                                string             `json:"type,omitempty"`
	Description                         string             `json:"description,omitempty"`
	Tags                                []string           `json:"tags,omitempty"`
	FolderId                            string             `json:"folderId,omitempty"`
	DashboardId                         string             `json:"dashboardId,omitempty"`
	PanelId                             string             `json:"panelId,omitempty"`
	Runbook                             string             `json:"runbook,omitempty"`
	Rca                                 bool               `json:"rca,omitempty"`
	RcaNotificationEndpointIds          []int              `json:"rcaNotificationEndpointIds,omitempty"`
	UseAlertNotificationEndpointsForRca bool               `json:"useAlertNotificationEndpointsForRca,omitempty"`
	LogAlert                            *LogAlertConfig    `json:"logAlert,omitempty"`
	MetricAlert                         *MetricAlertConfig `json:"metricAlert,omitempty"`
}

// UnifiedAlert represents the response from unified alert operations
type UnifiedAlert struct {
	Title                               string             `json:"title,omitempty"`
	Type                                string             `json:"type,omitempty"`
	Description                         string             `json:"description,omitempty"`
	Tags                                []string           `json:"tags,omitempty"`
	FolderId                            string             `json:"folderId,omitempty"`
	DashboardId                         string             `json:"dashboardId,omitempty"`
	PanelId                             string             `json:"panelId,omitempty"`
	Runbook                             string             `json:"runbook,omitempty"`
	Rca                                 bool               `json:"rca,omitempty"`
	RcaNotificationEndpointIds          []int              `json:"rcaNotificationEndpointIds,omitempty"`
	UseAlertNotificationEndpointsForRca bool               `json:"useAlertNotificationEndpointsForRca,omitempty"`
	LogAlert                            *LogAlertConfig    `json:"logAlert,omitempty"`
	MetricAlert                         *MetricAlertConfig `json:"metricAlert,omitempty"`
	Id                                  string             `json:"id,omitempty"`
	UpdatedAt                           float64            `json:"updatedAt,omitempty"`
	CreatedAt                           float64            `json:"createdAt,omitempty"`
	Enabled                             bool               `json:"enabled,omitempty"`
}

// LogAlertConfig represents the log alert configuration
type LogAlertConfig struct {
	Output                 LogAlertOutput `json:"output,omitempty"`
	SearchTimeFrameMinutes int            `json:"searchTimeFrameMinutes,omitempty"`
	SubComponents          []SubComponent `json:"subComponents,omitempty"`
	Correlations           Correlations   `json:"correlations,omitempty"`
	Schedule               Schedule       `json:"schedule,omitempty"`
}

// LogAlertOutput represents the output configuration for log alerts
type LogAlertOutput struct {
	Recipients                   Recipients `json:"recipients,omitempty"`
	SuppressNotificationsMinutes int        `json:"suppressNotificationsMinutes,omitempty"`
	Type                         string     `json:"type,omitempty"`
}

// Recipients represents notification recipients
type Recipients struct {
	Emails                  []string `json:"emails,omitempty"`
	NotificationEndpointIds []int    `json:"notificationEndpointIds,omitempty"`
}

// SubComponent represents a sub-component of a log alert
type SubComponent struct {
	QueryDefinition QueryDefinition     `json:"queryDefinition,omitempty"`
	Trigger         SubComponentTrigger `json:"trigger,omitempty"`
	Output          SubComponentOutput  `json:"output,omitempty"`
}

// QueryDefinition represents the query definition for a sub-component
type QueryDefinition struct {
	Query                    string      `json:"query,omitempty"`
	Filters                  BoolFilter  `json:"filters,omitempty"`
	GroupBy                  []string    `json:"groupBy,omitempty"`
	Aggregation              Aggregation `json:"aggregation,omitempty"`
	ShouldQueryOnAllAccounts bool        `json:"shouldQueryOnAllAccounts,omitempty"`
	AccountIdsToQueryOn      []int       `json:"accountIdsToQueryOn,omitempty"`
}

// BoolFilter represents boolean filters
type BoolFilter struct {
	Bool FilterLists `json:"bool,omitempty"`
}

// FilterLists represents the filter lists in a boolean filter
type FilterLists struct {
	Must    []map[string]interface{} `json:"must,omitempty"`
	Should  []map[string]interface{} `json:"should,omitempty"`
	Filter  []map[string]interface{} `json:"filter,omitempty"`
	MustNot []map[string]interface{} `json:"must_not,omitempty"`
}

// Aggregation represents aggregation configuration
type Aggregation struct {
	AggregationType    string `json:"aggregationType,omitempty"`
	FieldToAggregateOn string `json:"fieldToAggregateOn,omitempty"`
	ValueToAggregateOn string `json:"valueToAggregateOn,omitempty"`
}

// SubComponentTrigger represents the trigger configuration for a sub-component
type SubComponentTrigger struct {
	Operator               string             `json:"operator,omitempty"`
	SeverityThresholdTiers map[string]float32 `json:"severityThresholdTiers,omitempty"`
}

// SubComponentOutput represents the output configuration for a sub-component
type SubComponentOutput struct {
	Columns            []ColumnConfig `json:"columns,omitempty"`
	ShouldUseAllFields bool           `json:"shouldUseAllFields,omitempty"`
}

// ColumnConfig represents column configuration
type ColumnConfig struct {
	FieldName string `json:"fieldName,omitempty"`
	Regex     string `json:"regex,omitempty"`
	Sort      string `json:"sort,omitempty"`
}

// Correlations represents correlation configuration
type Correlations struct {
	CorrelationOperators []string            `json:"correlationOperators,omitempty"`
	Joins                []map[string]string `json:"joins,omitempty"`
}

// Schedule represents schedule configuration
type Schedule struct {
	CronExpression string `json:"cronExpression,omitempty"`
	Timezone       string `json:"timezone,omitempty"`
}

// MetricAlertConfig represents the metric alert configuration
type MetricAlertConfig struct {
	Severity   string        `json:"severity,omitempty"`
	Trigger    MetricTrigger `json:"trigger,omitempty"`
	Queries    []MetricQuery `json:"queries,omitempty"`
	Recipients Recipients    `json:"recipients,omitempty"`
}

// MetricTrigger represents the trigger configuration for metric alerts
type MetricTrigger struct {
	TriggerType            string  `json:"triggerType,omitempty"`
	MetricOperator         string  `json:"metricOperator,omitempty"`
	MinThreshold           float64 `json:"minThreshold,omitempty"`
	MaxThreshold           float64 `json:"maxThreshold,omitempty"`
	MathExpression         string  `json:"mathExpression,omitempty"`
	SearchTimeFrameMinutes int     `json:"searchTimeFrameMinutes,omitempty"`
}

// MetricQuery represents a metric query
type MetricQuery struct {
	RefId           string                `json:"refId,omitempty"`
	QueryDefinition MetricQueryDefinition `json:"queryDefinition,omitempty"`
}

// MetricQueryDefinition represents the query definition for a metric query
type MetricQueryDefinition struct {
	DatasourceUid string `json:"datasourceUid,omitempty"`
	PromqlQuery   string `json:"promqlQuery,omitempty"`
}

// New creates a new UnifiedAlertsClient
func New(apiToken, baseUrl string) (*UnifiedAlertsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}
	return &UnifiedAlertsClient{
		Client: client.New(apiToken, baseUrl),
	}, nil
}

// validateCreateUnifiedAlertRequest validates the create unified alert request
func validateCreateUnifiedAlertRequest(req CreateUnifiedAlert) error {
	if len(req.Title) == 0 {
		return fmt.Errorf("title must be set")
	}

	if len(req.Type) == 0 {
		return fmt.Errorf("type must be set")
	}

	validAlertTypes := []string{TypeLogAlert, TypeMetricAlert}
	if !logzio_client.Contains(validAlertTypes, req.Type) {
		return fmt.Errorf("type must be one of %v", validAlertTypes)
	}

	// Validate that the appropriate alert config is set based on type
	if req.Type == TypeLogAlert && req.LogAlert == nil {
		return fmt.Errorf("logAlert must be set when type is LOG_ALERT")
	}

	if req.Type == TypeMetricAlert && req.MetricAlert == nil {
		return fmt.Errorf("metricAlert must be set when type is METRIC_ALERT")
	}

	// Validate log alert if present
	if req.LogAlert != nil {
		if err := validateLogAlert(req.LogAlert); err != nil {
			return err
		}
	}

	// Validate metric alert if present
	if req.MetricAlert != nil {
		if err := validateMetricAlert(req.MetricAlert); err != nil {
			return err
		}
	}

	return nil
}

// validateLogAlert validates log alert configuration
func validateLogAlert(logAlert *LogAlertConfig) error {
	if logAlert.SubComponents == nil || len(logAlert.SubComponents) == 0 {
		return fmt.Errorf("logAlert.subComponents must not be empty")
	}

	validAggregationTypes := []string{AggregationTypeSum, AggregationTypeMin, AggregationTypeMax, AggregationTypeAvg, AggregationTypeCount, AggregationTypeUniqueCount, AggregationTypeNone}
	validOperators := []string{OperatorGreaterThanOrEquals, OperatorLessThanOrEquals, OperatorGreaterThan, OperatorLessThan, OperatorNotEquals, OperatorEquals}
	validSeverities := []string{SeverityInfo, SeverityLow, SeverityMedium, SeverityHigh, SeveritySevere}
	validSorts := []string{SortDesc, SortAsc}
	validOutputTypes := []string{OutputTypeJson, OutputTypeTable}

	// Validate output type if present
	if len(logAlert.Output.Type) > 0 {
		if !logzio_client.Contains(validOutputTypes, logAlert.Output.Type) {
			return fmt.Errorf("logAlert.output.type must be one of %v", validOutputTypes)
		}
	}

	// Validate each sub-component
	for i, subComponent := range logAlert.SubComponents {
		if len(subComponent.QueryDefinition.Query) == 0 {
			return fmt.Errorf("logAlert.subComponents[%d].queryDefinition.query must be set", i)
		}

		// Validate shouldQueryOnAllAccounts and accountIdsToQueryOn relationship
		if !subComponent.QueryDefinition.ShouldQueryOnAllAccounts {
			if subComponent.QueryDefinition.AccountIdsToQueryOn == nil || len(subComponent.QueryDefinition.AccountIdsToQueryOn) == 0 {
				return fmt.Errorf("logAlert.subComponents[%d].queryDefinition.accountIdsToQueryOn must be set when shouldQueryOnAllAccounts is false", i)
			}
		}

		if len(subComponent.QueryDefinition.Aggregation.AggregationType) > 0 {
			if !logzio_client.Contains(validAggregationTypes, subComponent.QueryDefinition.Aggregation.AggregationType) {
				return fmt.Errorf("logAlert.subComponents[%d].queryDefinition.aggregation.aggregationType must be one of %v", i, validAggregationTypes)
			}
		}

		if len(subComponent.Trigger.Operator) > 0 {
			if !logzio_client.Contains(validOperators, subComponent.Trigger.Operator) {
				return fmt.Errorf("logAlert.subComponents[%d].trigger.operator must be one of %v", i, validOperators)
			}
		}

		for severity := range subComponent.Trigger.SeverityThresholdTiers {
			if !logzio_client.Contains(validSeverities, severity) {
				return fmt.Errorf("logAlert.subComponents[%d].trigger.severityThresholdTiers contains invalid severity: %s, must be one of %v", i, severity, validSeverities)
			}
		}

		if subComponent.Output.Columns != nil && len(subComponent.Output.Columns) > 0 {
			for j, column := range subComponent.Output.Columns {
				if len(column.Sort) > 0 {
					if !logzio_client.Contains(validSorts, column.Sort) {
						return fmt.Errorf("logAlert.subComponents[%d].output.columns[%d].sort must be one of %v", i, j, validSorts)
					}
				}
			}
		}
	}

	return nil
}

// validateMetricAlert validates metric alert configuration
func validateMetricAlert(metricAlert *MetricAlertConfig) error {
	validSeverities := []string{SeverityInfo, SeverityLow, SeverityMedium, SeverityHigh, SeveritySevere}
	validTriggerTypes := []string{TriggerTypeThreshold, TriggerTypeMathExpression}
	validMetricOperators := []string{MetricOperatorAbove, MetricOperatorBelow, MetricOperatorWithinRange, MetricOperatorOutsideRange}

	if len(metricAlert.Severity) > 0 {
		if !logzio_client.Contains(validSeverities, metricAlert.Severity) {
			return fmt.Errorf("metricAlert.severity must be one of %v", validSeverities)
		}
	}

	if len(metricAlert.Trigger.TriggerType) > 0 {
		if !logzio_client.Contains(validTriggerTypes, metricAlert.Trigger.TriggerType) {
			return fmt.Errorf("metricAlert.trigger.triggerType must be one of %v", validTriggerTypes)
		}
	}

	if len(metricAlert.Trigger.MetricOperator) > 0 {
		if !logzio_client.Contains(validMetricOperators, metricAlert.Trigger.MetricOperator) {
			return fmt.Errorf("metricAlert.trigger.metricOperator must be one of %v", validMetricOperators)
		}
	}

	return nil
}

// validateUrlType validates the URL type parameter
func validateUrlType(urlType string) error {
	validUrlTypes := []string{UrlTypeLogs, UrlTypeMetrics}
	if !logzio_client.Contains(validUrlTypes, urlType) {
		return fmt.Errorf("alertType must be one of %v", validUrlTypes)
	}
	return nil
}
