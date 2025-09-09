package unified_dashboards_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedDashboards_GetDashboard(t *testing.T) {
	projClient, err := setupUnifiedProjectsIntegrationTest()
	assert.NoError(t, err)
	dashClient, err := setupUnifiedDashboardsIntegrationTest()
	assert.NoError(t, err)

	uniqueId := time.Now().Format("20060102150405")
	projName := "tf-client-it-get-dash-" + uniqueId

	// First create a project
	proj, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: projName})
	if assert.NoError(t, err) && assert.NotNil(t, proj) {
		defer projClient.DeleteProject(proj.Id)

		time.Sleep(2 * time.Second) // Allow for eventual consistency

		// Create a dashboard
		createReq := unified_dashboards.CreateDashboardRequest{
			Doc: map[string]interface{}{
				"title":       "IT Get Dashboard " + uniqueId,
				"description": "Integration test get dashboard",
				"panels":      []interface{}{},
			},
		}

		created, err := dashClient.CreateDashboard(proj.Id, createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer dashClient.DeleteDashboard(proj.Id, created.Uid)

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// Get the dashboard
			retrieved, err := dashClient.GetDashboard(proj.Id, created.Uid, nil)
			if assert.NoError(t, err) && assert.NotNil(t, retrieved) {
				assert.Equal(t, created.Uid, retrieved.Uid)
				assert.NotNil(t, retrieved.Doc)
				assert.Equal(t, "IT Get Dashboard "+uniqueId, retrieved.Doc["title"])
				assert.Equal(t, "Integration test get dashboard", retrieved.Doc["description"])
			}

			// Test with source parameter
			source := "grafana"
			retrievedWithSource, err := dashClient.GetDashboard(proj.Id, created.Uid, &source)
			if assert.NoError(t, err) && assert.NotNil(t, retrievedWithSource) {
				assert.Equal(t, created.Uid, retrievedWithSource.Uid)
			}
		}
	}
}
