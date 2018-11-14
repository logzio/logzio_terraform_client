package logzio_client

import (
	"os"
	"testing"
)

const ENV_LOGZIO_API_TOKEN string = "LOGZIO_API_TOKEN"

func containsAlert(alerts []AlertType, alertId int64) bool {
	for x := 0; x < len(alerts); x++ {
		if alerts[x].AlertId == alertId {
			return true
		}
	}
	return false
}

func GetApiToken(t *testing.T) string {
	api_token := os.Getenv(ENV_LOGZIO_API_TOKEN)
	if len(api_token) == 0 {
		t.Fatalf("%v could not get an API token from %v", "TestDeleteAlert", ENV_LOGZIO_API_TOKEN)
	}
	return api_token
}

func TestDeleteAlert(t *testing.T) {
	api_token := GetApiToken(t)

	var client *Client
	client = New(api_token)

	createAlert := createValidAlert()

	alert, err := client.CreateAlert(createAlert)
	if err != nil {
		t.Fatalf("%v should not have raised an error: %v", "DeleteAlert", err)
	}

	alertId := alert.AlertId
	client.DeleteAlert(alertId)

	alerts, err := client.ListAlerts()
	if containsAlert(alerts, alertId) {
		t.Fatalf("%v %d should have been deleted, but is returned by ListAlerts", "DeleteAlert", alertId)
	}
}
