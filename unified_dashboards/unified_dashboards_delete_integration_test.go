package unified_dashboards_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedDashboards_DeleteDashboard(t *testing.T) {
	projClient, err := setupUnifiedProjectsIntegrationTest()
	if !assert.NoError(t, err) || projClient == nil {
		return
	}
	dashClient, err := setupUnifiedDashboardsIntegrationTest()
	if !assert.NoError(t, err) || dashClient == nil {
		return
	}

	uniqueId := time.Now().Format("20060102150405")
	projName := "tf-client-it-del-dash-" + uniqueId

	// First create a project
	proj, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: projName})
	if assert.NoError(t, err) && assert.NotNil(t, proj) {
		defer projClient.DeleteProject(proj.Id)

		time.Sleep(10 * time.Second) // Allow for eventual consistency

		// Create a dashboard
		createReq := unified_dashboards.CreateDashboardRequest{
			Doc: map[string]interface{}{
				"title":       "IT Delete Dashboard " + uniqueId,
				"description": "Dashboard to be deleted",
				"panels":      []interface{}{},
			},
		}

		created, err := dashClient.CreateDashboard(proj.Id, createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// Delete the dashboard
			err = dashClient.DeleteDashboard(proj.Id, created.Uid)
			assert.NoError(t, err)

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// Verify the dashboard no longer exists by trying to get it
			_, err = dashClient.GetDashboard(proj.Id, created.Uid, nil)
			assert.Error(t, err, "Getting deleted dashboard should return an error")
		}
	}
}
