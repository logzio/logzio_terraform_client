package unified_alerts

import (
	"fmt"

	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	unifiedAlertsServiceEndpoint = "%s/v2/unified-alerts"
)

const (
	TypeLogAlert    = "LOG_ALERT"
	TypeMetricAlert = "METRIC_ALERT"

	UrlTypeLogs    = "logs"
	UrlTypeMetrics = "metrics"

	AggregationTypeSum         = "SUM"
	AggregationTypeMin         = "MIN"
	AggregationTypeMax         = "MAX"
	AggregationTypeAvg         = "AVG"
	AggregationTypeCount       = "COUNT"
	AggregationTypeUniqueCount = "UNIQUE_COUNT"
	AggregationTypeNone        = "NONE"
	AggregationTypePercentage  = "PERCENTAGE"
	AggregationTypePercentile  = "PERCENTILE"

	OperatorLessThan            = "LESS_THAN"
	OperatorGreaterThan         = "GREATER_THAN"
	OperatorLessThanOrEquals    = "LESS_THAN_OR_EQUALS"
	OperatorGreaterThanOrEquals = "GREATER_THAN_OR_EQUALS"
	OperatorEquals              = "EQUALS"
	OperatorNotEquals           = "NOT_EQUALS"

	SeverityInfo   = "INFO"
	SeverityLow    = "LOW"
	SeverityMedium = "MEDIUM"
	SeverityHigh   = "HIGH"
	SeveritySevere = "SEVERE"

	TriggerTypeThreshold = "threshold"
	TriggerTypeMath      = "math"

	OperatorTypeAbove        = "above"
	OperatorTypeBelow        = "below"
	OperatorTypeWithinRange  = "within_range"
	OperatorTypeOutsideRange = "outside_range"

	SortDesc = "DESC"
	SortAsc  = "ASC"

	OutputTypeJson  = "JSON"
	OutputTypeTable = "TABLE"

	createUnifiedAlertOperation = "CreateUnifiedAlert"
	getUnifiedAlertOperation    = "GetUnifiedAlert"
	updateUnifiedAlertOperation = "UpdateUnifiedAlert"
	deleteUnifiedAlertOperation = "DeleteUnifiedAlert"

	unifiedAlertResourceName = "unified_alert"
)

type UnifiedAlertsClient struct {
	*client.Client
}

type LinkedPanel struct {
	FolderId    string `json:"folderId,omitempty"`
	DashboardId string `json:"dashboardId,omitempty"`
	PanelId     string `json:"panelId,omitempty"`
}

type AlertConfiguration struct {
	Type string `json:"type,omitempty"`
	// Log alert fields
	SuppressNotificationsMinutes int            `json:"suppressNotificationsMinutes,omitempty"`
	AlertOutputTemplateType      string         `json:"alertOutputTemplateType,omitempty"`
	SearchTimeFrameMinutes       int            `json:"searchTimeFrameMinutes,omitempty"`
	SubComponents                []SubComponent `json:"subComponents,omitempty"`
	Correlations                 *Correlations  `json:"correlations,omitempty"`
	Schedule                     *Schedule      `json:"schedule,omitempty"`
	// Metric alert fields
	Severity string              `json:"severity,omitempty"`
	Trigger  *MetricAlertTrigger `json:"trigger,omitempty"`
	Queries  []MetricQuery       `json:"queries,omitempty"`
}

type MetricAlertTrigger struct {
	Type       string            `json:"type,omitempty"`
	Condition  *TriggerCondition `json:"condition,omitempty"`
	Expression string            `json:"expression,omitempty"`
}

type TriggerCondition struct {
	OperatorType string   `json:"operatorType,omitempty"`
	Threshold    *float64 `json:"threshold,omitempty"`
	From         *float64 `json:"from,omitempty"`
	To           *float64 `json:"to,omitempty"`
}

type CreateUnifiedAlert struct {
	Title                               string              `json:"title,omitempty"`
	Description                         string              `json:"description,omitempty"`
	Tags                                []string            `json:"tags,omitempty"`
	LinkedPanel                         *LinkedPanel        `json:"linkedPanel,omitempty"`
	Runbook                             string              `json:"runbook,omitempty"`
	Enabled                             *bool               `json:"enabled,omitempty"`
	Rca                                 bool                `json:"rca,omitempty"`
	RcaNotificationEndpointIds          []int               `json:"rcaNotificationEndpointIds,omitempty"`
	UseAlertNotificationEndpointsForRca bool                `json:"useAlertNotificationEndpointsForRca,omitempty"`
	Recipients                          *Recipients         `json:"recipients,omitempty"`
	AlertConfiguration                  *AlertConfiguration `json:"alertConfiguration,omitempty"`
}

type UnifiedAlert struct {
	Title                               string              `json:"title,omitempty"`
	Description                         string              `json:"description,omitempty"`
	Tags                                []string            `json:"tags,omitempty"`
	LinkedPanel                         *LinkedPanel        `json:"linkedPanel,omitempty"`
	Runbook                             string              `json:"runbook,omitempty"`
	Rca                                 bool                `json:"rca,omitempty"`
	RcaNotificationEndpointIds          []int               `json:"rcaNotificationEndpointIds,omitempty"`
	UseAlertNotificationEndpointsForRca bool                `json:"useAlertNotificationEndpointsForRca,omitempty"`
	Recipients                          *Recipients         `json:"recipients,omitempty"`
	AlertConfiguration                  *AlertConfiguration `json:"alertConfiguration,omitempty"`
	Id                                  string              `json:"id,omitempty"`
	UpdatedAt                           float64             `json:"updatedAt,omitempty"`
	CreatedAt                           float64             `json:"createdAt,omitempty"`
	CreatedBy                           string              `json:"createdBy,omitempty"`
	UpdatedBy                           string              `json:"updatedBy,omitempty"`
	Enabled                             bool                `json:"enabled,omitempty"`
}

type Recipients struct {
	Emails                  []string `json:"emails,omitempty"`
	NotificationEndpointIds []int    `json:"notificationEndpointIds,omitempty"`
}

type SubComponent struct {
	QueryDefinition QueryDefinition     `json:"queryDefinition"`
	Trigger         SubComponentTrigger `json:"trigger"`
	Output          SubComponentOutput  `json:"output"`
}

type QueryDefinition struct {
	Query                    string      `json:"query,omitempty"`
	Filters                  BoolFilter  `json:"filters"`
	GroupBy                  []string    `json:"groupBy,omitempty"`
	Aggregation              Aggregation `json:"aggregation"`
	ShouldQueryOnAllAccounts bool        `json:"shouldQueryOnAllAccounts,omitempty"`
	AccountIdsToQueryOn      []int       `json:"accountIdsToQueryOn,omitempty"`
}

type BoolFilter struct {
	Bool FilterLists `json:"bool"`
}

type FilterLists struct {
	Must    []map[string]any `json:"must,omitempty"`
	Should  []map[string]any `json:"should,omitempty"`
	Filter  []map[string]any `json:"filter,omitempty"`
	MustNot []map[string]any `json:"must_not,omitempty"`
}

type Aggregation struct {
	AggregationType    string `json:"aggregationType,omitempty"`
	FieldToAggregateOn string `json:"fieldToAggregateOn,omitempty"`
	ValueToAggregateOn string `json:"valueToAggregateOn,omitempty"`
}

type SubComponentTrigger struct {
	Operator               string             `json:"operator,omitempty"`
	SeverityThresholdTiers map[string]float32 `json:"severityThresholdTiers,omitempty"`
}

type SubComponentOutput struct {
	Columns            []ColumnConfig `json:"columns,omitempty"`
	ShouldUseAllFields bool           `json:"shouldUseAllFields,omitempty"`
}

type ColumnConfig struct {
	FieldName string `json:"fieldName,omitempty"`
	Regex     string `json:"regex,omitempty"`
	Sort      string `json:"sort,omitempty"`
}

type Correlations struct {
	CorrelationOperators []string            `json:"correlationOperators,omitempty"`
	Joins                []map[string]string `json:"joins,omitempty"`
}

type Schedule struct {
	CronExpression string `json:"cronExpression,omitempty"`
	Timezone       string `json:"timezone,omitempty"`
}

type MetricQuery struct {
	RefId           string                `json:"refId,omitempty"`
	QueryDefinition MetricQueryDefinition `json:"queryDefinition"`
}

type MetricQueryDefinition struct {
	AccountId   int32  `json:"accountId,omitempty"`
	PromqlQuery string `json:"promqlQuery,omitempty"`
}

func New(apiToken, baseUrl string) (*UnifiedAlertsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("base URL not defined")
	}
	return &UnifiedAlertsClient{
		Client: client.New(apiToken, baseUrl),
	}, nil
}

func validateCreateUnifiedAlertRequest(req CreateUnifiedAlert) error {
	if len(req.Title) == 0 {
		return fmt.Errorf("title must be set")
	}

	if req.AlertConfiguration == nil {
		return fmt.Errorf("alertConfiguration must be set")
	}

	if len(req.AlertConfiguration.Type) == 0 {
		return fmt.Errorf("alertConfiguration.type must be set")
	}

	validAlertTypes := []string{TypeLogAlert, TypeMetricAlert}
	if !logzio_client.Contains(validAlertTypes, req.AlertConfiguration.Type) {
		return fmt.Errorf("alertConfiguration.type must be one of %v", validAlertTypes)
	}

	if req.AlertConfiguration.Type == TypeLogAlert {
		if err := validateLogAlertConfiguration(req.AlertConfiguration); err != nil {
			return err
		}
	}

	if req.AlertConfiguration.Type == TypeMetricAlert {
		if err := validateMetricAlertConfiguration(req.AlertConfiguration); err != nil {
			return err
		}
	}

	return nil
}

func validateLogAlertConfiguration(config *AlertConfiguration) error {
	if len(config.SubComponents) == 0 {
		return fmt.Errorf("alertConfiguration.subComponents must not be empty")
	}

	validAggregationTypes := []string{AggregationTypeSum, AggregationTypeMin, AggregationTypeMax, AggregationTypeAvg, AggregationTypeCount, AggregationTypeUniqueCount, AggregationTypeNone, AggregationTypePercentage, AggregationTypePercentile}
	validOperators := []string{OperatorGreaterThanOrEquals, OperatorLessThanOrEquals, OperatorGreaterThan, OperatorLessThan, OperatorNotEquals, OperatorEquals}
	validSeverities := []string{SeverityInfo, SeverityLow, SeverityMedium, SeverityHigh, SeveritySevere}
	validSorts := []string{SortDesc, SortAsc}
	validOutputTypes := []string{OutputTypeJson, OutputTypeTable}

	if len(config.AlertOutputTemplateType) > 0 {
		if !logzio_client.Contains(validOutputTypes, config.AlertOutputTemplateType) {
			return fmt.Errorf("alertConfiguration.alertOutputTemplateType must be one of %v", validOutputTypes)
		}
	}

	for i, subComponent := range config.SubComponents {
		if config.AlertOutputTemplateType == OutputTypeTable {
			if len(subComponent.Output.Columns) == 0 {
				return fmt.Errorf("alertConfiguration.subComponents[%d].output.columns must be defined when alertOutputTemplateType is TABLE", i)
			}
		}

		if len(subComponent.QueryDefinition.Query) == 0 {
			return fmt.Errorf("alertConfiguration.subComponents[%d].queryDefinition.query must be set", i)
		}

		if !subComponent.QueryDefinition.ShouldQueryOnAllAccounts {
			if len(subComponent.QueryDefinition.AccountIdsToQueryOn) == 0 {
				return fmt.Errorf("alertConfiguration.subComponents[%d].queryDefinition.accountIdsToQueryOn must be set when shouldQueryOnAllAccounts is false", i)
			}
		}

		if len(subComponent.QueryDefinition.Aggregation.AggregationType) > 0 {
			if !logzio_client.Contains(validAggregationTypes, subComponent.QueryDefinition.Aggregation.AggregationType) {
				return fmt.Errorf("alertConfiguration.subComponents[%d].queryDefinition.aggregation.aggregationType must be one of %v", i, validAggregationTypes)
			}
		}

		if len(subComponent.Trigger.Operator) > 0 {
			if !logzio_client.Contains(validOperators, subComponent.Trigger.Operator) {
				return fmt.Errorf("alertConfiguration.subComponents[%d].trigger.operator must be one of %v", i, validOperators)
			}
		}

		for severity := range subComponent.Trigger.SeverityThresholdTiers {
			if !logzio_client.Contains(validSeverities, severity) {
				return fmt.Errorf("alertConfiguration.subComponents[%d].trigger.severityThresholdTiers contains invalid severity: %s, must be one of %v", i, severity, validSeverities)
			}
		}

		if len(subComponent.Output.Columns) > 0 {
			for j, column := range subComponent.Output.Columns {
				if len(column.Sort) > 0 {
					if !logzio_client.Contains(validSorts, column.Sort) {
						return fmt.Errorf("alertConfiguration.subComponents[%d].output.columns[%d].sort must be one of %v", i, j, validSorts)
					}
				}
			}
		}
	}

	return nil
}

func validateMetricAlertConfiguration(config *AlertConfiguration) error {
	validSeverities := []string{SeverityInfo, SeverityLow, SeverityMedium, SeverityHigh, SeveritySevere}
	validTriggerTypes := []string{TriggerTypeThreshold, TriggerTypeMath}
	validOperatorTypes := []string{OperatorTypeAbove, OperatorTypeBelow, OperatorTypeWithinRange, OperatorTypeOutsideRange}

	if len(config.Severity) > 0 {
		if !logzio_client.Contains(validSeverities, config.Severity) {
			return fmt.Errorf("alertConfiguration.severity must be one of %v", validSeverities)
		}
	}

	if config.Trigger == nil {
		return fmt.Errorf("alertConfiguration.trigger must be set for metric alerts")
	}

	if len(config.Trigger.Type) > 0 {
		if !logzio_client.Contains(validTriggerTypes, config.Trigger.Type) {
			return fmt.Errorf("alertConfiguration.trigger.type must be one of %v", validTriggerTypes)
		}
	}

	if config.Trigger.Type == TriggerTypeThreshold {
		if config.Trigger.Condition == nil {
			return fmt.Errorf("alertConfiguration.trigger.condition must be set for threshold triggers")
		}
		if len(config.Trigger.Condition.OperatorType) > 0 {
			if !logzio_client.Contains(validOperatorTypes, config.Trigger.Condition.OperatorType) {
				return fmt.Errorf("alertConfiguration.trigger.condition.operatorType must be one of %v", validOperatorTypes)
			}
		}
	}

	if config.Trigger.Type == TriggerTypeMath {
		if len(config.Trigger.Expression) == 0 {
			return fmt.Errorf("alertConfiguration.trigger.expression must be set for math triggers")
		}
	}

	if len(config.Queries) == 0 {
		return fmt.Errorf("alertConfiguration.queries must not be empty for metric alerts")
	}

	return nil
}

func validateUrlType(urlType string) error {
	validUrlTypes := []string{UrlTypeLogs, UrlTypeMetrics}
	if !logzio_client.Contains(validUrlTypes, urlType) {
		return fmt.Errorf("alertType must be one of %v", validUrlTypes)
	}
	return nil
}
