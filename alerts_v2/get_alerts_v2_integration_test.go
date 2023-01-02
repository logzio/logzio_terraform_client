package alerts_v2_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationAlertsV2_GetAlert(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.Title = "test alerts v2 get"

		alert, err := underTest.CreateAlert(createAlert)
		if assert.NoError(t, err) && assert.NotNil(t, alert) {
			time.Sleep(4 * time.Second)
			getAlert, err := underTest.GetAlert(alert.AlertId)
			assert.NoError(t, err)
			assert.NotNil(t, getAlert)
			assert.Equal(t, alert.AlertId, getAlert.AlertId)

			defer underTest.DeleteAlert(getAlert.AlertId)
		}
	}
}

func TestIntegrationAlertsV2_GetAlertWithSchedule(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.Title = "test alerts v2 get with schedule"
		createAlert.Schedule.CronExpression = "0 0/5 9-17 ? * * *"
		createAlert.Schedule.Timezone = "America/Costa_Rica"

		alert, err := underTest.CreateAlert(createAlert)
		if assert.NoError(t, err) && assert.NotNil(t, alert) {
			time.Sleep(4 * time.Second)
			getAlert, err := underTest.GetAlert(alert.AlertId)
			assert.NoError(t, err)
			assert.NotNil(t, getAlert)
			assert.Equal(t, alert.AlertId, getAlert.AlertId)
			assert.Equal(t, createAlert.Schedule.Timezone, getAlert.Schedule.Timezone)

			defer underTest.DeleteAlert(getAlert.AlertId)
		}
	}
}

func TestIntegrationAlertsV2_GetAlertIdNotExist(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		alert, err := underTest.GetAlert(int64(1234))
		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}
