package metrics_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationSubAccount_ListSubAccounts(t *testing.T) {
	underTest, _, err := setupMetricsAccountsIntegrationTest()

	if assert.NoError(t, err) {
		subAccounts, err := underTest.ListMetricsAccounts()
		assert.NoError(t, err)
		assert.NotNil(t, subAccounts)
	}
}
