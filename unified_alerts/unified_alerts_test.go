package unified_alerts_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/logzio/logzio_terraform_client/unified_alerts"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func fixture(path string) string {
	b, err := os.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setupUnifiedAlertsTest() (*unified_alerts.UnifiedAlertsClient, error, func()) {
	apiToken := "SOME_API_TOKEN"

	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	underTest, _ := unified_alerts.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func setupUnifiedAlertsIntegrationTest() (*unified_alerts.UnifiedAlertsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := unified_alerts.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, nil
}

func getCreateLogAlertType() unified_alerts.CreateUnifiedAlert {
	return unified_alerts.CreateUnifiedAlert{
		Title:       "Test Log Alert",
		Type:        unified_alerts.TypeLogAlert,
		Description: "Test log alert description",
		Tags:        []string{"test", "log"},
		FolderId:    "folder-123",
		LogAlert: &unified_alerts.LogAlertConfig{
			Output: unified_alerts.LogAlertOutput{
				Recipients: unified_alerts.Recipients{
					Emails:                  []string{"test@example.com"},
					NotificationEndpointIds: []int{1, 2},
				},
				SuppressNotificationsMinutes: 5,
				Type:                         unified_alerts.OutputTypeJson,
			},
			SearchTimeFrameMinutes: 15,
			SubComponents: []unified_alerts.SubComponent{
				{
					QueryDefinition: unified_alerts.QueryDefinition{
						Query: "loglevel:ERROR",
						Aggregation: unified_alerts.Aggregation{
							AggregationType: unified_alerts.AggregationTypeCount,
						},
						ShouldQueryOnAllAccounts: true,
					},
					Trigger: unified_alerts.SubComponentTrigger{
						Operator: unified_alerts.OperatorGreaterThan,
						SeverityThresholdTiers: map[string]float32{
							unified_alerts.SeverityHigh: 10.0,
							unified_alerts.SeverityInfo: 5.0,
						},
					},
					Output: unified_alerts.SubComponentOutput{
						ShouldUseAllFields: true,
					},
				},
			},
			Schedule: unified_alerts.Schedule{
				CronExpression: "0 */5 * * *",
				Timezone:       "UTC",
			},
		},
	}
}

func getCreateMetricAlertType() unified_alerts.CreateUnifiedAlert {
	return unified_alerts.CreateUnifiedAlert{
		Title:       "Test Metric Alert",
		Type:        unified_alerts.TypeMetricAlert,
		Description: "Test metric alert description",
		Tags:        []string{"test", "metric"},
		FolderId:    "folder-456",
		MetricAlert: &unified_alerts.MetricAlertConfig{
			Severity: unified_alerts.SeverityHigh,
			Trigger: unified_alerts.MetricTrigger{
				TriggerType:            unified_alerts.TriggerTypeThreshold,
				MetricOperator:         unified_alerts.MetricOperatorAbove,
				MinThreshold:           80.0,
				SearchTimeFrameMinutes: 5,
			},
			ConditionRefId: "A",
			Queries: []unified_alerts.MetricQuery{
				{
					RefId: "A",
					QueryDefinition: unified_alerts.MetricQueryDefinition{
						DatasourceUid: "prometheus-uid",
						PromqlQuery:   "up{job=\"api-server\"}",
					},
				},
			},
			Recipients: unified_alerts.Recipients{
				Emails:                  []string{"alerts@example.com"},
				NotificationEndpointIds: []int{3, 4},
			},
		},
	}
}

func TestNewWithEmptyBaseUrl(t *testing.T) {
	_, err := unified_alerts.New("any-api-token", "")
	if err == nil {
		t.Fatal("Expected error when baseUrl is empty")
	}
	if err.Error() != "Base URL not defined" {
		t.Fatalf("The expected error message to be '%s' but was '%s'",
			"Base URL not defined", err.Error())
	}
}

func TestNewWithEmptyApiToken(t *testing.T) {
	_, err := unified_alerts.New("", "any-base-url")

	if err == nil {
		t.Fatal("Expected error when API token is empty")
	}
	if err.Error() != "API token not defined" {
		t.Fatalf("The expected error message to be '%s' but was '%s'",
			"API token not defined", err.Error())
	}
}
