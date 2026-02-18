package unified_alerts_test

import (
	"os"
	"testing"

	"github.com/logzio/logzio_terraform_client/unified_alerts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationUnifiedAlerts_GetAlert(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	// First create an alert
	createLogAlert := getCreateLogAlertType()
	createdAlert, err := underTest.CreateUnifiedAlert(unified_alerts.UrlTypeLogs, createLogAlert)
	require.NoError(t, err, "Failed to create log alert for get test")
	require.NotNil(t, createdAlert, "Created alert should not be nil")

	// Cleanup
	defer func() {
		_, deleteErr := underTest.DeleteUnifiedAlert(unified_alerts.UrlTypeLogs, createdAlert.Id)
		if deleteErr != nil {
			t.Logf("Failed to cleanup alert: %s", deleteErr)
		}
	}()

	// Now get the alert
	alert, err := underTest.GetUnifiedAlert(unified_alerts.UrlTypeLogs, createdAlert.Id)
	require.NoError(t, err, "Failed to get alert")
	require.NotNil(t, alert, "Retrieved alert should not be nil")
	assert.Equal(t, createdAlert.Id, alert.Id)
	assert.Equal(t, createdAlert.Title, alert.Title)
}

func TestIntegrationUnifiedAlerts_GetAlertNotFound(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	_, err = underTest.GetUnifiedAlert(unified_alerts.UrlTypeLogs, "non-existent-id-12345")
	assert.Error(t, err)
}
