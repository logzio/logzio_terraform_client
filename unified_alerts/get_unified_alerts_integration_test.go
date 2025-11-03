//go:build integration
// +build integration

package unified_alerts_test

import (
	"os"
	"testing"

	"github.com/logzio/logzio_terraform_client/unified_alerts"
	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)
	assert.NotNil(t, createdAlert)

	// Cleanup
	defer func() {
		_, deleteErr := underTest.DeleteUnifiedAlert(unified_alerts.UrlTypeLogs, createdAlert.Id)
		if deleteErr != nil {
			t.Logf("Failed to cleanup alert: %s", deleteErr)
		}
	}()

	// Now get the alert
	alert, err := underTest.GetUnifiedAlert(unified_alerts.UrlTypeLogs, createdAlert.Id)
	assert.NoError(t, err)
	assert.NotNil(t, alert)
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
