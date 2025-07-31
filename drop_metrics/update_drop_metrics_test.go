package drop_metrics_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/drop_metrics"
	"github.com/stretchr/testify/assert"
)

func TestDropMetrics_UpdateDropMetric(t *testing.T) {
	underTest, mux, teardown := setupDropMetricsTest()
	defer teardown()

	mux.HandleFunc("/v1/metrics-management/drop-filters/1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_drop_metric.json"))
	})

	enabled := false
	updateReq := drop_metrics.UpdateDropMetric{
		AccountId: 1234,
		Enabled:   &enabled,
		Filter: drop_metrics.FilterObject{
			Operator: drop_metrics.OperatorAnd,
			Expression: []drop_metrics.FilterExpression{
				{
					Name:             "__name__",
					Value:            "UpdatedMetricName",
					ComparisonFilter: drop_metrics.ComparisonEq,
				},
				{
					Name:             "environment",
					Value:            "production",
					ComparisonFilter: drop_metrics.ComparisonEq,
				},
			},
		},
	}

	result, err := underTest.UpdateDropMetric(1, updateReq)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.Id)
	assert.Equal(t, int64(1234), result.AccountId)
	assert.False(t, result.Enabled)
	assert.Equal(t, "AND", result.Filter.Operator)
	assert.Len(t, result.Filter.Expression, 2)
	assert.Equal(t, "EQ", result.Filter.Expression[0].ComparisonFilter)
	assert.Equal(t, "EQ", result.Filter.Expression[1].ComparisonFilter)
}

func TestDropMetrics_UpdateDropMetricAPIFailed(t *testing.T) {
	underTest, mux, teardown := setupDropMetricsTest()
	defer teardown()

	mux.HandleFunc("/v1/metrics-management/drop-filters/1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("api_error.txt"))
	})

	enabled := true
	updateReq := drop_metrics.UpdateDropMetric{
		AccountId: 1234,
		Enabled:   &enabled,
		Filter: drop_metrics.FilterObject{
			Operator: drop_metrics.OperatorAnd,
			Expression: []drop_metrics.FilterExpression{
				{
					Name:             "__name__",
					Value:            "TestMetric",
					ComparisonFilter: drop_metrics.ComparisonEq,
				},
			},
		},
	}

	result, err := underTest.UpdateDropMetric(1, updateReq)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestDropMetrics_UpdateDropMetricInvalidId(t *testing.T) {
	underTest, _, teardown := setupDropMetricsTest()
	defer teardown()

	enabled := true
	updateReq := drop_metrics.UpdateDropMetric{
		AccountId: 1234,
		Enabled:   &enabled,
		Filter: drop_metrics.FilterObject{
			Operator: drop_metrics.OperatorAnd,
			Expression: []drop_metrics.FilterExpression{
				{
					Name:             "__name__",
					Value:            "TestMetric",
					ComparisonFilter: drop_metrics.ComparisonEq,
				},
			},
		},
	}

	result, err := underTest.UpdateDropMetric(0, updateReq)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "dropFilterId must be greater than 0")
}

func TestDropMetrics_UpdateDropMetricValidationError(t *testing.T) {
	underTest, _, teardown := setupDropMetricsTest()
	defer teardown()

	// Test with invalid accountId
	updateReq := drop_metrics.UpdateDropMetric{
		AccountId: 0,
		Filter: drop_metrics.FilterObject{
			Operator:   drop_metrics.OperatorAnd,
			Expression: []drop_metrics.FilterExpression{},
		},
	}

	result, err := underTest.UpdateDropMetric(1, updateReq)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "accountId must be set and greater than 0")
}
