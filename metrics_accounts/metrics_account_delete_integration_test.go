package metrics_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationMetricsAccount_DeleteMetricsAccount(t *testing.T) {
	underTest, email, err := setupMetricsAccountsIntegrationTest()
	if assert.NoError(t, err) {
		createMetricsAccount := getCreateOrUpdateMetricsAccount(email)
		createMetricsAccount.AccountName = createMetricsAccount.AccountName + "_delete"
		metricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
		if assert.NoError(t, err) && assert.NotNil(t, metricsAccount) {
			time.Sleep(2 * time.Second)
			defer func() {
				err = underTest.DeleteMetricsAccount(int64(metricsAccount.Id))
				assert.NoError(t, err)
			}()
		}
	}
}

func TestIntegrationMetricsAccount_DeleteMetricsAccountNotExists(t *testing.T) {
	underTest, _, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		err = underTest.DeleteMetricsAccount(int64(1234567))
		assert.Error(t, err)
	}
}
