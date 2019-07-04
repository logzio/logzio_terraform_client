package alerts

import (
	"github.com/jonboydell/logzio_client/client"
	"github.com/jonboydell/logzio_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAlert(t *testing.T) {
	api_token, _ := test_utils.GetApiToken()
	if len(api_token) == 0 {
		t.Fatalf("%v could not get an API token from %v", "TestDeleteAlert", client.ENV_LOGZIO_BASE_URL)
	}

	var underTest *Alerts
	underTest, err := New(api_token)
	assert.NoError(t, err)
	assert.NotNil(t, underTest)

	createAlert := createValidAlert()

	var alert *AlertType

	alert, err = underTest.CreateAlert(createAlert)
	assert.NoError(t, err)
	assert.NotNil(t, alert)

	alert, err = underTest.GetAlert(alert.AlertId)
	assert.NoError(t, err)
	assert.NotNil(t, alert)

	err = underTest.DeleteAlert(alert.AlertId)
	assert.NoError(t, err)

	_, err = underTest.GetAlert(12345)
	assert.Error(t, err)
}
