package sub_accounts_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/sub_accounts"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestSubAccount_UpdateValidSubAccount(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	subAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/time-based-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(subAccountId), 10))

		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		//assert.Contains(t, target, "email")
		assert.Contains(t, target, "accountName")
		assert.Contains(t, target, "maxDailyGB")
		assert.Contains(t, target, "retentionDays")
		assert.Contains(t, target, "searchable")
		assert.Contains(t, target, "accessible")
		assert.Contains(t, target, "sharingObjectsAccounts")
		assert.Contains(t, target, "docSizeSetting")
		assert.Contains(t, target, "utilizationSettings")

		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_subaccount.json"))
	})

	sharingAccounts := make([]interface{}, 2)
	sharingAccounts[0] = 1
	sharingAccounts[1] = 2

	utilizationSettings := make(map[string]interface{}, 2)
	utilizationSettings["a"] = "v"

	s := sub_accounts.SubAccount{
		//Email:                 "test.user@test.user",
		AccountName:           "some account name",
		MaxDailyGB:            10.5,
		RetentionDays:         10,
		Searchable:            true,
		Accessible:            false,
		SharingObjectAccounts: sharingAccounts,
		DocSizeSetting:        true,
		UtilizationSettings:   utilizationSettings,
	}

	err = underTest.UpdateSubAccount(subAccountId, s)
	assert.NoError(t, err)
}
