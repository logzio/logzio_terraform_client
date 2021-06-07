package sub_accounts_test

import (
	"github.com/logzio/logzio_terraform_client/sub_accounts"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationSubAccount_DeleteSubAccount(t *testing.T) {
	underTest, email, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err){
		createSubAccount := sub_accounts.SubAccountCreate{
			Email:                 email,
			AccountName:           "test_delete",
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
			defer func() {
				err = underTest.DeleteSubAccount(subAccount.Id)
				assert.NoError(t, err)
			}()

		}
	}
}

func TestIntegrationSubAccount_DeleteSubAccountNotExists(t *testing.T) {
	underTest, _, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		err = underTest.DeleteSubAccount(int64(1234567))
		assert.Error(t, err)
	}
}