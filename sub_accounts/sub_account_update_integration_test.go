package sub_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationSubAccount_UpdateSubAccount(t *testing.T) {
	underTest, email, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateSubAccount(email)

		subAccount, err := underTest.CreateSubAccount(createSubAccount)
		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			defer underTest.DeleteSubAccount(int64(subAccount.AccountId))
			time.Sleep(time.Second * 2)
			getSubAccount, err := underTest.GetSubAccount(int64(subAccount.AccountId))
			assert.NoError(t, err)
			assert.Equal(t, "tf_client_test", getSubAccount.AccountName)
			createSubAccount.AccountName = "test_after_update"
			time.Sleep(time.Second * 2)
			err = underTest.UpdateSubAccount(int64(subAccount.AccountId), createSubAccount)
			assert.NoError(t, err)
			// verify that the update was made
			time.Sleep(time.Second * 2)
			getSubAccount, err = underTest.GetSubAccount(int64(subAccount.AccountId))
			assert.NoError(t, err)
			assert.Equal(t, "test_after_update", getSubAccount.AccountName)
		}
	}
}

func TestIntegrationSubAccount_UpdateSubAccountIdNotExists(t *testing.T) {
	underTest, email, err := setupSubAccountsIntegrationTest()
	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateSubAccount(email)
		if assert.NoError(t, err) && assert.NotNil(t, createSubAccount) {
			err = underTest.UpdateSubAccount(int64(1234567), createSubAccount)
			assert.Error(t, err)
		}
	}
}
