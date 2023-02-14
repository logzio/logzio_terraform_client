package sub_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationSubAccount_GetSubAccount(t *testing.T) {
	underTest, email, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateSubAccount(email)
		createSubAccount.AccountName = createSubAccount.AccountName + "_get"

		subAccount, err := underTest.CreateSubAccount(createSubAccount)
		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			defer underTest.DeleteSubAccount(int64(subAccount.AccountId))
			time.Sleep(4 * time.Second)
			getSubAccount, err := underTest.GetSubAccount(int64(subAccount.AccountId))
			assert.NoError(t, err)
			assert.NotNil(t, getSubAccount)
			assert.Equal(t, subAccount.AccountId, getSubAccount.AccountId)
			assert.Equal(t, createSubAccount.AccountName, getSubAccount.AccountName)
		}
	}
}

func TestIntegrationSubAccount_GetSubAccountNotExists(t *testing.T) {
	underTest, _, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		subAccount, err := underTest.GetSubAccount(int64(1234567))
		assert.Error(t, err)
		assert.Nil(t, subAccount)
	}
}
