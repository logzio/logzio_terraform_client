package grafana_dashboards_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationGrafanaObjects_DeleteByUid(t *testing.T) {
	underTest, err := setupGrafanaObjectsIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createDashboard, err := getCreateDashboardIntegrationTests()
		if assert.NoError(t, err) {
			createDashboard.Dashboard["title"] = fmt.Sprintf("%s%s", createDashboard.Dashboard["title"], "_delete")
			createDashboard.Dashboard["title"] = fmt.Sprintf("%s%s", createDashboard.Dashboard["title"], "_delete")
			dashboard, err := underTest.CreateUpdateGrafanaDashboard(createDashboard)
			if assert.NoError(t, err) && assert.NotNil(t, dashboard) {
				time.Sleep(2 * time.Second)
				assert.NotEmpty(t, dashboard.Uid)
				defer func() {
					deleteRes, err := underTest.DeleteGrafanaDashboard(dashboard.Uid)
					assert.NoError(t, err)
					assert.NotNil(t, deleteRes)
				}()
			}
		}
	}
}

func TestIntegrationGrafanaObjects_DeleteByUidNotFound(t *testing.T) {
	underTest, err := setupGrafanaObjectsIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		deleteRes, err := underTest.DeleteGrafanaDashboard("not exist")
		assert.Error(t, err)
		assert.Nil(t, deleteRes)
		assert.Contains(t, err.Error(), "failed with missing grafana dashboard")
	}
}
