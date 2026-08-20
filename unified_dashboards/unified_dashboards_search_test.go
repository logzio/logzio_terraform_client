package unified_dashboards_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/stretchr/testify/assert"
)

func TestUnifiedDashboards_SearchDashboards(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/dashboards/search", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			// Pin the literal wire keys: decoding into the request struct
			// would accept a renamed field without complaint.
			jsonBytes, _ := io.ReadAll(r.Body)
			var target map[string]any
			assert.NoError(t, json.Unmarshal(jsonBytes, &target))
			filter, ok := target["filter"].(map[string]any)
			if assert.True(t, ok, "body should carry a filter object") {
				assert.Equal(t, "System", filter["searchTerm"])
				assert.Equal(t, []any{float64(2045352)}, filter["createdBy"])
			}
			pagination, ok := target["pagination"].(map[string]any)
			if assert.True(t, ok, "body should carry a pagination object") {
				assert.Equal(t, float64(1), pagination["pageNumber"])
				assert.Equal(t, float64(1), pagination["pageSize"])
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("search_dashboards.json"))
		})

		resp, err := underTest.SearchDashboards(unified_dashboards.SearchDashboardsRequest{
			Filter: &unified_dashboards.SearchDashboardsFilter{
				SearchTerm: "System",
				CreatedBy:  []int64{2045352},
			},
			Pagination: &unified_dashboards.SearchDashboardsPagination{PageNumber: 1, PageSize: 1},
		})
		assert.NoError(t, err)
		if assert.NotNil(t, resp) {
			// Total counts every match, not just the page that came back.
			assert.Equal(t, 3, resp.Total)
			assert.Equal(t, &unified_dashboards.SearchDashboardsPagination{PageNumber: 1, PageSize: 1}, resp.Pagination)
			if assert.Len(t, resp.Results, 1) {
				assert.Equal(t, "dashboard-1", resp.Results[0].Uid)
				assert.Equal(t, "row-4c81ee20", resp.Results[0].Id)
				assert.Equal(t, "project-1", resp.Results[0].ProjectId)
				assert.Equal(t, "Dashboard", resp.Results[0].Doc["kind"])
				assert.Equal(t, 2, resp.Results[0].Version)
			}
		}
	}
}

// An empty request is legal — it pages through every dashboard in the account.
func TestUnifiedDashboards_SearchDashboardsEmptyRequest(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/dashboards/search", func(w http.ResponseWriter, r *http.Request) {
			jsonBytes, _ := io.ReadAll(r.Body)
			// omitempty on both fields means the body stays a bare object.
			assert.JSONEq(t, "{}", string(jsonBytes))

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("search_dashboards.json"))
		})

		resp, err := underTest.SearchDashboards(unified_dashboards.SearchDashboardsRequest{})
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	}
}

func TestUnifiedDashboards_SearchDashboardsNoMatches(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/dashboards/search", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"results":[],"total":0,"pagination":{"pageNumber":1,"pageSize":20}}`)
		})

		resp, err := underTest.SearchDashboards(unified_dashboards.SearchDashboardsRequest{
			Filter: &unified_dashboards.SearchDashboardsFilter{SearchTerm: "no-such-dashboard"},
		})
		assert.NoError(t, err)
		if assert.NotNil(t, resp) {
			assert.Empty(t, resp.Results)
			assert.Equal(t, 0, resp.Total)
		}
	}
}

func TestUnifiedDashboards_SearchDashboardsAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/dashboards/search", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.SearchDashboards(unified_dashboards.SearchDashboardsRequest{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status code 500")
	}
}

func TestUnifiedDashboards_SearchDashboardsNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/dashboards/search", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err = underTest.SearchDashboards(unified_dashboards.SearchDashboardsRequest{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing unified dashboard")
	}
}
