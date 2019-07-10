package alerts_test

import (
	"github.com/jonboydell/logzio_client/alerts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAlert(t *testing.T) {
	underTest, err := setupAlertsTest()

	alert, err := underTest.CreateAlert(alerts.CreateAlertType{
		Title:       "test get alert",
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
	assert.NoError(t, err)
	if assert.NoError(t, err) && assert.NotNil(t, alert) {
		v, err := underTest.GetAlert(alert.AlertId)
		assert.NoError(t, err)
		assert.NotNil(t, v)
	}

	if assert.NoError(t, err) && assert.NotZero(t, alert) {
		err = underTest.DeleteAlert(alert.AlertId)
		assert.NoError(t, err)
	}
	_, err = underTest.GetAlert(12345)
	assert.Error(t, err)
}

func TestGetAlert_DeleteAlertNotExisting(t *testing.T) {
	underTest, err := setupAlertsTest()
	_, err = underTest.GetAlert(12345)
	assert.Error(t, err)
}