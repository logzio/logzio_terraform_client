package unified_dashboards_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedDashboards_UpdateDashboard(t *testing.T) {
	projClient, err := setupUnifiedProjectsIntegrationTest()
	assert.NoError(t, err)
	dashClient, err := setupUnifiedDashboardsIntegrationTest()
	assert.NoError(t, err)

	uniqueId := time.Now().Format("20060102150405")
	projName := "tf-client-it-upd-dash-" + uniqueId

	// First create a project
	proj, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: projName})
	if assert.NoError(t, err) && assert.NotNil(t, proj) {
		defer projClient.DeleteProject(proj.Id)

		time.Sleep(2 * time.Second) // Allow for eventual consistency

		// Create a dashboard
		createReq := unified_dashboards.CreateDashboardRequest{
			Doc: map[string]interface{}{
				"title":       "IT Update Dashboard " + uniqueId,
				"description": "Original description",
				"panels":      []interface{}{},
			},
		}

		created, err := dashClient.CreateDashboard(proj.Id, createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer dashClient.DeleteDashboard(proj.Id, created.Uid)

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// Update the dashboard
			updateReq := unified_dashboards.UpdateDashboardRequest{
				Doc: map[string]interface{}{
					"title":       "IT Updated Dashboard " + uniqueId,
					"description": "Updated description for integration test",
					"panels":      []interface{}{},
					"refresh":     "30s",
				},
			}

			updated, err := dashClient.UpdateDashboard(proj.Id, created.Uid, updateReq)
			if assert.NoError(t, err) && assert.NotNil(t, updated) {
				assert.Equal(t, created.Uid, updated.Uid)
				assert.NotNil(t, updated.Doc)
				assert.Equal(t, "IT Updated Dashboard "+uniqueId, updated.Doc["title"])
				assert.Equal(t, "Updated description for integration test", updated.Doc["description"])
				assert.Equal(t, "30s", updated.Doc["refresh"])
				assert.NotEmpty(t, updated.UpdatedAt)
			}
		}
	}
}
