package unified_dashboards_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedDashboards_ListDashboards(t *testing.T) {
	projClient, err := setupUnifiedProjectsIntegrationTest()
	assert.NoError(t, err)
	dashClient, err := setupUnifiedDashboardsIntegrationTest()
	assert.NoError(t, err)

	uniqueId := time.Now().Format("20060102150405")
	projName := "tf-client-it-list-dash-" + uniqueId

	// First create a project
	proj, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: projName})
	if assert.NoError(t, err) && assert.NotNil(t, proj) {
		defer projClient.DeleteProject(proj.Id)

		time.Sleep(2 * time.Second) // Allow for eventual consistency

		// Create a dashboard
		createReq := unified_dashboards.CreateDashboardRequest{
			Doc: map[string]interface{}{
				"title":       "IT List Dashboard " + uniqueId,
				"description": "Dashboard for list test",
				"panels":      []interface{}{},
			},
		}

		created, err := dashClient.CreateDashboard(proj.Id, createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer dashClient.DeleteDashboard(proj.Id, created.Uid)

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// List dashboards and verify our dashboard appears
			dashboards, err := dashClient.ListDashboards()
			if assert.NoError(t, err) {
				found := false
				for _, dashboard := range dashboards {
					if dashboard.Uid == created.Uid {
						found = true
						assert.Equal(t, "IT List Dashboard "+uniqueId, dashboard.Doc["title"])
						break
					}
				}
				assert.True(t, found, "Created dashboard should appear in list")
			}
		}
	}
}
