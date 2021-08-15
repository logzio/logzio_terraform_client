package sub_accounts_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestSubAccount_GetDetailedValidSubAccount(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	subAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/time-based-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(subAccountId), 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_detailed_subaccount.json"))
	})

	subAccount, err := underTest.GetDetailedSubAccount(subAccountId)
	assert.NoError(t, err)
	assert.NotNil(t, subAccount)
	assert.Equal(t, int32(1234555), subAccount.SubAccountRelation.OwnerAccountId)
	assert.Equal(t, int32(1234567), subAccount.SubAccountRelation.SubAccountId)
	assert.True(t, subAccount.SubAccountRelation.Searchable)
	assert.False(t, subAccount.SubAccountRelation.Accessible)
	assert.Equal(t, int64(1626964508000), subAccount.SubAccountRelation.CreatedDate)
	assert.Equal(t, int64(1626964508000), subAccount.SubAccountRelation.LastUpdatedDate)
	assert.Equal(t, int32(32123), subAccount.SubAccountRelation.LastUpdaterUserId)
	assert.Equal(t, "SUB_ACCOUNT", subAccount.SubAccountRelation.Type)
	assert.Equal(t, int32(1234567), subAccount.Account.AccountId)
	assert.Equal(t, "testAccount", subAccount.Account.AccountName)
	assert.True(t, subAccount.Account.Active)
	assert.Equal(t, "testIndex", subAccount.Account.EsIndexPrefix)
	assert.Equal(t, float32(5.0), subAccount.Account.MaxDailyGB)
	assert.Zero(t, subAccount.Account.ReservedDailyGB)
	assert.False(t, subAccount.Account.Flexible)
	assert.Equal(t, int32(5), subAccount.Account.RetentionDays)
	assert.Equal(t, "testToken", subAccount.Account.AccountToken)
	assert.Zero(t, len(subAccount.SharingObjectsAccounts))
	assert.Equal(t, int32(5), subAccount.UtilizationSettings.FrequencyMinutes)
	assert.True(t, subAccount.UtilizationSettings.UtilizationEnabled)
	assert.Zero(t, len(subAccount.DailyUsagesList.Usage))
	assert.True(t, subAccount.DocSizeSetting)
}

func TestSubAccount_GetDetailedSubAccountAPIFail(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	subAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/time-based-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(subAccountId), 10))
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("get_detailed_subaccount_failed.txt"))
	})

	subAccount, err := underTest.GetDetailedSubAccount(subAccountId)
	assert.Error(t, err)
	assert.Nil(t, subAccount)
}
