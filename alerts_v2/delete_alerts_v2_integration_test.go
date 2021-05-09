package alerts_v2_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationAlertsV2_DeleteAlert(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.Title = "test alerts v2 delete alert"

		alert, err := underTest.CreateAlert(createAlert)
		if assert.NoError(t, err) && assert.NotNil(t, alert) {
			time.Sleep(4 * time.Second)
			defer func() {
				err = underTest.DeleteAlert(alert.AlertId)
				assert.NoError(t, err)
			}()
		}
	}
}

func TestIntegrationAlertsV2_DeleteAlertIdNotExist(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		err = underTest.DeleteAlert(int64(1234567))
		assert.Error(t, err)
	}
}
