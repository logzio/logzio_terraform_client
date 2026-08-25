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
		errDst := dashClient.DeleteDashboard(dst.Id, created.Uid)
		if errDst != nil {
			if errSrc := dashClient.DeleteDashboard(src.Id, created.Uid); errSrc != nil {
				t.Logf("cleanup: failed to delete dashboard %s from either folder (dst: %v; src: %v)", created.Uid, errDst, errSrc)
			}
		}
	}()

	time.Sleep(2 * time.Second) // Allow for eventual consistency

	// Update once before moving. A freshly created dashboard has id == uid ==
	// name, which makes any assertion about the move response ambiguous — it
	// would pass whichever identifier the server actually returned. One update
	// forks the version row off the uid so the assertions below can tell them
	// apart.
	updated, err := dashClient.UpdateDashboard(src.Id, created.Uid, unified_dashboards.UpdateDashboardRequest{
		Doc: itDashboardDoc("it-move-dashboard-"+uniqueId, "IT Move Dashboard "+uniqueId+" v2"),
	})
	if !assert.NoError(t, err) || !assert.NotNil(t, updated) {
		return
	}
	assert.Equal(t, created.Uid, updated.Uid, "the uid must survive an update")
	assert.NotEqual(t, created.Uid, updated.Id, "the version-row id must fork away from the uid on update")

	time.Sleep(2 * time.Second) // Allow for eventual consistency

	moved, err := dashClient.MoveDashboard(unified_dashboards.MoveDashboardRequest{
		DashboardId:  created.Uid,
		OldProjectId: src.Id,
		NewProjectId: dst.Id,
	})
	if assert.NoError(t, err) && assert.NotNil(t, moved) {
		// The move response's "id" is the current version-row id, not the uid.
		assert.Equal(t, updated.Id, moved.Id, "move should acknowledge with the current version-row id")
		assert.NotEqual(t, created.Uid, moved.Id, "move must not be read as returning the uid")

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
