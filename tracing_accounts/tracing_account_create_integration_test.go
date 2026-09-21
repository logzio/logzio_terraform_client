package tracing_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationTracingAccount_CreateTracingAccount(t *testing.T) {
	underTest, email, err := setupTracingAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createTracingAccount := getCreateOrUpdateTracingAccount(email)
		createTracingAccount.AccountName = createTracingAccount.AccountName + "_create"

		tracingAccount, err := underTest.CreateTracingAccount(createTracingAccount)
		if assert.NoError(t, err) && assert.NotNil(t, tracingAccount) {
			defer func() {
				// a failed cleanup leaks a real account, so surface it instead of ignoring it
				assert.NoError(t, underTest.DeleteTracingAccount(int64(tracingAccount.AccountId)),
					"cleanup of tracing account %d failed", tracingAccount.AccountId)
			}()
			assert.NotZero(t, tracingAccount.AccountId)
			assert.Equal(t, createTracingAccount.AccountName, tracingAccount.AccountName)
		}
	}
}
