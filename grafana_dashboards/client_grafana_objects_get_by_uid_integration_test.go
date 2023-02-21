package grafana_dashboards_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationGrafanaObjects_GetByUid(t *testing.T) {
	underTest, err := setupGrafanaObjectsIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createDashboard, err := getCreateDashboardIntegrationTests()
		if assert.NoError(t, err) {
			createDashboard.Dashboard["title"] = fmt.Sprintf("%s%s", createDashboard.Dashboard["title"], "_get")
			createDashboard.Dashboard["uid"] = fmt.Sprintf("%s%s", createDashboard.Dashboard["uid"], "_get")
			dashboard, err := underTest.CreateUpdateGrafanaDashboard(createDashboard)
			if assert.NoError(t, err) && assert.NotNil(t, dashboard) {
				time.Sleep(2 * time.Second)
				assert.NotEmpty(t, dashboard.Uid)
				defer underTest.DeleteGrafanaDashboard(dashboard.Uid)
				getDashboard, err := underTest.GetGrafanaDashboard(dashboard.Uid)
				assert.NoError(t, err)
				assert.NotNil(t, getDashboard)
				assert.Equal(t, dashboard.Uid, getDashboard.Dashboard["uid"].(string))
				assert.Equal(t, createDashboard.FolderId, getDashboard.Meta["folderId"].(float64))
				assert.Equal(t, createDashboard.Dashboard["title"].(string), getDashboard.Dashboard["title"].(string))
			}
		}
	}
}

func TestIntegrationGrafanaObjects_GetByUidNotFound(t *testing.T) {
	underTest, err := setupGrafanaObjectsIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		get, err := underTest.GetGrafanaDashboard("not exist")
		assert.Error(t, err)
		assert.Nil(t, get)
		assert.Contains(t, err.Error(), "failed with missing grafana dashboard")
	}
}
