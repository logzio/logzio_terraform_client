package alerts_v2_test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestIntegrationAlertsV2_EnableAlert(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.Title = "test alerts v2 enable alert"
		createAlert.Enabled = strconv.FormatBool(false)

		alert, err := underTest.CreateAlert(createAlert)
		if assert.NoError(t, err) && assert.NotNil(t, alert) {
			time.Sleep(4 * time.Second)
			enabledAlert, err := underTest.EnableAlert(*alert)
			assert.NoError(t, err)
			assert.NotNil(t, enabledAlert)
			assert.True(t, enabledAlert.Enabled)

			defer underTest.DeleteAlert(enabledAlert.AlertId)
		}
	}
}

func TestIntegrationAlertsV2_EnableAlertIdNotExist(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		alertType := getAlertType()

		alert, err := underTest.EnableAlert(alertType)
		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}