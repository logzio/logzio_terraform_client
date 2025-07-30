package drop_metrics

import (
	"fmt"

	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
)

const dropMetricsServiceEndpoint = "%s/v1/metrics-management/drop-filters"

const (
	// Comparison operations
	ComparisonEq           = "eq"
	ComparisonNotEq        = "not_eq"
	ComparisonRegexMatch   = "regex_match"
	ComparisonRegexNoMatch = "regex_no_match"

	// Logical operations
	OperatorAnd = "and"

	// Filter types
	FilterTypeSingle  = "single"
	FilterTypeComplex = "complex"

	// Operation names for logging
	createDropMetricOperation     = "CreateDropMetric"
	bulkCreateDropMetricOperation = "BulkCreateDropMetric"
	getDropMetricOperation        = "GetDropMetric"
	searchDropMetricsOperation    = "SearchDropMetrics"
	enableDropMetricOperation     = "EnableDropMetric"
	disableDropMetricOperation    = "DisableDropMetric"
	deleteDropMetricOperation     = "DeleteDropMetric"
	updateDropMetricOperation     = "UpdateDropMetric"
	deleteBySearchOperation       = "DeleteDropMetricsBySearch"
	bulkDeleteDropMetricOperation = "BulkDeleteDropMetric"

	resourceName = "drop_metric"
)

type DropMetricsClient struct {
	*client.Client
}

// CreateDropMetric represents the request payload for creating a drop metric filter
type CreateDropMetric struct {
	AccountId int64        `json:"accountId"`
	Enabled   *bool        `json:"enabled,omitempty"`
	Filter    FilterObject `json:"filter"`
}

// FilterObject represents a filter that can be either simple or complex
type FilterObject struct {
	Operator   string             `json:"operator,omitempty"`
	Expression []FilterExpression `json:"expression,omitempty"`
}

// FilterExpression represents a single filter expression that can be nested
type FilterExpression struct {
	// For single filters
	Name             string `json:"name,omitempty"`
	Value            string `json:"value,omitempty"`
	ComparisonFilter string `json:"comparisonFilter,omitempty"`
	FilterType       string `json:"filterType,omitempty"`

	// For nested complex filters
	Operator   string             `json:"operator,omitempty"`
	Expression []FilterExpression `json:"expression,omitempty"`
}

// DropMetric represents the response object for a drop metric filter
type DropMetric struct {
	Id         int64        `json:"id"`
	AccountId  int64        `json:"accountId"`
	Enabled    bool         `json:"enabled"`
	Filter     FilterObject `json:"filter"`
	CreatedAt  string       `json:"createdAt,omitempty"`
	CreatedBy  string       `json:"createdBy,omitempty"`
	ModifiedAt string       `json:"modifiedAt,omitempty"`
	ModifiedBy string       `json:"modifiedBy,omitempty"`
}

// SearchDropMetricsRequest represents the request payload for searching drop metrics
type SearchDropMetricsRequest struct {
	Filter     *SearchFilter `json:"filter,omitempty"`
	Pagination *Pagination   `json:"pagination,omitempty"`
}

// SearchFilter represents search criteria for drop metrics
type SearchFilter struct {
	AccountIds  []int64  `json:"accountIds,omitempty"`
	MetricNames []string `json:"metricNames,omitempty"`
	Enabled     *bool    `json:"enabled,omitempty"`
}

// Pagination represents pagination parameters
type Pagination struct {
	PageNumber int `json:"pageNumber,omitempty"`
	PageSize   int `json:"pageSize,omitempty"`
}

// New creates a new entry point into the drop metrics functions
func New(apiToken, baseUrl string) (*DropMetricsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	return &DropMetricsClient{
		Client: client.New(apiToken, baseUrl),
	}, nil
}

// validateCreateDropMetricRequest validates the create drop metric request
func validateCreateDropMetricRequest(req CreateDropMetric) error {
	if req.AccountId <= 0 {
		return fmt.Errorf("accountId must be set and greater than 0")
	}

	if err := validateFilterObject(req.Filter); err != nil {
		return fmt.Errorf("filter validation failed: %w", err)
	}

	return nil
}

// validateFilterObject validates a filter object
func validateFilterObject(filter FilterObject) error {
	if len(filter.Operator) == 0 {
		return fmt.Errorf("operator must be set")
	}

	validOperators := []string{OperatorAnd}
	if !logzio_client.Contains(validOperators, filter.Operator) {
		return fmt.Errorf("operator must be one of %v", validOperators)
	}

	if len(filter.Expression) == 0 {
		return fmt.Errorf("expression must not be empty")
	}

	for i, expr := range filter.Expression {
		if err := validateFilterExpression(expr); err != nil {
			return fmt.Errorf("expression[%d] validation failed: %w", i, err)
		}
	}

	return nil
}

// validateFilterExpression validates a filter expression
func validateFilterExpression(expr FilterExpression) error {
	// Check if this is a nested complex filter
	if len(expr.Operator) > 0 {
		// This is a complex filter with nested expressions
		validOperators := []string{OperatorAnd}
		if !logzio_client.Contains(validOperators, expr.Operator) {
			return fmt.Errorf("nested operator must be one of %v", validOperators)
		}

		if len(expr.Expression) == 0 {
			return fmt.Errorf("nested expression must not be empty")
		}

		for i, nestedExpr := range expr.Expression {
			if err := validateFilterExpression(nestedExpr); err != nil {
				return fmt.Errorf("nested expression[%d] validation failed: %w", i, err)
			}
		}
	} else {
		// This is a single filter
		if len(expr.Name) == 0 {
			return fmt.Errorf("name must be set for single filter")
		}

		if len(expr.Value) == 0 {
			return fmt.Errorf("value must be set for single filter")
		}

		if len(expr.ComparisonFilter) == 0 {
			return fmt.Errorf("comparisonFilter must be set for single filter")
		}

		validComparisons := []string{ComparisonEq, ComparisonNotEq, ComparisonRegexMatch, ComparisonRegexNoMatch}
		if !logzio_client.Contains(validComparisons, expr.ComparisonFilter) {
			return fmt.Errorf("comparisonFilter must be one of %v", validComparisons)
		}
	}

	return nil
}

// validateSearchDropMetricsRequest validates the search request
func validateSearchDropMetricsRequest(req SearchDropMetricsRequest) error {
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
