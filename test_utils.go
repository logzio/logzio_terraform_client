package logzio_client

import (
	"os"
	"testing"
)

const ENV_LOGZIO_API_TOKEN string = "LOGZIO_API_TOKEN"

func containsAlert(alerts []AlertType, alertId int64) bool {
	for x := 0; x < len(alerts); x++ {
		if alerts[x].AlertId == alertId {
			return true
		}
	}
	return false
}

func getApiToken(t *testing.T) string {
	api_token := os.Getenv(ENV_LOGZIO_API_TOKEN)
	if len(api_token) == 0 {
		t.Fatalf("%v could not get an API token from %v", "TestDeleteAlert", ENV_LOGZIO_API_TOKEN)
	}
	return api_token
}

func alertToCreateAlert(alert *AlertType) (CreateAlertType) {
	target := CreateAlertType{
		Title:alert.Title,
		Description:alert.Description,
		QueryString:alert.QueryString,
		Filter:alert.Filter,
		Operation:alert.Operation,
		SeverityThresholdTiers:alert.SeverityThresholdTiers,
		SearchTimeFrameMinutes:alert.SearchTimeFrameMinutes,
		NotificationEmails:alert.NotificationEmails,
		IsEnabled:alert.IsEnabled,
		SuppressNotificationMinutes:alert.SuppressNotificationMinutes,
		ValueAggregationType:alert.ValueAggregationType,
		ValueAggregationField:alert.ValueAggregationField,
		GroupByAggregationFields:alert.GroupByAggregationFields,
		AlertNotificationEndpoints:alert.AlertNotificationEndpoints,
	}
	return target
}