package unified_dashboards_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/stretchr/testify/assert"
)

func TestUnifiedDashboards_UpdateDashboard(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/dashboards/dashboard-1", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)

			jsonBytes, _ := io.ReadAll(r.Body)
			var target unified_dashboards.UpdateDashboardRequest
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.Equal(t, "Dashboard", target.Doc["kind"])

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("update_dashboard.json"))
		})

		updated, err := underTest.UpdateDashboard("project-1", "dashboard-1", unified_dashboards.UpdateDashboardRequest{
			Doc: map[string]interface{}{
				"kind":     "Dashboard",
				"metadata": map[string]interface{}{"name": "system-overview"},
				"spec":     map[string]interface{}{"display": map[string]interface{}{"name": "Updated System Overview"}},
			},
		})
		assert.NoError(t, err)
		if assert.NotNil(t, updated) {
			assert.Equal(t, "dashboard-1", updated.Uid)
			assert.Equal(t, "project-1", updated.ProjectId)
			assert.Equal(t, 2, updated.Version)
		}
	}
}

func TestUnifiedDashboards_UpdateDashboardAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/dashboards/dashboard-1", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.UpdateDashboard("project-1", "dashboard-1", unified_dashboards.UpdateDashboardRequest{
			Doc: map[string]interface{}{"title": "x"},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status code 500")
	}
}

func TestUnifiedDashboards_UpdateDashboardNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/dashboards/missing", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err = underTest.UpdateDashboard("project-1", "missing", unified_dashboards.UpdateDashboardRequest{
			Doc: map[string]interface{}{"title": "x"},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing unified dashboard")
	}
}

func TestUnifiedDashboards_UpdateDashboardValidation(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err = underTest.UpdateDashboard("", "dashboard-1", unified_dashboards.UpdateDashboardRequest{Doc: map[string]interface{}{"x": 1}})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "folderId must be set")

		_, err = underTest.UpdateDashboard("project-1", "", unified_dashboards.UpdateDashboardRequest{Doc: map[string]interface{}{"x": 1}})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "uid must be set")

		_, err = underTest.UpdateDashboard("project-1", "dashboard-1", unified_dashboards.UpdateDashboardRequest{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "doc must be set")
	}
}
