package unified_projects_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnifiedProjects_RenameProject(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/rename", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)

			jsonBytes, _ := io.ReadAll(r.Body)
			var target map[string]string
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.Equal(t, "renamed-project", target["newProjectName"])

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("rename_project.json"))
		})

		renamed, err := underTest.RenameProject("project-1", "renamed-project")
		assert.NoError(t, err)
		assert.NotNil(t, renamed)
		assert.Equal(t, "project-1", renamed.Id)
		assert.Equal(t, "renamed-project", renamed.Name)
	}
}

func TestUnifiedProjects_RenameProjectAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/rename", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.RenameProject("project-1", "renamed-project")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status code 500")
	}
}

func TestUnifiedProjects_RenameProjectNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/missing/rename", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err = underTest.RenameProject("missing", "renamed-project")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing unified project")
	}
}

func TestUnifiedProjects_RenameProjectValidation(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err = underTest.RenameProject("", "x")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id must be set")

		_, err = underTest.RenameProject("project-1", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "newName must be set")
	}
}
