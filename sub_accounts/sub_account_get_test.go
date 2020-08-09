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
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_subaccount.json"))
	})

	subAccount, err := underTest.GetSubAccount(subAccountId)
	assert.NoError(t, err)
	assert.NotNil(t, subAccount)

	assert.Equal(t, "404 errors", subAccount.AccountName)
	assert.Equal(t, float32(100), subAccount.MaxDailyGB)
	assert.Equal(t, int32(5), subAccount.RetentionDays)
	assert.True(t, subAccount.Searchable)
	assert.True(t, subAccount.DocSizeSetting)
	assert.False(t, subAccount.Accessible)
	sharingAccountObjects := subAccount.SharingObjectAccounts[0].(map[string]interface{})
	assert.Equal(t, 2, len(sharingAccountObjects))
	assert.Equal(t, float64(88888), sharingAccountObjects["accountId"].(float64))
	assert.Equal(t, "dev group 8", sharingAccountObjects["accountName"].(string))

	//	"sharingObjectsAccounts": [
	//	{
	//		"accountId": 88888,
	//		"accountName": "dev group 8"
	//	}
	//],
	//"utilizationSettings": {
	//"frequencyMinutes": 5,
	//"utilizationEnabled": true
	//}
	//}

}
