package alerts_test

import (
	"fmt"
	"github.com/jonboydell/logzio_client/alerts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAlert(t *testing.T) {
	underTest, err := setupAlertsTest()

	var alert *alerts.AlertType

	if assert.NoError(t, err) {
		createAlert := createValidAlert()

		alert, err = underTest.CreateAlert(createAlert)
		if assert.NoError(t, err) {
			underTest.DeleteAlert(alert.AlertId)
		}

		alertId := fmt.Sprintf("%d", alert.AlertId)

		if len(alertId) == 0 {
			t.Fatalf("%s should have a length > 0: %v", alertId, err)
		}

		createAlert = createValidAlert()
		createAlert.Filter = "{\"bool\":{\"must\":[{\"match\":{\"type\":\"mytype\"}}],\"must_not\":[]}}"
		alert, err = underTest.CreateAlert(createAlert)
		if assert.NoError(t, err) {
			underTest.DeleteAlert(alert.AlertId)
		}

		createAlert = createValidAlert()
		createAlert.Title = ""
		alert, err = underTest.CreateAlert(createAlert)
		assert.Error(t, err)

		createAlert = createValidAlert()
		createAlert.Filter = "This is my filter"
		alert, err = underTest.CreateAlert(createAlert)
		assert.Error(t, err)

		createAlert = createValidAlert()
		createAlert.Operation = ""
		alert, err = underTest.CreateAlert(createAlert)
		assert.Error(t, err)

		createAlert = createValidAlert()
		createAlert.ValueAggregationType = ""
		alert, err = underTest.CreateAlert(createAlert)
		assert.Error(t, err)

		createAlert = createValidAlert()
		createAlert.ValueAggregationField = ""
		alert, err = underTest.CreateAlert(createAlert)
		assert.Error(t, err)

		createAlert = createValidAlert()
		createAlert.ValueAggregationType = alerts.AggregationTypeNone
		createAlert.ValueAggregationField = nil
		createAlert.GroupByAggregationFields = []interface{}{"my_field"}
		alert, err = underTest.CreateAlert(createAlert)
		assert.Error(t, err)

		createAlert = createValidAlert()
		createAlert.ValueAggregationType = alerts.AggregationTypeCount
		createAlert.ValueAggregationField = "hello"
		createAlert.GroupByAggregationFields = []interface{}{"my_field"}
		alert, err = underTest.CreateAlert(createAlert)
		assert.Error(t, err)

		createAlert = createValidAlert()
		createAlert.NotificationEmails = nil
		alert, err = underTest.CreateAlert(createAlert)
		assert.Error(t, err)

		createAlert = createValidAlert()
		createAlert.QueryString = ""
		alert, err = underTest.CreateAlert(createAlert)
		assert.Error(t, err)

		createAlert = createValidAlert()
		createAlert.SeverityThresholdTiers = []alerts.SeverityThresholdType{
			alerts.SeverityThresholdType{
				Severity:  "TEST",
				Threshold: 10,
			},
		}
		alert, err = underTest.CreateAlert(createAlert)
		assert.Error(t, err)
	}
}
