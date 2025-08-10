package drop_metrics_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropMetrics_BulkDeleteDropMetrics(t *testing.T) {
	underTest, mux, teardown := setupDropMetricsTest()
	defer teardown()

	mux.HandleFunc("/v1/metrics-management/drop-filters/bulk/delete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	dropFilterIds := []int64{1, 2, 3}
	err := underTest.BulkDeleteDropMetrics(dropFilterIds)
	assert.NoError(t, err)
}

func TestDropMetrics_BulkDeleteDropMetricsEmptyArray(t *testing.T) {
	underTest, _, teardown := setupDropMetricsTest()
	defer teardown()

	err := underTest.BulkDeleteDropMetrics([]int64{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "dropFilterIds array cannot be empty")
}

func TestDropMetrics_BulkDeleteDropMetricsInvalidId(t *testing.T) {
	underTest, _, teardown := setupDropMetricsTest()
	defer teardown()

	dropFilterIds := []int64{1, 0, 3}
	err := underTest.BulkDeleteDropMetrics(dropFilterIds)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "dropFilterIds[1] must be greater than 0")
}

func TestDropMetrics_BulkDeleteDropMetricsAPIFailed(t *testing.T) {
	underTest, mux, teardown := setupDropMetricsTest()
	defer teardown()

	mux.HandleFunc("/v1/metrics-management/drop-filters/bulk/delete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("api_error.txt"))
	})

	dropFilterIds := []int64{1, 2, 3}
	err := underTest.BulkDeleteDropMetrics(dropFilterIds)
	assert.Error(t, err)
}

func TestDropMetrics_BulkDeleteDropMetricsNotFound(t *testing.T) {
	underTest, mux, teardown := setupDropMetricsTest()
	defer teardown()

	mux.HandleFunc("/v1/metrics-management/drop-filters/bulk/delete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found")
	})

	dropFilterIds := []int64{999, 998}
	err := underTest.BulkDeleteDropMetrics(dropFilterIds)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed with missing drop_metric")
}
