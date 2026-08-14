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
		mux.HandleFunc("/perses-public/api/v1/projects/system-metrics", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)

			jsonBytes, _ := io.ReadAll(r.Body)
			var target unified_projects.UpdateProjectRequest
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.Equal(t, "System Metrics Updated", target.DisplayName)
			assert.Equal(t, "Updated description", target.Description)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("update_project.json"))
		})

		updated, err := underTest.UpdateProject("system-metrics", unified_projects.UpdateProjectRequest{
			DisplayName: "System Metrics Updated",
			Description: "Updated description",
		})
		assert.NoError(t, err)
		assert.NotNil(t, updated)
		assert.Equal(t, "project-1", updated.Id)
		assert.Equal(t, "system-metrics", updated.Name)
	}
}

func TestUnifiedProjects_UpdateProjectAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/system-metrics", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.UpdateProject("system-metrics", unified_projects.UpdateProjectRequest{DisplayName: "x"})
		assert.Error(t, err)
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

		_, err = underTest.UpdateProject("missing", unified_projects.UpdateProjectRequest{DisplayName: "x"})
		assert.Error(t, err)
	}
}

func TestUnifiedProjects_UpdateProjectNoName(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err = underTest.UpdateProject("", unified_projects.UpdateProjectRequest{DisplayName: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name must be set")
	}
}

func TestUnifiedProjects_UpdateProjectEmptyRequest(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err = underTest.UpdateProject("system-metrics", unified_projects.UpdateProjectRequest{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "displayName or description must be set")
	}
}
