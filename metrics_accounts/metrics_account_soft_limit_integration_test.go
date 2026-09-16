package metrics_accounts_test

import (
	"github.com/logzio/logzio_terraform_client/metrics_accounts"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationMetricsAccount_GetMetricsAccountSoftLimit(t *testing.T) {
	underTest, email, err := setupMetricsAccountsConsumptionIntegrationTest()

	if assert.NoError(t, err) {
		createMetricsAccount := getCreateOrUpdateMetricsAccount(email)
		createMetricsAccount.AccountName = createMetricsAccount.AccountName + "_sl_get"

		metricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
		if assert.NoError(t, err) && assert.NotNil(t, metricsAccount) {
			defer underTest.DeleteMetricsAccount(int64(metricsAccount.Id))
			time.Sleep(4 * time.Second)

			softLimit, err := underTest.GetMetricsAccountSoftLimit(int64(metricsAccount.Id))
			assert.NoError(t, err)
			if assert.NotNil(t, softLimit) {
				assert.Equal(t, metricsAccount.Id, softLimit.MetricsAccountId)
			}
		}
	}
}

func TestIntegrationMetricsAccount_UpdateMetricsAccountSoftLimit(t *testing.T) {
	underTest, email, err := setupMetricsAccountsConsumptionIntegrationTest()

	if assert.NoError(t, err) {
		createMetricsAccount := getCreateOrUpdateMetricsAccount(email)
		createMetricsAccount.AccountName = createMetricsAccount.AccountName + "_sl_update"

		metricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
		if assert.NoError(t, err) && assert.NotNil(t, metricsAccount) {
			defer underTest.DeleteMetricsAccount(int64(metricsAccount.Id))
			time.Sleep(4 * time.Second)

			// above the plan UTS of 100, so the limiter actually applies it
			updated, err := underTest.UpdateMetricsAccountSoftLimit(int64(metricsAccount.Id),
				metrics_accounts.UpdateMetricsAccountSoftLimit{SoftLimitUniqueMetrics: 500})
			assert.NoError(t, err)
			if assert.NotNil(t, updated) && assert.NotNil(t, updated.SoftLimitUniqueMetrics) {
				assert.Equal(t, int32(500), *updated.SoftLimitUniqueMetrics)
			}

			// verify that the update was made
			time.Sleep(time.Second * 2)
			softLimit, err := underTest.GetMetricsAccountSoftLimit(int64(metricsAccount.Id))
			assert.NoError(t, err)
			if assert.NotNil(t, softLimit) && assert.NotNil(t, softLimit.SoftLimitUniqueMetrics) {
				assert.Equal(t, int32(500), *softLimit.SoftLimitUniqueMetrics)
				assert.Equal(t, metricsAccount.Id, softLimit.MetricsAccountId)
			}
		}
	}
}

func TestIntegrationMetricsAccount_GetMetricsAccountSoftLimitNotExists(t *testing.T) {
	underTest, _, err := setupMetricsAccountsConsumptionIntegrationTest()

	if assert.NoError(t, err) {
		softLimit, err := underTest.GetMetricsAccountSoftLimit(int64(1234567))
		assert.Error(t, err)
		assert.Nil(t, softLimit)
	}
}

func TestIntegrationMetricsAccount_UpdateMetricsAccountSoftLimitNotExists(t *testing.T) {
	underTest, _, err := setupMetricsAccountsConsumptionIntegrationTest()

	if assert.NoError(t, err) {
		softLimit, err := underTest.UpdateMetricsAccountSoftLimit(int64(1234567),
			metrics_accounts.UpdateMetricsAccountSoftLimit{SoftLimitUniqueMetrics: 500})
		assert.Error(t, err)
		assert.Nil(t, softLimit)
	}
}
