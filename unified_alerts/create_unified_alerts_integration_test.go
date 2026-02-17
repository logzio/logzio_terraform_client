//go:build integration
// +build integration

package unified_alerts_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_alerts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test configuration constants
var (
	testFolderId    = os.Getenv("LOGZIO_UNIFIED_FOLDER_ID")
	testDashboardId = os.Getenv("LOGZIO_UNIFIED_DASHBOARD_ID")
	testPanelId     = os.Getenv("LOGZIO_UNIFIED_PANEL_ID")
	testAccountId   = getTestAccountId()
)

func getTestAccountId() int32 {
	accountIdStr := os.Getenv("LOGZIO_UNIFIED_ACCOUNT_ID")
	if accountIdStr == "" {
		return 0
	}
	accountId, err := strconv.Atoi(accountIdStr)
	if err != nil {
		return 0
	}
	return int32(accountId)
}

// getTestNotificationEndpoint returns the notification endpoint ID as an integer
func getTestNotificationEndpoint() int {
	endpointStr := os.Getenv("LOGZIO_UNIFIED_NOTIFICATION_ENDPOINT_ID")
	if endpointStr == "" {
		return 0
	}
	endpoint, err := strconv.Atoi(endpointStr)
	if err != nil {
		return 0
	}
	return endpoint
}

// generateUniqueTitle generates a unique title with timestamp
func generateUniqueTitle(base string) string {
	return fmt.Sprintf("%s - %d", base, time.Now().Unix())
}

func TestIntegrationUnifiedAlerts_CreateLogAlert(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	createLogAlert := getCreateLogAlertType()

	alert, err := underTest.CreateUnifiedAlert(unified_alerts.UrlTypeLogs, createLogAlert)
	require.NoError(t, err, "Failed to create log alert")
	require.NotNil(t, alert, "Alert should not be nil")
	assert.NotEmpty(t, alert.Id)
	assert.NotNil(t, alert.AlertConfiguration)
	assert.Equal(t, unified_alerts.TypeLogAlert, alert.AlertConfiguration.Type)

	// Cleanup
	defer func() {
		_, deleteErr := underTest.DeleteUnifiedAlert(unified_alerts.UrlTypeLogs, alert.Id)
		if deleteErr != nil {
			t.Logf("Failed to cleanup alert: %s", deleteErr)
		}
	}()
}

func TestIntegrationUnifiedAlerts_CreateMetricAlert(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	createMetricAlert := getCreateMetricAlertType()

	alert, err := underTest.CreateUnifiedAlert(unified_alerts.UrlTypeMetrics, createMetricAlert)
	require.NoError(t, err, "Failed to create metric alert")
	require.NotNil(t, alert, "Alert should not be nil")
	assert.NotEmpty(t, alert.Id)
	assert.NotNil(t, alert.AlertConfiguration)
	assert.Equal(t, unified_alerts.TypeMetricAlert, alert.AlertConfiguration.Type)
}

// TestIntegrationUnifiedAlerts_CreateMetricAlert_OneQueryNoAI tests creating a metric alert with one query and no AI
func TestIntegrationUnifiedAlerts_CreateMetricAlert_OneQueryNoAI(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	threshold := 80.0
	// One query with simple promql, trigger if bigger than X
	createMetricAlert := unified_alerts.CreateUnifiedAlert{
		Title:       generateUniqueTitle("Metric Alert - One Query No AI"),
		Description: "Alert when CPU usage exceeds threshold",
		Tags:        []string{"cpu", "infrastructure", "test"},
		LinkedPanel: &unified_alerts.LinkedPanel{
			FolderId:    testFolderId,
			DashboardId: testDashboardId,
			PanelId:     testPanelId,
		},
		Recipients: &unified_alerts.Recipients{
			Emails:                  []string{"alerts@example.com"},
			NotificationEndpointIds: []int{getTestNotificationEndpoint()},
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
						PromqlQuery: "avg(cpu_usage_percent)",
					},
				},
			},
		},
	}

	alert, err := underTest.CreateUnifiedAlert(unified_alerts.UrlTypeMetrics, createMetricAlert)
	require.NoError(t, err, "Failed to create one query no AI metric alert")
	require.NotNil(t, alert, "Alert should not be nil")
	assert.NotEmpty(t, alert.Id)
	assert.NotNil(t, alert.AlertConfiguration)
	assert.Equal(t, unified_alerts.TypeMetricAlert, alert.AlertConfiguration.Type)
	assert.Contains(t, alert.Title, "Metric Alert - One Query No AI")
	t.Logf("Created metric alert (one query, no AI) with ID: %s, Title: %s", alert.Id, alert.Title)
}

// TestIntegrationUnifiedAlerts_CreateMetricAlert_TwoQueriesNoAI tests creating a metric alert with two queries and no AI
func TestIntegrationUnifiedAlerts_CreateMetricAlert_TwoQueriesNoAI(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	// Two queries with math expression $A > $B
	createMetricAlert := unified_alerts.CreateUnifiedAlert{
		Title:       generateUniqueTitle("Metric Alert - Two Queries No AI"),
		Description: "Alert when request rate exceeds error rate threshold",
		Tags:        []string{"requests", "errors", "test"},
		LinkedPanel: &unified_alerts.LinkedPanel{
			FolderId:    testFolderId,
			DashboardId: testDashboardId,
			PanelId:     testPanelId,
		},
		Recipients: &unified_alerts.Recipients{
			Emails:                  []string{"team@example.com"},
			NotificationEndpointIds: []int{getTestNotificationEndpoint()},
		},
		AlertConfiguration: &unified_alerts.AlertConfiguration{
			Type:     unified_alerts.TypeMetricAlert,
			Severity: unified_alerts.SeverityMedium,
			Trigger: &unified_alerts.MetricAlertTrigger{
				Type:       unified_alerts.TriggerTypeMath,
				Expression: "$A > $B",
			},
			Queries: []unified_alerts.MetricQuery{
				{
					RefId: "A",
					QueryDefinition: unified_alerts.MetricQueryDefinition{
						AccountId:   testAccountId,
						PromqlQuery: "rate(http_requests_total[5m])",
					},
				},
				{
					RefId: "B",
					QueryDefinition: unified_alerts.MetricQueryDefinition{
						AccountId:   testAccountId,
						PromqlQuery: "rate(http_errors_total[5m])",
					},
				},
			},
		},
	}

	alert, err := underTest.CreateUnifiedAlert(unified_alerts.UrlTypeMetrics, createMetricAlert)
	require.NoError(t, err, "Failed to create two queries no AI metric alert")
	require.NotNil(t, alert, "Alert should not be nil")
	assert.NotEmpty(t, alert.Id)
	assert.NotNil(t, alert.AlertConfiguration)
	assert.Equal(t, unified_alerts.TypeMetricAlert, alert.AlertConfiguration.Type)
	assert.Contains(t, alert.Title, "Metric Alert - Two Queries No AI")
	t.Logf("Created metric alert (two queries, no AI) with ID: %s, Title: %s", alert.Id, alert.Title)
}

// TestIntegrationUnifiedAlerts_CreateMetricAlert_OneQueryWithAI tests creating a metric alert with one query and AI enabled
func TestIntegrationUnifiedAlerts_CreateMetricAlert_OneQueryWithAI(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	threshold := 90.0
	// One query with AI enabled, same notification endpoints for alert and AI
	createMetricAlert := unified_alerts.CreateUnifiedAlert{
		Title:       generateUniqueTitle("Metric Alert - One Query With AI"),
		Description: "Alert with AI analysis when memory usage is high",
		Tags:        []string{"memory", "ai", "test"},
		LinkedPanel: &unified_alerts.LinkedPanel{
			FolderId:    testFolderId,
			DashboardId: testDashboardId,
			PanelId:     testPanelId,
		},
		Rca:                                 true,
		UseAlertNotificationEndpointsForRca: true,
		Recipients: &unified_alerts.Recipients{
			Emails:                  []string{"oncall@example.com"},
			NotificationEndpointIds: []int{getTestNotificationEndpoint()},
		},
		AlertConfiguration: &unified_alerts.AlertConfiguration{
			Type:     unified_alerts.TypeMetricAlert,
			Severity: unified_alerts.SeveritySevere,
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
						PromqlQuery: "avg(memory_usage_percent)",
					},
				},
			},
		},
	}

	alert, err := underTest.CreateUnifiedAlert(unified_alerts.UrlTypeMetrics, createMetricAlert)
	require.NoError(t, err, "Failed to create one query with AI metric alert")
	require.NotNil(t, alert, "Alert should not be nil")
	assert.NotEmpty(t, alert.Id)
	assert.NotNil(t, alert.AlertConfiguration)
	assert.Equal(t, unified_alerts.TypeMetricAlert, alert.AlertConfiguration.Type)
	assert.Contains(t, alert.Title, "Metric Alert - One Query With AI")
	assert.True(t, alert.Rca)
	assert.True(t, alert.UseAlertNotificationEndpointsForRca)
	t.Logf("Created metric alert (one query, with AI) with ID: %s, Title: %s", alert.Id, alert.Title)
}

// TestIntegrationUnifiedAlerts_CreateMetricAlert_TwoQueriesWithAI tests creating a metric alert with two queries and AI enabled
func TestIntegrationUnifiedAlerts_CreateMetricAlert_TwoQueriesWithAI(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	// Two queries with math expression and AI enabled
	createMetricAlert := unified_alerts.CreateUnifiedAlert{
		Title:       generateUniqueTitle("Metric Alert - Two Queries With AI"),
		Description: "Alert with AI when disk usage exceeds available space",
		Tags:        []string{"disk", "storage", "ai", "test"},
		LinkedPanel: &unified_alerts.LinkedPanel{
			FolderId:    testFolderId,
			DashboardId: testDashboardId,
			PanelId:     testPanelId,
		},
		Rca:                                 true,
		UseAlertNotificationEndpointsForRca: true,
		Recipients: &unified_alerts.Recipients{
			Emails:                  []string{"storage-team@example.com"},
			NotificationEndpointIds: []int{getTestNotificationEndpoint()},
		},
		AlertConfiguration: &unified_alerts.AlertConfiguration{
			Type:     unified_alerts.TypeMetricAlert,
			Severity: unified_alerts.SeverityHigh,
			Trigger: &unified_alerts.MetricAlertTrigger{
				Type:       unified_alerts.TriggerTypeMath,
				Expression: "$A > $B",
			},
			Queries: []unified_alerts.MetricQuery{
				{
					RefId: "A",
					QueryDefinition: unified_alerts.MetricQueryDefinition{
						AccountId:   testAccountId,
						PromqlQuery: "disk_used_bytes",
					},
				},
				{
					RefId: "B",
					QueryDefinition: unified_alerts.MetricQueryDefinition{
						AccountId:   testAccountId,
						PromqlQuery: "disk_available_bytes * 0.9",
					},
				},
			},
		},
	}

	alert, err := underTest.CreateUnifiedAlert(unified_alerts.UrlTypeMetrics, createMetricAlert)
	require.NoError(t, err, "Failed to create two queries with AI metric alert")
	require.NotNil(t, alert, "Alert should not be nil")
	assert.NotEmpty(t, alert.Id)
	assert.NotNil(t, alert.AlertConfiguration)
	assert.Equal(t, unified_alerts.TypeMetricAlert, alert.AlertConfiguration.Type)
	assert.Contains(t, alert.Title, "Metric Alert - Two Queries With AI")
	assert.True(t, alert.Rca)
	assert.True(t, alert.UseAlertNotificationEndpointsForRca)
	t.Logf("Created metric alert (two queries, with AI) with ID: %s, Title: %s", alert.Id, alert.Title)
}

// TestIntegrationUnifiedAlerts_CreateMetricAlert_OneQueryWithAI_DifferentEndpoints tests creating a metric alert with one query and AI with different notification endpoints
func TestIntegrationUnifiedAlerts_CreateMetricAlert_OneQueryWithAI_DifferentEndpoints(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	threshold := 100.0
	// One query with AI enabled, different notification endpoints for alert and AI
	createMetricAlert := unified_alerts.CreateUnifiedAlert{
		Title:       generateUniqueTitle("Metric Alert - One Query With AI Different Endpoints"),
		Description: "Alert with AI analysis using different endpoints when network latency is high",
		Tags:        []string{"network", "latency", "ai", "test"},
		LinkedPanel: &unified_alerts.LinkedPanel{
			FolderId:    testFolderId,
			DashboardId: testDashboardId,
			PanelId:     testPanelId,
		},
		Rca:                        true,
		RcaNotificationEndpointIds: []int{getTestNotificationEndpoint()},
		Recipients: &unified_alerts.Recipients{
			Emails:                  []string{"network-team@example.com"},
			NotificationEndpointIds: []int{getTestNotificationEndpoint()},
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
						PromqlQuery: "avg(network_latency_ms)",
					},
				},
			},
		},
	}

	alert, err := underTest.CreateUnifiedAlert(unified_alerts.UrlTypeMetrics, createMetricAlert)
	require.NoError(t, err, "Failed to create one query with AI (different endpoints) metric alert")
	require.NotNil(t, alert, "Alert should not be nil")
	assert.NotEmpty(t, alert.Id)
	assert.NotNil(t, alert.AlertConfiguration)
	assert.Equal(t, unified_alerts.TypeMetricAlert, alert.AlertConfiguration.Type)
	assert.Contains(t, alert.Title, "Metric Alert - One Query With AI Different Endpoints")
	assert.True(t, alert.Rca)
	assert.False(t, alert.UseAlertNotificationEndpointsForRca)
	assert.NotEmpty(t, alert.RcaNotificationEndpointIds)
	t.Logf("Created metric alert (one query, with AI, different endpoints) with ID: %s, Title: %s", alert.Id, alert.Title)
}

// TestIntegrationUnifiedAlerts_CreateMetricAlert_TwoQueriesWithAI_DifferentEndpoints tests creating a metric alert with two queries and AI with different notification endpoints
func TestIntegrationUnifiedAlerts_CreateMetricAlert_TwoQueriesWithAI_DifferentEndpoints(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	// Two queries with math expression and AI enabled, different notification endpoints
	createMetricAlert := unified_alerts.CreateUnifiedAlert{
		Title:       generateUniqueTitle("Metric Alert - Two Queries With AI Different Endpoints"),
		Description: "Alert with AI using different endpoints when response time exceeds error threshold",
		Tags:        []string{"performance", "response-time", "ai", "test"},
		LinkedPanel: &unified_alerts.LinkedPanel{
			FolderId:    testFolderId,
			DashboardId: testDashboardId,
			PanelId:     testPanelId,
		},
		Rca:                        true,
		RcaNotificationEndpointIds: []int{getTestNotificationEndpoint()},
		Recipients: &unified_alerts.Recipients{
			Emails:                  []string{"performance-team@example.com"},
			NotificationEndpointIds: []int{getTestNotificationEndpoint()},
		},
		AlertConfiguration: &unified_alerts.AlertConfiguration{
			Type:     unified_alerts.TypeMetricAlert,
			Severity: unified_alerts.SeverityMedium,
			Trigger: &unified_alerts.MetricAlertTrigger{
				Type:       unified_alerts.TriggerTypeMath,
				Expression: "$A > $B",
			},
			Queries: []unified_alerts.MetricQuery{
				{
					RefId: "A",
					QueryDefinition: unified_alerts.MetricQueryDefinition{
						AccountId:   testAccountId,
						PromqlQuery: "avg(http_response_time_ms)",
					},
				},
				{
					RefId: "B",
					QueryDefinition: unified_alerts.MetricQueryDefinition{
						AccountId:   testAccountId,
						PromqlQuery: "scalar(http_response_time_threshold_ms)",
					},
				},
			},
		},
	}

	alert, err := underTest.CreateUnifiedAlert(unified_alerts.UrlTypeMetrics, createMetricAlert)
	require.NoError(t, err, "Failed to create two queries with AI (different endpoints) metric alert")
	require.NotNil(t, alert, "Alert should not be nil")
	assert.NotEmpty(t, alert.Id)
	assert.NotNil(t, alert.AlertConfiguration)
	assert.Equal(t, unified_alerts.TypeMetricAlert, alert.AlertConfiguration.Type)
	assert.Contains(t, alert.Title, "Metric Alert - Two Queries With AI Different Endpoints")
	assert.True(t, alert.Rca)
	assert.False(t, alert.UseAlertNotificationEndpointsForRca)
	assert.NotEmpty(t, alert.RcaNotificationEndpointIds)
	t.Logf("Created metric alert (two queries, with AI, different endpoints) with ID: %s, Title: %s", alert.Id, alert.Title)
}
