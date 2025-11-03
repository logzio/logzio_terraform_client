//go:build integration
// +build integration

package unified_alerts_test

import (
	"os"
	"testing"

	"github.com/logzio/logzio_terraform_client/unified_alerts"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedAlerts_DeleteAlert(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedAlertsIntegrationTest()
	if err != nil {
		t.Fatalf("setupUnifiedAlertsIntegrationTest failed: %s", err)
	}

	// First create an alert
	createLogAlert := getCreateLogAlertType()
	createdAlert, err := underTest.CreateUnifiedAlert(createLogAlert)
	assert.NoError(t, err)
	assert.NotNil(t, createdAlert)

	// Delete the alert
	deletedAlert, err := underTest.DeleteUnifiedAlert(unified_alerts.UrlTypeLogs, createdAlert.Id)
	assert.NoError(t, err)
	assert.NotNil(t, deletedAlert)
	assert.Equal(t, createdAlert.Id, deletedAlert.Id)

	// Verify it's deleted by trying to get it (should fail)
	_, err = underTest.GetUnifiedAlert(unified_alerts.UrlTypeLogs, createdAlert.Id)
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
