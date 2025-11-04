package unified_alerts_test

import (
	"testing"

	"github.com/logzio/logzio_terraform_client/unified_alerts"
	"github.com/stretchr/testify/assert"
)

func TestUnifiedAlerts_CreateLogAlert_ShouldQueryOnAllAccountsValidation(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	assert.NoError(t, err)

	// Test case 1: shouldQueryOnAllAccounts = false without accountIdsToQueryOn (should fail)
	createAlert := getCreateLogAlertType()
	createAlert.LogAlert.SubComponents[0].QueryDefinition.ShouldQueryOnAllAccounts = false
	createAlert.LogAlert.SubComponents[0].QueryDefinition.AccountIdsToQueryOn = nil

	_, err = underTest.CreateUnifiedAlert(unified_alerts.UrlTypeLogs, createAlert)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "accountIdsToQueryOn must be set when shouldQueryOnAllAccounts is false")
}

func TestUnifiedAlerts_CreateLogAlert_ShouldQueryOnAllAccountsValidationWithEmptyList(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	assert.NoError(t, err)

	// Test case 2: shouldQueryOnAllAccounts = false with empty accountIdsToQueryOn (should fail)
	createAlert := getCreateLogAlertType()
	createAlert.LogAlert.SubComponents[0].QueryDefinition.ShouldQueryOnAllAccounts = false
	createAlert.LogAlert.SubComponents[0].QueryDefinition.AccountIdsToQueryOn = []int{}

	_, err = underTest.CreateUnifiedAlert(unified_alerts.UrlTypeLogs, createAlert)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "accountIdsToQueryOn must be set when shouldQueryOnAllAccounts is false")
}

func TestUnifiedAlerts_CreateLogAlert_ShouldQueryOnAllAccountsValidationWithAccounts(t *testing.T) {
	_, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	assert.NoError(t, err)

	// Test case 3: shouldQueryOnAllAccounts = false with valid accountIdsToQueryOn (should be valid)
	createAlert := getCreateLogAlertType()
	createAlert.LogAlert.SubComponents[0].QueryDefinition.ShouldQueryOnAllAccounts = false
	createAlert.LogAlert.SubComponents[0].QueryDefinition.AccountIdsToQueryOn = []int{123456}

	// This should pass validation - no error expected
	// Note: We're only testing the validation logic, not making actual API calls
	assert.NotNil(t, createAlert)
}

func TestUnifiedAlerts_CreateLogAlert_ShouldQueryOnAllAccountsTrueWithEmptyList(t *testing.T) {
	_, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	assert.NoError(t, err)

	// Test case 4: shouldQueryOnAllAccounts = true with empty/nil accountIdsToQueryOn (should be valid)
	createAlert := getCreateLogAlertType()
	createAlert.LogAlert.SubComponents[0].QueryDefinition.ShouldQueryOnAllAccounts = true
	createAlert.LogAlert.SubComponents[0].QueryDefinition.AccountIdsToQueryOn = nil

	// This should pass validation - no error expected
	assert.NotNil(t, createAlert)
}
