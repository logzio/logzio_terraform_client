package metrics_accounts_test

import (
	"github.com/logzio/logzio_terraform_client/metrics_accounts"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// distinctive so the value read back cannot be a default or another test's account
const planUtsForSoftLimitTest = 1234

func TestIntegrationMetricsAccount_GetMetricsAccountSoftLimit(t *testing.T) {
	underTest, email, err := setupMetricsAccountsConsumptionIntegrationTest()

	if assert.NoError(t, err) {
		createMetricsAccount := getCreateOrUpdateMetricsAccount(email)
		createMetricsAccount.AccountName = createMetricsAccount.AccountName + "_sl_get"
		*createMetricsAccount.PlanUts = planUtsForSoftLimitTest

		metricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
		if assert.NoError(t, err) && assert.NotNil(t, metricsAccount) {
			t.Logf("CREATED metrics account id=%d name=%s planUts=%d (kept, not deleted)", metricsAccount.Id, metricsAccount.AccountName, metricsAccount.PlanUts)
			assert.Equal(t, int32(planUtsForSoftLimitTest), metricsAccount.PlanUts)
			time.Sleep(4 * time.Second)

			softLimit, err := underTest.GetMetricsAccountSoftLimit(int64(metricsAccount.Id))
			assert.NoError(t, err)
			if assert.NotNil(t, softLimit) {
				assert.Equal(t, metricsAccount.Id, softLimit.MetricsAccountId)
				// a consumption owner gets maxUniqueMetrics = the plan UTS at creation, so the
				// soft limit must come back as the value the account was created with
				if assert.NotNil(t, softLimit.SoftLimitUniqueMetrics) {
					assert.Equal(t, int32(planUtsForSoftLimitTest), *softLimit.SoftLimitUniqueMetrics)
				}
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
			t.Logf("CREATED metrics account id=%d name=%s planUts=%d (kept, not deleted)", metricsAccount.Id, metricsAccount.AccountName, metricsAccount.PlanUts)
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
