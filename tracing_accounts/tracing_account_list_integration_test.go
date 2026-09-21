package tracing_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationTracingAccount_ListTracingAccounts(t *testing.T) {
	underTest, email, err := setupTracingAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createTracingAccount := getCreateOrUpdateTracingAccount(email)
		createTracingAccount.AccountName = createTracingAccount.AccountName + "_list"

		tracingAccount, err := underTest.CreateTracingAccount(createTracingAccount)
		if assert.NoError(t, err) && assert.NotNil(t, tracingAccount) {
			defer func() {
				// a failed cleanup leaks a real account, so surface it instead of ignoring it
				assert.NoError(t, underTest.DeleteTracingAccount(int64(tracingAccount.AccountId)),
					"cleanup of tracing account %d failed", tracingAccount.AccountId)
			}()
			time.Sleep(4 * time.Second)

			tracingAccounts, err := underTest.ListTracingAccounts()
			assert.NoError(t, err)
			assert.NotEmpty(t, tracingAccounts)

			found := false
			for _, account := range tracingAccounts {
				if account.AccountId == tracingAccount.AccountId {
					found = true
					break
				}
			}
			assert.True(t, found, "created tracing account should appear in the list")
		}
	}
}
