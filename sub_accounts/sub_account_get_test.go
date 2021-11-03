package sub_accounts_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestSubAccount_GetValidSubAccount(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	subAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/time-based-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(subAccountId), 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_subaccount.json"))
	})

	subAccount, err := underTest.GetSubAccount(subAccountId)
	assert.NoError(t, err)
	assert.NotNil(t, subAccount)
	assert.Equal(t, int32(1234567), subAccount.AccountId)
	assert.Empty(t, subAccount.Email)
	assert.Equal(t, "testAccount", subAccount.AccountName)
	assert.Equal(t, float32(1.0), subAccount.MaxDailyGB)
	assert.Zero(t, subAccount.ReservedDailyGB)
	assert.False(t, subAccount.Flexible)
	assert.Equal(t, int32(5), subAccount.RetentionDays)
	assert.True(t, subAccount.Searchable)
	assert.False(t, subAccount.Accessible)
	assert.True(t, subAccount.DocSizeSetting)
	assert.Zero(t, len(subAccount.SharingObjectsAccounts))
	assert.Equal(t, int32(5), subAccount.UtilizationSettings.FrequencyMinutes)
	assert.True(t, subAccount.UtilizationSettings.UtilizationEnabled)
}

func TestSubAccount_GetValidSubAccountNotFound(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	subAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/time-based-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(subAccountId), 10))
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fixture("get_subaccount_not_found.txt"))
	})

	subAccount, err := underTest.GetSubAccount(subAccountId)
	assert.Error(t, err)
	assert.Nil(t, subAccount)
	assert.Contains(t, err.Error(), "failed with missing sub account")
}
