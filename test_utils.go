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

func alertToCreateAlert(alert *AlertType) CreateAlertType {
	target := CreateAlertType{
		AlertNotificationEndpoints:   alert.AlertNotificationEndpoints,
		Description:                  alert.Description,
		Filter:                       alert.Filter,
		GroupByAggregationFields:     alert.GroupByAggregationFields,
		IsEnabled:                    alert.IsEnabled,
		NotificationEmails:           alert.NotificationEmails,
		Operation:                    alert.Operation,
		QueryString:                  alert.QueryString,
		SearchTimeFrameMinutes:       alert.SearchTimeFrameMinutes,
		SeverityThresholdTiers:       alert.SeverityThresholdTiers,
		SuppressNotificationsMinutes: alert.SuppressNotificationsMinutes,
		Title:                        alert.Title,
		ValueAggregationType:         alert.ValueAggregationType,
		ValueAggregationField:        alert.ValueAggregationField,
	}
	return target
}
