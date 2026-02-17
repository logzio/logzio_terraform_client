package unified_alerts_test

import (
	"os"
	"testing"

	"github.com/logzio/logzio_terraform_client/unified_alerts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationUnifiedAlerts_UpdateAlert(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	createLogAlert := getCreateLogAlertType()
	createdAlert, err := underTest.CreateUnifiedAlert(unified_alerts.UrlTypeLogs, createLogAlert)
	require.NoError(t, err, "Failed to create log alert for update test")
	require.NotNil(t, createdAlert, "Created alert should not be nil")

	defer func() {
		_, deleteErr := underTest.DeleteUnifiedAlert(unified_alerts.UrlTypeLogs, createdAlert.Id)
		if deleteErr != nil {
			t.Logf("Failed to cleanup alert: %s", deleteErr)
		}
	}()

	updateAlert := getCreateLogAlertType()
	updateAlert.Title = "Updated Integration Test Alert"
	updateAlert.Description = "Updated description"

	updatedAlert, err := underTest.UpdateUnifiedAlert(unified_alerts.UrlTypeLogs, createdAlert.Id, updateAlert)
	require.NoError(t, err, "Failed to update alert")
	require.NotNil(t, updatedAlert, "Updated alert should not be nil")
	assert.Equal(t, createdAlert.Id, updatedAlert.Id)
	assert.Equal(t, "Updated Integration Test Alert", updatedAlert.Title)
	assert.Equal(t, "Updated description", updatedAlert.Description)
}

func TestIntegrationUnifiedAlerts_UpdateMetricAlert(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	createMetricAlert := getCreateMetricAlertType()
	createdAlert, err := underTest.CreateUnifiedAlert(unified_alerts.UrlTypeMetrics, createMetricAlert)
	require.NoError(t, err, "Failed to create metric alert for update test")
	require.NotNil(t, createdAlert, "Created metric alert should not be nil")

	updateAlert := getCreateMetricAlertType()
	updateAlert.Title = "Updated Integration Test Metric Alert"
	updateAlert.Description = "Updated metric alert description"

	updatedAlert, err := underTest.UpdateUnifiedAlert(unified_alerts.UrlTypeMetrics, createdAlert.Id, updateAlert)
	require.NoError(t, err, "Failed to update metric alert")
	require.NotNil(t, updatedAlert, "Updated metric alert should not be nil")
	assert.Equal(t, createdAlert.Id, updatedAlert.Id)
	assert.Equal(t, "Updated Integration Test Metric Alert", updatedAlert.Title)
	assert.Equal(t, "Updated metric alert description", updatedAlert.Description)
}

func TestIntegrationUnifiedAlerts_UpdateAlertNotFound(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	updateAlert := getCreateLogAlertType()
	_, err = underTest.UpdateUnifiedAlert(unified_alerts.UrlTypeLogs, "non-existent-id-12345", updateAlert)
	assert.Error(t, err)
}
