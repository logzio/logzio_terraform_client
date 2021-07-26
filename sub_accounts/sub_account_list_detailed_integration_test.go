package sub_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationSubAccount_ListDetailedSubAccounts(t *testing.T) {
	underTest, _, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		listDetailed, err := underTest.ListDetailedSubAccounts()
		assert.NoError(t, err)
		assert.NotNil(t, listDetailed)
	}
}
