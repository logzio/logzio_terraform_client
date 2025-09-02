package metrics_rollup_rules_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/metrics_rollup_rules"
	"github.com/stretchr/testify/assert"
)

func TestCreateRollupRuleSuccess(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("create_metrics_rollup_rule.json"))
		})

		request := metrics_rollup_rules.CreateUpdateRollupRule{
			AccountId:               1,
			MetricName:              "cpu",
			MetricType:              metrics_rollup_rules.MetricTypeGauge,
			RollupFunction:          metrics_rollup_rules.AggLast,
			LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
			Labels:                  []string{"x"},
		}

		res, err := underTest.CreateRollupRule(request)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "abc", res.Id)
		assert.Equal(t, "my-rollup-rule", res.Name)
		assert.Equal(t, int64(1), res.AccountId)
		assert.Equal(t, metrics_rollup_rules.MetricTypeGauge, res.MetricType)
		assert.Equal(t, metrics_rollup_rules.AggLast, res.RollupFunction)
		assert.Equal(t, metrics_rollup_rules.LabelsExcludeBy, res.LabelsEliminationMethod)
		assert.Equal(t, []string{"x"}, res.Labels)
		assert.False(t, res.IsDeleted)

	}
}

func TestCreateRollupRuleValidation(t *testing.T) {
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	_, err := underTest.CreateRollupRule(metrics_rollup_rules.CreateUpdateRollupRule{})
	assert.Error(t, err)
}

func TestCreateRollupRuleApiFailed(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		request := metrics_rollup_rules.CreateUpdateRollupRule{
			AccountId:               1,
			MetricName:              "cpu",
			MetricType:              metrics_rollup_rules.MetricTypeGauge,
			RollupFunction:          metrics_rollup_rules.AggLast,
			LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
			Labels:                  []string{"x"},
		}

		res, err := underTest.CreateRollupRule(request)
		assert.Error(t, err)
		assert.Nil(t, res)
	}
}

func TestCreateRollupRuleWithName(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("create_metrics_rollup_rule.json"))
		})

		request := metrics_rollup_rules.CreateUpdateRollupRule{
			AccountId:               1,
			Name:                    "test-rollup-rule",
			MetricName:              "cpu",
			MetricType:              metrics_rollup_rules.MetricTypeGauge,
			RollupFunction:          metrics_rollup_rules.AggLast,
			LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
			Labels:                  []string{"x"},
		}

		res, err := underTest.CreateRollupRule(request)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "my-rollup-rule", res.Name)
	}
}

func TestCreateRollupRuleCounterWithRollupFunction(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("create_metrics_rollup_rule.json"))
		})

		request := metrics_rollup_rules.CreateUpdateRollupRule{
			AccountId:               1,
			MetricName:              "counter_metric",
			MetricType:              metrics_rollup_rules.MetricTypeCounter,
			RollupFunction:          metrics_rollup_rules.AggSum, // Should succeed for COUNTER with SUM
			LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
			Labels:                  []string{"x"},
		}

		res, err := underTest.CreateRollupRule(request)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	}
}

func TestCreateRollupRuleCounterWithoutRollupFunction(t *testing.T) {
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	request := metrics_rollup_rules.CreateUpdateRollupRule{
		AccountId:  1,
		MetricName: "counter_metric",
		MetricType: metrics_rollup_rules.MetricTypeCounter,
		// No RollupFunction - should fail for COUNTER (now required)
		LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
		Labels:                  []string{"x"},
	}

	res, err := underTest.CreateRollupRule(request)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "rollupFunction must be set for COUNTER metrics")
}

func TestCreateRollupRuleCounterWithSum(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("create_metrics_rollup_rule.json"))
		})

		request := metrics_rollup_rules.CreateUpdateRollupRule{
			AccountId:               1,
			MetricName:              "counter_metric",
			MetricType:              metrics_rollup_rules.MetricTypeCounter,
			RollupFunction:          metrics_rollup_rules.AggSum, // SUM is required and supported for COUNTER
			LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
			Labels:                  []string{"x"},
		}

		res, err := underTest.CreateRollupRule(request)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	}
}

func TestCreateRollupRuleCounterWithNonSum(t *testing.T) {
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	request := metrics_rollup_rules.CreateUpdateRollupRule{
		AccountId:               1,
		MetricName:              "counter_metric",
		MetricType:              metrics_rollup_rules.MetricTypeCounter,
		RollupFunction:          metrics_rollup_rules.AggMax, // Should fail - only SUM allowed for COUNTER
		LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
		Labels:                  []string{"x"},
	}

	res, err := underTest.CreateRollupRule(request)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "for COUNTER metrics, rollupFunction must be SUM")
}

func TestCreateRollupRuleDeltaCounterWithSum(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("create_metrics_rollup_rule.json"))
		})

		request := metrics_rollup_rules.CreateUpdateRollupRule{
			AccountId:               1,
			MetricName:              "delta_counter_metric",
			MetricType:              metrics_rollup_rules.MetricTypeDeltaCounter,
			RollupFunction:          metrics_rollup_rules.AggSum, // SUM is required and supported for DELTA_COUNTER
			LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
			Labels:                  []string{"x"},
		}

		res, err := underTest.CreateRollupRule(request)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	}
}

func TestCreateRollupRuleCumulativeCounterWithNonSum(t *testing.T) {
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	request := metrics_rollup_rules.CreateUpdateRollupRule{
		AccountId:               1,
		MetricName:              "cumulative_counter_metric",
		MetricType:              metrics_rollup_rules.MetricTypeCumulativeCounter,
		RollupFunction:          metrics_rollup_rules.AggLast, // Should fail - only SUM allowed for CUMULATIVE_COUNTER
		LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
		Labels:                  []string{"x"},
	}

	res, err := underTest.CreateRollupRule(request)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "for CUMULATIVE_COUNTER metrics, rollupFunction must be SUM")
}

func TestCreateRollupRuleGaugeWithoutRollupFunction(t *testing.T) {
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	request := metrics_rollup_rules.CreateUpdateRollupRule{
		AccountId:               1,
		MetricName:              "gauge_metric",
		MetricType:              metrics_rollup_rules.MetricTypeGauge,
		LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
		Labels:                  []string{"x"},
	}

	res, err := underTest.CreateRollupRule(request)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "rollupFunction must be set for GAUGE metrics")
}

func TestCreateRollupRuleMeasurementWithoutRollupFunction(t *testing.T) {
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	request := metrics_rollup_rules.CreateUpdateRollupRule{
		AccountId:  1,
		MetricName: "measurement_metric",
		MetricType: metrics_rollup_rules.MetricTypeMeasurement,
		// No RollupFunction - should fail for MEASUREMENT
		LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
		Labels:                  []string{"x"},
	}

	res, err := underTest.CreateRollupRule(request)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "rollupFunction must be set for MEASUREMENT metrics")
}

func TestCreateRollupRuleCounterJsonContainsSum(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			// Read and check the request body
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)
			bodyStr := string(body)

			// Should contain rollupFunction with SUM value for COUNTER
			assert.Contains(t, bodyStr, "rollupFunction")
			assert.Contains(t, bodyStr, "SUM")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("create_metrics_rollup_rule.json"))
		})

		request := metrics_rollup_rules.CreateUpdateRollupRule{
			AccountId:               1,
			MetricName:              "counter_metric",
			MetricType:              metrics_rollup_rules.MetricTypeCounter,
			RollupFunction:          metrics_rollup_rules.AggSum, // SUM required for COUNTER
			LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
			Labels:                  []string{"x"},
		}

		res, err := underTest.CreateRollupRule(request)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	}
}

func TestCreateRollupRuleWithMeasurementTypeInvalidAggregation(t *testing.T) {
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	request := metrics_rollup_rules.CreateUpdateRollupRule{
		AccountId:               1,
		Name:                    "measurement-rollup-rule",
		MetricName:              "temperature_measurement",
		MetricType:              metrics_rollup_rules.MetricTypeMeasurement,
		RollupFunction:          metrics_rollup_rules.AggMedian, // Invalid for MEASUREMENT
		LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
		Labels:                  []string{"sensor_id"},
	}

	res, err := underTest.CreateRollupRule(request)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "invalid aggregation function for MEASUREMENT metric type")
}

func TestCreateRollupRuleWithMeasurementTypeValidAggregations(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("create_metrics_rollup_rule.json"))
		})

		request := metrics_rollup_rules.CreateUpdateRollupRule{
			AccountId:               1,
			Name:                    "measurement-rollup-rule-sum",
			MetricName:              "temperature_measurement",
			MetricType:              metrics_rollup_rules.MetricTypeMeasurement,
			RollupFunction:          metrics_rollup_rules.AggSum,
			LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
			Labels:                  []string{"sensor_id"},
		}

		res, err := underTest.CreateRollupRule(request)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	}
}
