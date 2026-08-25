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

func TestUnifiedDashboards_MoveDashboard(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/dashboards/move", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)

			jsonBytes, _ := io.ReadAll(r.Body)
			var target map[string]any
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.Equal(t, "dashboard-1", target["dashboardId"])
			assert.Equal(t, "project-1", target["oldProjectId"])
			assert.Equal(t, "project-2", target["newProjectId"])

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("move_dashboard.json"))
		})

		moved, err := underTest.MoveDashboard(unified_dashboards.MoveDashboardRequest{
			DashboardId:  "dashboard-1",
			OldProjectId: "project-1",
			NewProjectId: "project-2",
		})
		assert.NoError(t, err)
		if assert.NotNil(t, moved) {
			assert.Equal(t, "dashboard-1", moved.Id)
		}
	}
}

func TestUnifiedDashboards_MoveDashboardAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/dashboards/move", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.MoveDashboard(unified_dashboards.MoveDashboardRequest{
			DashboardId: "dashboard-1", OldProjectId: "project-1", NewProjectId: "project-2",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status code 500")
	}
}

func TestUnifiedDashboards_MoveDashboardNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/dashboards/move", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err = underTest.MoveDashboard(unified_dashboards.MoveDashboardRequest{
			DashboardId: "missing", OldProjectId: "project-1", NewProjectId: "project-2",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing unified dashboard")
	}
}

func TestUnifiedDashboards_MoveDashboardEmptyResponse(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/dashboards/move", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "{}")
		})

		_, err = underTest.MoveDashboard(unified_dashboards.MoveDashboardRequest{
			DashboardId: "dashboard-1", OldProjectId: "project-1", NewProjectId: "project-2",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "contained no dashboard id")
	}
}

func TestUnifiedDashboards_MoveDashboardValidation(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err = underTest.MoveDashboard(unified_dashboards.MoveDashboardRequest{OldProjectId: "p1", NewProjectId: "p2"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "dashboardId must be set")

		_, err = underTest.MoveDashboard(unified_dashboards.MoveDashboardRequest{DashboardId: "d1", NewProjectId: "p2"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "oldProjectId must be set")

		_, err = underTest.MoveDashboard(unified_dashboards.MoveDashboardRequest{DashboardId: "d1", OldProjectId: "p1"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "newProjectId must be set")
	}
}
