package logzio_client

import (
	"testing"
)

func TestUpdateAlert(t *testing.T) {
	api_token := getApiToken(t)

	var client *Client
	client = New(api_token)

	createAlert := createValidAlert()

	var alert *AlertType
	var err error
	alerts := []int64{}

	alert, err = client.CreateAlert(createAlert)
	if err != nil {
		t.Fatalf("%q should not have raised an error: %v", "CreateAlert", err)
	}
	alerts = append(alerts, alert.AlertId)

	updateAlert := alertToCreateAlert(alert)
	alert, err = client.UpdateAlert(alert.AlertId, updateAlert)

	// clean up any created alerts
	for x := 0; x < len(alerts); x++ {
		client.DeleteAlert(alerts[x])
	}
}
