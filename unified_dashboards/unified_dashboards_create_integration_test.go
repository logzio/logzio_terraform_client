package unified_dashboards_test

import (
	"os"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

// itDashboardDoc builds a minimal valid Perses Dashboard document — the API
// requires the full envelope (kind/metadata/spec).
func itDashboardDoc(name, displayName string) map[string]interface{} {
	return map[string]interface{}{
		"kind":     "Dashboard",
		"metadata": map[string]interface{}{"name": name},
		"spec": map[string]interface{}{
			"display":  map[string]interface{}{"name": displayName},
			"duration": "1h",
			"panels":   map[string]interface{}{},
			"layouts":  []interface{}{},
		},
	}
}

func TestIntegrationUnifiedDashboards_CreateDashboard(t *testing.T) {
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
	projName := "tf-client-it-dash-" + uniqueId

	proj, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: projName})
	if assert.NoError(t, err) && assert.NotNil(t, proj) {
		defer func() {
			if err := projClient.DeleteProject(proj.Id); err != nil {
				t.Logf("cleanup: failed to delete project %s: %v", proj.Id, err)
			}
		}()

		time.Sleep(2 * time.Second) // Allow for eventual consistency

		created, err := dashClient.CreateDashboard(proj.Id, unified_dashboards.CreateDashboardRequest{
			Doc: itDashboardDoc("it-dashboard-"+uniqueId, "IT Dashboard "+uniqueId),
		})
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer func() {
				if err := dashClient.DeleteDashboard(proj.Id, created.Uid); err != nil {
					t.Logf("cleanup: failed to delete dashboard %s: %v", created.Uid, err)
				}
			}()

			assert.NotEmpty(t, created.Uid)
			assert.Equal(t, proj.Id, created.ProjectId)
			assert.Equal(t, "Dashboard", created.Doc["kind"])
			assert.NotEmpty(t, created.CreatedAt)

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// Live-verify the embedded-dashboards mapping of the projects list.
			items, err := projClient.ListProjects(true)
			if assert.NoError(t, err) {
				found := false
				for _, item := range items {
					if item.Project.Id == proj.Id {
						found = true
						if assert.Len(t, item.Dashboards, 1, "created dashboard should be embedded in its project's list entry") {
							assert.Equal(t, proj.Id, item.Dashboards[0].ProjectId)
							assert.Equal(t, created.Uid, item.Dashboards[0].Id, "embedded dashboard Id should be the addressable uid")
						}
						break
					}
				}
				assert.True(t, found, "created project should appear in the withDashboards list")
			}
		}
	}
}
