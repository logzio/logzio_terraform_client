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

func getCreateDashboardRequest() unified_dashboards.CreateDashboardRequest {
	return unified_dashboards.CreateDashboardRequest{
		Doc: map[string]interface{}{
			"kind":     "Dashboard",
			"metadata": map[string]interface{}{"name": "cpu-usage-dashboard"},
			"spec": map[string]interface{}{
				"display":  map[string]interface{}{"name": "CPU Usage Dashboard"},
				"duration": "1h",
				"panels":   map[string]interface{}{},
				"layouts":  []interface{}{},
			},
		},
	}
}

func TestUnifiedDashboards_CreateDashboard(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/dashboards", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			jsonBytes, _ := io.ReadAll(r.Body)
			var target map[string]interface{}
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			// Dashboards wrap the Perses document in a "doc" key (projects send it bare) — pin the wrapper literally.
			doc, ok := target["doc"].(map[string]interface{})
			if assert.True(t, ok, `request body must nest the document under "doc"`) {
				assert.Equal(t, "Dashboard", doc["kind"])
				assert.Equal(t, "cpu-usage-dashboard", doc["metadata"].(map[string]interface{})["name"])
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("create_dashboard.json"))
		})

		created, err := underTest.CreateDashboard("project-1", getCreateDashboardRequest())
		assert.NoError(t, err)
		if assert.NotNil(t, created) {
			assert.Equal(t, "dashboard-new", created.Uid)
			assert.Equal(t, "project-1", created.ProjectId)
			assert.Equal(t, "Dashboard", created.Doc["kind"])
			assert.Equal(t, 1, created.Version)
		}
	}
}

func TestUnifiedDashboards_CreateDashboardAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/dashboards", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.CreateDashboard("project-1", getCreateDashboardRequest())
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status code 500")
	}
}

func TestUnifiedDashboards_CreateDashboardNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/missing/dashboards", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err = underTest.CreateDashboard("missing", getCreateDashboardRequest())
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing unified dashboard")
	}
}

func TestUnifiedDashboards_CreateDashboardValidation(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err = underTest.CreateDashboard("", getCreateDashboardRequest())
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "folderId must be set")

		_, err = underTest.CreateDashboard("project-1", unified_dashboards.CreateDashboardRequest{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "doc must be set")
	}
}

func TestUnifiedDashboards_CreateDashboardEmptyResponse(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/dashboards", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "{}")
		})

		_, err = underTest.CreateDashboard("project-1", getCreateDashboardRequest())
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "contained no dashboard uid")
	}
}
