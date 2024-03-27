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
