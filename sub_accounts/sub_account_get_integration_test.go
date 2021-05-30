package sub_accounts_test

import (
	"github.com/logzio/logzio_terraform_client/sub_accounts"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationSubAccount_GetSubAccount(t *testing.T) {
	underTest, email, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := sub_accounts.SubAccountCreate{
			Email:                 email,
			AccountName:           "test_get",
			MaxDailyGB:            1,
			RetentionDays:         1,
			Searchable:            false,
			Accessible:            true,
			DocSizeSetting:        false,
			SharingObjectAccounts: []int32{},
		}

		subAccount, err := underTest.CreateSubAccount(createSubAccount)
		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			defer underTest.DeleteSubAccount(subAccount.Id)
			time.Sleep(4 * time.Second)
			getSubAccount, err := underTest.GetSubAccount(subAccount.Id)

			assert.NoError(t, err)
			assert.NotNil(t, getSubAccount)
			assert.Equal(t, subAccount.Id, getSubAccount.Id)
			assert.Equal(t, subAccount.AccountName, getSubAccount.AccountName)
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