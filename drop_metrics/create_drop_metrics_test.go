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

	enabled := true
	createReq := drop_metrics.CreateDropMetric{
		AccountId: 1234,
		Enabled:   &enabled,
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
	assert.True(t, result.Enabled)
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

	enabled := true
	createReq := drop_metrics.CreateDropMetric{
		AccountId: 1234,
		Enabled:   &enabled,
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
	createReq := drop_metrics.CreateDropMetric{
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
