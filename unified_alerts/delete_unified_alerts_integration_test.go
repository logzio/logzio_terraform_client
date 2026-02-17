package unified_alerts_test

import (
	"os"
	"testing"

	"github.com/logzio/logzio_terraform_client/unified_alerts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationUnifiedAlerts_DeleteAlert(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	createLogAlert := getCreateLogAlertType()
	createdAlert, err := underTest.CreateUnifiedAlert(unified_alerts.UrlTypeLogs, createLogAlert)
	require.NoError(t, err, "Failed to create log alert for delete test")
	require.NotNil(t, createdAlert, "Created alert should not be nil")

	deletedAlert, err := underTest.DeleteUnifiedAlert(unified_alerts.UrlTypeLogs, createdAlert.Id)
	require.NoError(t, err, "Failed to delete alert")
	require.NotNil(t, deletedAlert, "Deleted alert response should not be nil")
	assert.Equal(t, createdAlert.Id, deletedAlert.Id)

	_, err = underTest.GetUnifiedAlert(unified_alerts.UrlTypeLogs, createdAlert.Id)
	assert.Error(t, err)
}

func TestIntegrationUnifiedAlerts_DeleteMetricAlert(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}
	if os.Getenv("LOGZIO_UNIFIED_ACCOUNT_ID") == "" {
		t.Skip("LOGZIO_UNIFIED_ACCOUNT_ID not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	createMetricAlert := getCreateMetricAlertType()
	createdAlert, err := underTest.CreateUnifiedAlert(unified_alerts.UrlTypeMetrics, createMetricAlert)
	require.NoError(t, err, "Failed to create metric alert for delete test")
	require.NotNil(t, createdAlert, "Created metric alert should not be nil")

	deletedAlert, err := underTest.DeleteUnifiedAlert(unified_alerts.UrlTypeMetrics, createdAlert.Id)
	require.NoError(t, err, "Failed to delete metric alert")
	require.NotNil(t, deletedAlert, "Deleted metric alert response should not be nil")
	assert.Equal(t, createdAlert.Id, deletedAlert.Id)

	_, err = underTest.GetUnifiedAlert(unified_alerts.UrlTypeMetrics, createdAlert.Id)
	assert.Error(t, err)
}

func TestIntegrationUnifiedAlerts_DeleteAlertNotFound(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	_, err = underTest.DeleteUnifiedAlert(unified_alerts.UrlTypeLogs, "non-existent-id-12345")
	assert.Error(t, err)
}
