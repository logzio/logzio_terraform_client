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
