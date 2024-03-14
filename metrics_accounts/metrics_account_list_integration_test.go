package metrics_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationMetricsAccount_ListMetricsAccounts(t *testing.T) {
	underTest, _, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		metricsAccounts, err := underTest.ListMetricsAccounts()
		assert.NoError(t, err)
		assert.NotNil(t, metricsAccounts)
	}
}
