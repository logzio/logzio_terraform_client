package alerts_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/alerts"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestAlerts_UpdateAlert(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	alertId := int64(1234567)

	mux.HandleFunc("/v1/alerts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(alertId, 10))

		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		assert.Contains(t, target, "title")
		assert.Contains(t, target, "description")
		assert.Contains(t, target, "query_string")
		assert.Contains(t, target, "severityThresholdTiers")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_alert.json"))
	})

	updatedAlert, err := underTest.UpdateAlert(alertId, alerts.CreateAlertType{
		Title:       "test update alert updated ",
		Description: "this is my description updated",
		QueryString: "loglevel:ERROR",
		Filter:      "",
		Operation:   alerts.OperatorGreaterThan,
		SeverityThresholdTiers: []alerts.SeverityThresholdType{
			alerts.SeverityThresholdType{
				alerts.SeverityHigh,
				10,
			},
		},
		SearchTimeFrameMinutes:       0,
		NotificationEmails:           []interface{}{},
		IsEnabled:                    true,
		SuppressNotificationsMinutes: 0,
		ValueAggregationType:         alerts.AggregationTypeCount,
		ValueAggregationField:        nil,
		GroupByAggregationFields:     []interface{}{"my_field"},
		AlertNotificationEndpoints:   []interface{}{},
	})

	assert.NoError(t, err)
	assert.NotNil(t, updatedAlert)
}

func TestAlerts_UpdateAlert_NotExist(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	alertId := int64(1234567)

	mux.HandleFunc("/v1/alerts/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_alert_not_exist.txt"))
	})

	updatedAlert, err := underTest.UpdateAlert(alertId, alerts.CreateAlertType{
		Title:       "test update alert updated ",
		Description: "this is my description updated",
		QueryString: "loglevel:ERROR",
		Filter:      "",
		Operation:   alerts.OperatorGreaterThan,
		SeverityThresholdTiers: []alerts.SeverityThresholdType{
			alerts.SeverityThresholdType{
				alerts.SeverityHigh,
				10,
			},
		},
		SearchTimeFrameMinutes:       0,
		NotificationEmails:           []interface{}{},
		IsEnabled:                    true,
		SuppressNotificationsMinutes: 0,
		ValueAggregationType:         alerts.AggregationTypeCount,
		ValueAggregationField:        nil,
		GroupByAggregationFields:     []interface{}{"my_field"},
		AlertNotificationEndpoints:   []interface{}{},
	})

	assert.Error(t, err)
	assert.Nil(t, updatedAlert)
}
