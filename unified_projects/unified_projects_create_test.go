package unified_projects_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestUnifiedProjects_CreateProject(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			jsonBytes, _ := io.ReadAll(r.Body)
			var target unified_projects.CreateProjectRequest
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotEmpty(t, target.Name)
			assert.Equal(t, "new-project", target.Name)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, fixture("create_project.json"))
		})

		req := unified_projects.CreateProjectRequest{
			Name: "new-project",
		}

		project, err := underTest.CreateProject(req)
		assert.NoError(t, err)
		assert.Equal(t, "project-new", project.Id)
		assert.Equal(t, "new-project", project.Name)
		assert.Equal(t, "new-project", project.DisplayName)
		assert.NotNil(t, project.CreatedAt)
	}
}

func TestUnifiedProjects_CreateProjectAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		req := unified_projects.CreateProjectRequest{
			Name: "new-project",
		}

		_, err = underTest.CreateProject(req)
		assert.Error(t, err)
	}
}

func TestUnifiedProjects_CreateProjectNoName(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		req := unified_projects.CreateProjectRequest{
			Name: "",
		}

		_, err = underTest.CreateProject(req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name must be set")
	}
}
