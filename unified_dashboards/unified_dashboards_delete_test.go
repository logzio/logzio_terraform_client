package unified_dashboards_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnifiedDashboards_DeleteDashboard(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/dashboards/dashboard-1", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			w.WriteHeader(http.StatusNoContent)
		})

		err = underTest.DeleteDashboard("project-1", "dashboard-1")
		assert.NoError(t, err)
	}
}

func TestUnifiedDashboards_DeleteDashboardAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/dashboards/dashboard-1", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		err = underTest.DeleteDashboard("project-1", "dashboard-1")
		assert.Error(t, err)
	}
}

func TestUnifiedDashboards_DeleteDashboardNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/dashboards/missing", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		err = underTest.DeleteDashboard("project-1", "missing")
		assert.Error(t, err)
	}
}

func TestUnifiedDashboards_DeleteDashboardValidation(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		err = underTest.DeleteDashboard("", "dashboard-1")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "folderId must be set")

		err = underTest.DeleteDashboard("project-1", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "uid must be set")
	}
}
