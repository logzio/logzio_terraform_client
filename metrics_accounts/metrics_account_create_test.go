package metrics_accounts_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/metrics_accounts"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestMetricsAccount_CreateValidMetricsAccount(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/account-management/metrics-accounts", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target metrics_accounts.CreateOrUpdateMetricsAccount
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Email)
			assert.NotEmpty(t, target.AccountName)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("create_metrics_account.json"))
		})

		createMetricsAccount := getCreateOrUpdateMetricsAccount("test.user@test.user")
		createMetricsAccount.AccountName = createMetricsAccount.AccountName + "_test_create"
		metricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
		assert.NoError(t, err)
		assert.NotNil(t, metricsAccount)
		assert.NotEmpty(t, metricsAccount.Id)
		assert.NotEmpty(t, metricsAccount.Token)
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

		createMetricsAccount := getCreateOrUpdateMetricsAccount("test.user@test.user")
		createMetricsAccount.AccountName = createMetricsAccount.AccountName + "_test_create"
		MetricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
		assert.Error(t, err)
		assert.Nil(t, MetricsAccount)
	}
}

func TestMetricsAccount_CreateMetricsAccountNoEmail(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	defer teardown()

	createMetricsAccount := getCreateOrUpdateMetricsAccount("")
	createMetricsAccount.AccountName = createMetricsAccount.AccountName + "_test_create_no_email"
	MetricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
	assert.Error(t, err)
	assert.Nil(t, MetricsAccount)
}

func TestMetricsAccount_CreateMetricsAccountNoAccountName(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/account-management/metrics-accounts", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target metrics_accounts.CreateOrUpdateMetricsAccount
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Email)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("create_metrics_account_generatename.json"))
		})

		createMetricsAccount := getCreateOrUpdateMetricsAccount("test.user@test.user")
		createMetricsAccount.AccountName = ""
		MetricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
		assert.NoError(t, err)
		assert.NotNil(t, MetricsAccount)
		assert.NotEmpty(t, MetricsAccount.AccountName)
	}
}

func TestMetricsAccount_CreateMetricsAccountNoPlanUts(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	defer teardown()

	createMetricsAccount := getCreateOrUpdateMetricsAccount("test.user@test.user")
	createMetricsAccount.AccountName = createMetricsAccount.AccountName + "_test_create_no_planuts"
	*createMetricsAccount.PlanUts = 0
	MetricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
	assert.Error(t, err)
	assert.Nil(t, MetricsAccount)
}

func TestMetricsAccount_CreateMetricsAccountNoSharingAccount(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	defer teardown()

	createMetricsAccount := getCreateOrUpdateMetricsAccount("test.user@test.user")
	createMetricsAccount.AccountName = createMetricsAccount.AccountName + "_test_create_no_sharing"
	createMetricsAccount.AuthorizedAccountsIds = nil
	MetricsAccount, err := underTest.CreateMetricsAccount(createMetricsAccount)
	assert.Error(t, err)
	assert.Nil(t, MetricsAccount)
}
