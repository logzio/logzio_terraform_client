package sub_accounts_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/sub_accounts"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSubAccount_CreateValidSubAccount(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/account-management/time-based-accounts", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target sub_accounts.CreateOrUpdateSubAccount
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Email)
			assert.NotEmpty(t, target.AccountName)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("create_subaccount.json"))
		})

		createSubAccount := getCreatrOrUpdateSubAccount("test.user@test.user")
		createSubAccount.AccountName = createSubAccount.AccountName + "_test_create"
		subAccount, err := underTest.CreateSubAccount(createSubAccount)
		assert.NoError(t, err)
		assert.NotNil(t, subAccount)
		assert.NotEmpty(t, subAccount.AccountId)
		assert.NotEmpty(t, subAccount.AccountToken)
	}
}

func TestSubAccount_CreateValidSubAccountAPIFail(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/account-management/time-based-accounts", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("create_subaccount_failed.txt"))
		})

		createSubAccount := getCreatrOrUpdateSubAccount("test.user@test.user")
		createSubAccount.AccountName = createSubAccount.AccountName + "_test_create"
		subAccount, err := underTest.CreateSubAccount(createSubAccount)
		assert.Error(t, err)
		assert.Nil(t, subAccount)
	}
}

func TestSubAccount_CreateSubAccountNoEmail(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	defer teardown()

	createSubAccount := getCreatrOrUpdateSubAccount("")
	createSubAccount.AccountName = createSubAccount.AccountName + "_test_create_no_email"
	subAccount, err := underTest.CreateSubAccount(createSubAccount)
	assert.Error(t, err)
	assert.Nil(t, subAccount)
}

func TestSubAccount_CreateSubAccountNoAccountName(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	defer teardown()

	createSubAccount := getCreatrOrUpdateSubAccount("test.user@test.user")
	createSubAccount.AccountName = ""
	subAccount, err := underTest.CreateSubAccount(createSubAccount)
	assert.Error(t, err)
	assert.Nil(t, subAccount)
}

func TestSubAccount_CreateSubAccountNoRetentionDays(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	defer teardown()

	createSubAccount := getCreatrOrUpdateSubAccount("test.user@test.user")
	createSubAccount.AccountName = createSubAccount.AccountName + "_test_create_no_retention"
	createSubAccount.RetentionDays = 0
	subAccount, err := underTest.CreateSubAccount(createSubAccount)
	assert.Error(t, err)
	assert.Nil(t, subAccount)
}

func TestSubAccount_CreateSubAccountNoSharingAccount(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	defer teardown()

	createSubAccount := getCreatrOrUpdateSubAccount("test.user@test.user")
	createSubAccount.AccountName = createSubAccount.AccountName + "_test_create_no_sharing"
	createSubAccount.SharingObjectsAccounts = nil
	subAccount, err := underTest.CreateSubAccount(createSubAccount)
	assert.Error(t, err)
	assert.Nil(t, subAccount)
}
