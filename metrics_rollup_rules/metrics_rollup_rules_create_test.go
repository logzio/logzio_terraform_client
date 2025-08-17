package metrics_rollup_rules_test

import (
	"fmt"
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
