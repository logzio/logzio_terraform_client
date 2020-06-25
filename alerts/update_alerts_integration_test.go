// +build integration

package alerts_test

import (
	"github.com/logzio/logzio_terraform_client/alerts"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationAlerts_UpdateAlert(t *testing.T) {
	underTest, err := setupAlertsIntegrationTest()

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

			time.Sleep(3 * time.Second)
			if assert.NoError(t, err) && assert.NotZero(t, alert) {
				defer underTest.DeleteAlert(alert.AlertId)
			}
		}
	}
}
