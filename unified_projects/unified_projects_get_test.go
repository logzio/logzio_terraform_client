package unified_projects_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestUnifiedProjects_GetProject(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("get_project.json"))
		})

		project, err := underTest.GetProject("project-1")
		assert.NoError(t, err)
		if assert.NotNil(t, project) {
			assert.Equal(t, "project-1", project.Id)
			assert.Equal(t, "System Metrics", project.Name)
			assert.Equal(t, "Project", project.Doc["kind"])
			assert.Equal(t, "system-metrics", project.MetadataName())
			assert.Empty(t, (&unified_projects.ProjectSummary{}).MetadataName(), "MetadataName must be nil-safe")
		}
	}
}

func TestUnifiedProjects_GetProjectAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.GetProject("project-1")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status code 500")
	}
}

func TestUnifiedProjects_GetProjectNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/missing", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err = underTest.GetProject("missing")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing unified project")
	}
}

func TestUnifiedProjects_GetProjectNoId(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err = underTest.GetProject("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id must be set")
	}
}
