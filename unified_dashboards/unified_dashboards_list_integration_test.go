package unified_dashboards_test

import (
	"os"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedDashboards_ListDashboards(t *testing.T) {
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
	projName := "tf-client-it-list-dash-" + uniqueId

	proj, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: projName})
	if assert.NoError(t, err) && assert.NotNil(t, proj) {
		defer func() {
			if err := projClient.DeleteProject(proj.Id); err != nil {
				t.Logf("cleanup: failed to delete project %s: %v", proj.Id, err)
			}
		}()

		time.Sleep(2 * time.Second) // Allow for eventual consistency

		created, err := dashClient.CreateDashboard(proj.Id, unified_dashboards.CreateDashboardRequest{
			Doc: itDashboardDoc("it-list-dashboard-"+uniqueId, "IT List Dashboard "+uniqueId),
		})
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer func() {
				if err := dashClient.DeleteDashboard(proj.Id, created.Uid); err != nil {
					t.Logf("cleanup: failed to delete dashboard %s: %v", created.Uid, err)
				}
			}()

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			dashboards, err := dashClient.ListDashboards()
			if assert.NoError(t, err) {
				found := false
				for _, dashboard := range dashboards {
					if dashboard.Uid == created.Uid {
						found = true
						assert.Equal(t, proj.Id, dashboard.ProjectId)
						assert.Equal(t, "Dashboard", dashboard.Doc["kind"])
						break
					}
				}
				assert.True(t, found, "Created dashboard should appear in list")
			}
		}
	}
}
