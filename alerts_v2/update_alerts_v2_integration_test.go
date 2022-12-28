package alerts_v2_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationAlertsV2_UpdateAlert(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.Title = "test alerts v2 update"

		alert, err := underTest.CreateAlert(createAlert)
		if assert.NoError(t, err) && assert.NotNil(t, alert) {
			time.Sleep(4 * time.Second)
			createAlert.Title = "test alerts v2 update change title"
			updated, err := underTest.UpdateAlert(alert.AlertId, createAlert)
			assert.NoError(t, err)
			assert.NotNil(t, updated)
			assert.Equal(t, createAlert.Title, updated.Title)

			time.Sleep(4 * time.Second)
			defer underTest.DeleteAlert(updated.AlertId)
		}

	}
}

func TestIntegrationAlertsV2_UpdateAlertAddSchedule(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.Title = "test alerts v2 update add schedule"

		alert, err := underTest.CreateAlert(createAlert)
		if assert.NoError(t, err) && assert.NotNil(t, alert) {
			time.Sleep(4 * time.Second)
			createAlert.Schedule.CronExpression = "0 0/5 9-17 ? * * *"
			createAlert.Schedule.Timezone = "Europe/Madrid"
			updated, err := underTest.UpdateAlert(alert.AlertId, createAlert)
			assert.NoError(t, err)
			assert.NotNil(t, updated)
			assert.Equal(t, createAlert.Schedule.Timezone, updated.Schedule.Timezone)

			time.Sleep(4 * time.Second)
			defer underTest.DeleteAlert(updated.AlertId)
		}

	}
}

func TestIntegrationAlertsV2_UpdateAlertRemoveSchedule(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.Title = "test alerts v2 update remove schedule"
		createAlert.Schedule.CronExpression = "0 0/5 9-17 ? * * *"
		createAlert.Schedule.Timezone = "Europe/Madrid"

		alert, err := underTest.CreateAlert(createAlert)
		if assert.NoError(t, err) && assert.NotNil(t, alert) {
			time.Sleep(4 * time.Second)
			createAlert.Schedule.CronExpression = ""
			createAlert.Schedule.Timezone = ""
			updated, err := underTest.UpdateAlert(alert.AlertId, createAlert)
			assert.NoError(t, err)
			assert.NotNil(t, updated)
			assert.Equal(t, createAlert.Schedule.CronExpression, updated.Schedule.CronExpression)
			assert.Equal(t, "UTC", updated.Schedule.Timezone)

			time.Sleep(4 * time.Second)
			defer underTest.DeleteAlert(updated.AlertId)
		}

	}
}
