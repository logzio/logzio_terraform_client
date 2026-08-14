package unified_projects_test

import (
	"fmt"
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
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "system", r.URL.Query().Get("query"))
			assert.Equal(t, "10", r.URL.Query().Get("limit"))
			assert.Equal(t, "1", r.URL.Query().Get("page"))

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("search_projects.json"))
		})

		results, err := underTest.SearchProjects(unified_projects.SearchProjectsRequest{
			Query: "system",
			Limit: 10,
			Page:  1,
		})
		assert.NoError(t, err)
		assert.Len(t, results, 1)
		assert.Equal(t, "project-1", results[0].Id)
		assert.Equal(t, "system-metrics", results[0].Name)
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
	}
}

func TestUnifiedProjects_SearchProjectsNoQuery(t *testing.T) {
	underTest, err, teardown := setupUnifiedProjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err = underTest.SearchProjects(unified_projects.SearchProjectsRequest{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "query must be set")
	}
}
