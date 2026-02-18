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

	createAlert := getCreateLogAlertType()
	createAlert.AlertConfiguration.SubComponents[0].QueryDefinition.ShouldQueryOnAllAccounts = false
	createAlert.AlertConfiguration.SubComponents[0].QueryDefinition.AccountIdsToQueryOn = nil

	_, err = underTest.CreateUnifiedAlert(unified_alerts.UrlTypeLogs, createAlert)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "accountIdsToQueryOn must be set when shouldQueryOnAllAccounts is false")
}

func TestUnifiedAlerts_CreateLogAlert_ShouldQueryOnAllAccountsValidationWithEmptyList(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	assert.NoError(t, err)

	createAlert := getCreateLogAlertType()
	createAlert.AlertConfiguration.SubComponents[0].QueryDefinition.ShouldQueryOnAllAccounts = false
	createAlert.AlertConfiguration.SubComponents[0].QueryDefinition.AccountIdsToQueryOn = []int{}

	_, err = underTest.CreateUnifiedAlert(unified_alerts.UrlTypeLogs, createAlert)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "accountIdsToQueryOn must be set when shouldQueryOnAllAccounts is false")
}

func TestUnifiedAlerts_CreateLogAlert_ShouldQueryOnAllAccountsValidationWithAccounts(t *testing.T) {
	_, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	assert.NoError(t, err)

	createAlert := getCreateLogAlertType()
	createAlert.AlertConfiguration.SubComponents[0].QueryDefinition.ShouldQueryOnAllAccounts = false
	createAlert.AlertConfiguration.SubComponents[0].QueryDefinition.AccountIdsToQueryOn = []int{123456}

	assert.NotNil(t, createAlert)
}

func TestUnifiedAlerts_CreateLogAlert_ShouldQueryOnAllAccountsTrueWithEmptyList(t *testing.T) {
	_, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	assert.NoError(t, err)

	createAlert := getCreateLogAlertType()
	createAlert.AlertConfiguration.SubComponents[0].QueryDefinition.ShouldQueryOnAllAccounts = true
	createAlert.AlertConfiguration.SubComponents[0].QueryDefinition.AccountIdsToQueryOn = nil

	assert.NotNil(t, createAlert)
}
