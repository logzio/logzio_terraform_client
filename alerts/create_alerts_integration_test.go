package alerts_test

import (
	"github.com/jonboydell/logzio_client/alerts"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationAlerts_CreateAlert(t *testing.T) {
	underTest, err := setupAlertsIntegrationTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "test create alert",
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
		if assert.NoError(t, err) && assert.NotZero(t, alert) {
			defer underTest.DeleteAlert(alert.AlertId)
		}
	}
}

func TestIntegrationAlerts_CreateAlertWithFilter(t *testing.T) {
	underTest, err := setupAlertsIntegrationTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "test create alert with filter",
			Description: "this is my description",
			QueryString: "loglevel:ERROR",
			Filter:      "{\"bool\":{\"must\":[{\"match\":{\"type\":\"mytype\"}}],\"must_not\":[]}}",
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
		if assert.NoError(t, err) && assert.NotZero(t, alert) {
			defer underTest.DeleteAlert(alert.AlertId)
		}
	}
}

func TestIntegrationAlerts_CreateAlertInvalidFilter(t *testing.T) {
	underTest, err := setupAlertsIntegrationTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "test alert with invalid filter",
			Description: "this is my description",
			QueryString: "loglevel:ERROR",
			Filter:      "Invalid Filter",
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
		assert.Nil(t, alert)
	}
}

func TestIntegrationAlerts_CreateAlertInvalidAggregationType(t *testing.T) {
	underTest, err := setupAlertsIntegrationTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "test alert with invalid agg type",
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
			ValueAggregationType:         "",
			ValueAggregationField:        nil,
			GroupByAggregationFields:     []interface{}{"my_field"},
			AlertNotificationEndpoints:   []interface{}{},
		})

		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}


func TestIntegrationAlerts_CreateAlertInvaldValueAggregationField(t *testing.T) {
	underTest, err := setupAlertsIntegrationTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "test alert with invalid agg field",
			Description: "this is my description",
			QueryString: "loglevel:ERROR",
			Filter:      "{\"bool\":{\"must\":[{\"match\":{\"type\":\"mytype\"}}],\"must_not\":[]}}",
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
			ValueAggregationField:        "",
			GroupByAggregationFields:     []interface{}{"my_field"},
			AlertNotificationEndpoints:   []interface{}{},
		})

		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}

func TestIntegrationAlerts_CreateAlertInvalidValueAggregationTypeNone(t *testing.T) {
	underTest, err := setupAlertsIntegrationTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "test alert with none agg type",
			Description: "this is my description",
			QueryString: "loglevel:ERROR",
			Filter:      "{\"bool\":{\"must\":[{\"match\":{\"type\":\"mytype\"}}],\"must_not\":[]}}",
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
			ValueAggregationType:         alerts.AggregationTypeNone,
			ValueAggregationField:        nil,
			GroupByAggregationFields:     []interface{}{"my_field"},
			AlertNotificationEndpoints:   []interface{}{},
		})

		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}

func TestIntegrationAlerts_CreateAlertInvalidValueAggregationTypeCount(t *testing.T) {
	underTest, err := setupAlertsIntegrationTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "test alert with invalid agg type count",
			Description: "this is my description",
			QueryString: "loglevel:ERROR",
			Filter:      "{\"bool\":{\"must\":[{\"match\":{\"type\":\"mytype\"}}],\"must_not\":[]}}",
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
			ValueAggregationField:        "hello",
			GroupByAggregationFields:     []interface{}{"my_field"},
			AlertNotificationEndpoints:   []interface{}{},
		})

		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}


func TestIntegrationAlerts_CreateAlertNoNotifications(t *testing.T) {
	underTest, err := setupAlertsIntegrationTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "test create alert with no notification",
			Description: "this is my description",
			QueryString: "loglevel:ERROR",
			Filter:      "{\"bool\":{\"must\":[{\"match\":{\"type\":\"mytype\"}}],\"must_not\":[]}}",
			Operation:   alerts.OperatorGreaterThan,
			SeverityThresholdTiers: []alerts.SeverityThresholdType{
				alerts.SeverityThresholdType{
					alerts.SeverityHigh,
					10,
				},
			},
			SearchTimeFrameMinutes:       0,
			NotificationEmails:           nil,
			IsEnabled:                    true,
			SuppressNotificationsMinutes: 0,
			ValueAggregationType:         alerts.AggregationTypeCount,
			ValueAggregationField:        nil,
			GroupByAggregationFields:     []interface{}{"my_field"},
			AlertNotificationEndpoints:   []interface{}{},
		})

		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}

func TestIntegrationAlerts_CreateAlertNoQueryString(t *testing.T) {
	underTest, err := setupAlertsIntegrationTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "test create alert with no query",
			Description: "this is my description",
			QueryString: "",
			Filter:      "{\"bool\":{\"must\":[{\"match\":{\"type\":\"mytype\"}}],\"must_not\":[]}}",
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
		assert.Nil(t, alert)
	}
}

func TestIntegrationAlerts_CreateAlertInvalidSeverity(t *testing.T) {
	underTest, err := setupAlertsIntegrationTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "test create alert with invalid severity",
			Description: "this is my description",
			QueryString: "loglevel:ERROR",
			Filter:      "{\"bool\":{\"must\":[{\"match\":{\"type\":\"mytype\"}}],\"must_not\":[]}}",
			Operation:   alerts.OperatorGreaterThan,
			SeverityThresholdTiers: []alerts.SeverityThresholdType{
				alerts.SeverityThresholdType{
					"TEST",
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
		assert.Nil(t, alert)
	}
}