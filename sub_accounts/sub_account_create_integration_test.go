package sub_accounts_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestIntegrationSubAccount_CreateSubAccount(t *testing.T) {
	underTest, email, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateSubAccount(email)
		createSubAccount.AccountName = createSubAccount.AccountName + "_create"

		subAccount, err := underTest.CreateSubAccount(createSubAccount)
		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteSubAccount(int64(subAccount.AccountId))
			assert.NotEmpty(t, subAccount.AccountToken)
			assert.NotEmpty(t, subAccount.AccountId)
		}
	}
}

func TestIntegrationSubAccount_CreateSubAccountWithSharingAccount(t *testing.T) {
	underTest, email, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		accountId, err := test_utils.GetAccountId()
		assert.NoError(t, err)

		createSubAccount := getCreateOrUpdateSubAccount(email)
		createSubAccount.AccountName = createSubAccount.AccountName + "_create_with_sharing_account"
		createSubAccount.SharingObjectsAccounts = []int32{int32(accountId)}
		subAccount, err := underTest.CreateSubAccount(createSubAccount)
		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteSubAccount(int64(subAccount.AccountId))
			assert.NotEmpty(t, subAccount.AccountToken)
			assert.NotEmpty(t, subAccount.AccountId)
		}
	}
}

func TestIntegrationSubAccount_CreateSubAccountWithUtilization(t *testing.T) {
	underTest, email, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateSubAccount(email)
		createSubAccount.AccountName = createSubAccount.AccountName + "_create_utilization"
		createSubAccount.UtilizationSettings.UtilizationEnabled = strconv.FormatBool(true)
		createSubAccount.UtilizationSettings.FrequencyMinutes = 5

		subAccount, err := underTest.CreateSubAccount(createSubAccount)
		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteSubAccount(int64(subAccount.AccountId))
			assert.NotEmpty(t, subAccount.AccountToken)
			assert.NotEmpty(t, subAccount.AccountId)
		}
	}
}

/*
TODO: Add the following test as part of the automated testings.
This test can only be run when the main test account can be set to 'flexible'.
Since the account is non-flexible, and needs manual action to make it flexible,
this test will be tested locally by uncommenting it, and not as part of the automated tests.
*/
//func TestIntegrationSubAccount_CreateSubAccountFlexible(t *testing.T) {
//	underTest, email, err := setupSubAccountsIntegrationTest()
//
//	if assert.NoError(t, err) {
//		createSubAccount := getCreateOrUpdateSubAccount(email)
//		createSubAccount.AccountName = createSubAccount.AccountName + "_create_flexible"
//		createSubAccount.Flexible = strconv.FormatBool(true)
//		createSubAccount.ReservedDailyGB = new(float32)
//		*createSubAccount.ReservedDailyGB = 1
//		subAccount, err := underTest.CreateSubAccount(createSubAccount)
//
//		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
//			time.Sleep(4 * time.Second)
//			defer underTest.DeleteSubAccount(int64(subAccount.AccountId))
//			assert.NotEmpty(t, subAccount.AccountToken)
//			assert.NotEmpty(t, subAccount.AccountId)
//		}
//	}
//}
//
//func TestIntegrationSubAccount_CreateSubAccountFlexibleReservedZero(t *testing.T) {
//	underTest, email, err := setupSubAccountsIntegrationTest()
//
//	if assert.NoError(t, err) {
//		createSubAccount := getCreateOrUpdateSubAccount(email)
//		createSubAccount.AccountName = createSubAccount.AccountName + "_create_flexible"
//		createSubAccount.Flexible = strconv.FormatBool(true)
//		createSubAccount.ReservedDailyGB = new(float32)
//		*createSubAccount.ReservedDailyGB = 0
//		subAccount, err := underTest.CreateSubAccount(createSubAccount)
//
//		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
//			time.Sleep(4 * time.Second)
//			defer underTest.DeleteSubAccount(int64(subAccount.AccountId))
//			assert.NotEmpty(t, subAccount.AccountToken)
//			assert.NotEmpty(t, subAccount.AccountId)
//		}
//	}
//}

func TestIntegrationSubAccount_CreateSubAccountInvalidMail(t *testing.T) {
	underTest, _, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateSubAccount("invalid@mail.test")
		subAccount, err := underTest.CreateSubAccount(createSubAccount)

		assert.Error(t, err)
		assert.Nil(t, subAccount)
	}
}

func TestIntegrationSubAccount_CreateSubAccountNoMail(t *testing.T) {
	underTest, _, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateSubAccount("")
		subAccount, err := underTest.CreateSubAccount(createSubAccount)

		assert.Error(t, err)
		assert.Nil(t, subAccount)
	}
}

func TestIntegrationSubAccount_CreateSubAccountNoAccountName(t *testing.T) {
	underTest, email, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateSubAccount(email)
		createSubAccount.AccountName = ""
		subAccount, err := underTest.CreateSubAccount(createSubAccount)

		assert.Error(t, err)
		assert.Nil(t, subAccount)
	}
}

func TestIntegrationSubAccount_CreateSubAccountNoRetention(t *testing.T) {
	underTest, email, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateSubAccount(email)
		createSubAccount.AccountName = createSubAccount.AccountName + "_no_retention"
		createSubAccount.RetentionDays = 0
		subAccount, err := underTest.CreateSubAccount(createSubAccount)

		assert.Error(t, err)
		assert.Nil(t, subAccount)
	}
}

func TestIntegrationSubAccount_CreateSubAccountWarmRetention(t *testing.T) {
	underTest, email, err := setupSubAccountsWarmIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateSubAccount(email)
		createSubAccount.AccountName = createSubAccount.AccountName + "_create"
		createSubAccount.RetentionDays = 4
		warmRetention := int32(2)
		createSubAccount.SnapSearchRetentionDays = &warmRetention
		createSubAccount.ReservedDailyGB = new(float32)
		createSubAccount.Flexible = "true"

		subAccount, err := underTest.CreateSubAccount(createSubAccount)
		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteSubAccount(int64(subAccount.AccountId))
			assert.NotEmpty(t, subAccount.AccountToken)
			assert.NotEmpty(t, subAccount.AccountId)
		}
	}
}
