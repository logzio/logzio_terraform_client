package tracing_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationTracingAccount_UpdateTracingAccount(t *testing.T) {
	underTest, email, err := setupTracingAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createTracingAccount := getCreateOrUpdateTracingAccount(email)
		createTracingAccount.AccountName = createTracingAccount.AccountName + "_update"

		tracingAccount, err := underTest.CreateTracingAccount(createTracingAccount)
		if assert.NoError(t, err) && assert.NotNil(t, tracingAccount) {
			defer func() {
				// a failed cleanup leaks a real account, so surface it instead of ignoring it
				assert.NoError(t, underTest.DeleteTracingAccount(int64(tracingAccount.AccountId)),
					"cleanup of tracing account %d failed", tracingAccount.AccountId)
			}()
			time.Sleep(time.Second * 2)

			createTracingAccount.AccountName = "test_after_update"
			updated, err := underTest.UpdateTracingAccount(int64(tracingAccount.AccountId), createTracingAccount)
			assert.NoError(t, err)
			assert.NotNil(t, updated)

			// verify that the update was made
			time.Sleep(time.Second * 2)
			getTracingAccount, err := underTest.GetTracingAccount(int64(tracingAccount.AccountId))
			assert.NoError(t, err)
			if assert.NotNil(t, getTracingAccount) {
				assert.Equal(t, "test_after_update", getTracingAccount.AccountName)
			}
		}
	}
}
