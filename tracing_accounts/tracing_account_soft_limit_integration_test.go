package tracing_accounts_test

import (
	"github.com/logzio/logzio_terraform_client/tracing_accounts"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationTracingAccount_GetTracingAccountSoftLimit(t *testing.T) {
	underTest, email, err := setupTracingAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createTracingAccount := getCreateOrUpdateTracingAccount(email)
		createTracingAccount.AccountName = createTracingAccount.AccountName + "_sl_get"

		tracingAccount, err := underTest.CreateTracingAccount(createTracingAccount)
		if assert.NoError(t, err) && assert.NotNil(t, tracingAccount) {
			defer underTest.DeleteTracingAccount(int64(tracingAccount.AccountId))
			time.Sleep(4 * time.Second)

			softLimit, err := underTest.GetTracingAccountSoftLimit(int64(tracingAccount.AccountId))
			assert.NoError(t, err)
			if assert.NotNil(t, softLimit) {
				assert.Equal(t, tracingAccount.AccountId, softLimit.AccountId)
				// the soft limit is maxDailyGB under another name, so create already set it
				if assert.NotNil(t, softLimit.SoftLimitGB) {
					assert.Equal(t, createTracingAccount.MaxDailyGB, *softLimit.SoftLimitGB)
				}
			}
		}
	}
}

func TestIntegrationTracingAccount_UpdateTracingAccountSoftLimit(t *testing.T) {
	underTest, email, err := setupTracingAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createTracingAccount := getCreateOrUpdateTracingAccount(email)
		createTracingAccount.AccountName = createTracingAccount.AccountName + "_sl_update"

		tracingAccount, err := underTest.CreateTracingAccount(createTracingAccount)
		if assert.NoError(t, err) && assert.NotNil(t, tracingAccount) {
			defer underTest.DeleteTracingAccount(int64(tracingAccount.AccountId))
			time.Sleep(4 * time.Second)

			updated, err := underTest.UpdateTracingAccountSoftLimit(int64(tracingAccount.AccountId),
				tracing_accounts.UpdateTracingAccountSoftLimit{
					TracingAccountId: tracingAccount.AccountId,
					SoftLimitGB:      2,
				})
			assert.NoError(t, err)
			if assert.NotNil(t, updated) && assert.NotNil(t, updated.SoftLimitGB) {
				assert.Equal(t, float32(2), *updated.SoftLimitGB)
			}

			// verify that the update was made
			time.Sleep(time.Second * 2)
			softLimit, err := underTest.GetTracingAccountSoftLimit(int64(tracingAccount.AccountId))
			assert.NoError(t, err)
			if assert.NotNil(t, softLimit) && assert.NotNil(t, softLimit.SoftLimitGB) {
				assert.Equal(t, float32(2), *softLimit.SoftLimitGB)
			}

			// the soft limit and maxDailyGB are the same underlying value
			getTracingAccount, err := underTest.GetTracingAccount(int64(tracingAccount.AccountId))
			assert.NoError(t, err)
			if assert.NotNil(t, getTracingAccount) && assert.NotNil(t, getTracingAccount.MaxDailyGB) {
				assert.Equal(t, float32(2), *getTracingAccount.MaxDailyGB)
			}
		}
	}
}

func TestIntegrationTracingAccount_UpdateTracingAccountSoftLimitMismatchedId(t *testing.T) {
	underTest, email, err := setupTracingAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createTracingAccount := getCreateOrUpdateTracingAccount(email)
		createTracingAccount.AccountName = createTracingAccount.AccountName + "_sl_mismatch"

		tracingAccount, err := underTest.CreateTracingAccount(createTracingAccount)
		if assert.NoError(t, err) && assert.NotNil(t, tracingAccount) {
			defer underTest.DeleteTracingAccount(int64(tracingAccount.AccountId))
			time.Sleep(4 * time.Second)

			// the API rejects a body id that does not match the path id
			softLimit, err := underTest.UpdateTracingAccountSoftLimit(int64(tracingAccount.AccountId),
				tracing_accounts.UpdateTracingAccountSoftLimit{
					TracingAccountId: tracingAccount.AccountId + 1,
					SoftLimitGB:      2,
				})
			assert.Error(t, err)
			assert.Nil(t, softLimit)
		}
	}
}

func TestIntegrationTracingAccount_GetTracingAccountSoftLimitNotExists(t *testing.T) {
	underTest, _, err := setupTracingAccountsIntegrationTest()

	if assert.NoError(t, err) {
		softLimit, err := underTest.GetTracingAccountSoftLimit(int64(1234567))
		assert.Error(t, err)
		assert.Nil(t, softLimit)
	}
}
