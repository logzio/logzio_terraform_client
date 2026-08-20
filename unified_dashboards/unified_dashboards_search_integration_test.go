package unified_dashboards_test

import (
	"os"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedDashboards_SearchDashboards(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	projClient, err := setupUnifiedProjectsIntegrationTest()
	if !assert.NoError(t, err) || projClient == nil {
		return
	}
	dashClient, err := setupUnifiedDashboardsIntegrationTest()
	if !assert.NoError(t, err) || dashClient == nil {
		return
	}

	uniqueId := time.Now().Format("20060102150405")

	proj, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: "tf-client-it-dashsearch-" + uniqueId})
	if !assert.NoError(t, err) || !assert.NotNil(t, proj) {
		return
	}
	defer func() {
		if err := projClient.DeleteProject(proj.Id); err != nil {
			t.Logf("cleanup: failed to delete project %s: %v", proj.Id, err)
		}
	}()

	time.Sleep(2 * time.Second) // Allow for eventual consistency

	// metadata.name and the display name are deliberately different strings:
	// SearchTerm matches the former only, and the assertions below rely on
	// being able to tell them apart.
	dashName := "it-dashsearch-" + uniqueId
	displayName := "zzdisplayonly-" + uniqueId
	created, err := dashClient.CreateDashboard(proj.Id, unified_dashboards.CreateDashboardRequest{
		Doc: itDashboardDoc(dashName, displayName),
	})
	if !assert.NoError(t, err) || !assert.NotNil(t, created) {
		return
	}
	defer func() {
		if err := dashClient.DeleteDashboard(proj.Id, created.Uid); err != nil {
			t.Logf("cleanup: failed to delete dashboard %s: %v", created.Uid, err)
		}
	}()

	time.Sleep(3 * time.Second) // Allow for eventual consistency

	// Unlike the projects search, SearchTerm really does narrow the result set.
	resp, err := dashClient.SearchDashboards(unified_dashboards.SearchDashboardsRequest{
		Filter: &unified_dashboards.SearchDashboardsFilter{SearchTerm: dashName},
	})
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		if assert.NotNil(t, resp.Pagination, "server should echo a pagination block") {
			assert.Equal(t, 1, resp.Pagination.PageNumber)
			assert.Positive(t, resp.Pagination.PageSize)
		}
		found := false
		for _, d := range resp.Results {
			if d.Uid == created.Uid {
				found = true
				assert.Equal(t, proj.Id, d.ProjectId)
				assert.Equal(t, "Dashboard", d.Doc["kind"])
				break
			}
		}
		assert.True(t, found, "the created dashboard should match its own name")
		assert.Positive(t, resp.Total)
	}

	// A term that matches nothing comes back empty rather than unfiltered —
	// this is what proves the filter is applied server-side at all.
	none, err := dashClient.SearchDashboards(unified_dashboards.SearchDashboardsRequest{
		Filter: &unified_dashboards.SearchDashboardsFilter{SearchTerm: "no-such-dashboard-" + uniqueId},
	})
	if assert.NoError(t, err) && assert.NotNil(t, none) {
		assert.Empty(t, none.Results)
		assert.Equal(t, 0, none.Total)
	}

	// SearchTerm reads metadata.name only. The display name is unique to this
	// dashboard and still matches nothing, which is what stops the endpoint
	// being mistaken for a full-text search over the document.
	byDisplay, err := dashClient.SearchDashboards(unified_dashboards.SearchDashboardsRequest{
		Filter: &unified_dashboards.SearchDashboardsFilter{SearchTerm: displayName},
	})
	if assert.NoError(t, err) && assert.NotNil(t, byDisplay) {
		assert.Empty(t, byDisplay.Results, "searchTerm must not match the display name")
	}

	// Pagination is honoured, and Total stays the unpaginated match count.
	paged, err := dashClient.SearchDashboards(unified_dashboards.SearchDashboardsRequest{
		Pagination: &unified_dashboards.SearchDashboardsPagination{PageNumber: 1, PageSize: 1},
	})
	if assert.NoError(t, err) && assert.NotNil(t, paged) {
		assert.LessOrEqual(t, len(paged.Results), 1, "pageSize=1 should return at most one dashboard")
		if assert.NotNil(t, paged.Pagination) {
			assert.Equal(t, 1, paged.Pagination.PageSize)
		}
		assert.GreaterOrEqual(t, paged.Total, len(paged.Results))
	}
}
