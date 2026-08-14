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
	if !assert.NoError(t, err) {
		return
	}
	dashClient, err := setupUnifiedDashboardsIntegrationTest()
	if !assert.NoError(t, err) {
		return
	}

	uniqueId := time.Now().Format("20060102150405")

	src, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: "tf-client-it-move-src-" + uniqueId})
	if !assert.NoError(t, err) || !assert.NotNil(t, src) {
		return
	}
	defer projClient.DeleteProject(src.Id)

	dst, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: "tf-client-it-move-dst-" + uniqueId})
	if !assert.NoError(t, err) || !assert.NotNil(t, dst) {
		return
	}
	defer projClient.DeleteProject(dst.Id)

	time.Sleep(2 * time.Second)

	created, err := dashClient.CreateDashboard(src.Id, unified_dashboards.CreateDashboardRequest{
		Doc: map[string]interface{}{"title": "IT Move Dashboard " + uniqueId, "panels": []interface{}{}},
	})
	if !assert.NoError(t, err) || !assert.NotNil(t, created) {
		return
	}
	defer dashClient.DeleteDashboard(dst.Id, created.Uid) // it ends up in dst

	time.Sleep(2 * time.Second)

	moved, err := dashClient.MoveDashboard(unified_dashboards.MoveDashboardRequest{
		Uid:            created.Uid,
		TargetFolderId: dst.Id,
	})
	if assert.NoError(t, err) && assert.NotNil(t, moved) {
		assert.Equal(t, created.Uid, moved.Uid)

		time.Sleep(2 * time.Second)
		retrieved, err := dashClient.GetDashboard(dst.Id, created.Uid)
		assert.NoError(t, err)
		assert.NotNil(t, retrieved)
	}
}
