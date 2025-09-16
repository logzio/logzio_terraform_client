package sub_accounts_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func TestIntegrationSubAccount_UpdateSubAccountWarmRetention(t *testing.T) {
	underTest, email, err := setupSubAccountsWarmIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateSubAccount(email)
		createSubAccount.ReservedDailyGB = new(float32)
		createSubAccount.Flexible = "true"

		subAccount, err := underTest.CreateSubAccount(createSubAccount)
		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			defer underTest.DeleteSubAccount(int64(subAccount.AccountId))
			time.Sleep(time.Second * 2)
			getSubAccount, err := underTest.GetSubAccount(int64(subAccount.AccountId))
			assert.NoError(t, err)
			assert.Equal(t, "tf_client_test", getSubAccount.AccountName)
			createSubAccount.AccountName = "test_after_update"
			createSubAccount.RetentionDays = 4
			warmRetention := int32(2)
			createSubAccount.SnapSearchRetentionDays = &warmRetention

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

func TestIntegrationSubAccount_UpdateSubAccountWithSoftLimitGB(t *testing.T) {
	underTest, email, err := setupSubAccountConsumptionIntegrationTest()

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
			softLimitGB := float32(1)
			createSubAccount.SoftLimitGB = &softLimitGB
			time.Sleep(time.Second * 2)
			err = underTest.UpdateSubAccount(int64(subAccount.AccountId), createSubAccount)
			assert.NoError(t, err)
			// verify that the update was made
			time.Sleep(time.Second * 2)
			getSubAccount, err = underTest.GetSubAccount(int64(subAccount.AccountId))
			assert.NoError(t, err)
			assert.Equal(t, "test_after_update", getSubAccount.AccountName)
			assert.Equal(t, softLimitGB, getSubAccount.SoftLimitGB)
		}
	}
}
