package unified_dashboards_test

import (
	"os"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedDashboards_UpdateDashboard(t *testing.T) {
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
	projName := "tf-client-it-upd-dash-" + uniqueId

	proj, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: projName})
	if assert.NoError(t, err) && assert.NotNil(t, proj) {
		defer func() {
			if err := projClient.DeleteProject(proj.Id); err != nil {
				t.Logf("cleanup: failed to delete project %s: %v", proj.Id, err)
			}
		}()

		time.Sleep(2 * time.Second) // Allow for eventual consistency

		dashName := "it-upd-dashboard-" + uniqueId
		created, err := dashClient.CreateDashboard(proj.Id, unified_dashboards.CreateDashboardRequest{
			Doc: itDashboardDoc(dashName, "IT Update Dashboard "+uniqueId),
		})
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer func() {
				if err := dashClient.DeleteDashboard(proj.Id, created.Uid); err != nil {
					t.Logf("cleanup: failed to delete dashboard %s: %v", created.Uid, err)
				}
			}()

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			updated, err := dashClient.UpdateDashboard(proj.Id, created.Uid, unified_dashboards.UpdateDashboardRequest{
				Doc: itDashboardDoc(dashName, "IT Updated Dashboard "+uniqueId),
			})
			if assert.NoError(t, err) && assert.NotNil(t, updated) {
				assert.Equal(t, created.Uid, updated.Uid)
				assert.Greater(t, updated.Version, created.Version)

				spec := updated.Doc["spec"].(map[string]interface{})
				display := spec["display"].(map[string]interface{})
				assert.Equal(t, "IT Updated Dashboard "+uniqueId, display["name"])
			}
		}
	}
}
