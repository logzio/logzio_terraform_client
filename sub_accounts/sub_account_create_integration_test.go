package sub_accounts_test

import (
	"github.com/logzio/logzio_terraform_client/sub_accounts"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationSubAccount_CreateSubAccount(t *testing.T) {
	underTest, email, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := sub_accounts.SubAccountCreate{
			Email:                 email,
			AccountName:           "test_create",
			MaxDailyGB:            1,
			RetentionDays:         1,
			Searchable:            false,
			Accessible:            true,
			DocSizeSetting:        false,
			SharingObjectAccounts: []int32{},
		}

		subAccount, err := underTest.CreateSubAccount(createSubAccount)

		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteSubAccount(subAccount.Id)
			assert.NotEmpty(t, subAccount.Token)
			assert.NotEmpty(t, subAccount.AccountId)
		}
	}
}

func TestIntegrationSubAccount_CreateSubAccountWithSharingAccount(t *testing.T) {
	underTest, email, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		accountId, err := test_utils.GetAccountId()
		assert.NoError(t, err)

		createSubAccount := sub_accounts.SubAccountCreate{
			Email:                 email,
			AccountName:           "test_create_sharing_account",
			MaxDailyGB:            1,
			RetentionDays:         1,
			Searchable:            false,
			Accessible:            true,
			DocSizeSetting:        false,
			SharingObjectAccounts: []int32{int32(accountId)},
		}

		subAccount, err := underTest.CreateSubAccount(createSubAccount)

		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteSubAccount(subAccount.Id)
			assert.NotEmpty(t, subAccount.Token)
			assert.NotEmpty(t, subAccount.AccountId)
		}
	}
}

func TestIntegrationSubAccount_CreateSubAccountInvalidMail(t *testing.T) {
	underTest, _, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := sub_accounts.SubAccountCreate{
			Email:          "invalid@mail.com",
			AccountName:    "test_create",
			MaxDailyGB:     1,
			RetentionDays:  1,
			Searchable:     false,
			Accessible:     true,
			DocSizeSetting: false,
		}

		subAccount, err := underTest.CreateSubAccount(createSubAccount)

		assert.Error(t, err)
		assert.Nil(t, subAccount)
	}
}
