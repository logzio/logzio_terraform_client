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

func TestUnifiedProjects_SearchProjects(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/search", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			jsonBytes, _ := io.ReadAll(r.Body)
			var target map[string]interface{}
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.Equal(t, "system", target["query"])
			assert.Equal(t, float64(10), target["limit"])
			assert.Equal(t, float64(1), target["page"])

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("search_projects.json"))
		})

		resp, err := underTest.SearchProjects(unified_projects.SearchProjectsRequest{
			Query: "system",
			Limit: 10,
			Page:  1,
		})
		assert.NoError(t, err)
		if assert.NotNil(t, resp) {
			assert.Equal(t, 184, resp.Total)
			if assert.Len(t, resp.Results, 1) {
				assert.Equal(t, "project-1", resp.Results[0].Project.Id)
				assert.Equal(t, "System Metrics", resp.Results[0].Project.Name)
			}
		}
	}
}

func TestUnifiedProjects_SearchProjectsOmitsUnsetFields(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/search", func(w http.ResponseWriter, r *http.Request) {
			jsonBytes, _ := io.ReadAll(r.Body)
			var target map[string]interface{}
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			_, hasLimit := target["limit"]
			_, hasPage := target["page"]
			assert.False(t, hasLimit, "unset limit must be omitted")
			assert.False(t, hasPage, "unset page must be omitted")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("search_projects.json"))
		})

		_, err = underTest.SearchProjects(unified_projects.SearchProjectsRequest{Query: "system"})
		assert.NoError(t, err)
	}
}

func TestUnifiedProjects_SearchProjectsAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/search", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.SearchProjects(unified_projects.SearchProjectsRequest{Query: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status code 500")
	}
}

func TestUnifiedProjects_SearchProjectsNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/search", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err = underTest.SearchProjects(unified_projects.SearchProjectsRequest{Query: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing unified project")
	}
}
