// +build integration

package alerts_test

import (
	"github.com/logzio/logzio_terraform_client/alerts"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationAlerts_DeleteAlert(t *testing.T) {
	underTest, err := setupAlertsIntegrationTest()

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
					{
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

func TestIntegrationAlerts_DeleteMissingAlert(t *testing.T) {
	underTest, err := setupAlertsIntegrationTest()

	if assert.NoError(t, err) {
		err = underTest.DeleteAlert(int64(1234567))
		assert.Error(t, err)
	}
}
