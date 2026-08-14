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
			"title":   "CPU Usage Dashboard",
			"panels":  []interface{}{map[string]interface{}{"id": 1, "title": "CPU Usage", "type": "graph"}},
			"refresh": "1m",
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
			var target unified_dashboards.CreateDashboardRequest
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.Equal(t, "CPU Usage Dashboard", target.Doc["title"])

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, fixture("create_dashboard.json"))
		})

		created, err := underTest.CreateDashboard("project-1", getCreateDashboardRequest())
		assert.NoError(t, err)
		assert.NotNil(t, created)
		assert.Equal(t, "dashboard-new", created.Uid)
		assert.Equal(t, "CPU Usage Dashboard", created.Doc["title"])
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
