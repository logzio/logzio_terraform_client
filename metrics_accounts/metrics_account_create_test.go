package metrics_accounts_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/metrics_accounts"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMetricsAccount_CreateValidMetricsAccount(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/account-management/metrics-accounts", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target metrics_accounts.CreateOrUpdateMetricsAccount
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Email)
			assert.NotEmpty(t, target.AccountName)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("create_metrics_account.json"))
		})

		createSubAccount := getCreateOrUpdateMetricsAccount("test.user@test.user")
		createSubAccount.AccountName = createSubAccount.AccountName + "_test_create"
		subAccount, err := underTest.CreateMetricsAccount(createSubAccount)
		assert.NoError(t, err)
		assert.NotNil(t, subAccount)
		assert.NotEmpty(t, subAccount.Id)
		assert.NotEmpty(t, subAccount.Token)
	}
}

func TestMetricsAccount_CreateValidMetricsAccountAPIFail(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/account-management/metrics-accounts", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("create_metrics_account_failed.txt"))
		})

		createSubAccount := getCreateOrUpdateMetricsAccount("test.user@test.user")
		createSubAccount.AccountName = createSubAccount.AccountName + "_test_create"
		subAccount, err := underTest.CreateMetricsAccount(createSubAccount)
		assert.Error(t, err)
		assert.Nil(t, subAccount)
	}
}

func TestMetricsAccount_CreateMetricsAccountNoEmail(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	defer teardown()

	createSubAccount := getCreateOrUpdateMetricsAccount("")
	createSubAccount.AccountName = createSubAccount.AccountName + "_test_create_no_email"
	subAccount, err := underTest.CreateMetricsAccount(createSubAccount)
	assert.Error(t, err)
	assert.Nil(t, subAccount)
}

func TestMetricsAccount_CreateMetricsAccountNoAccountName(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/account-management/metrics-accounts", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target metrics_accounts.CreateOrUpdateMetricsAccount
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Email)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("create_metrics_account_generatename.json"))
		})

		createSubAccount := getCreateOrUpdateMetricsAccount("test.user@test.user")
		createSubAccount.AccountName = ""
		subAccount, err := underTest.CreateMetricsAccount(createSubAccount)
		assert.NoError(t, err)
		assert.NotNil(t, subAccount)
		assert.NotEmpty(t, subAccount.AccountName)
	}
}

func TestMetricsAccount_CreateMetricsAccountNoPlanUts(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	defer teardown()

	createSubAccount := getCreateOrUpdateMetricsAccount("test.user@test.user")
	createSubAccount.AccountName = createSubAccount.AccountName + "_test_create_no_planuts"
	createSubAccount.PlanUts = -1
	subAccount, err := underTest.CreateMetricsAccount(createSubAccount)
	assert.Error(t, err)
	assert.Nil(t, subAccount)
}

func TestMetricsAccount_CreateMetricsAccountNoSharingAccount(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	defer teardown()

	createSubAccount := getCreateOrUpdateMetricsAccount("test.user@test.user")
	createSubAccount.AccountName = createSubAccount.AccountName + "_test_create_no_sharing"
	createSubAccount.AuthorizedAccountsIds = nil
	subAccount, err := underTest.CreateMetricsAccount(createSubAccount)
	assert.Error(t, err)
	assert.Nil(t, subAccount)
}
