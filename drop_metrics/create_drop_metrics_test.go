package drop_metrics_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/drop_metrics"
	"github.com/stretchr/testify/assert"
)

func TestDropMetrics_CreateDropMetric(t *testing.T) {
	underTest, mux, teardown := setupDropMetricsTest()
	defer teardown()

	mux.HandleFunc("/v1/metrics-management/drop-filters", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_drop_metric.json"))
	})

	active := true
	createReq := drop_metrics.CreateUpdateDropMetric{
		AccountId: 1234,
		Active:    &active,
		Filter: drop_metrics.FilterObject{
			Operator: drop_metrics.OperatorAnd,
			Expression: []drop_metrics.FilterExpression{
				{
					Name:             "__name__",
					Value:            "CpuUsage",
					ComparisonFilter: drop_metrics.ComparisonEq,
				},
				{
					Name:             "service",
					Value:            "metrics-service",
					ComparisonFilter: drop_metrics.ComparisonEq,
				},
			},
		},
	}

	result, err := underTest.CreateDropMetric(createReq)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.Id)
	assert.Equal(t, int64(1234), result.AccountId)
	assert.Equal(t, "my-drop-filter", result.Name)
	assert.True(t, result.Active)
	assert.Equal(t, "DROP_BEFORE_PROCESSING", result.DropPolicy)
	assert.Equal(t, "AND", result.Filter.Operator)
	assert.Len(t, result.Filter.Expression, 2)
}

func TestDropMetrics_CreateDropMetricAPIFailed(t *testing.T) {
	underTest, mux, teardown := setupDropMetricsTest()
	defer teardown()

	mux.HandleFunc("/v1/metrics-management/drop-filters", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("api_error.txt"))
	})

	active := true
	createReq := drop_metrics.CreateUpdateDropMetric{
		AccountId: 1234,
		Active:    &active,
		Filter: drop_metrics.FilterObject{
			Operator: drop_metrics.OperatorAnd,
			Expression: []drop_metrics.FilterExpression{
				{
					Name:             "__name__",
					Value:            "CpuUsage",
					ComparisonFilter: drop_metrics.ComparisonEq,
				},
			},
		},
	}

	result, err := underTest.CreateDropMetric(createReq)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestDropMetrics_CreateDropMetricValidationError(t *testing.T) {
	underTest, _, teardown := setupDropMetricsTest()
	defer teardown()

	// Test with invalid accountId
	createReq := drop_metrics.CreateUpdateDropMetric{
		AccountId: 0,
		Filter: drop_metrics.FilterObject{
			Operator:   drop_metrics.OperatorAnd,
			Expression: []drop_metrics.FilterExpression{},
		},
	}

	result, err := underTest.CreateDropMetric(createReq)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "accountId must be set and greater than 0")
}

func TestDropMetrics_SearchWithSearchTerm(t *testing.T) {
	underTest, mux, teardown := setupDropMetricsTest()
	defer teardown()

	mux.HandleFunc("/v1/metrics-management/drop-filters/search", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("search_drop_metrics.json"))
	})

	searchReq := drop_metrics.SearchDropMetricsRequest{
		Filter: &drop_metrics.SearchFilter{
			AccountIds: []int64{1234},
			SearchTerm: "cpu-filter",
		},
		Pagination: &drop_metrics.Pagination{
			PageNumber: 0,
			PageSize:   10,
		},
	}

	results, err := underTest.SearchDropMetrics(searchReq)
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Len(t, results, 1)
	assert.Equal(t, "searchable-drop-filter", results[0].Name)
	assert.Equal(t, "DROP_BEFORE_PROCESSING", results[0].DropPolicy)
}

func TestDropMetrics_CreateDropMetricWithName(t *testing.T) {
	underTest, mux, teardown := setupDropMetricsTest()
	defer teardown()

	mux.HandleFunc("/v1/metrics-management/drop-filters", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_drop_metric.json"))
	})

	active := true
	createReq := drop_metrics.CreateUpdateDropMetric{
		AccountId: 1234,
		Name:      "test-drop-filter",
		Active:    &active,
		Filter: drop_metrics.FilterObject{
			Operator: drop_metrics.OperatorAnd,
			Expression: []drop_metrics.FilterExpression{
				{
					Name:             "__name__",
					Value:            "CpuUsage",
					ComparisonFilter: drop_metrics.ComparisonEq,
				},
			},
		},
	}

	result, err := underTest.CreateDropMetric(createReq)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "my-drop-filter", result.Name)
}

func TestDropMetrics_CreateDropMetricWithDropPolicy(t *testing.T) {
	underTest, mux, teardown := setupDropMetricsTest()
	defer teardown()

	mux.HandleFunc("/v1/metrics-management/drop-filters", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_drop_metric.json")) // Uses DROP_BEFORE_STORING
	})

	active := true
	createReq := drop_metrics.CreateUpdateDropMetric{
		AccountId:  1234,
		Name:       "test-drop-filter",
		Active:     &active,
		DropPolicy: drop_metrics.DropPolicyBeforeStoring,
		Filter: drop_metrics.FilterObject{
			Operator: drop_metrics.OperatorAnd,
			Expression: []drop_metrics.FilterExpression{
				{
					Name:             "__name__",
					Value:            "CpuUsage",
					ComparisonFilter: drop_metrics.ComparisonEq,
				},
			},
		},
	}

	result, err := underTest.CreateDropMetric(createReq)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "DROP_BEFORE_STORING", result.DropPolicy)
}

func TestDropMetrics_CreateDropMetricInvalidDropPolicy(t *testing.T) {
	underTest, _, teardown := setupDropMetricsTest()
	defer teardown()

	active := true
	createReq := drop_metrics.CreateUpdateDropMetric{
		AccountId:  1234,
		Name:       "test-drop-filter",
		Active:     &active,
		DropPolicy: "INVALID_POLICY",
		Filter: drop_metrics.FilterObject{
			Operator: drop_metrics.OperatorAnd,
			Expression: []drop_metrics.FilterExpression{
				{
					Name:             "__name__",
					Value:            "CpuUsage",
					ComparisonFilter: drop_metrics.ComparisonEq,
				},
			},
		},
	}

	result, err := underTest.CreateDropMetric(createReq)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "dropPolicy must be one of")
}
