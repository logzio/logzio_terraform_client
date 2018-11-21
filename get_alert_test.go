package logzio_client

import (
	"testing"
)

func TestGetAlert(t *testing.T) {
	api_token := getApiToken(t)
	if len(api_token) == 0 {
		t.Fatalf("%v could not get an API token from %v", "TestDeleteAlert", ENV_LOGZIO_API_TOKEN)
	}

	var client *Client
	client = New(api_token)

	createAlert := createValidAlert()

	var alert *AlertType
	var err error
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
