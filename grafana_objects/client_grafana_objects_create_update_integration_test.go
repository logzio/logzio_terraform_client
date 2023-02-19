package grafana_objects_test

import (
	"github.com/logzio/logzio_terraform_client/grafana_objects"
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
			createDashboard.Dashboard.Title += "_create"
			createDashboard.Dashboard.Uid += "_create"
			dashboard, err := underTest.CreateUpdate(createDashboard)
			if assert.NoError(t, err) && assert.NotNil(t, dashboard) {
				time.Sleep(2 * time.Second)
				assert.NotEmpty(t, dashboard.Uid)
				defer underTest.Delete(dashboard.Uid)
				assert.Equal(t, grafana_objects.GrafanaSuccessStatus, dashboard.Status)
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
			createDashboard.Dashboard.Title += "_before_update"
			createDashboard.Dashboard.Uid += "_update"
			dashboard, err := underTest.CreateUpdate(createDashboard)
			if assert.NoError(t, err) && assert.NotNil(t, dashboard) {
				time.Sleep(2 * time.Second)
				assert.NotEmpty(t, dashboard.Uid)
				defer underTest.Delete(dashboard.Uid)
				assert.Equal(t, grafana_objects.GrafanaSuccessStatus, dashboard.Status)
				createDashboard.Dashboard.Title = "after_update"
				createDashboard.Dashboard.Id = dashboard.Id
				createDashboard.Dashboard.Uid = dashboard.Uid
				updated, err := underTest.CreateUpdate(createDashboard)
				if assert.NoError(t, err) && assert.NotNil(t, updated) {
					assert.Equal(t, grafana_objects.GrafanaSuccessStatus, updated.Status)
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
			createDashboard.Dashboard.Title += "_create_invalid"
			createDashboard.Dashboard.Uid += "_create_invalid"
			createDashboard.Dashboard.Panels = nil
			dashboard, err := underTest.CreateUpdate(createDashboard)
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
			createDashboard.Dashboard.Title += "_update_not_found"
			createDashboard.Dashboard.Uid += "_update_not_found"
			createDashboard.Dashboard.Id = 1
			createDashboard.Overwrite = true
			dashboard, err := underTest.CreateUpdate(createDashboard)
			assert.Error(t, err)
			assert.Nil(t, dashboard)
			assert.Contains(t, err.Error(), "failed with missing grafana dashboard")
		}
	}
}
