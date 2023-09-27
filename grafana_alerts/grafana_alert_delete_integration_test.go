package grafana_alerts_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationGrafanaAlert_DeleteGrafanaAlert(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createGrafanaAlert := getGrafanaAlertRuleObject()
		grafanaAlert, err := underTest.CreateGrafanaAlertRule(createGrafanaAlert)
		if assert.NoError(t, err) && assert.NotNil(t, grafanaAlert) && assert.NotEmpty(t, grafanaAlert.Uid) {
			time.Sleep(2 * time.Second)
			defer func() {
				err = underTest.DeleteGrafanaAlertRule(grafanaAlert.Uid)
				assert.NoError(t, err)
			}()
		}
	}
}

func TestIntegrationGrafanaAlert_DeleteGrafanaAlertEmptyUid(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		err = underTest.DeleteGrafanaAlertRule("")
		assert.Error(t, err)
	}
}

func TestIntegrationGrafanaAlert_DeleteGrafanaAlertUidNotFound(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		err = underTest.DeleteGrafanaAlertRule("uidNotFound")
		assert.Error(t, err)
	}
}
