package metrics_rollup_rules_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/metrics_rollup_rules"
	"github.com/stretchr/testify/assert"
)

func TestUpdateRollupRuleSuccess(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath+"/abc", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("update_metrics_rollup_rule.json"))
		})

		request := metrics_rollup_rules.CreateUpdateRollupRule{
			MetricName:              "cpu2",
			MetricType:              metrics_rollup_rules.MetricTypeCounter,
			RollupFunction:          metrics_rollup_rules.AggSum,
			LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
			Labels:                  []string{"host", "region"},
		}

		res, err := underTest.UpdateRollupRule("abc", request)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "abc", res.Id)
		assert.Equal(t, "updated-rollup-rule", res.Name)
		assert.Equal(t, "cpu2", res.MetricName)
	}
}

func TestUpdateRollupRuleValidation(t *testing.T) {
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	_, err := underTest.UpdateRollupRule("", metrics_rollup_rules.CreateUpdateRollupRule{MetricName: "x"})
	assert.Error(t, err)

	_, err = underTest.UpdateRollupRule("abc", metrics_rollup_rules.CreateUpdateRollupRule{})
	assert.Error(t, err)
}

func TestUpdateRollupRuleCounterWithSum(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath+"/abc", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("update_metrics_rollup_rule.json"))
		})

		request := metrics_rollup_rules.CreateUpdateRollupRule{
			MetricName:              "counter_metric",
			MetricType:              metrics_rollup_rules.MetricTypeCounter,
			RollupFunction:          metrics_rollup_rules.AggSum,
			LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
			Labels:                  []string{"host", "region"},
		}

		res, err := underTest.UpdateRollupRule("abc", request)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	}
}

func TestUpdateRollupRuleCounterWithNonSum(t *testing.T) {
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	request := metrics_rollup_rules.CreateUpdateRollupRule{
		MetricName:              "counter_metric",
		MetricType:              metrics_rollup_rules.MetricTypeCounter,
		RollupFunction:          metrics_rollup_rules.AggMax,
		LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
		Labels:                  []string{"host", "region"},
	}

	res, err := underTest.UpdateRollupRule("abc", request)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "for COUNTER metrics, rollupFunction must be SUM")
}

func TestUpdateRollupRuleCounterWithoutRollupFunction(t *testing.T) {
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	request := metrics_rollup_rules.CreateUpdateRollupRule{
		MetricName:              "counter_metric",
		MetricType:              metrics_rollup_rules.MetricTypeCounter,
		LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
		Labels:                  []string{"host", "region"},
	}

	res, err := underTest.UpdateRollupRule("abc", request)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "rollupFunction must be set for COUNTER metrics")
}

func TestUpdateRollupRuleApiFailed(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath+"/abc", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		req := metrics_rollup_rules.CreateUpdateRollupRule{MetricName: "cpu2"}
		res, err := underTest.UpdateRollupRule("abc", req)
		assert.Error(t, err)
		assert.Nil(t, res)
	}
}
