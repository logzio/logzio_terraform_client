package metrics_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationMetricsAccount_DeleteMetricsAccount(t *testing.T) {
	underTest, email, err := setupMetricsAccountsIntegrationTest()
	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateMetricsAccount(email)
		createSubAccount.AccountName = createSubAccount.AccountName + "_delete"
		subAccount, err := underTest.CreateMetricsAccount(createSubAccount)
		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			time.Sleep(2 * time.Second)
			defer func() {
				err = underTest.DeleteMetricsAccount(int64(subAccount.Id))
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
