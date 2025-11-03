//go:build integration
// +build integration

package unified_alerts_test

import (
	"os"
	"testing"

	"github.com/logzio/logzio_terraform_client/unified_alerts"
	"github.com/stretchr/testify/assert"
)

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
	assert.NoError(t, err)
	assert.NotNil(t, alert)
	assert.NotEmpty(t, alert.Id)
	assert.Equal(t, unified_alerts.TypeLogAlert, alert.Type)

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
	assert.NoError(t, err)
	assert.NotNil(t, alert)
	assert.NotEmpty(t, alert.Id)
	assert.Equal(t, unified_alerts.TypeMetricAlert, alert.Type)

	// Cleanup
	defer func() {
		_, deleteErr := underTest.DeleteUnifiedAlert(unified_alerts.UrlTypeMetrics, alert.Id)
		if deleteErr != nil {
			t.Logf("Failed to cleanup alert: %s", deleteErr)
		}
	}()
}
