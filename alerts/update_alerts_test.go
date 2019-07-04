package alerts

import (
	"github.com/jonboydell/logzio_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateAlert(t *testing.T) {
	var client *Alerts
	var alert *AlertType

	api_token, _ := test_utils.GetApiToken()

	client, err := New(api_token)
	assert.NoError(t, err)

	createAlert := createValidAlert()
	if assert.NotNil(t, client) {
		alert, err = client.CreateAlert(createAlert)
		assert.NoError(t, err)
		assert.NotNil(t, alert)
		client.DeleteAlert(alert.AlertId)
	}
}
