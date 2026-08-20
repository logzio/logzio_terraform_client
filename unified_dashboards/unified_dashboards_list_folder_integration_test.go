package unified_dashboards_test

import (
	"os"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedDashboards_ListFolderDashboards(t *testing.T) {
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

	proj, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: "tf-client-it-folderlist-" + uniqueId})
	if !assert.NoError(t, err) || !assert.NotNil(t, proj) {
		return
	}
	defer func() {
		if err := projClient.DeleteProject(proj.Id); err != nil {
			t.Logf("cleanup: failed to delete project %s: %v", proj.Id, err)
		}
	}()

	time.Sleep(2 * time.Second) // Allow for eventual consistency

	// A folder with no dashboards yet answers 200 with an empty list.
	empty, err := dashClient.ListFolderDashboards(proj.Id)
	assert.NoError(t, err)
	assert.Empty(t, empty, "a fresh folder should list no dashboards")

	created, err := dashClient.CreateDashboard(proj.Id, unified_dashboards.CreateDashboardRequest{
		Doc: itDashboardDoc("it-folderlist-"+uniqueId, "IT Folder List "+uniqueId),
	})
	if !assert.NoError(t, err) || !assert.NotNil(t, created) {
		return
	}
	defer func() {
		if err := dashClient.DeleteDashboard(proj.Id, created.Uid); err != nil {
			t.Logf("cleanup: failed to delete dashboard %s: %v", created.Uid, err)
		}
	}()

	// Update once so the version row forks off the uid; otherwise the two
	// fields are equal and the assertions below could not tell them apart.
	updated, err := dashClient.UpdateDashboard(proj.Id, created.Uid, unified_dashboards.UpdateDashboardRequest{
		Doc: itDashboardDoc("it-folderlist-"+uniqueId, "IT Folder List "+uniqueId+" v2"),
	})
	if !assert.NoError(t, err) || !assert.NotNil(t, updated) {
		return
	}

	time.Sleep(2 * time.Second) // Allow for eventual consistency

	listed, err := dashClient.ListFolderDashboards(proj.Id)
	if assert.NoError(t, err) && assert.Len(t, listed, 1, "the folder should hold exactly the dashboard we created") {
		assert.Equal(t, created.Uid, listed[0].Uid, "Uid is the stable handle")
		assert.Equal(t, updated.Id, listed[0].Id, "Id tracks the current version row")
		assert.NotEqual(t, listed[0].Uid, listed[0].Id)
		assert.Equal(t, proj.Id, listed[0].ProjectId)
		assert.Equal(t, "Dashboard", listed[0].Doc["kind"])
		assert.Equal(t, 2, listed[0].Version)
	}

	// An id that matches no folder is not an error on this route.
	none, err := dashClient.ListFolderDashboards("00000000-0000-0000-0000-000000000000")
	assert.NoError(t, err, "an unknown folder returns 200 with an empty list, not 404")
	assert.Empty(t, none)
}
