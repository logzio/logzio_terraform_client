package logzio_client

import (
	"os"
	"testing"
)

func containsAlert(alerts []AlertType, alertId int64) bool {
	for x := 0; x < len(alerts); x++  {
		if alerts[x].AlertId == alertId {
			return true
		}
	}
	return false
}

func TestDeleteAlert(t *testing.T) {
	api_token := os.Getenv("LOGZIO_API_TOKEN")

	var client *Client
	client = New(api_token)

	createAlert := createValidAlert()

	alert, err := client.CreateAlert(createAlert)
	if err != nil {
		t.Fatalf("%q should not have raised an error: %v", "CreateAlert", err)
	}

	alertId := alert.AlertId
	client.DeleteAlert(alertId)

	alerts, err := client.ListAlerts()
	if containsAlert(alerts, alertId) {
		t.Fatalf("Alert %d should have been deleted, but is returned by ListAlerts", alertId)
	}
}

