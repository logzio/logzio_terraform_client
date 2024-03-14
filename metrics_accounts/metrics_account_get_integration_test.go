package metrics_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationMetricsAccount_GetMetricsAccount(t *testing.T) {
	underTest, email, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createMetricsAccount := getCreateOrUpdateMetricsAccount(email)
		createMetricsAccount.AccountName = createMetricsAccount.AccountName + "_get"

		metricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
		if assert.NoError(t, err) && assert.NotNil(t, metricsAccount) {
			defer underTest.DeleteMetricsAccount(int64(metricsAccount.Id))
			time.Sleep(4 * time.Second)
			getMetricsAccount, err := underTest.GetMetricsAccount(int64(metricsAccount.Id))
			assert.NoError(t, err)
			assert.NotNil(t, getMetricsAccount)
			assert.Equal(t, metricsAccount.Id, getMetricsAccount.Id)
			assert.Equal(t, createMetricsAccount.AccountName, getMetricsAccount.AccountName)
		}
	}
}

func TestIntegrationMetricsAccount_GetMetricsAccountNotExists(t *testing.T) {
	underTest, _, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		metricsAccount, err := underTest.GetMetricsAccount(int64(1234567))
		assert.Error(t, err)
		assert.Nil(t, metricsAccount)
	}
}
