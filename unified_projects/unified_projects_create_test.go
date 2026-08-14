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
			var target map[string]interface{}
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			// The API expects a Perses Project envelope — pin the wire shape literally.
			assert.Equal(t, "Project", target["kind"])
			assert.Equal(t, "new-project", target["metadata"].(map[string]interface{})["name"])
			display := target["spec"].(map[string]interface{})["display"].(map[string]interface{})
			assert.Equal(t, "new-project", display["name"])

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("create_project.json"))
		})

		project, err := underTest.CreateProject(unified_projects.CreateProjectRequest{Name: "new-project"})
		assert.NoError(t, err)
		if assert.NotNil(t, project) {
			assert.Equal(t, "project-new", project.Id)
			assert.Equal(t, "new-project", project.Name)
			assert.Equal(t, "Project", project.Doc["kind"])
			assert.NotEmpty(t, project.CreatedAt)
		}
	}
}

func TestUnifiedProjects_CreateProjectWithDisplayName(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects", func(w http.ResponseWriter, r *http.Request) {
			jsonBytes, _ := io.ReadAll(r.Body)
			var target map[string]interface{}
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			display := target["spec"].(map[string]interface{})["display"].(map[string]interface{})
			assert.Equal(t, "Pretty Name", display["name"])
			assert.Equal(t, "some description", display["description"])

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("create_project.json"))
		})

		project, err := underTest.CreateProject(unified_projects.CreateProjectRequest{
			Name:        "new-project",
			DisplayName: "Pretty Name",
			Description: "some description",
		})
		assert.NoError(t, err)
		assert.NotNil(t, project)
	}
}

func TestUnifiedProjects_CreateProjectAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.CreateProject(unified_projects.CreateProjectRequest{Name: "new-project"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status code 500")
	}
}

func TestUnifiedProjects_CreateProjectNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err = underTest.CreateProject(unified_projects.CreateProjectRequest{Name: "new-project"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing unified project")
	}
}

func TestUnifiedProjects_CreateProjectEmptyResponse(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "{}")
		})

		_, err = underTest.CreateProject(unified_projects.CreateProjectRequest{Name: "new-project"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "contained no project id")
	}
}

func TestUnifiedProjects_CreateProjectNoName(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err = underTest.CreateProject(unified_projects.CreateProjectRequest{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name must be set")
	}
}
