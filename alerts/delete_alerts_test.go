package alerts_test

import (
	"github.com/jonboydell/logzio_client/alerts"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDeleteAlert(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		// create alert
		alert, err := underTest.CreateAlert(
			alerts.CreateAlertType{
				Title:       "this is my deletable alert",
				Description: "this is my description",
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
		time.Sleep(3 * time.Second)
		if assert.NoError(t, err) {
			defer underTest.DeleteAlert(alert.AlertId)
		}
	}
}

func TestDeleteMissingAlert(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		err = underTest.DeleteAlert(12345)
		assert.Error(t, err)
	}
}
