package unified_alerts_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/logzio/logzio_terraform_client/unified_alerts"
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
		Title:       "TF client Log Alert",
		Description: "Test log alert description",
		Tags:        []string{"test", "log"},
		Recipients: &unified_alerts.Recipients{
			Emails:                  []string{"test@example.com"},
			NotificationEndpointIds: []int{},
		},
		AlertConfiguration: &unified_alerts.AlertConfiguration{
			Type:                         unified_alerts.TypeLogAlert,
			SuppressNotificationsMinutes: 5,
			AlertOutputTemplateType:      unified_alerts.OutputTypeJson,
			SearchTimeFrameMinutes:       15,
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
			Schedule: &unified_alerts.Schedule{
				CronExpression: "0 0 0 * * ?",
				Timezone:       "UTC",
			},
		},
	}
}

func getCreateMetricAlertType() unified_alerts.CreateUnifiedAlert {
	uniqueId := fmt.Sprintf("%d", time.Now().UnixNano())
	threshold := 80.0
	return unified_alerts.CreateUnifiedAlert{
		Title:       "TF client Metric Alert" + uniqueId,
		Description: "Test metric alert description",
		Tags:        []string{"test", "metric"},
		Recipients: &unified_alerts.Recipients{
			Emails:                  []string{"alerts@example.com"},
			NotificationEndpointIds: []int{},
		},
		AlertConfiguration: &unified_alerts.AlertConfiguration{
			Type:     unified_alerts.TypeMetricAlert,
			Severity: unified_alerts.SeverityHigh,
			Trigger: &unified_alerts.MetricAlertTrigger{
				Type: unified_alerts.TriggerTypeThreshold,
				Condition: &unified_alerts.TriggerCondition{
					OperatorType: unified_alerts.OperatorTypeAbove,
					Threshold:    &threshold,
				},
			},
			Queries: []unified_alerts.MetricQuery{
				{
					RefId: "A",
					QueryDefinition: unified_alerts.MetricQueryDefinition{
						AccountId:   testAccountId,
						PromqlQuery: "up{job=\"api-server\"}",
					},
				},
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
