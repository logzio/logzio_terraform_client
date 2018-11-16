package logzio_client

import (
	"fmt"
	"testing"
)

func createValidAlert() CreateAlertType {
	return CreateAlertType{
		Title:       "this is my title",
		Description: "this is my description",
		QueryString: "loglevel:ERROR",
		Filter:      "",
		Operation:   OperatorGreaterThan,
		SeverityThresholdTiers:      []SeverityThresholdType{
			SeverityThresholdType{
				SeverityHigh,
				10,
			},
		},
		SearchTimeFrameMinutes:      0,
		NotificationEmails:          []interface{}{},
		IsEnabled:                   true,
		SuppressNotificationMinutes: 0,
		ValueAggregationType:        AggregationTypeCount,
		ValueAggregationField:       nil,
		GroupByAggregationFields:    []interface{}{"my_field"},
		AlertNotificationEndpoints:  []interface{}{},
	}
}

func TestCreateAlert(t *testing.T) {

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
			Severity:"TEST",
			Threshold:10,
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
