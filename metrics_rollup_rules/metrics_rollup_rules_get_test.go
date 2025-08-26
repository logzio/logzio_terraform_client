package metrics_rollup_rules_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRollupRuleSuccess(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath+"/abc", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("get_metrics_rollup_rule.json"))
		})

		res, err := underTest.GetRollupRule("abc")
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "abc", res.Id)
		assert.Equal(t, "get-rollup-rule", res.Name)
	}
}

func TestGetRollupRuleValidation(t *testing.T) {
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	_, err := underTest.GetRollupRule("")
	assert.Error(t, err)
}

func TestGetRollupRuleApiFailed(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath+"/abc", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		res, err := underTest.GetRollupRule("abc")
		assert.Error(t, err)
		assert.Nil(t, res)
	}
}

func TestGetRollupRuleNotFound(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath+"/missing", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			w.WriteHeader(http.StatusNotFound)
		})

		res, err := underTest.GetRollupRule("missing")
		assert.Error(t, err)
		assert.Nil(t, res)
	}
}
