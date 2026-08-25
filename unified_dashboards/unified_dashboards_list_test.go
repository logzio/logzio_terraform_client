package unified_dashboards_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnifiedDashboards_ListDashboards(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/dashboards", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("list_dashboards.json"))
		})

		dashboards, err := underTest.ListDashboards()
		assert.NoError(t, err)
		if assert.Len(t, dashboards, 2) {
			assert.Equal(t, "dashboard-1", dashboards[0].Uid)
			assert.Equal(t, "project-1", dashboards[0].ProjectId)
			assert.Equal(t, "Dashboard", dashboards[0].Doc["kind"])
			assert.Equal(t, 1, dashboards[0].Version)
			assert.Equal(t, "dashboard-2", dashboards[1].Uid)
		}
	}
}

func TestUnifiedDashboards_ListDashboardsAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/dashboards", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.ListDashboards()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status code 500")
	}
}

func TestUnifiedDashboards_ListDashboardsNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/dashboards", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err = underTest.ListDashboards()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing unified dashboard")
	}
}
