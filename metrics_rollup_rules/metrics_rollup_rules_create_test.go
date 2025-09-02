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
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	request := metrics_rollup_rules.CreateUpdateRollupRule{
		AccountId:               1,
		MetricName:              "counter_metric",
		MetricType:              metrics_rollup_rules.MetricTypeCounter,
		RollupFunction:          metrics_rollup_rules.AggSum, // Should fail for COUNTER
		LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
		Labels:                  []string{"x"},
	}

	res, err := underTest.CreateRollupRule(request)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "rollupFunction is supported only for GAUGE and MEASUREMENT metrics")
}

func TestCreateRollupRuleCounterWithoutRollupFunction(t *testing.T) {
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
			LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
			Labels:                  []string{"x"},
		}

		res, err := underTest.CreateRollupRule(request)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	}
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

func TestCreateRollupRuleCounterJsonOmitsRollupFunction(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			// Read and check the request body
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)
			bodyStr := string(body)

			// Should not contain rollupFunction in JSON when metric type is COUNTER
			assert.NotContains(t, bodyStr, "rollupFunction")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("create_metrics_rollup_rule.json"))
		})

		request := metrics_rollup_rules.CreateUpdateRollupRule{
			AccountId:  1,
			MetricName: "counter_metric",
			MetricType: metrics_rollup_rules.MetricTypeCounter,
			// No RollupFunction - should be omitted from JSON for COUNTER
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
