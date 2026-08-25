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
		if assert.Len(t, projects, 2) {
			assert.Equal(t, "project-1", projects[0].Project.Id)
			assert.Equal(t, "System Metrics", projects[0].Project.Name)
			assert.Equal(t, "Project", projects[0].Project.Doc["kind"])
			// dashboards are a sibling of "project" in each list entry
			if assert.Len(t, projects[0].Dashboards, 1) {
				// uid and id are deliberately different values in the fixture:
				// they coincide only on a never-updated dashboard, so a decode
				// that dropped one of them would still pass if they matched.
				assert.Equal(t, "1f96a105-8ec3-4242-b074-0f57f37e7fbb", projects[0].Dashboards[0].Uid)
				assert.Equal(t, "3da41d03-ca61-436d-be45-69047d4f84be", projects[0].Dashboards[0].Id)
				assert.Equal(t, "project-1", projects[0].Dashboards[0].ProjectId)
			}
			assert.Empty(t, projects[1].Dashboards)
		}
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
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.ListProjects(false)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status code 500")
	}
}

func TestUnifiedProjects_ListProjectsNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err = underTest.ListProjects(false)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing unified project")
	}
}
