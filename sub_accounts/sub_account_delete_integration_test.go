package sub_accounts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationSubAccount_DeleteSubAccount(t *testing.T) {
	underTest, email, err := setupSubAccountsIntegrationTest()
	if assert.NoError(t, err) {
		createSubAccount := getCreatrOrUpdateSubAccount(email)
		createSubAccount.AccountName = createSubAccount.AccountName + "_delete"
		subAccount, err := underTest.CreateSubAccount(createSubAccount)
		if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
			time.Sleep(2 * time.Second)
			defer func() {
				err = underTest.DeleteSubAccount(int64(subAccount.AccountId))
				assert.NoError(t, err)
			}()
		}
	}
}

func TestIntegrationSubAccount_DeleteSubAccountNotExists(t *testing.T) {
	underTest, _, err := setupSubAccountsIntegrationTest()

	if assert.NoError(t, err) {
		err = underTest.DeleteSubAccount(int64(1234567))
		assert.Error(t, err)
	}
}
