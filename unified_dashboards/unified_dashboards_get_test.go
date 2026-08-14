package unified_dashboards_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnifiedDashboards_GetDashboard(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/dashboards/dashboard-1", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("get_dashboard.json"))
		})

		dashboard, err := underTest.GetDashboard("project-1", "dashboard-1")
		assert.NoError(t, err)
		assert.NotNil(t, dashboard)
		assert.Equal(t, "dashboard-1", dashboard.Uid)
		assert.Equal(t, "System Overview", dashboard.Doc["title"])
		assert.Equal(t, 1, dashboard.Version)
	}
}

func TestUnifiedDashboards_GetDashboardAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/dashboards/dashboard-1", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.GetDashboard("project-1", "dashboard-1")
		assert.Error(t, err)
	}
}

func TestUnifiedDashboards_GetDashboardNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/dashboards/missing", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err = underTest.GetDashboard("project-1", "missing")
		assert.Error(t, err)
	}
}

func TestUnifiedDashboards_GetDashboardValidation(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err = underTest.GetDashboard("", "dashboard-1")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "folderId must be set")

		_, err = underTest.GetDashboard("project-1", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "uid must be set")
	}
}
