package alerts

import (
	"fmt"
	"github.com/jonboydell/logzio_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAlert(t *testing.T) {
	api_token := test_utils.GetApiToken()

	var client *Alerts
	client, err := New(api_token)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	createAlert := createValidAlert()

	var alert *AlertType
	alerts := []int64{}

	if assert.NotNil(t, client) {
		alert, err = client.CreateAlert(createAlert)
		if err != nil {
			t.Fatalf("%q should not have raised an error: %v", "CreateAlert", err)
		}
		alerts = append(alerts, alert.AlertId)

		alertId := fmt.Sprintf("%d", alert.AlertId)

		if len(alertId) == 0 {
			t.Fatalf("%s should have a length > 0: %v", alertId, err)
		}

		createAlert = createValidAlert()
		createAlert.Filter = "{\"bool\":{\"must\":[{\"match\":{\"type\":\"mytype\"}}],\"must_not\":[]}}"
		alert, err = client.CreateAlert(createAlert)
		if err != nil {
			t.Fatalf("%q should not have raised an error: %v", "CreateAlert", err)
		}
		alerts = append(alerts, alert.AlertId)

		createAlert = createValidAlert()
		createAlert.Title = ""
		alert, err = client.CreateAlert(createAlert)
		if err == nil {
			t.Fatalf("should have raised an error for blank title: %v", err)
		}

		createAlert = createValidAlert()
		createAlert.Filter = "This is my filter"
		alert, err = client.CreateAlert(createAlert)
		if err == nil {
			t.Fatalf("should have raised an error for invalid use of filter: %v", err)
		}

		createAlert = createValidAlert()
		createAlert.Operation = ""
		alert, err = client.CreateAlert(createAlert)
		if err == nil {
			t.Fatalf("should have raised an error for invalid use of operation: %v", err)
		}

		createAlert = createValidAlert()
		createAlert.ValueAggregationType = ""
		alert, err = client.CreateAlert(createAlert)
		if err == nil {
			t.Fatalf("should have raised an error for invalid use of valueAggregationType: %v", err)
		}

		createAlert = createValidAlert()
		createAlert.ValueAggregationField = ""
		alert, err = client.CreateAlert(createAlert)
		if err == nil {
			t.Fatalf("should have raised an error for invalid use of valueAggregationField: %v", err)
		}

		createAlert = createValidAlert()
		createAlert.ValueAggregationType = AggregationTypeNone
		createAlert.ValueAggregationField = nil
		createAlert.GroupByAggregationFields = []interface{}{"my_field"}
		alert, err = client.CreateAlert(createAlert)
		if err == nil {
			t.Fatalf("should have raised an error for invalid use of groupByAggregationFields: %v", err)
		}

		createAlert = createValidAlert()
		createAlert.ValueAggregationType = AggregationTypeCount
		createAlert.ValueAggregationField = "hello"
		createAlert.GroupByAggregationFields = []interface{}{"my_field"}
		alert, err = client.CreateAlert(createAlert)
		if err == nil {
			t.Fatalf("should have raised an error for invalid use of valueAggregationField: %v", err)
		}

		createAlert = createValidAlert()
		createAlert.NotificationEmails = nil
		alert, err = client.CreateAlert(createAlert)
		if err == nil {
			t.Fatalf("should have raised an error for invalid use of notificationEmails: %v", err)
		}

		createAlert = createValidAlert()
		createAlert.QueryString = ""
		alert, err = client.CreateAlert(createAlert)
		if err == nil {
			t.Fatalf("should have raised an error for invalid use of queryString: %v", err)
		}

		createAlert = createValidAlert()
		createAlert.SeverityThresholdTiers = []SeverityThresholdType{
			SeverityThresholdType{
				Severity:  "TEST",
				Threshold: 10,
			},
		}
		alert, err = client.CreateAlert(createAlert)
		if err == nil {
			t.Fatalf("should have raised an error for invalid severity: %v", err)
		}

		// clean up any created alerts
		for x := 0; x < len(alerts); x++ {
			client.DeleteAlert(alerts[x])
		}
	}
}
