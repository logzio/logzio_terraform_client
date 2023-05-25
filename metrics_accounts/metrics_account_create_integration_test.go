package metrics_accounts_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationMetricsAccount_CreateMetricsAccount(t *testing.T) {
	underTest, email, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateMetricsAccount(email)
		createSubAccount.AccountName = createSubAccount.AccountName + "_create"

		subAccount, err := underTest.CreateMetricsAccount(createSubAccount)
		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteMetricsAccount(int64(subAccount.Id))
			assert.NotEmpty(t, subAccount.Token)
			assert.NotEmpty(t, subAccount.Id)
		}
	}
}

func TestIntegrationMetricsAccount_CreateMetricsAccountWithSharingAccount(t *testing.T) {
	underTest, email, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		accountId, err := test_utils.GetAccountId()
		assert.NoError(t, err)

		createSubAccount := getCreateOrUpdateMetricsAccount(email)
		createSubAccount.AccountName = createSubAccount.AccountName + "_create_with_sharing_account"
		createSubAccount.AuthorizedAccountsIds = []int32{int32(accountId)}
		subAccount, err := underTest.CreateMetricsAccount(createSubAccount)
		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteMetricsAccount(int64(subAccount.Id))
			assert.NotEmpty(t, subAccount.Token)
			assert.NotEmpty(t, subAccount.Id)
		}
	}
}

func TestIntegrationMetricsAccount_CreateMetricsAccountInvalidMail(t *testing.T) {
	underTest, _, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateMetricsAccount("invalid@mail.test")
		subAccount, err := underTest.CreateMetricsAccount(createSubAccount)

		assert.Error(t, err)
		assert.Nil(t, subAccount)
	}
}

func TestIntegrationMetricsAccount_CreateMetricsAccountNoMail(t *testing.T) {
	underTest, _, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateMetricsAccount("")
		subAccount, err := underTest.CreateMetricsAccount(createSubAccount)

		assert.Error(t, err)
		assert.Nil(t, subAccount)
	}
}

func TestIntegrationMetricsAccount_CreateMetricsAccountNoAccountName(t *testing.T) { //TODO rewrite
	underTest, email, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createSubAccount := getCreateOrUpdateMetricsAccount(email)
		createSubAccount.AccountName = ""
		subAccount, err := underTest.CreateMetricsAccount(createSubAccount)

		assert.Error(t, err)
		assert.Nil(t, subAccount)
	}
}
