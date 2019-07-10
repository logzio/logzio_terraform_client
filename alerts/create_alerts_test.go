package alerts_test

import (
	"github.com/jonboydell/logzio_client/alerts"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateAlert(t *testing.T) {
	underTest, err := setupAlertsTest()

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
		time.Sleep(3 * time.Seconds)
		if assert.NoError(t, err) && assert.NotZero(t, alert) {
			err = underTest.DeleteAlert(alert.AlertId)
			assert.NoError(t, err)
		}
	}
}

func TestCreateAlertWithFilter(t *testing.T) {
	underTest, err := setupAlertsTest()

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
		time.Sleep(3000)

		if assert.NoError(t, err) && assert.NotZero(t, alert) {
			err = underTest.DeleteAlert(alert.AlertId)
			assert.NoError(t, err)
		}
	}
}

func TestCreateAlertWithNoTitle(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "",
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

		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}

func TestCreateAlertWithInvalidFilter(t *testing.T) {
	underTest, err := setupAlertsTest()

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

func TestCreateAlertWithInvalidValueAggregationType(t *testing.T) {
	underTest, err := setupAlertsTest()

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


func TestCreateAlertWithInvalidValueAggregationField(t *testing.T) {
	underTest, err := setupAlertsTest()

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

func TestCreateAlertWithInvalidValueAggregationTypeNone(t *testing.T) {
	underTest, err := setupAlertsTest()

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

func TestCreateAlertWithInvalidValueAggregationTypeCount(t *testing.T) {
	underTest, err := setupAlertsTest()

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


func TestCreateAlertWithNoNotifications(t *testing.T) {
	underTest, err := setupAlertsTest()

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

func TestCreateAlertWithNoQuery(t *testing.T) {
	underTest, err := setupAlertsTest()

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

func TestCreateAlertWithInvalidSeverity(t *testing.T) {
	underTest, err := setupAlertsTest()

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