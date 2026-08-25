package unified_dashboards_test

import (
	"os"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedDashboards_DeleteDashboard(t *testing.T) {
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
	projName := "tf-client-it-del-dash-" + uniqueId

	proj, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: projName})
	if assert.NoError(t, err) && assert.NotNil(t, proj) {
		defer func() {
			if err := projClient.DeleteProject(proj.Id); err != nil {
				t.Logf("cleanup: failed to delete project %s: %v", proj.Id, err)
			}
		}()

		time.Sleep(2 * time.Second) // Allow for eventual consistency

		created, err := dashClient.CreateDashboard(proj.Id, unified_dashboards.CreateDashboardRequest{
			Doc: itDashboardDoc("it-del-dashboard-"+uniqueId, "IT Delete Dashboard "+uniqueId),
		})
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			time.Sleep(2 * time.Second) // Allow for eventual consistency

			err = dashClient.DeleteDashboard(proj.Id, created.Uid)
			assert.NoError(t, err)

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// Verify the dashboard no longer exists; the error must be the
			// not-found classification, not some other failure.
			_, err = dashClient.GetDashboard(proj.Id, created.Uid)
			if assert.Error(t, err, "Getting deleted dashboard should return an error") {
				assert.Contains(t, err.Error(), "failed with missing unified dashboard")
			}
		}
	}
}
