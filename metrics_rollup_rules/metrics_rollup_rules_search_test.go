package metrics_rollup_rules_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/metrics_rollup_rules"
	"github.com/stretchr/testify/assert"
)

func TestSearchRollupRulesSuccess(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath+"/search", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `[{"id":"a","accountId":1,"metricType":"GAUGE","rollupFunction":"LAST","labelsEliminationMethod":"EXCLUDE_BY","labels":["x"],"isDeleted":false,"dropOriginalMetric":false,"version":1}]`)
		})

		res, err := underTest.SearchRollupRules(metrics_rollup_rules.SearchRollupRulesRequest{})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res, 1)
		assert.Equal(t, "a", res[0].Id)
	}
}

func TestSearchRollupRulesApiFailed(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath+"/search", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		res, err := underTest.SearchRollupRules(metrics_rollup_rules.SearchRollupRulesRequest{})
		assert.Error(t, err)
		assert.Nil(t, res)
	}
}

func TestSearchRollupRulesWithSearchTerm(t *testing.T) {
	underTest, err, teardown := setupMetricsRollupRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(metricsRollupRulesPath+"/search", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `[{"id":"a","name":"cpu-rollup-rule","accountId":1,"metricType":"GAUGE","rollupFunction":"LAST","labelsEliminationMethod":"EXCLUDE_BY","labels":["x"],"isDeleted":false,"dropOriginalMetric":false,"version":1}]`)
		})

		searchReq := metrics_rollup_rules.SearchRollupRulesRequest{
			Filter: &metrics_rollup_rules.SearchFilter{
				AccountIds: []int64{1},
				SearchTerm: "cpu-rollup", // Test the new searchTerm field
			},
			Pagination: &metrics_rollup_rules.Pagination{
				PageNumber: 0,
				PageSize:   10,
			},
		}

		res, err := underTest.SearchRollupRules(searchReq)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res, 1)
		assert.Equal(t, "cpu-rollup-rule", res[0].Name)
	}
}
