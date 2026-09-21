package tracing_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationTracingAccount_DeleteTracingAccount(t *testing.T) {
	underTest, email, err := setupTracingAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createTracingAccount := getCreateOrUpdateTracingAccount(email)
		createTracingAccount.AccountName = createTracingAccount.AccountName + "_delete"

		tracingAccount, err := underTest.CreateTracingAccount(createTracingAccount)
		if assert.NoError(t, err) && assert.NotNil(t, tracingAccount) {
			time.Sleep(4 * time.Second)

			err = underTest.DeleteTracingAccount(int64(tracingAccount.AccountId))
			assert.NoError(t, err)

			time.Sleep(time.Second * 2)
			getTracingAccount, err := underTest.GetTracingAccount(int64(tracingAccount.AccountId))
			assert.Error(t, err)
			assert.Nil(t, getTracingAccount)
		}
	}
}
