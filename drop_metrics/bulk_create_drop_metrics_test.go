package drop_metrics_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/drop_metrics"
	"github.com/stretchr/testify/assert"
)

func TestDropMetrics_BulkCreateDropMetrics(t *testing.T) {
	underTest, mux, teardown := setupDropMetricsTest()
	defer teardown()

	mux.HandleFunc("/v1/metrics-management/drop-filters/bulk/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("bulk_create_drop_metrics.json"))
	})

	active := true
	requests := []drop_metrics.CreateDropMetric{
		{
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
		},
		{
			AccountId: 1235,
			Active:    &active,
			Filter: drop_metrics.FilterObject{
				Operator: drop_metrics.OperatorAnd,
				Expression: []drop_metrics.FilterExpression{
					{
						Name:             "__name__",
						Value:            "MemoryUsage",
						ComparisonFilter: drop_metrics.ComparisonEq,
					},
				},
			},
		},
	}

	results, err := underTest.BulkCreateDropMetrics(requests)
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Len(t, results, 2)
	assert.Equal(t, int64(1), results[0].Id)
	assert.Equal(t, int64(2), results[1].Id)
}

func TestDropMetrics_BulkCreateDropMetricsEmptyArray(t *testing.T) {
	underTest, _, teardown := setupDropMetricsTest()
	defer teardown()

	results, err := underTest.BulkCreateDropMetrics([]drop_metrics.CreateDropMetric{})
	assert.Error(t, err)
	assert.Nil(t, results)
	assert.Contains(t, err.Error(), "requests array cannot be empty")
}

func TestDropMetrics_BulkCreateDropMetricsAPIFailed(t *testing.T) {
	underTest, mux, teardown := setupDropMetricsTest()
	defer teardown()

	mux.HandleFunc("/v1/metrics-management/drop-filters/bulk/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("api_error.txt"))
	})

	active := true
	requests := []drop_metrics.CreateDropMetric{
		{
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
		},
	}

	results, err := underTest.BulkCreateDropMetrics(requests)
	assert.Error(t, err)
	assert.Nil(t, results)
}
