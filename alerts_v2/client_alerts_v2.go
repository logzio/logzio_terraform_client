package alerts_v2

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"strconv"
)

const (
	alertsServiceEndpoint         string = "%s/v2/alerts"
)

const (
	AggregationTypeSum         string = "SUM"
	AggregationTypeMin         string = "MIN"
	AggregationTypeMax         string = "MAX"
	AggregationTypeAvg         string = "AVG"
	AggregationTypeCount       string = "COUNT"
	AggregationTypeUniqueCount string = "UNIQUE_COUNT"
	AggregationTypeNone        string = "NONE"

	OperatorLessThan            string = "LESS_THAN"
	OperatorGreaterThan         string = "GREATER_THAN"
	OperatorLessThanOrEquals    string = "LESS_THAN_OR_EQUALS"
	OperatorGreaterThanOrEquals string = "GREATER_THAN_OR_EQUALS"
	OperatorEquals              string = "EQUALS"
	OperatorNotEquals           string = "NOT_EQUALS"

	SeverityInfo   string = "INFO"
	SeverityLow    string = "LOW"
	SeverityMedium string = "MEDIUM"
	SeverityHigh   string = "HIGH"
	SeveritySevere string = "SEVERE"

	SortDesc string = "DESC"
	SortAsc  string = "ASC"

	OutputTypeJson  string = "JSON"
	OutputTypeTable string = "TABLE"

	CorrelationOperatorAnd string = "AND"

	createAlertOperation = "CreateAlertV2"
	deleteAlertOperation = "DeleteAlertV2"
	disableAlertOperation = "DisableAlertV2"
	enableAlertOperation = "EnableAlertV2"
	getAlertOperation = "GetAlertV2"
)

type AlertsV2Client struct {
	*client.Client
}

type CreateAlertType struct {
	Title                  string              `json:"title,omitempty"`
	Description            string              `json:"description,omitempty"`
	Tags                   []string            `json:"tags,omitempty"`
	Output                 AlertOutput         `json:"output,omitempty"`
	SearchTimeFrameMinutes int                 `json:"searchTimeFrameMinutes,omitempty"`
	SubComponents          []SubAlert          `json:"subComponents,omitempty"`
	Correlations           SubAlertCorrelation `json:"correlations,omitempty"`
	Enabled                string              `json:"enabled,omitempty"`
}

type AlertOutput struct {
	Recipients                   AlertRecipients `json:"recipients,omitempty"`
	SuppressNotificationsMinutes int             `json:"suppressNotificationsMinutes,omitempty"`
	Type                         string          `json:"type,omitempty"`
}

type AlertRecipients struct {
	Emails                  []string `json:"emails,omitempty"`
	NotificationEndpointIds []int    `json:"notificationEndpointIds,omitempty"`
}

type SubAlert struct {
	QueryDefinition AlertQuery     `json:"queryDefinition,omitempty"`
	Trigger         AlertTrigger   `json:"trigger,omitempty"`
	Output          SubAlertOutput `json:"output,omitempty"`
}

type AlertQuery struct {
	Query                    string         `json:"query,omitempty"`
	Filters                  BoolFilter     `json:"filters,omitempty"`
	GroupBy                  []string       `json:"groupBy,omitempty"`
	Aggregation              AggregationObj `json:"aggregation,omitempty"`
	ShouldQueryOnAllAccounts bool           `json:"shouldQueryOnAllAccounts,omitempty"`
	AccountIdsToQueryOn      []int          `json:"accountIdsToQueryOn,omitempty,omitempty"`
}

type BoolFilter struct {
	Bool FilterLists `json:"bool,omitempty"`
}

type FilterLists struct {
	Must    []map[string]interface{} `json:"must,omitempty"`
	MustNot []map[string]interface{} `json:"must_not,omitempty"`
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
	Output                 AlertOutput         `json:"output,omitempty"`
	SearchTimeFrameMinutes int                 `json:"searchTimeFrameMinutes"`
	SubComponents          []SubAlert          `json:"subComponents"`
	Correlations           SubAlertCorrelation `json:"correlations,omitempty"`
}

func New(apiToken, baseUrl string) (*AlertsV2Client, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
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

	if alert.SubComponents == nil || len(alert.SubComponents) == 0 {
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

		if subComponent.Output.Columns != nil && len(subComponent.Output.Columns) > 0 {
			for _, column := range subComponent.Output.Columns {
				if !logzio_client.Contains(validSorts, column.Sort) {
					return fmt.Errorf("sort must be one of %s", validSorts)
				}
			}
		}
	}

	return nil
}
