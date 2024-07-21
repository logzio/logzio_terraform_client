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
		createMetricsAccount := getCreateOrUpdateMetricsAccount(email)
		createMetricsAccount.AccountName = createMetricsAccount.AccountName + "_create"

		metricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
		if assert.NoError(t, err) && assert.NotNil(t, metricsAccount) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteMetricsAccount(int64(metricsAccount.Id))
			assert.NotEmpty(t, metricsAccount.Token)
			assert.NotEmpty(t, metricsAccount.Id)
		}
	}
}

func TestIntegrationMetricsAccount_CreateMetricsAccountWithSharingAccount(t *testing.T) {
	underTest, email, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		accountId, err := test_utils.GetAccountId()
		assert.NoError(t, err)

		createMetricsAccount := getCreateOrUpdateMetricsAccount(email)
		createMetricsAccount.AccountName = createMetricsAccount.AccountName + "_create_with_sharing_account"
		createMetricsAccount.AuthorizedAccountsIds = []int32{int32(accountId)}
		metricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
		if assert.NoError(t, err) && assert.NotNil(t, metricsAccount) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteMetricsAccount(int64(metricsAccount.Id))
			assert.NotEmpty(t, metricsAccount.Token)
			assert.NotEmpty(t, metricsAccount.Id)
		}
	}
}

func TestIntegrationMetricsAccount_CreateMetricsAccountInvalidMail(t *testing.T) {
	underTest, _, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createMetricsAccount := getCreateOrUpdateMetricsAccount("invalid@mail.test")
		metricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)

		assert.Error(t, err)
		assert.Nil(t, metricsAccount)
	}
}

func TestIntegrationMetricsAccount_CreateMetricsAccountNoMail(t *testing.T) {
	underTest, _, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createMetricsAccount := getCreateOrUpdateMetricsAccount("")
		metricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)

		assert.Error(t, err)
		assert.Nil(t, metricsAccount)
	}
}

func TestIntegrationMetricsAccount_CreateMetricsAccountNoAccountName(t *testing.T) {
	underTest, email, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		createMetricsAccount := getCreateOrUpdateMetricsAccount(email)
		createMetricsAccount.AccountName = ""
		metricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)

		if assert.NoError(t, err) && assert.NotNil(t, metricsAccount) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteMetricsAccount(int64(metricsAccount.Id))
			assert.NotEmpty(t, metricsAccount.Token)
			assert.NotEmpty(t, metricsAccount.Id)
			assert.Equal(t, metricsAccount.AccountName, "IntegrationsTeamTesting_metrics")
		}
	}
}
