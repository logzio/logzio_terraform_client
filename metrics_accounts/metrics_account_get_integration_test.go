package metrics_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationMetricsAccount_GetMetricsAccount(t *testing.T) {
	underTest, email, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateMetricsAccount(email)
		createSubAccount.AccountName = createSubAccount.AccountName + "_get"

		subAccount, err := underTest.CreateMetricsAccount(createSubAccount)
		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			defer underTest.DeleteMetricsAccount(int64(subAccount.Id))
			time.Sleep(4 * time.Second)
			getSubAccount, err := underTest.GetMetricsAccount(int64(subAccount.Id))
			assert.NoError(t, err)
			assert.NotNil(t, getSubAccount)
			assert.Equal(t, subAccount.Id, getSubAccount.Id)
			assert.Equal(t, createSubAccount.AccountName, getSubAccount.AccountName)
		}
	}
}

func TestIntegrationMetricsAccount_GetMetricsAccountNotExists(t *testing.T) {
	underTest, _, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		subAccount, err := underTest.GetMetricsAccount(int64(1234567))
		assert.Error(t, err)
		assert.Nil(t, subAccount)
	}
}
