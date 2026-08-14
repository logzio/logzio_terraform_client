package unified_dashboards_test

import (
	"os"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedDashboards_MoveDashboard(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	// POST /perses-public/api/v1/dashboards/move is documented at
	// api-docs.logz.io but was not deployed on the api.logz.io gateway as of
	// 2026-08-14 (the route 404s). Delete the skip below once the endpoint is
	// live; the scenario is kept ready. MoveDashboard's client-side behavior
	// is covered by unit tests.
	t.Skip("dashboards/move endpoint not deployed on the public gateway (verified 2026-08-14)")

	projClient, err := setupUnifiedProjectsIntegrationTest()
	if !assert.NoError(t, err) || projClient == nil {
		return
	}
	dashClient, err := setupUnifiedDashboardsIntegrationTest()
	if !assert.NoError(t, err) || dashClient == nil {
		return
	}

	uniqueId := time.Now().Format("20060102150405")

	src, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: "tf-client-it-move-src-" + uniqueId})
	if !assert.NoError(t, err) || !assert.NotNil(t, src) {
		return
	}
	defer func() {
		if err := projClient.DeleteProject(src.Id); err != nil {
			t.Logf("cleanup: failed to delete project %s: %v", src.Id, err)
		}
	}()

	dst, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: "tf-client-it-move-dst-" + uniqueId})
	if !assert.NoError(t, err) || !assert.NotNil(t, dst) {
		return
	}
	defer func() {
		if err := projClient.DeleteProject(dst.Id); err != nil {
			t.Logf("cleanup: failed to delete project %s: %v", dst.Id, err)
		}
	}()

	time.Sleep(2 * time.Second) // Allow for eventual consistency

	created, err := dashClient.CreateDashboard(src.Id, unified_dashboards.CreateDashboardRequest{
		Doc: itDashboardDoc("it-move-dashboard-"+uniqueId, "IT Move Dashboard "+uniqueId),
	})
	if !assert.NoError(t, err) || !assert.NotNil(t, created) {
		return
	}
	defer func() {
		// Best-effort cleanup from both folders — the dashboard's location
		// depends on whether the move succeeded.
		if err := dashClient.DeleteDashboard(dst.Id, created.Uid); err != nil {
			if err := dashClient.DeleteDashboard(src.Id, created.Uid); err != nil {
				t.Logf("cleanup: failed to delete dashboard %s from either folder: %v", created.Uid, err)
			}
		}
	}()

	time.Sleep(2 * time.Second) // Allow for eventual consistency

	moved, err := dashClient.MoveDashboard(unified_dashboards.MoveDashboardRequest{
		Uid:            created.Uid,
		TargetFolderId: dst.Id,
	})
	if assert.NoError(t, err) && assert.NotNil(t, moved) {
		assert.Equal(t, created.Uid, moved.Uid)

		time.Sleep(2 * time.Second) // Allow for eventual consistency

		// Present in the destination…
		retrieved, err := dashClient.GetDashboard(dst.Id, created.Uid)
		assert.NoError(t, err)
		assert.NotNil(t, retrieved)

		// …and gone from the source (a copy would leave it behind).
		_, err = dashClient.GetDashboard(src.Id, created.Uid)
		if assert.Error(t, err, "moved dashboard should be gone from the source folder") {
			assert.Contains(t, err.Error(), "failed with missing unified dashboard")
		}
	}
}
