package logzio_client

import (
	"testing"
)

func TestListAlerts(t *testing.T) {

	api_token := GetApiToken(t)
	if len(api_token) == 0 {
		t.Fatalf("%v could not get an API token from %v", "TestDeleteAlert", ENV_LOGZIO_API_TOKEN)
	}

	var client *Client
	client = New(api_token)

	_, err := client.ListAlerts()

	if err != nil {
		t.Fatalf("%q should not have raised an error: %v", "ListAlerts", err)
	}
}
