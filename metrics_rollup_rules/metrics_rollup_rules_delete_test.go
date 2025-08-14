package metrics_rollup_rules_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteRollupRuleSuccess(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/metrics-management/rollup-rules/abc", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			w.WriteHeader(http.StatusOK)
		})

		err := underTest.DeleteRollupRule("abc")
		assert.NoError(t, err)
	}
}

func TestDeleteRollupRuleValidation(t *testing.T) {
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	err := underTest.DeleteRollupRule("")
	assert.Error(t, err)
}

func TestDeleteRollupRuleApiFailed(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/metrics-management/rollup-rules/abc", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		err := underTest.DeleteRollupRule("abc")
		assert.Error(t, err)
	}
}
