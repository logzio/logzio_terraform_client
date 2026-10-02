package sub_accounts_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/sub_accounts"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationSubAccount_UpdateWarmRetentionMainAccount(t *testing.T) {
	underTest, _, err := setupSubAccountsWarmIntegrationTest()

	if assert.NoError(t, err) {
		mainAccount, err := getWarmMainAccount(underTest)
		if assert.NoError(t, err) {
			// setting the value the main account already has exercises the endpoint without changing the account
			retentionDetails, err := underTest.UpdateWarmRetention(int64(mainAccount.AccountId),
				sub_accounts.UpdateWarmRetention{SnapSearchRetentionDays: mainAccount.SnapSearchRetentionDays})
			if assert.NoError(t, err) {
				updated := findRetentionDetails(retentionDetails, mainAccount.AccountId)
				if assert.NotNil(t, updated) && assert.NotNil(t, updated.SnapSearchRetentionDays) {
					assert.Equal(t, mainAccount.SnapSearchRetentionDays, *updated.SnapSearchRetentionDays)
				}
			}
		}
	}
}

func TestIntegrationSubAccount_UpdateWarmRetentionSubAccount(t *testing.T) {
	underTest, email, err := setupSubAccountsWarmIntegrationTest()

	if assert.NoError(t, err) {
		mainAccount, err := getWarmMainAccount(underTest)
		if assert.NoError(t, err) {
			createSubAccount := getWarmCreateSubAccount(email, mainAccount)
			createSubAccount.AccountName = createSubAccount.AccountName + "_warm_update"

			subAccount, err := underTest.CreateSubAccount(createSubAccount)
			if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
				time.Sleep(4 * time.Second)
				defer underTest.DeleteSubAccount(int64(subAccount.AccountId))

				retentionDetails, err := underTest.UpdateWarmRetention(int64(subAccount.AccountId),
					sub_accounts.UpdateWarmRetention{SnapSearchRetentionDays: 1})
				if assert.NoError(t, err) {
					updated := findRetentionDetails(retentionDetails, subAccount.AccountId)
					if assert.NotNil(t, updated) && assert.NotNil(t, updated.SnapSearchRetentionDays) {
						assert.Equal(t, int32(1), *updated.SnapSearchRetentionDays)
					}
				}
				// verify that the update was made
				time.Sleep(time.Second * 2)
				getSubAccount, err := underTest.GetSubAccount(int64(subAccount.AccountId))
				assert.NoError(t, err)
				assert.Equal(t, int32(1), getSubAccount.SnapSearchRetentionDays)

				// 0 turns warm tier off for a sub account
				time.Sleep(time.Second * 2)
				retentionDetails, err = underTest.UpdateWarmRetention(int64(subAccount.AccountId),
					sub_accounts.UpdateWarmRetention{SnapSearchRetentionDays: 0})
				if assert.NoError(t, err) {
					updated := findRetentionDetails(retentionDetails, subAccount.AccountId)
					if assert.NotNil(t, updated) && assert.NotNil(t, updated.SnapSearchRetentionDays) {
						assert.Equal(t, int32(0), *updated.SnapSearchRetentionDays)
					}
				}
			}
		}
	}
}

func TestIntegrationSubAccount_UpdateWarmRetentionAboveMainAccount(t *testing.T) {
	underTest, email, err := setupSubAccountsWarmIntegrationTest()

	if assert.NoError(t, err) {
		mainAccount, err := getWarmMainAccount(underTest)
		if assert.NoError(t, err) {
			createSubAccount := getWarmCreateSubAccount(email, mainAccount)
			createSubAccount.AccountName = createSubAccount.AccountName + "_warm_above"

			subAccount, err := underTest.CreateSubAccount(createSubAccount)
			if assert.NoError(t, err) && assert.NotNil(t, subAccount) {
				time.Sleep(4 * time.Second)
				defer underTest.DeleteSubAccount(int64(subAccount.AccountId))

				retentionDetails, err := underTest.UpdateWarmRetention(int64(subAccount.AccountId),
					sub_accounts.UpdateWarmRetention{SnapSearchRetentionDays: mainAccount.SnapSearchRetentionDays + 1})
				assert.Nil(t, retentionDetails)
				if assert.Error(t, err) {
					assert.Contains(t, err.Error(), "INVALID_WARM_RETENTION")
				}
			}
		}
	}
}

// getWarmMainAccount returns the main account of the warm API token, which must have warm tier for the endpoint to work
func getWarmMainAccount(underTest *sub_accounts.SubAccountClient) (sub_accounts.SubAccount, error) {
	accounts, err := underTest.ListSubAccounts()
	if err != nil {
		return sub_accounts.SubAccount{}, err
	}
	for _, account := range accounts {
		if account.IsOwner {
			if account.SnapSearchRetentionDays <= 0 {
				return sub_accounts.SubAccount{}, fmt.Errorf("main account %d has no warm tier", account.AccountId)
			}
			return account, nil
		}
	}
	return sub_accounts.SubAccount{}, fmt.Errorf("main account is missing from the accounts list")
}

// getWarmCreateSubAccount returns a sub account without warm retention that the main account can hold: same hot
// retention and volume mode, and little volume for a fixed main account
func getWarmCreateSubAccount(email string, mainAccount sub_accounts.SubAccount) sub_accounts.CreateOrUpdateSubAccount {
	createSubAccount := getCreateOrUpdateSubAccount(email)
	createSubAccount.RetentionDays = mainAccount.RetentionDays
	if mainAccount.Flexible {
		createSubAccount.ReservedDailyGB = new(float32)
		createSubAccount.Flexible = "true"
	} else {
		*createSubAccount.MaxDailyGB = 0.01
	}
	return createSubAccount
}

func findRetentionDetails(retentionDetails []sub_accounts.AccountRetentionDetails, accountId int32) *sub_accounts.AccountRetentionDetails {
	for i := range retentionDetails {
		if retentionDetails[i].AccountId == accountId {
			return &retentionDetails[i]
		}
	}
	return nil
}
