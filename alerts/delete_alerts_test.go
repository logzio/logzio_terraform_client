package alerts

import (
	"github.com/jonboydell/logzio_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteAlert(t *testing.T) {
	var client *Alerts
	client, err := New(test_utils.GetApiToken())
	assert.NoError(t, err)

	if assert.NotNil(t, client) {
		// create alert
		createAlert := createValidAlert()
		alert, err := client.CreateAlert(createAlert)
		assert.NoError(t, err)
		assert.NotNil(t, alert)

		// delete alert
		alertId := alert.AlertId
		client.DeleteAlert(alertId)

		// make sure alert is really deleted
		alert, err = client.GetAlert(alertId)
		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}

func TestDeleteMissingAlert(t *testing.T) {
	var client *Alerts
	client, err := New(test_utils.GetApiToken())
	assert.NoError(t, err)

	if assert.NotNil(t, client) {
		// delete alert that doesn't exist
		err = client.DeleteAlert(12345)
		assert.Error(t, err)
	}
}