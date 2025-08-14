package metrics_rollup_rules_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBulkDeleteRollupRulesSuccess(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/metrics-management/rollup-rules/bulk/delete", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusNoContent)
		})

		err = underTest.BulkDeleteRollupRules([]string{"a", "b"})
		assert.NoError(t, err)
	}
}

func TestBulkDeleteRollupRulesApiFailed(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/metrics-management/rollup-rules/bulk/delete", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		err := underTest.BulkDeleteRollupRules([]string{"a", "b"})
		assert.Error(t, err)
	}
}

func TestBulkDeleteRollupRulesValidation(t *testing.T) {
	underTest, _, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	err := underTest.BulkDeleteRollupRules([]string{})
	assert.Error(t, err)
}
