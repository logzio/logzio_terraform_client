package metrics_rollup_rules_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/logzio/logzio_terraform_client/metrics_rollup_rules"
)

func TestBulkCreateRollupRulesSuccess(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath+"/bulk/create", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("bulk_create_metrics_rollup_rules.json"))
		})

		requests := []metrics_rollup_rules.CreateUpdateRollupRule{
			{
				AccountId:               1,
				MetricName:              "cpu",
				MetricType:              metrics_rollup_rules.MetricTypeGauge,
				RollupFunction:          metrics_rollup_rules.AggLast,
				LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
				Labels:                  []string{"x"},
			},
			{
				AccountId:               1,
				MetricName:              "memory",
				MetricType:              metrics_rollup_rules.MetricTypeGauge,
				RollupFunction:          metrics_rollup_rules.AggCount,
				LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
				Labels:                  []string{"y"},
			},
		}

		res, err := underTest.BulkCreateRollupRules(requests)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res, 2)
		assert.Equal(t, "a", res[0].Id)
		assert.Equal(t, "bulk-rollup-rule-1", res[0].Name)
		assert.Equal(t, "b", res[1].Id)
		assert.Equal(t, "bulk-rollup-rule-2", res[1].Name)
	}
}

func TestBulkCreateRollupRulesApiFailed(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath+"/bulk/create", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		requests := []metrics_rollup_rules.CreateUpdateRollupRule{
			{
				AccountId:               1,
				MetricName:              "cpu",
				MetricType:              metrics_rollup_rules.MetricTypeGauge,
				RollupFunction:          metrics_rollup_rules.AggLast,
				LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
				Labels:                  []string{"x"},
			},
		}

		res, err := underTest.BulkCreateRollupRules(requests)
		assert.Error(t, err)
		assert.Nil(t, res)
	}
}

func TestBulkCreateRollupRulesNameTooLong(t *testing.T) {
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	// Create a name that exceeds 256 characters
	longName := strings.Repeat("a", 257)

	requests := []metrics_rollup_rules.CreateUpdateRollupRule{
		{
			AccountId:               1,
			Name:                    longName,
			MetricName:              "cpu",
			MetricType:              metrics_rollup_rules.MetricTypeGauge,
			RollupFunction:          metrics_rollup_rules.AggLast,
			LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
			Labels:                  []string{"x"},
		},
	}

	res, err := underTest.BulkCreateRollupRules(requests)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "name must not exceed 256 characters")
}

func TestBulkCreateRollupRulesNotFound(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath+"/bulk/create", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusNotFound)
		})

		requests := []metrics_rollup_rules.CreateUpdateRollupRule{{
			AccountId:               1,
			MetricName:              "cpu",
			MetricType:              metrics_rollup_rules.MetricTypeGauge,
			RollupFunction:          metrics_rollup_rules.AggLast,
			LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
			Labels:                  []string{"x"},
		}}

		res, err := underTest.BulkCreateRollupRules(requests)
		assert.Error(t, err)
		assert.Nil(t, res)
	}
}
