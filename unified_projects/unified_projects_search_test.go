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
			var target map[string]any
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			// Pin the server's real search schema literally.
			filter, ok := target["filter"].(map[string]any)
			if assert.True(t, ok, `body must carry a "filter" object`) {
				assert.Equal(t, "system", filter["searchTerm"])
				assert.Equal(t, []any{float64(123)}, filter["createdBy"])
			}
			pagination, ok := target["pagination"].(map[string]any)
			if assert.True(t, ok, `body must carry a "pagination" object`) {
				assert.Equal(t, float64(1), pagination["pageNumber"])
				assert.Equal(t, float64(10), pagination["pageSize"])
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("search_projects.json"))
		})

		resp, err := underTest.SearchProjects(unified_projects.SearchProjectsRequest{
			Filter:     &unified_projects.SearchProjectsFilter{SearchTerm: "system", CreatedBy: []int64{123}},
			Pagination: &unified_projects.SearchProjectsPagination{PageNumber: 1, PageSize: 10},
		})
		assert.NoError(t, err)
		if assert.NotNil(t, resp) {
			assert.Equal(t, 1, resp.Total)
			assert.Equal(t, &unified_projects.SearchProjectsPagination{PageNumber: 1, PageSize: 10}, resp.Pagination)
			if assert.Len(t, resp.Results, 1) {
				assert.Equal(t, "project-1", resp.Results[0].Project.Id)
				assert.Equal(t, "System Metrics", resp.Results[0].Project.Name)
				if assert.Len(t, resp.Results[0].Dashboards, 1) {
					// See the list test: uid and id differ on purpose.
					assert.Equal(t, "1f96a105-8ec3-4242-b074-0f57f37e7fbb", resp.Results[0].Dashboards[0].Uid)
					assert.Equal(t, "3da41d03-ca61-436d-be45-69047d4f84be", resp.Results[0].Dashboards[0].Id)
					assert.Equal(t, "project-1", resp.Results[0].Dashboards[0].ProjectId)
				}
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
			var target map[string]any
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			_, hasFilter := target["filter"]
			_, hasPagination := target["pagination"]
			assert.False(t, hasFilter, "unset filter must be omitted")
			assert.False(t, hasPagination, "unset pagination must be omitted")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("search_projects.json"))
		})

		_, err = underTest.SearchProjects(unified_projects.SearchProjectsRequest{})
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

		_, err = underTest.SearchProjects(unified_projects.SearchProjectsRequest{})
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

		_, err = underTest.SearchProjects(unified_projects.SearchProjectsRequest{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing unified project")
	}
}
