package grafana_alerts_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationGrafanaAlert_GetGrafanaAlert(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaAlert := getGrafanaAlertRuleObject()
		grafanaAlert, err := underTest.CreateGrafanaAlertRule(createGrafanaAlert)
		if assert.NoError(t, err) && assert.NotNil(t, grafanaAlert) && assert.NotEmpty(t, grafanaAlert.Uid) {
			defer underTest.DeleteGrafanaAlertRule(grafanaAlert.Uid)
			time.Sleep(4 * time.Second)
			getAlert, err := underTest.GetGrafanaAlertRule(grafanaAlert.Uid)
			assert.NoError(t, err)
			assert.NotNil(t, getAlert)
			assert.Equal(t, grafanaAlert.Annotations, getAlert.Annotations)
			assert.Equal(t, grafanaAlert.Uid, getAlert.Uid)
			assert.Equal(t, grafanaAlert.Id, getAlert.Id)
			assert.Equal(t, grafanaAlert.Title, getAlert.Title)
			assert.Equal(t, grafanaAlert.RuleGroup, getAlert.RuleGroup)
			assert.Equal(t, grafanaAlert.FolderUID, getAlert.FolderUID)
			assert.Equal(t, grafanaAlert.Data, getAlert.Data)
			assert.Equal(t, grafanaAlert.OrgID, getAlert.OrgID)
			assert.Equal(t, grafanaAlert.Condition, getAlert.Condition)
			assert.Equal(t, grafanaAlert.ExecErrState, getAlert.ExecErrState)
			assert.Equal(t, grafanaAlert.For, getAlert.For)
			assert.Equal(t, grafanaAlert.Labels, getAlert.Labels)
			assert.Equal(t, grafanaAlert.NoDataState, getAlert.NoDataState)
		}
	}
}

func TestIntegrationGrafanaAlert_GetGrafanaAlertUidNotExists(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		getAlert, err := underTest.GetGrafanaAlertRule("someUid")
		assert.Error(t, err)
		assert.Nil(t, getAlert)
		assert.Contains(t, err.Error(), "failed with missing grafana alert rule")
	}
}
