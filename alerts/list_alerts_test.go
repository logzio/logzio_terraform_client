package alerts

import (
	"github.com/jonboydell/logzio_client/client"
	"github.com/jonboydell/logzio_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListAlerts(t *testing.T) {
	api_token, _ := test_utils.GetApiToken()
	if len(api_token) == 0 {
		t.Fatalf("%v could not get an API token from %v", "TestDeleteAlert", client.ENV_LOGZIO_BASE_URL)
	}

	var client *Alerts
	client, err := New(api_token)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	_, err = client.ListAlerts()

	if err != nil {
		t.Fatalf("%q should not have raised an error: %v", "ListAlerts", err)
	}
}
