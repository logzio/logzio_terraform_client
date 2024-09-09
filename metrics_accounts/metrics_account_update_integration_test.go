package metrics_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationMetricsAccount_UpdateMetricsAccount(t *testing.T) {
	underTest, email, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createMetricsAccount := getCreateOrUpdateMetricsAccount(email)
		metricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
		if assert.NoError(t, err) && assert.NotNil(t, metricsAccount) {
			defer underTest.DeleteMetricsAccount(int64(metricsAccount.Id))
			time.Sleep(time.Second * 2)
			getMetricsAccount, err := underTest.GetMetricsAccount(int64(metricsAccount.Id))
			assert.NoError(t, err)
			assert.Equal(t, "tf_client_test", getMetricsAccount.AccountName)
			createMetricsAccount.AccountName = "test_after_update"
			createMetricsAccount.PlanUts = nil // make sure update works without affecting curr UTS
			time.Sleep(time.Second * 2)
			err = underTest.UpdateMetricsAccount(int64(metricsAccount.Id), createMetricsAccount)
			assert.NoError(t, err)
			// verify that the update was made
			time.Sleep(time.Second * 2)
			getMetricsAccount, err = underTest.GetMetricsAccount(int64(metricsAccount.Id))
			assert.NoError(t, err)
			assert.Equal(t, "test_after_update", getMetricsAccount.AccountName)
		}
	}
}
func TestIntegrationMetricsAccount_UpdateMetricsAccountPlanUts(t *testing.T) {
	underTest, email, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createMetricsAccount := getCreateOrUpdateMetricsAccount(email)
		metricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
		if assert.NoError(t, err) && assert.NotNil(t, metricsAccount) {
			defer underTest.DeleteMetricsAccount(int64(metricsAccount.Id))
			time.Sleep(time.Second * 2)
			// Update the plan_uts field without changing the account_name
			createMetricsAccount.PlanUts = new(int32)
			*createMetricsAccount.PlanUts = 200
			err = underTest.UpdateMetricsAccount(int64(metricsAccount.Id), createMetricsAccount)
			assert.Error(t, err)
			// Check if the error message is as expected
			assert.Contains(t, err.Error(), "Sub account with name (TEST_AUTOMATION_ACCOUNT) already exists on account")
		}
	}
}

func TestIntegrationMetricsAccount_UpdateMetricsAccountWithAuthorizedAccounts(t *testing.T) {
	underTest, email, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createMetricsAccount := getCreateOrUpdateMetricsAccount(email)
		// Add an authorized account
		createMetricsAccount.AuthorizedAccountsIds = append(createMetricsAccount.AuthorizedAccountsIds, 123456)
		metricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
		if assert.NoError(t, err) && assert.NotNil(t, metricsAccount) {
			defer underTest.DeleteMetricsAccount(int64(metricsAccount.Id))
			time.Sleep(time.Second * 2)
			// Update the plan_uts field
			createMetricsAccount.PlanUts = new(int32)
			*createMetricsAccount.PlanUts = 200
			err = underTest.UpdateMetricsAccount(int64(metricsAccount.Id), createMetricsAccount)
			assert.Error(t, err)
			// Check if the error message is as expected
			assert.Contains(t, err.Error(), "Sub account is not found for the owner account")
		}
	}
}

func TestIntegrationMetricsAccount_UpdateMetricsAccountIdNotExists(t *testing.T) {
	underTest, email, err := setupMetricsAccountsIntegrationTest()
	if assert.NoError(t, err) {
		createMetricsAccount := getCreateOrUpdateMetricsAccount(email)
		if assert.NoError(t, err) && assert.NotNil(t, createMetricsAccount) {
			err = underTest.UpdateMetricsAccount(int64(1234567), createMetricsAccount)
			assert.Error(t, err)
		}
	}
}
