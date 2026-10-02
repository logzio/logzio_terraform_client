package sub_accounts_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/sub_accounts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Setting the main account's warm retention to the value it already has exercises the endpoint for the main
// account without changing it.
func TestIntegrationSubAccount_UpdateWarmRetentionMainAccount(t *testing.T) {
	underTest, _, err := setupSubAccountsWarmIntegrationTest()
	require.NoError(t, err)

	mainAccount := getWarmMainAccount(t, underTest)

	retentionDetails, err := underTest.UpdateWarmRetention(int64(mainAccount.AccountId),
		sub_accounts.UpdateWarmRetention{SnapSearchRetentionDays: mainAccount.SnapSearchRetentionDays})
	require.NoError(t, err)

	updated := findRetentionDetails(t, retentionDetails, mainAccount.AccountId)
	require.NotNil(t, updated.SnapSearchRetentionDays)
	assert.Equal(t, mainAccount.SnapSearchRetentionDays, *updated.SnapSearchRetentionDays)
}

func TestIntegrationSubAccount_UpdateWarmRetentionSubAccount(t *testing.T) {
	underTest, email, err := setupSubAccountsWarmIntegrationTest()
	require.NoError(t, err)

	mainAccount := getWarmMainAccount(t, underTest)
	subAccountId := createWarmTestSubAccount(t, underTest, email, mainAccount, "_warm_update")
	defer underTest.DeleteSubAccount(subAccountId)

	retentionDetails, err := underTest.UpdateWarmRetention(subAccountId, sub_accounts.UpdateWarmRetention{SnapSearchRetentionDays: 1})
	require.NoError(t, err)
	updated := findRetentionDetails(t, retentionDetails, int32(subAccountId))
	require.NotNil(t, updated.SnapSearchRetentionDays)
	assert.Equal(t, int32(1), *updated.SnapSearchRetentionDays)

	assert.Eventually(t, func() bool {
		subAccount, getErr := underTest.GetSubAccount(subAccountId)
		return getErr == nil && subAccount.SnapSearchRetentionDays == 1
	}, 30*time.Second, 2*time.Second, "GetSubAccount should read back the updated warm retention")

	// 0 turns warm tier off for a sub account.
	retentionDetails, err = underTest.UpdateWarmRetention(subAccountId, sub_accounts.UpdateWarmRetention{SnapSearchRetentionDays: 0})
	require.NoError(t, err)
	updated = findRetentionDetails(t, retentionDetails, int32(subAccountId))
	require.NotNil(t, updated.SnapSearchRetentionDays)
	assert.Equal(t, int32(0), *updated.SnapSearchRetentionDays)
}

func TestIntegrationSubAccount_UpdateWarmRetentionAboveMainAccount(t *testing.T) {
	underTest, email, err := setupSubAccountsWarmIntegrationTest()
	require.NoError(t, err)

	mainAccount := getWarmMainAccount(t, underTest)
	subAccountId := createWarmTestSubAccount(t, underTest, email, mainAccount, "_warm_above")
	defer underTest.DeleteSubAccount(subAccountId)

	retentionDetails, err := underTest.UpdateWarmRetention(subAccountId,
		sub_accounts.UpdateWarmRetention{SnapSearchRetentionDays: mainAccount.SnapSearchRetentionDays + 1})
	assert.Error(t, err)
	assert.Nil(t, retentionDetails)
	assert.Contains(t, err.Error(), "INVALID_WARM_RETENTION")
}

// getWarmMainAccount returns the main account of the warm API token, failing the test when it has no warm tier,
// which the endpoint needs.
func getWarmMainAccount(t *testing.T, underTest *sub_accounts.SubAccountClient) sub_accounts.SubAccount {
	accounts, err := underTest.ListSubAccounts()
	require.NoError(t, err)
	for _, account := range accounts {
		if account.IsOwner {
			require.Greater(t, account.SnapSearchRetentionDays, int32(0), "the warm test account must have warm tier")
			return account
		}
	}
	require.FailNow(t, "the warm API token's main account is missing from the accounts list")
	return sub_accounts.SubAccount{}
}

// createWarmTestSubAccount creates a sub account with no warm retention, the main account's hot retention and volume
// mode, and returns its id.
func createWarmTestSubAccount(t *testing.T, underTest *sub_accounts.SubAccountClient, email string,
	mainAccount sub_accounts.SubAccount, nameSuffix string) int64 {
	createSubAccount := getCreateOrUpdateSubAccount(email)
	createSubAccount.AccountName = createSubAccount.AccountName + nameSuffix
	createSubAccount.RetentionDays = mainAccount.RetentionDays
	if mainAccount.Flexible {
		createSubAccount.Flexible = "true"
		createSubAccount.ReservedDailyGB = new(float32)
	} else {
		// a fixed main account must have this much volume left to hand out, so keep it small
		*createSubAccount.MaxDailyGB = 0.01
	}

	subAccount, err := underTest.CreateSubAccount(createSubAccount)
	require.NoError(t, err)
	require.NotNil(t, subAccount)
	time.Sleep(4 * time.Second)
	return int64(subAccount.AccountId)
}

func findRetentionDetails(t *testing.T, retentionDetails []sub_accounts.AccountRetentionDetails, accountId int32) sub_accounts.AccountRetentionDetails {
	for _, details := range retentionDetails {
		if details.AccountId == accountId {
			return details
		}
	}
	require.FailNowf(t, "missing retention details", "account %d is missing from the response", accountId)
	return sub_accounts.AccountRetentionDetails{}
}
