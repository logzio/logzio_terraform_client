package sub_accounts_test

import (
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client/sub_accounts"
	"github.com/jonboydell/logzio_client/users"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSubAccount_CreateValidSubAccount(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/account-management/time-based-accounts", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)

		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		assert.Contains(t, target, "email")
		assert.Contains(t, target, "accountName")
		assert.Contains(t, target, "maxDailyGB")
		assert.Contains(t, target, "retentionDays")
		assert.Contains(t, target, "searchable")
		assert.Contains(t, target, "accessible")
		assert.Contains(t, target, "sharingObjectsAccounts")
		assert.Contains(t, target, "docSizeSetting")
		assert.Contains(t, target, "utilizationSettings")

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_subaccount.json"))
		w.WriteHeader(http.StatusOK)
	})

	subAccount := sub_accounts.SubAccount {
	}

	user, err := underTest.CreateSubAccount(subAccount)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}