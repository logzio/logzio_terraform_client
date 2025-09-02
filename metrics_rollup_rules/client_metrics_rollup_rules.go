package metrics_rollup_rules

import (
	"fmt"

	"github.com/logzio/logzio_terraform_client/client"
)

const (
	metricsRollupRulesServiceEndpoint = "%s/v1/metrics-management/rollup-rules"

	operationCreateMetricsRollupRule      = "CreateMetricsRollupRule"
	operationBulkCreateMetricsRollupRules = "BulkCreateMetricsRollupRules"
	operationGetMetricsRollupRule         = "GetMetricsRollupRule"
	operationSearchMetricsRollupRules     = "SearchMetricsRollupRules"
	operationUpdateMetricsRollupRule      = "UpdateMetricsRollupRule"
	operationDeleteMetricsRollupRule      = "DeleteMetricsRollupRule"
	operationBulkDeleteMetricsRollupRules = "BulkDeleteMetricsRollupRules"

	resourceName = "metrics_rollup_rule"
)

type MetricsRollupRulesClient struct {
	*client.Client
}

type MetricType string

const (
	MetricTypeGauge             MetricType = "GAUGE"
	MetricTypeCounter           MetricType = "COUNTER"
	MetricTypeDeltaCounter      MetricType = "DELTA_COUNTER"
	MetricTypeCumulativeCounter MetricType = "CUMULATIVE_COUNTER"
	MetricTypeMeasurement       MetricType = "MEASUREMENT"
)

type AggregationFunction string

const (
	AggLast   AggregationFunction = "LAST"
	AggMin    AggregationFunction = "MIN"
	AggMax    AggregationFunction = "MAX"
	AggMean   AggregationFunction = "MEAN"
	AggMedian AggregationFunction = "MEDIAN"
	AggCount  AggregationFunction = "COUNT"
	AggSum    AggregationFunction = "SUM"
	AggSumSq  AggregationFunction = "SUMSQ"
	AggStdev  AggregationFunction = "STDEV"
	AggP10    AggregationFunction = "P10"
	AggP20    AggregationFunction = "P20"
	AggP25    AggregationFunction = "P25"
	AggP30    AggregationFunction = "P30"
	AggP40    AggregationFunction = "P40"
	AggP50    AggregationFunction = "P50"
	AggP60    AggregationFunction = "P60"
	AggP70    AggregationFunction = "P70"
	AggP75    AggregationFunction = "P75"
	AggP80    AggregationFunction = "P80"
	AggP90    AggregationFunction = "P90"
	AggP95    AggregationFunction = "P95"
	AggP99    AggregationFunction = "P99"
	AggP999   AggregationFunction = "P999"
	AggP9999  AggregationFunction = "P9999"
)

type LabelsRemovalMethod string

const (
	LabelsExcludeBy LabelsRemovalMethod = "EXCLUDE_BY"
	LabelsGroupBy   LabelsRemovalMethod = "GROUP_BY"
)

type Comparison string

const (
	CmpEq           Comparison = "EQ"
	CmpNotEq        Comparison = "NOT_EQ"
	CmpRegexMatch   Comparison = "REGEX_MATCH"
	CmpRegexNoMatch Comparison = "REGEX_NO_MATCH"
)

type SingleFilter struct {
	Comparison Comparison `json:"comparison"`
	Name       string     `json:"name"`
	Value      string     `json:"value"`
}

type ComplexFilter struct {
	Expression []SingleFilter `json:"expression"`
}

type RollupRule struct {
	Id                      string              `json:"id"`
	Name                    string              `json:"name,omitempty"`
	AccountId               int64               `json:"accountId"`
	MetricName              string              `json:"metricName,omitempty"`
	MetricType              MetricType          `json:"metricType"`
	RollupFunction          AggregationFunction `json:"rollupFunction"`
	LabelsEliminationMethod LabelsRemovalMethod `json:"labelsEliminationMethod"`
	Labels                  []string            `json:"labels"`
	Namespaces              []string            `json:"namespaces,omitempty"`
	ClusterId               string              `json:"clusterId,omitempty"`
	IsDeleted               bool                `json:"isDeleted,omitempty"`
	DropOriginalMetric      bool                `json:"dropOriginalMetric,omitempty"`
	DropPolicyRuleId        *string             `json:"dropPolicyRuleId,omitempty"`
	Filter                  *ComplexFilter      `json:"filter,omitempty"`
	NewMetricNameTemplate   *string             `json:"newMetricNameTemplate,omitempty"`
	Version                 int                 `json:"version,omitempty"`
}

// CreateUpdateRollupRule represents the request payload for creating a rollup rule
type CreateUpdateRollupRule struct {
	AccountId               int64               `json:"accountId"`
	Name                    string              `json:"name,omitempty"`
	MetricName              string              `json:"metricName,omitempty"`
	MetricType              MetricType          `json:"metricType"`
	RollupFunction          AggregationFunction `json:"rollupFunction,omitempty"`
	LabelsEliminationMethod LabelsRemovalMethod `json:"labelsEliminationMethod"`
	Labels                  []string            `json:"labels"`
	Filter                  *ComplexFilter      `json:"filter,omitempty"`
	NewMetricNameTemplate   *string             `json:"newMetricNameTemplate,omitempty"`
	DropOriginalMetric      *bool               `json:"dropOriginalMetric,omitempty"`
}

type SortField struct {
	Field string `json:"field"`
}

type Pagination struct {
	PageNumber int `json:"pageNumber,omitempty"`
	PageSize   int `json:"pageSize,omitempty"`
}

type SearchFilter struct {
	AccountIds  []int64  `json:"accountIds,omitempty"`
	MetricNames []string `json:"metricNames,omitempty"`
	SearchTerm  string   `json:"searchTerm,omitempty"`
}

type SearchRollupRulesRequest struct {
	Filter     *SearchFilter `json:"filter,omitempty"`
	Pagination *Pagination   `json:"pagination,omitempty"`
	Sort       []SortField   `json:"sort,omitempty"`
}

type SearchRollupRulesResponse struct {
	Results    []RollupRule `json:"results"`
	Pagination *Pagination  `json:"pagination,omitempty"`
}

// New creates a new entry point into the metrics rollup rules functions
func New(apiToken, baseUrl string) (*MetricsRollupRulesClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	return &MetricsRollupRulesClient{
		Client: client.New(apiToken, baseUrl),
	}, nil
}

// validateCreateUpdateRollupRuleRequest validates the create rollup rule request
func validateCreateUpdateRollupRuleRequest(req CreateUpdateRollupRule) error {
	if req.MetricType == "" {
		return fmt.Errorf("metricType must be set")
	}

	// rollupFunction requirement based on metric type
	switch req.MetricType {
	case MetricTypeMeasurement:
		if req.RollupFunction == "" {
			return fmt.Errorf("rollupFunction must be set for MEASUREMENT metrics")
		}
	case MetricTypeGauge:
		if req.RollupFunction == "" {
			return fmt.Errorf("rollupFunction must be set for GAUGE metrics")
		}
	default:
		// For non-supported types, rollupFunction must not be provided
		if req.RollupFunction != "" {
			return fmt.Errorf("rollupFunction is supported only for GAUGE and MEASUREMENT metrics")
		}
	}

	if req.LabelsEliminationMethod == "" {
		return fmt.Errorf("labelsEliminationMethod must be set")
	}

	if len(req.Labels) == 0 {
		return fmt.Errorf("labels must not be empty")
	}

	// Either "metricName" or "filter" should be part of the payload
	if len(req.MetricName) == 0 && req.Filter == nil {
		return fmt.Errorf("either 'metricName' or 'filter' must be provided")
	}

	// Only validate rollupFunction enum if it's provided
	if req.RollupFunction != "" {
		err := validateEnums(req.MetricType, req.RollupFunction, req.LabelsEliminationMethod)
		if err != nil {
			return err
		}

		err = validateMeasurementTypeAggregation(req.MetricType, req.RollupFunction)
		if err != nil {
			return err
		}
	} else {
		// Still validate metric type and labels elimination method even if rollupFunction is empty
		err := isValidMetricType(req.MetricType)
		if err != nil {
			return err
		}
		err = isValidLabelsRemovalMethod(req.LabelsEliminationMethod)
		if err != nil {
			return err
		}
	}

	err = validateMeasurementTypeAggregation(req.MetricType, req.RollupFunction)
	if err != nil {
		return err
	}

	return nil
}

// validateSearchRollupRuleRequest validates the search request
func validateSearchRollupRuleRequest(req SearchRollupRulesRequest) error {
	if req.Pagination != nil {
		if req.Pagination.PageNumber < 0 {
			return fmt.Errorf("pageNumber must be non-negative")
		}
		if req.Pagination.PageSize < 0 {
			return fmt.Errorf("pageSize must be non-negative")
		}
	}

	return nil
}

// validateEnums validates the metric rollup's enums
func validateEnums(metricType MetricType, agg AggregationFunction, method LabelsRemovalMethod) error {
	err := isValidMetricType(metricType)
	if err != nil {
		return err
	}
	err = isValidAggregationFunction(agg)
	if err != nil {
		return err
	}
	err = isValidLabelsRemovalMethod(method)
	if err != nil {
		return err
	}
	return nil
}

// GetValidMetricType returns a list of valid metric types
func GetValidMetricType() []MetricType {
	return []MetricType{
		MetricTypeGauge,
		MetricTypeCounter,
		MetricTypeDeltaCounter,
		MetricTypeCumulativeCounter,
		MetricTypeMeasurement,
	}
}

// GetValidAggregationFunctions returns a list of valid aggregation functions
func GetValidAggregationFunctions() []AggregationFunction {
	return []AggregationFunction{
		AggLast,
		AggMin,
		AggMax,
		AggMean,
		AggMedian,
		AggCount,
		AggSum,
		AggSumSq,
		AggStdev,
		AggP10,
		AggP20,
		AggP25,
		AggP30,
		AggP40,
		AggP50,
		AggP60,
		AggP70,
		AggP75,
		AggP80,
		AggP90,
		AggP95,
		AggP99,
		AggP999,
		AggP9999,
	}
}

// GetValidLabelsRemovalMethods returns a list of valid labels removal methods
func GetValidLabelsRemovalMethods() []LabelsRemovalMethod {
	return []LabelsRemovalMethod{
		LabelsExcludeBy,
		LabelsGroupBy,
	}
}

// isValidMetricType checks if the provided metric type is valid
func isValidMetricType(metricType MetricType) error {
	validMetricTypes := GetValidMetricType()

	for _, validType := range validMetricTypes {
		if validType == metricType {
			return nil
		}
	}
	return fmt.Errorf("invalid metric type. metric type must be one of: %s", validMetricTypes)
}

// isValidAggregationFunction checks if the provided aggregation function is valid
func isValidAggregationFunction(agg AggregationFunction) error {
	validAggs := GetValidAggregationFunctions()

	for _, validAgg := range validAggs {
		if validAgg == agg {
			return nil
		}
	}
	return fmt.Errorf("invalid aggregation function. aggregation function must be one of: %s", validAggs)
}

// isValidLabelsRemovalMethod checks if the provided labels removal method is valid
func isValidLabelsRemovalMethod(method LabelsRemovalMethod) error {
	validMethods := GetValidLabelsRemovalMethods()

	for _, validMethod := range validMethods {
		if validMethod == method {
			return nil
		}
	}
	return fmt.Errorf("invalid labels elimination method. method must be one of: %s", validMethods)
}

// validateMeasurementTypeAggregation validates that MEASUREMENT metric type uses only allowed aggregation functions
func validateMeasurementTypeAggregation(metricType MetricType, rollupFunction AggregationFunction) error {
	if metricType != MetricTypeMeasurement {
		return nil // No additional validation needed for non-MEASUREMENT types
	}

	validAggregationsForMeasurement := []AggregationFunction{
		AggSum,
		AggMin,
		AggMax,
		AggCount,
		AggSumSq,
		AggMean,
		AggLast,
	}

	for _, validAgg := range validAggregationsForMeasurement {
		if rollupFunction == validAgg {
			return nil
		}
	}

	return fmt.Errorf("invalid aggregation function for MEASUREMENT metric type. For MEASUREMENT metrics, rollup_function must be one of: %s", validAggregationsForMeasurement)
}

// validateRollupRuleId checks if the provided rollup rule ID is valid
func validateRollupRuleId(rollupRuleId string) error {
	if len(rollupRuleId) == 0 {
		return fmt.Errorf("rollupRuleId must be set")
	}
	return nil
}
