package unified_dashboards_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedDashboards_CreateDashboard(t *testing.T) {
	projClient, err := setupUnifiedProjectsIntegrationTest()
	assert.NoError(t, err)
	dashClient, err := setupUnifiedDashboardsIntegrationTest()
	assert.NoError(t, err)

	uniqueId := time.Now().Format("20060102150405")
	projName := "tf-client-it-dash-" + uniqueId

	// First create a project
	proj, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: projName})
	if assert.NoError(t, err) && assert.NotNil(t, proj) {
		defer projClient.DeleteProject(proj.Id)

		time.Sleep(2 * time.Second) // Allow for eventual consistency

		// Create a dashboard
		createReq := unified_dashboards.CreateDashboardRequest{
			Doc: map[string]interface{}{
				"title":       "IT Dashboard " + uniqueId,
				"description": "Integration test dashboard",
				"panels":      []interface{}{},
				"refresh":     "1m",
			},
		}

		created, err := dashClient.CreateDashboard(proj.Id, createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer dashClient.DeleteDashboard(proj.Id, created.Uid)

			assert.NotEmpty(t, created.Uid)
			assert.NotNil(t, created.Doc)
			assert.Equal(t, "IT Dashboard "+uniqueId, created.Doc["title"])
			assert.Equal(t, "Integration test dashboard", created.Doc["description"])
			assert.NotNil(t, created.CreatedAt)
		}
	}
}

// Helper function to setup unified projects integration test
func setupUnifiedProjectsIntegrationTest() (*unified_projects.ProjectsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := unified_projects.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}
