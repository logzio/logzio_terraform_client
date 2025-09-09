package unified_projects_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnifiedProjects_ListProjects(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "", r.URL.Query().Get("withDashboards"))

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("list_projects.json"))
		})

		projects, err := underTest.ListProjects(false)
		assert.NoError(t, err)
		assert.Len(t, projects, 2)
		assert.Equal(t, "project-1", projects[0].Project.Id)
		assert.Equal(t, "system-metrics", projects[0].Project.Name)
		assert.Equal(t, "System Metrics", projects[0].Project.DisplayName)
		assert.Equal(t, "System monitoring dashboards", *projects[0].Project.Description)
		assert.Len(t, projects[0].Project.Dashboards, 2)
		assert.Equal(t, "dashboard-1", projects[0].Project.Dashboards[0].Uid)
		assert.Equal(t, "CPU Usage", projects[0].Project.Dashboards[0].Title)
	}
}

func TestUnifiedProjects_ListProjectsWithDashboards(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "true", r.URL.Query().Get("withDashboards"))

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("list_projects.json"))
		})

		projects, err := underTest.ListProjects(true)
		assert.NoError(t, err)
		assert.Len(t, projects, 2)
	}
}

func TestUnifiedProjects_ListProjectsAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.ListProjects(false)
		assert.Error(t, err)
	}
}

func TestUnifiedProjects_ListProjectsNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err = underTest.ListProjects(false)
		assert.Error(t, err)
	}
}
