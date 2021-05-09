package alerts_v2_test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestIntegrationAlertsV2_DisableAlert(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.Title = "test alerts v2 disable alert"
		createAlert.Enabled = strconv.FormatBool(true)

		alert, err := underTest.CreateAlert(createAlert)
		if assert.NoError(t, err) && assert.NotNil(t, alert) {
			time.Sleep(4 * time.Second)
			disabledAlert, err := underTest.DisableAlert(*alert)
			assert.NoError(t, err)
			assert.NotNil(t, disabledAlert)
			assert.False(t, disabledAlert.Enabled)

			defer underTest.DeleteAlert(disabledAlert.AlertId)
		}
	}
}

func TestIntegrationAlertsV2_DisableAlertIdNotExist(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		alertType := getAlertType()

		alert, err := underTest.DisableAlert(alertType)
		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}
