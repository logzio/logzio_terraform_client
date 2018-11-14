package logzio_client

import (
	"os"
	"testing"
)

func TestListAlerts(t *testing.T) {
	api_token := os.Getenv("LOGZIO_API_TOKEN")
	if len(api_token) == 0 {
		t.Fatalf("%q could not get an API token from LOGZIO_API_TOKEN", "TestDeleteAlert")
	}

	var client *Client
	client = New(api_token)

	_, err := client.ListAlerts()

	if err != nil {
		t.Fatalf("%q should not have raised an error: %v", "ListAlerts", err)
	}
}