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
func itDashboardDoc(name, displayName string) map[string]any {
	return map[string]any{
		"kind":     "Dashboard",
		"metadata": map[string]any{"name": name},
		"spec": map[string]any{
			"display":  map[string]any{"name": displayName},
			"duration": "1h",
			"panels":   map[string]any{},
			"layouts":  []any{},
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

			// Update once before verifying identifiers. On a brand-new
			// dashboard the server sets id == uid == name, so an assertion
			// made here would hold for either field and prove nothing about
			// which one the embedded payload actually carries.
			updated, err := dashClient.UpdateDashboard(proj.Id, created.Uid, unified_dashboards.UpdateDashboardRequest{
				Doc: itDashboardDoc("it-dashboard-"+uniqueId, "IT Dashboard "+uniqueId+" v2"),
			})
			if !assert.NoError(t, err) || !assert.NotNil(t, updated) {
				return
			}
			assert.Equal(t, created.Uid, updated.Uid, "the uid must survive an update")
			assert.NotEqual(t, updated.Uid, updated.Id, "the version-row id must fork away from the uid on update")

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// Live-verify the embedded-dashboards mapping of the projects list.
			items, err := projClient.ListProjects(true)
			if assert.NoError(t, err) {
				found := false
				for _, item := range items {
					if item.Project.Id == proj.Id {
						found = true
						if assert.Len(t, item.Dashboards, 1, "created dashboard should be embedded in its project's list entry") {
							embedded := item.Dashboards[0]
							assert.Equal(t, proj.Id, embedded.ProjectId)
							assert.Equal(t, created.Uid, embedded.Uid, "embedded Uid is the addressable handle")
							assert.Equal(t, updated.Id, embedded.Id, "embedded Id is the version-row id")
							assert.NotEqual(t, embedded.Uid, embedded.Id, "the two must not be conflated")

							// The distinction is not cosmetic: only the uid resolves.
							byUid, err := dashClient.GetDashboard(proj.Id, embedded.Uid)
							assert.NoError(t, err, "embedded Uid must address the dashboard")
							assert.NotNil(t, byUid)

							_, err = dashClient.GetDashboard(proj.Id, embedded.Id)
							if assert.Error(t, err, "embedded Id must not address the dashboard") {
								assert.Contains(t, err.Error(), "failed with missing unified dashboard")
							}
						}
						break
					}
				}
				assert.True(t, found, "created project should appear in the withDashboards list")
			}
		}
	}
}
