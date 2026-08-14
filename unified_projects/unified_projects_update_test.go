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

func TestUnifiedProjects_UpdateProject(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)

			jsonBytes, _ := io.ReadAll(r.Body)
			var target map[string]interface{}
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.Equal(t, "Project", target["kind"])
			assert.Equal(t, "system-metrics", target["metadata"].(map[string]interface{})["name"])
			display := target["spec"].(map[string]interface{})["display"].(map[string]interface{})
			assert.Equal(t, "System Metrics Updated", display["name"])
			assert.Equal(t, "Updated description", display["description"])

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("update_project.json"))
		})

		updated, err := underTest.UpdateProject("project-1", unified_projects.UpdateProjectRequest{
			Name:        "system-metrics",
			DisplayName: "System Metrics Updated",
			Description: "Updated description",
		})
		assert.NoError(t, err)
		if assert.NotNil(t, updated) {
			assert.Equal(t, "project-1", updated.Id)
			assert.Equal(t, "System Metrics Updated", updated.Name)
		}
	}
}

func TestUnifiedProjects_UpdateProjectOmitsEmptyDescription(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1", func(w http.ResponseWriter, r *http.Request) {
			jsonBytes, _ := io.ReadAll(r.Body)
			var target map[string]interface{}
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			display := target["spec"].(map[string]interface{})["display"].(map[string]interface{})
			_, hasDescription := display["description"]
			assert.False(t, hasDescription, "empty description must be omitted (the PUT replaces the document, omitting clears it)")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("update_project.json"))
		})

		_, err = underTest.UpdateProject("project-1", unified_projects.UpdateProjectRequest{
			Name:        "system-metrics",
			DisplayName: "System Metrics Updated",
		})
		assert.NoError(t, err)
	}
}

func TestUnifiedProjects_UpdateProjectAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.UpdateProject("project-1", unified_projects.UpdateProjectRequest{
			Name:        "system-metrics",
			DisplayName: "x",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status code 500")
	}
}

func TestUnifiedProjects_UpdateProjectNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/missing", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err = underTest.UpdateProject("missing", unified_projects.UpdateProjectRequest{
			Name:        "system-metrics",
			DisplayName: "x",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing unified project")
	}
}

func TestUnifiedProjects_UpdateProjectValidation(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err = underTest.UpdateProject("", unified_projects.UpdateProjectRequest{Name: "n", DisplayName: "d"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id must be set")

		_, err = underTest.UpdateProject("project-1", unified_projects.UpdateProjectRequest{DisplayName: "d"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name must be set")

		_, err = underTest.UpdateProject("project-1", unified_projects.UpdateProjectRequest{Name: "n"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "displayName must be set")
	}
}
