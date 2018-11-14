package logzio_client

import (
	"log"
	"os"
	"testing"
)

func TestListAlerts(t *testing.T) {

	api_token := os.Getenv("LOGZIO_API_TOKEN")

	var client *Client
	client = New(api_token)

	alerts, err := client.ListAlerts()

	if err != nil {
		t.Fatalf("%q should not have raised an error: %v", "ListAlerts", err)
	}

	log.Print(alerts)

}