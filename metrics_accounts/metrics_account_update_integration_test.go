package metrics_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationMetricsAccount_UpdateMetricsAccount(t *testing.T) {
	underTest, email, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateMetricsAccount(email)

		subAccount, err := underTest.CreateMetricsAccount(createSubAccount)
		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			defer underTest.DeleteMetricsAccount(int64(subAccount.Id))
			time.Sleep(time.Second * 2)
			getSubAccount, err := underTest.GetMetricsAccount(int64(subAccount.Id))
			assert.NoError(t, err)
			assert.Equal(t, "tf_client_test", getSubAccount.AccountName)
			createSubAccount.AccountName = "test_after_update"
			time.Sleep(time.Second * 2)
			err = underTest.UpdateMetricsAccount(int64(subAccount.Id), createSubAccount)
			assert.NoError(t, err)
			// verify that the update was made
			time.Sleep(time.Second * 2)
			getSubAccount, err = underTest.GetMetricsAccount(int64(subAccount.Id))
			assert.NoError(t, err)
			assert.Equal(t, "test_after_update", getSubAccount.AccountName)
		}
	}
}

func TestIntegrationMetricsAccount_UpdateMetricsAccountIdNotExists(t *testing.T) {
	underTest, email, err := setupMetricsAccountsIntegrationTest()
	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateMetricsAccount(email)
		if assert.NoError(t, err) && assert.NotNil(t, createSubAccount) {
			err = underTest.UpdateMetricsAccount(int64(1234567), createSubAccount)
			assert.Error(t, err)
		}
	}
}
