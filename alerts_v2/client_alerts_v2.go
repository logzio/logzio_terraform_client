package alerts_v2

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"strconv"
)

const (
	alertsServiceEndpoint = "%s/v2/alerts"
)

const (
	AggregationTypeSum         = "SUM"
	AggregationTypeMin         = "MIN"
	AggregationTypeMax         = "MAX"
	AggregationTypeAvg         = "AVG"
	AggregationTypeCount       = "COUNT"
	AggregationTypeUniqueCount = "UNIQUE_COUNT"
	AggregationTypeNone        = "NONE"

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

	SortDesc = "DESC"
	SortAsc  = "ASC"

	OutputTypeJson  = "JSON"
	OutputTypeTable = "TABLE"

	CorrelationOperatorAnd = "AND"

	createAlertOperation  = "CreateAlertV2"
	deleteAlertOperation  = "DeleteAlertV2"
	disableAlertOperation = "DisableAlertV2"
	enableAlertOperation  = "EnableAlertV2"
	getAlertOperation     = "GetAlertV2"
	listAlertOperation    = "ListAlertsV2"
	updateAlertOperation  = "UpdateAlertV2"

	alertResourceName = "alert"
)

type AlertsV2Client struct {
	*client.Client
}

type CreateAlertType struct {
	Title                  string              `json:"title,omitempty"`
	Description            string              `json:"description,omitempty"`
	Tags                   []string            `json:"tags,omitempty"`
	Output                 AlertOutput         `json:"output"`
	SearchTimeFrameMinutes int                 `json:"searchTimeFrameMinutes,omitempty"`
	SubComponents          []SubAlert          `json:"subComponents,omitempty"`
	Correlations           SubAlertCorrelation `json:"correlations"`
	Enabled                string              `json:"enabled,omitempty"`
	Schedule               ScheduleObj         `json:"schedule"`
}

type AlertOutput struct {
	Recipients                   AlertRecipients `json:"recipients"`
	SuppressNotificationsMinutes int             `json:"suppressNotificationsMinutes,omitempty"`
	Type                         string          `json:"type,omitempty"`
}

type AlertRecipients struct {
	Emails                  []string `json:"emails,omitempty"`
	NotificationEndpointIds []int    `json:"notificationEndpointIds,omitempty"`
}

type SubAlert struct {
	QueryDefinition AlertQuery     `json:"queryDefinition"`
	Trigger         AlertTrigger   `json:"trigger"`
	Output          SubAlertOutput `json:"output"`
}

type AlertQuery struct {
	Query                    string         `json:"query,omitempty"`
	Filters                  BoolFilter     `json:"filters"`
	GroupBy                  []string       `json:"groupBy,omitempty"`
	Aggregation              AggregationObj `json:"aggregation"`
	ShouldQueryOnAllAccounts bool           `json:"shouldQueryOnAllAccounts,omitempty"`
	AccountIdsToQueryOn      []int          `json:"accountIdsToQueryOn,omitempty"`
}

type BoolFilter struct {
	Bool FilterLists `json:"bool"`
}

type FilterLists struct {
	Must    []map[string]any `json:"must,omitempty"`
	MustNot []map[string]any `json:"must_not,omitempty"`
}

type AggregationObj struct {
	AggregationType    string `json:"aggregationType,omitempty"`
	FieldToAggregateOn string `json:"fieldToAggregateOn,omitempty"`
}

type AlertTrigger struct {
	Operator               string             `json:"operator,omitempty"`
	SeverityThresholdTiers map[string]float32 `json:"severityThresholdTiers,omitempty"`
}

type SubAlertOutput struct {
	Columns            []ColumnConfig `json:"columns,omitempty"`
	ShouldUseAllFields bool           `json:"shouldUseAllFields,omitempty"`
}

type ColumnConfig struct {
	FieldName string `json:"fieldName,omitempty"`
	Regex     string `json:"regex,omitempty"`
	Sort      string `json:"sort,omitempty"`
}

type SubAlertCorrelation struct {
	CorrelationOperators []string            `json:"correlationOperators,omitempty"`
	Joins                []map[string]string `json:"joins,omitempty"`
}

type ScheduleObj struct {
	CronExpression string `json:"cronExpression,omitempty"`
	Timezone       string `json:"timezone,omitempty"`
}

type AlertType struct {
	AlertId                int64               `json:"id"`
	UpdatedAt              string              `json:"updatedAt"`
	UpdatedBy              string              `json:"updatedBy"`
	CreatedAt              string              `json:"createdAt"`
	CreatedBy              string              `json:"createdBy"`
	Enabled                bool                `json:"enabled"`
	Title                  string              `json:"title"`
	Description            string              `json:"description,omitempty"`
	Tags                   []string            `json:"tags,omitempty"`
	Output                 AlertOutput         `json:"output"`
	SearchTimeFrameMinutes int                 `json:"searchTimeFrameMinutes"`
	SubComponents          []SubAlert          `json:"subComponents"`
	Correlations           SubAlertCorrelation `json:"correlations"`
	Schedule               ScheduleObj         `json:"schedule"`
}

func New(apiToken, baseUrl string) (*AlertsV2Client, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("base URL not defined")
	}
	c := &AlertsV2Client{
		Client: client.New(apiToken, baseUrl),
	}
	return c, nil
}

func validateCreateAlertRequest(alert CreateAlertType) error {
	if len(alert.Title) == 0 {
		return fmt.Errorf("title must be set")
	}

	if len(alert.SubComponents) == 0 {
		return fmt.Errorf("subComponents must be not empty")
	}

	if len(alert.Enabled) > 0 {
		if alert.Enabled != strconv.FormatBool(true) && alert.Enabled != strconv.FormatBool(false) {
			return fmt.Errorf("enabled field must be %s or %s", strconv.FormatBool(true), strconv.FormatBool(false))
		}
	}

	validAggregationTypes := []string{AggregationTypeSum, AggregationTypeMin, AggregationTypeMax, AggregationTypeAvg, AggregationTypeCount, AggregationTypeUniqueCount, AggregationTypeNone}
	validOperations := []string{OperatorGreaterThanOrEquals, OperatorLessThanOrEquals, OperatorGreaterThan, OperatorLessThan, OperatorNotEquals, OperatorEquals}
	validSeverities := []string{SeverityInfo, SeverityLow, SeverityMedium, SeverityHigh, SeveritySevere}
	validSorts := []string{SortDesc, SortAsc}

	for _, subComponent := range alert.SubComponents {
		if len(subComponent.QueryDefinition.Query) == 0 {
			return fmt.Errorf("query string must be set")
		}

		if !logzio_client.Contains(validAggregationTypes, subComponent.QueryDefinition.Aggregation.AggregationType) {
			return fmt.Errorf("valueAggregationType must be one of %s", validAggregationTypes)
		}

		if !logzio_client.Contains(validOperations, subComponent.Trigger.Operator) {
			return fmt.Errorf("operation must be one of %s", validOperations)
		}

		for severity := range subComponent.Trigger.SeverityThresholdTiers {
			if !logzio_client.Contains(validSeverities, severity) {
				return fmt.Errorf("severity must be one of %s", validSeverities)
			}
		}

		if len(subComponent.Output.Columns) > 0 {
			for _, column := range subComponent.Output.Columns {
				if len(column.Sort) > 0 {
					if !logzio_client.Contains(validSorts, column.Sort) {
						return fmt.Errorf("sort must be one of %s", validSorts)
					}
				}
			}
		}
	}

	return nil
}
