package security_rules

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	securityRulesServiceEndpoint = "%s/v2/security/rules"

	SecurityRulesOutputTypeJson  = "JSON"
	SecurityRulesOutputTypeTable = "TABLE"

	SecurityRuleAggregationTypeSum         = "SUM"
	SecurityRuleAggregationTypeMin         = "MIN"
	SecurityRuleAggregationTypeMax         = "MAX"
	SecurityRuleAggregationTypeAvg         = "AVG"
	SecurityRuleAggregationTypeCount       = "COUNT"
	SecurityRuleAggregationTypeUniqueCount = "UNIQUE_COUNT"
	SecurityRuleAggregationTypeNone        = "NONE"

	SecurityRuleOperatorLessThan            = "LESS_THAN"
	SecurityRuleOperatorGreaterThan         = "GREATER_THAN"
	SecurityRuleOperatorLessThanOrEquals    = "LESS_THAN_OR_EQUALS"
	SecurityRuleOperatorGreaterThanOrEquals = "GREATER_THAN_OR_EQUALS"
	SecurityRuleOperatorEquals              = "EQUALS"
	SecurityRuleOperatorNotEquals           = "NOT_EQUALS"

	SecurityRuleSeverityThresholdInfo   = "INFO"
	SecurityRuleSeverityThresholdLow    = "LOW"
	SecurityRuleSeverityThresholdMedium = "MEDIUM"
	SecurityRuleSeverityThresholdHigh   = "HIGH"
	SecurityRuleSeverityThresholdSevere = "SEVERE"

	SecurityRuleOutputSortDesc = "DESC"
	SecurityRuleOutputSortAsc  = "ASC"

	SecurityRuleCorrelationOperatorAnd = "AND"

	securityRuleResourceName = "security rule"

	securityRuleOperationCreate           = "CreateUpdateSecurityRule"
	securityRuleOperationRetrieve         = "RetrieveSecurityRule"
	securityRuleOperationUpdate           = "UpdateSecurityRule"
	securityRuleOperationDelete           = "DeleteSecurityRule"
	securityRuleOperationRetrieveFiltered = "RetrieveFilteredSecurityRules"
	securityRuleOperationEnable           = "EnableSecurityRule"
	securityRuleOperationDisable          = "DisableSecurityRule"

	RetrieveFilteredSortByFieldSeverity  = "SEVERITY"
	RetrieveFilteredSortByFieldName      = "NAME"
	RetrieveFilteredSortByFieldCreatedAt = "CREATED_AT"
	RetrieveFilteredSortByFieldUpdatedAt = "UPDATED_AT"
)

type SecurityRulesClient struct {
	*client.Client
}

type CreateUpdateSecurityRule struct {
	Title                  string             `json:"title,omitempty"`
	Description            string             `json:"description,omitempty"`
	Tags                   []string           `json:"tags,omitempty"`
	Output                 RuleOutput         `json:"output,omitempty"`
	SearchTimeFrameMinutes int32              `json:"searchTimeFrameMinutes,omitempty"`
	SubComponents          []SubRule          `json:"subComponents,omitempty"` // required
	Correlations           SubRuleCorrelation `json:"correlations,omitempty"`
	Enabled                *bool              `json:"enabled,omitempty"` // boolean - defined as a pointer because omitempty automatically omits false value
}

type SecurityRule struct {
	Id                     int32              `json:"id"`
	UpdatedAt              string             `json:"updatedAt"`
	UpdatedBy              string             `json:"updatedBy"`
	CreatedAt              string             `json:"createdAt"`
	CreatedBy              string             `json:"createdBy"`
	Enabled                bool               `json:"enabled"`
	Title                  string             `json:"title"`
	Description            string             `json:"description"`
	Tags                   []string           `json:"tags"`
	Output                 RuleOutput         `json:"output"`
	SearchTimeFrameMinutes int32              `json:"searchTimeFrameMinutes"`
	SubComponents          []SubRule          `json:"subComponents"`
	Correlations           SubRuleCorrelation `json:"correlations"`
	Schedule               ScheduleObj        `json:"schedule,omitempty"`
	Protected              *bool              `json:"protected,omitempty"`
}

type ScheduleObj struct {
	Cron     string `json:"cron,omitempty"`
	Timezone string `json:"timezone,omitempty"`
}

type RuleOutput struct {
	Recipients                   RuleRecipients `json:"recipients,omitempty"`
	SuppressNotificationsMinutes int32          `json:"suppressNotificationsMinutes,omitempty"`
	Type                         string         `json:"type"`
}

type RuleRecipients struct {
	Emails                  []string `json:"emails,omitempty"`
	NotificationEndpointIds []int32  `json:"notificationEndpointIds,omitempty"`
}

type SubRule struct {
	QueryDefinition RuleQuery     `json:"queryDefinition,omitempty"`
	Trigger         RuleTrigger   `json:"trigger,omitempty"`
	Output          SubRuleOutput `json:"output,omitempty"`
}

type RuleQuery struct {
	Query                    string         `json:"query,omitempty"`
	Filters                  BoolFilter     `json:"filters,omitempty"`
	GroupBy                  []string       `json:"groupBy,omitempty"`
	Aggregation              AggregationObj `json:"aggregation,omitempty"`
	ShouldQueryOnAllAccounts *bool          `json:"shouldQueryOnAllAccounts,omitempty"` // boolean - defined as a pointer because omitempty automatically omits false value
	AccountIdsToQueryOn      []int32        `json:"accountIdsToQueryOn,omitempty"`
}

type BoolFilter struct {
	Bool FilterLists `json:"bool,omitempty,omitempty"`
}

type FilterLists struct {
	Must    []map[string]interface{} `json:"must,omitempty"`
	MustNot []map[string]interface{} `json:"must_not,omitempty"`
}

type AggregationObj struct {
	AggregationType    string `json:"aggregationType,omitempty"`
	FieldToAggregateOn string `json:"fieldToAggregateOn,omitempty"`
}

type RuleTrigger struct {
	Operator               string             `json:"operator,omitempty"`
	SeverityThresholdTiers map[string]float32 `json:"severityThresholdTiers,omitempty"`
}

type SubRuleOutput struct {
	Columns []ColumnConfig `json:"columns,omitempty"`
}

type ColumnConfig struct {
	FieldName string `json:"fieldName,omitempty"`
	Regex     string `json:"regex,omitempty"`
	Sort      string `json:"sort,omitempty"`
}

type SubRuleCorrelation struct {
	CorrelationOperators []string            `json:"correlationOperators,omitempty"`
	Joins                []map[string]string `json:"joins,omitempty"`
}

type RetrieveFiltered struct {
	Filter     SecurityRuleFilter     `json:"filter,omitempty"`
	Sort       SecurityRuleSort       `json:"sort,omitempty"`
	Pagination SecurityRulePagination `json:"pagination,omitempty"`
}

type SecurityRuleFilter struct {
	Search             string   `json:"search,omitempty"`
	Severities         []string `json:"severities,omitempty"`
	UpdatedBy          []string `json:"updatedBy,omitempty"`
	CreatedBy          []string `json:"createdBy,omitempty"`
	EnabledState       []bool   `json:"enabledState,omitempty"`
	EmailNotifications []string `json:"emailNotifications,omitempty"`
	Tags               []string `json:"tags,omitempty"`
}

type SecurityRuleSort struct {
	SortByField string `json:"sortByField,omitempty"`
	Descending  *bool  `json:"descending,omitempty"` // boolean - defined as a pointer because omitempty automatically omits false value
}

type SecurityRulePagination struct {
	PageNumber int32 `json:"pageNumber,omitempty"`
	PageSize   int32 `json:"pageSize,omitempty"`
}

type FilteredSecurityRules struct {
	Total      int32                  `json:"total"`
	Results    []SecurityRule         `json:"results"`
	Pagination SecurityRulePagination `json:"pagination"`
	Schedule   ScheduleObj            `json:"schedule,omitempty"`
}

func New(apiToken, baseUrl string) (*SecurityRulesClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}
	c := &SecurityRulesClient{
		Client: client.New(apiToken, baseUrl),
	}
	return c, nil
}

func GetValidOutputTypes() []string {
	return []string{
		SecurityRulesOutputTypeJson,
		SecurityRulesOutputTypeTable,
	}
}

func GetValidAggregationTypes() []string {
	return []string{
		SecurityRuleAggregationTypeSum,
		SecurityRuleAggregationTypeMin,
		SecurityRuleAggregationTypeMax,
		SecurityRuleAggregationTypeAvg,
		SecurityRuleAggregationTypeCount,
		SecurityRuleAggregationTypeUniqueCount,
		SecurityRuleAggregationTypeNone,
	}
}

func GetValidOperators() []string {
	return []string{
		SecurityRuleOperatorLessThan,
		SecurityRuleOperatorGreaterThan,
		SecurityRuleOperatorLessThanOrEquals,
		SecurityRuleOperatorGreaterThanOrEquals,
		SecurityRuleOperatorEquals,
		SecurityRuleOperatorNotEquals,
	}
}

func GetValidThresholds() []string {
	return []string{
		SecurityRuleSeverityThresholdInfo,
		SecurityRuleSeverityThresholdLow,
		SecurityRuleSeverityThresholdMedium,
		SecurityRuleSeverityThresholdHigh,
		SecurityRuleSeverityThresholdSevere,
	}
}

func GetValidOutputSorts() []string {
	return []string{
		SecurityRuleOutputSortDesc,
		SecurityRuleOutputSortAsc,
	}
}

func GetValidCorrelationOperators() []string {
	return []string{
		SecurityRuleCorrelationOperatorAnd,
	}
}

func GetValidFilteredSorts() []string {
	return []string{
		RetrieveFilteredSortByFieldSeverity,
		RetrieveFilteredSortByFieldName,
		RetrieveFilteredSortByFieldCreatedAt,
		RetrieveFilteredSortByFieldUpdatedAt,
	}
}

func validateCreateUpdateSecurityRule(ruleRequest CreateUpdateSecurityRule) error {
	validOutputs := GetValidOutputTypes()
	validAggregations := GetValidAggregationTypes()
	validOperators := GetValidOperators()
	validThresholds := GetValidThresholds()
	validOutputSorts := GetValidOutputSorts()
	validCorrelationOperators := GetValidCorrelationOperators()

	if len(ruleRequest.Title) == 0 {
		return fmt.Errorf("title must be set")
	}

	if ruleRequest.SubComponents == nil || len(ruleRequest.SubComponents) == 0 {
		return fmt.Errorf("sub components must be set")
	}

	if !logzio_client.Contains(validOutputs, ruleRequest.Output.Type) {
		return fmt.Errorf("invalid output type. should be one of: %s", validOutputs)
	}

	for _, subComponent := range ruleRequest.SubComponents {
		if len(subComponent.QueryDefinition.Aggregation.AggregationType) > 0 {
			if !logzio_client.Contains(validAggregations, subComponent.QueryDefinition.Aggregation.AggregationType) {
				return fmt.Errorf("invalid aggregation type. should be one of: %s", validAggregations)
			}
		}

		if len(subComponent.Trigger.Operator) > 0 {
			if !logzio_client.Contains(validOperators, subComponent.Trigger.Operator) {
				return fmt.Errorf("invalid operator. should be one of: %s", validOperators)
			}
		}

		for threshold := range subComponent.Trigger.SeverityThresholdTiers {
			if !logzio_client.Contains(validThresholds, threshold) {
				return fmt.Errorf("invalid severity threshold tiers. should be one of: %s", validThresholds)
			}
		}

		for _, column := range subComponent.Output.Columns {
			if len(column.Sort) > 0 &&
				!logzio_client.Contains(validOutputSorts, column.Sort) {
				return fmt.Errorf("invalid sort. should be one of: %s", validOutputSorts)
			}
		}
	}

	for _, correlationOperator := range ruleRequest.Correlations.CorrelationOperators {
		if !logzio_client.Contains(validCorrelationOperators, correlationOperator) {
			return fmt.Errorf("invalid correlation operator. should be one of: %s", validCorrelationOperators)
		}
	}

	return nil
}
