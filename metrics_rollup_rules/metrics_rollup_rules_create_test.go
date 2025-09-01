package metrics_rollup_rules_test

import (
	"fmt"
	"net/http"
	"strings"
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

func TestCreateRollupRuleNameTooLong(t *testing.T) {
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	longName := strings.Repeat("a", 257)

	request := metrics_rollup_rules.CreateUpdateRollupRule{
		AccountId:               1,
		Name:                    longName,
		MetricName:              "cpu",
		MetricType:              metrics_rollup_rules.MetricTypeGauge,
		RollupFunction:          metrics_rollup_rules.AggLast,
		LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
		Labels:                  []string{"x"},
	}

	res, err := underTest.CreateRollupRule(request)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "name must not exceed 256 characters")
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
