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
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_detailed_subaccount.json"))
	})

	subAccount, err := underTest.GetDetailedSubAccount(subAccountId)
	assert.NoError(t, err)
	assert.NotNil(t, subAccount)

	assert.Equal(t, "testAccount", subAccount.Account.AccountName)
	assert.Equal(t, int64(12345), subAccount.Account.AccountId)
	assert.Equal(t, "testToken", subAccount.Account.AccountToken)
	assert.Equal(t, "testIndex", subAccount.Account.EsIndexPrefix)
	assert.Equal(t, int64(12), subAccount.Account.MaxDailyGB)
	assert.Equal(t, int64(123), subAccount.Account.RetentionDays)
	assert.Equal(t, float64(5), subAccount.UtilizationSettings["frequencyMinutes"])
	assert.True(t, subAccount.SubAccountRelation.Searchable)
	assert.Equal(t, int64(1584532201), subAccount.SubAccountRelation.CreatedDate)
	assert.Equal(t, int64(1584532202), subAccount.SubAccountRelation.LastUpdatedDate)
	assert.True(t, subAccount.DocSizeSetting)
	assert.False(t, subAccount.SubAccountRelation.Accessible)
	sharingAccountObjects := subAccount.SharingObjectAccounts[0].(map[string]interface{})
	assert.Equal(t, 7, len(sharingAccountObjects))
}
