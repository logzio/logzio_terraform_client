package alerts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteAlert(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		// create alert
		createAlert := createValidAlert()
		alert, err := underTest.CreateAlert(createAlert)
		assert.NoError(t, err)
		assert.NotNil(t, alert)

		// delete alert
		alertId := alert.AlertId
		underTest.DeleteAlert(alertId)

		// make sure alert is really deleted
		alert, err = underTest.GetAlert(alertId)
		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}

func TestDeleteMissingAlert(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		// delete alert that doesn't exist
		err = underTest.DeleteAlert(12345)
		assert.Error(t, err)
	}
}
