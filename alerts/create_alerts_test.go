package alerts_test

import (
	"github.com/jonboydell/logzio_client/alerts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAlert(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "this is my title",
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
		if assert.NoError(t, err) {
			underTest.DeleteAlert(alert.AlertId)
		}
	}
}

func TestCreateAlertWithFilter(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "this is my title",
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

		if assert.NoError(t, err) {
			underTest.DeleteAlert(alert.AlertId)
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
		if alert != nil {
			err = underTest.DeleteAlert(alert.AlertId)
			assert.Nil(t, alert)
		}
	}
}

func TestCreateAlertWithInvalidFilter(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "this is my title",
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
		if alert != nil {
			err = underTest.DeleteAlert(alert.AlertId)
			assert.Nil(t, alert)
		}
	}
}

func TestCreateAlertWithInvalidValueAggregationType(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "this is my title",
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
		if alert != nil {
			err = underTest.DeleteAlert(alert.AlertId)
			assert.Nil(t, alert)
		}
	}
}


func TestCreateAlertWithInvalidValueAggregationField(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "this is my title",
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
		if alert != nil {
			err = underTest.DeleteAlert(alert.AlertId)
			assert.Nil(t, alert)
		}
	}
}

func TestCreateAlertWithInvalidValueAggregationTypeNone(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "this is my title",
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
		if alert != nil {
			err = underTest.DeleteAlert(alert.AlertId)
			assert.Nil(t, alert)
		}
	}
}

func TestCreateAlertWithInvalidValueAggregationTypeCount(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "this is my title",
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
		if alert != nil {
			err = underTest.DeleteAlert(alert.AlertId)
			assert.Nil(t, alert)
		}
	}
}


func TestCreateAlertWithNoNotifications(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "this is my title",
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
		if alert != nil {
			err = underTest.DeleteAlert(alert.AlertId)
			assert.Nil(t, alert)
		}
	}
}

func TestCreateAlertWithNoQuery(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "this is my title",
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
		if alert != nil {
			err = underTest.DeleteAlert(alert.AlertId)
			assert.Nil(t, alert)
		}
	}
}

func TestCreateAlertWithInvalidSeverity(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(alerts.CreateAlertType{
			Title:       "this is my title",
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
		if alert != nil {
			err = underTest.DeleteAlert(alert.AlertId)
			assert.Nil(t, alert)
		}
	}
}