package sub_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationSubAccount_GetDetailedSubAccount(t *testing.T) {
	underTest, email, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateSubAccount(email)
		createSubAccount.AccountName = createSubAccount.AccountName + "_get_detailed"

		subAccount, err := underTest.CreateSubAccount(createSubAccount)
		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			defer underTest.DeleteSubAccount(int64(subAccount.AccountId))
			time.Sleep(4 * time.Second)
			getSubAccount, err := underTest.GetDetailedSubAccount(int64(subAccount.AccountId))
			assert.NoError(t, err)
			assert.NotNil(t, getSubAccount)
			assert.Equal(t, subAccount.AccountId, getSubAccount.Account.AccountId)
		}
	}
}

func TestIntegrationSubAccount_GetDetailedSubAccountNotExists(t *testing.T) {
	underTest, _, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		subAccount, err := underTest.GetDetailedSubAccount(int64(1234567))
		assert.Error(t, err)
		assert.Nil(t, subAccount)
	}
}
