package unified_dashboards_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedDashboards_SearchDashboards(t *testing.T) {
	projClient, err := setupUnifiedProjectsIntegrationTest()
	assert.NoError(t, err)
	dashClient, err := setupUnifiedDashboardsIntegrationTest()
	assert.NoError(t, err)

	uniqueId := time.Now().Format("20060102150405")
	projName := "tf-client-it-search-dash-" + uniqueId

	// First create a project
	proj, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: projName})
	if assert.NoError(t, err) && assert.NotNil(t, proj) {
		defer projClient.DeleteProject(proj.Id)

		time.Sleep(2 * time.Second) // Allow for eventual consistency

		// Create a dashboard with unique searchable content
		createReq := unified_dashboards.CreateDashboardRequest{
			Doc: map[string]interface{}{
				"title":       "IT Search Dashboard " + uniqueId,
				"description": "Dashboard for search test with unique id " + uniqueId,
				"panels":      []interface{}{},
			},
		}

		created, err := dashClient.CreateDashboard(proj.Id, createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer dashClient.DeleteDashboard(proj.Id, created.Uid)

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// Search for dashboards using the unique ID
			searchReq := unified_dashboards.SearchDashboardsRequest{
				Query: stringPtr(uniqueId),
				Limit: intPtr(10),
				Page:  intPtr(1),
			}

			results, err := dashClient.SearchDashboards(searchReq)
			if assert.NoError(t, err) {
				found := false
				for _, dashboard := range results {
					if dashboard.Uid == created.Uid {
						found = true
						assert.Equal(t, "IT Search Dashboard "+uniqueId, dashboard.Doc["title"])
						break
					}
				}
				assert.True(t, found, "Created dashboard should appear in search results")
			}
		}
	}
}

// Helper functions for creating pointers
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
