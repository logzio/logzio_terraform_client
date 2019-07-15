package alerts_test

import (
	"github.com/jonboydell/logzio_client/alerts"
	"github.com/jonboydell/logzio_client/client"
	"github.com/jonboydell/logzio_client/test_utils"
)

func setupAlertsIntegrationTest() (*alerts.AlertsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := alerts.New(apiToken)
	underTest.BaseUrl = client.GetLogzIoBaseUrl()
	return underTest, nil
}

func createValidAlert() alerts.CreateAlertType {
	return alerts.CreateAlertType{
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
	}
}

func createUpdateAlert() alerts.CreateAlertType {
	return alerts.CreateAlertType{
		Title:       "this is my updated title",
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
	}
}
