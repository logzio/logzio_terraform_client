package sub_accounts_test

import (
	"github.com/logzio/logzio_terraform_client/sub_accounts"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationSubAccount_UpdateSubAccount(t *testing.T) {
	underTest, email, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err){
		createSubAccount := sub_accounts.SubAccountCreate{
			Email:                 email,
			AccountName:           "test_before_update",
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
			assert.Equal(t, "test_before_update", subAccount.AccountName)
			subAccount.AccountName = "test_after_update"
			time.Sleep(4 * time.Second)
			err := underTest.UpdateSubAccount(subAccount.Id, *subAccount)
			assert.NoError(t, err)
		}
	}
}