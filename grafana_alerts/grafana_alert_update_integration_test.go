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
			grafanaAlert.Title = "changed"
			err = underTest.UpdateGrafanaAlertRule(*grafanaAlert)
			assert.NoError(t, err)
			// verify that the update was made
			time.Sleep(time.Second * 4)
			getAlert, err := underTest.GetGrafanaAlertRule(grafanaAlert.Uid)
			assert.NoError(t, err)
			assert.Equal(t, grafanaAlert.Title, getAlert.Title)
		}
	}
}

func TestIntegrationGrafanaAlert_UpdateGrafanaAlertIdNotFound(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()

	if assert.NoError(t, err) {
		request := getGrafanaAlertRuleObject()
		request.Uid = "not-exist"
		err = underTest.UpdateGrafanaAlertRule(request)
		assert.Error(t, err)
	}
}
