package grafana_dashboards_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/grafana_dashboards"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationGrafanaObjects_CreateDashboard(t *testing.T) {
	underTest, err := setupGrafanaObjectsIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createDashboard, err := getCreateDashboardIntegrationTests()
		if assert.NoError(t, err) {
			createDashboard.Dashboard["title"] = fmt.Sprintf("%s%s", createDashboard.Dashboard["title"], "_create")
			createDashboard.Dashboard["uid"] = fmt.Sprintf("%s%s", createDashboard.Dashboard["uid"], "_create")
			dashboard, err := underTest.CreateUpdateGrafanaDashboard(createDashboard)
			if assert.NoError(t, err) && assert.NotNil(t, dashboard) {
				time.Sleep(2 * time.Second)
				assert.NotEmpty(t, dashboard.Uid)
				defer underTest.DeleteGrafanaDashboard(dashboard.Uid)
				assert.Equal(t, grafana_dashboards.GrafanaSuccessStatus, dashboard.Status)
			}
		}
	}
}

func TestIntegrationGrafanaObjects_UpdateDashboard(t *testing.T) {
	underTest, err := setupGrafanaObjectsIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createDashboard, err := getCreateDashboardIntegrationTests()
		if assert.NoError(t, err) {
			createDashboard.Dashboard["title"] = fmt.Sprintf("%s%s", createDashboard.Dashboard["title"], "_before_update")
			createDashboard.Dashboard["uid"] = fmt.Sprintf("%s%s", createDashboard.Dashboard["uid"], "_update")
			dashboard, err := underTest.CreateUpdateGrafanaDashboard(createDashboard)
			if assert.NoError(t, err) && assert.NotNil(t, dashboard) {
				time.Sleep(2 * time.Second)
				assert.NotEmpty(t, dashboard.Uid)
				defer underTest.DeleteGrafanaDashboard(dashboard.Uid)
				assert.Equal(t, grafana_dashboards.GrafanaSuccessStatus, dashboard.Status)
				createDashboard.Dashboard["title"] = fmt.Sprintf("%s%s", createDashboard.Dashboard["title"], "after_update")
				createDashboard.Dashboard["uid"] = dashboard.Uid
				createDashboard.Dashboard["id"] = dashboard.Id
				updated, err := underTest.CreateUpdateGrafanaDashboard(createDashboard)
				if assert.NoError(t, err) && assert.NotNil(t, updated) {
					assert.Equal(t, grafana_dashboards.GrafanaSuccessStatus, updated.Status)
				}
			}
		}
	}
}

func TestIntegrationGrafanaObjects_CreateUpdateDashboardInvalidPayload(t *testing.T) {
	underTest, err := setupGrafanaObjectsIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createDashboard, err := getCreateDashboardIntegrationTests()
		if assert.NoError(t, err) {
			createDashboard.Dashboard["title"] = fmt.Sprintf("%s%s", createDashboard.Dashboard["title"], "_create_invalid")
			createDashboard.Dashboard["uid"] = fmt.Sprintf("%s%s", createDashboard.Dashboard["uid"], "_create_invalid")
			createDashboard.Dashboard["panels"] = nil
			dashboard, err := underTest.CreateUpdateGrafanaDashboard(createDashboard)
			assert.Error(t, err)
			assert.Nil(t, dashboard)
		}
	}
}

func TestIntegrationGrafanaObjects_UpdateDashboardNotFound(t *testing.T) {
	underTest, err := setupGrafanaObjectsIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createDashboard, err := getCreateDashboardIntegrationTests()
		if assert.NoError(t, err) {
			createDashboard.Dashboard["title"] = fmt.Sprintf("%s%s", createDashboard.Dashboard["title"], "_update_not_found")
			createDashboard.Dashboard["uid"] = fmt.Sprintf("%s%s", createDashboard.Dashboard["uid"], "_update_not_found")
			createDashboard.Dashboard["id"] = 1
			createDashboard.Overwrite = true
			dashboard, err := underTest.CreateUpdateGrafanaDashboard(createDashboard)
			assert.Error(t, err)
			assert.Nil(t, dashboard)
			assert.Contains(t, err.Error(), "failed with missing grafana dashboard")
		}
	}
}
