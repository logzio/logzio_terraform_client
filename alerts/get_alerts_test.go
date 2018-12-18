package alerts

import (
	"github.com/jonboydell/logzio_client/client"
	"github.com/jonboydell/logzio_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAlert(t *testing.T) {
	api_token := test_utils.GetApiToken()
	if len(api_token) == 0 {
		t.Fatalf("%v could not get an API token from %v", "TestDeleteAlert", client.ENV_LOGZIO_BASE_URL)
	}

	var client *Alerts
	client, err := New(api_token)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	createAlert := createValidAlert()

	var alert *AlertType
	alerts := []int64{}

	alert, err = client.CreateAlert(createAlert)
	if err != nil {
		t.Fatalf("%v should not have raised an error: %v", "CreateAlert", err)
	}
	alerts = append(alerts, alert.AlertId)

	if alert.AlertId <= 0 {
		t.Fatalf("%d have a value > 0: %v", alert.AlertId, err)
	}

	_, err = client.GetAlert(alert.AlertId)

	if err != nil {
		t.Fatalf("%v should not have raised an error: %v", "GetAlert", err)
	}

	var alertId int64 = 12345
	_, err = client.GetAlert(alertId)
	if err == nil {
		t.Fatalf("%v should have raised an error, alert %d not found: %v", "GetAlert", alertId, err)
	}

	// clean up any created alerts
	for x := 0; x < len(alerts); x++ {
		client.DeleteAlert(alerts[x])
	}
}
