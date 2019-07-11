package alerts_test

import (
	"github.com/jonboydell/logzio_client/alerts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateAlert(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "test update alert",
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

		if assert.NoError(t, err) && assert.NotNil(t, alert) {
			updatedAlert, err := underTest.UpdateAlert(alert.AlertId, alerts.CreateAlertType{
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

			err = underTest.DeleteAlert(alert.AlertId)
			assert.NoError(t, err)
		}
	}
}