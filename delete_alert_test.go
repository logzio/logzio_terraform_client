package logzio_client

import (
	"os"
	"testing"
)

func containsAlert(alerts []AlertType, alertId int64) bool {
	for x := 0; x < len(alerts); x++ {
		if alerts[x].AlertId == alertId {
			return true
		}
	}
	return false
}

func TestDeleteAlert(t *testing.T) {
	api_token := os.Getenv("LOGZIO_API_TOKEN")
	if len(api_token) == 0 {
		t.Fatalf("%q could not get an API token from LOGZIO_API_TOKEN", "TestDeleteAlert")
	}

	var client *Client
	client = New(api_token)

	createAlert := createValidAlert()

	alert, err := client.CreateAlert(createAlert)
	if err != nil {
		t.Fatalf("%q should not have raised an error: %v", "DeleteAlert", err)
	}

	alertId := alert.AlertId
	client.DeleteAlert(alertId)

	alerts, err := client.ListAlerts()
	if containsAlert(alerts, alertId) {
		t.Fatalf("%q %d should have been deleted, but is returned by ListAlerts", "DeleteAlert", alertId)
	}
}
