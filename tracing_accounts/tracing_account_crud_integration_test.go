package tracing_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationTracingAccount_CreateTracingAccount(t *testing.T) {
	underTest, email, err := setupTracingAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createTracingAccount := getCreateOrUpdateTracingAccount(email)
		createTracingAccount.AccountName = createTracingAccount.AccountName + "_create"

		tracingAccount, err := underTest.CreateTracingAccount(createTracingAccount)
		if assert.NoError(t, err) && assert.NotNil(t, tracingAccount) {
			defer underTest.DeleteTracingAccount(int64(tracingAccount.AccountId))
			assert.NotZero(t, tracingAccount.AccountId)
			assert.Equal(t, createTracingAccount.AccountName, tracingAccount.AccountName)
		}
	}
}

func TestIntegrationTracingAccount_GetTracingAccount(t *testing.T) {
	underTest, email, err := setupTracingAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createTracingAccount := getCreateOrUpdateTracingAccount(email)
		createTracingAccount.AccountName = createTracingAccount.AccountName + "_get"

		tracingAccount, err := underTest.CreateTracingAccount(createTracingAccount)
		if assert.NoError(t, err) && assert.NotNil(t, tracingAccount) {
			defer underTest.DeleteTracingAccount(int64(tracingAccount.AccountId))
			time.Sleep(4 * time.Second)

			getTracingAccount, err := underTest.GetTracingAccount(int64(tracingAccount.AccountId))
			assert.NoError(t, err)
			if assert.NotNil(t, getTracingAccount) {
				assert.Equal(t, tracingAccount.AccountId, getTracingAccount.AccountId)
				assert.Equal(t, createTracingAccount.AccountName, getTracingAccount.AccountName)
			}
		}
	}
}

func TestIntegrationTracingAccount_ListTracingAccounts(t *testing.T) {
	underTest, email, err := setupTracingAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createTracingAccount := getCreateOrUpdateTracingAccount(email)
		createTracingAccount.AccountName = createTracingAccount.AccountName + "_list"

		tracingAccount, err := underTest.CreateTracingAccount(createTracingAccount)
		if assert.NoError(t, err) && assert.NotNil(t, tracingAccount) {
			defer underTest.DeleteTracingAccount(int64(tracingAccount.AccountId))
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

func TestIntegrationTracingAccount_UpdateTracingAccount(t *testing.T) {
	underTest, email, err := setupTracingAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createTracingAccount := getCreateOrUpdateTracingAccount(email)
		createTracingAccount.AccountName = createTracingAccount.AccountName + "_update"

		tracingAccount, err := underTest.CreateTracingAccount(createTracingAccount)
		if assert.NoError(t, err) && assert.NotNil(t, tracingAccount) {
			defer underTest.DeleteTracingAccount(int64(tracingAccount.AccountId))
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
