package grafana_alerts_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationGrafanaAlert_UpdateGrafanaAlert(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()

	if assert.NoError(t, err) {
		createGrafanaAlert := getGrafanaAlertRuleObject()
		createGrafanaAlert.Title = fmt.Sprintf("%s_%s", createGrafanaAlert.Title, "update")
		grafanaAlert, err := underTest.CreateGrafanaAlertRule(createGrafanaAlert)
		if assert.NoError(t, err) && assert.NotNil(t, grafanaAlert) && assert.NotEmpty(t, grafanaAlert.Uid) {
			defer underTest.DeleteGrafanaAlertRule(grafanaAlert.Uid)
			time.Sleep(time.Second * 2)
			createGrafanaAlert.Condition = "B"
			err = underTest.UpdateGrafanaAlertRule(*grafanaAlert)
			assert.NoError(t, err)
			// verify that the update was made
			time.Sleep(time.Second * 4)
			getFolder, err := underTest.GetGrafanaAlertRule(grafanaAlert.Uid)
			assert.NoError(t, err)
			assert.Equal(t, createGrafanaAlert.Condition, getFolder.Condition)
		}
	}
}
