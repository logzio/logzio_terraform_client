package sub_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationSubAccount_ListSubAccounts(t *testing.T) {
	underTest, _, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		subAccounts, err := underTest.ListSubAccounts()
		assert.NoError(t, err)
		assert.NotNil(t, subAccounts)
	}
}
