package tracing_accounts_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/tracing_accounts"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strconv"
	"testing"
)

func TestTracingAccount_New(t *testing.T) {
	_, err := tracing_accounts.New("", "https://api.logz.io")
	assert.Error(t, err)

	_, err = tracing_accounts.New("SOME_API_TOKEN", "")
	assert.Error(t, err)

	underTest, err := tracing_accounts.New("SOME_API_TOKEN", "https://api.logz.io")
	assert.NoError(t, err)
	assert.NotNil(t, underTest)
}

func TestTracingAccount_Create(t *testing.T) {
	underTest, err, teardown := setupTracingAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/account-management/consumption/tracing-accounts", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)

		body, readErr := io.ReadAll(r.Body)
		assert.NoError(t, readErr)
		var sent map[string]any
		assert.NoError(t, json.Unmarshal(body, &sent))
		// the wire field names are the contract with the API
		assert.Equal(t, "tf_client_test_tracing", sent["accountName"])
		assert.Equal(t, float64(5), sent["maxDailyGB"])
		assert.Equal(t, "some@email.test", sent["email"])
		assert.NotNil(t, sent["authorizedAccountIds"])

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_tracing_account.json"))
	})

	account, err := underTest.CreateTracingAccount(validCreateTracingAccount())
	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, int32(7654321), account.AccountId)
	assert.Equal(t, "tf_client_test_tracing", account.AccountName)
	assert.Equal(t, int32(14), account.Retention)
	assert.Len(t, account.AuthorizedAccounts, 1)
}

func TestTracingAccount_CreateValidation(t *testing.T) {
	underTest, err, teardown := setupTracingAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	called := false
	mux.HandleFunc("/v1/account-management/consumption/tracing-accounts", func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	noEmail := validCreateTracingAccount()
	noEmail.Email = ""
	_, err = underTest.CreateTracingAccount(noEmail)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "email must be set")

	noName := validCreateTracingAccount()
	noName.AccountName = ""
	_, err = underTest.CreateTracingAccount(noName)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "accountName must be set")

	// nil rather than an empty slice would serialise as null and be rejected by the API
	nilAccounts := validCreateTracingAccount()
	nilAccounts.AuthorizedAccountIds = nil
	_, err = underTest.CreateTracingAccount(nilAccounts)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "authorizedAccountIds must be initialized")

	negative := validCreateTracingAccount()
	negative.MaxDailyGB = -1
	_, err = underTest.CreateTracingAccount(negative)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "maxDailyGB should be non-negative")

	assert.False(t, called, "no request should be sent when validation fails")
}

func TestTracingAccount_Get(t *testing.T) {
	underTest, err, teardown := setupTracingAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	tracingAccountId := int64(7654321)

	mux.HandleFunc("/v1/account-management/consumption/tracing-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(tracingAccountId, 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_tracing_account.json"))
	})

	account, err := underTest.GetTracingAccount(tracingAccountId)
	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, int32(7654321), account.AccountId)
	assert.NotNil(t, account.MaxDailyGB)
	assert.Equal(t, float32(5), *account.MaxDailyGB)
}

func TestTracingAccount_GetNotFound(t *testing.T) {
	underTest, err, teardown := setupTracingAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/account-management/consumption/tracing-accounts/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fixture("tracing_account_not_found.txt"))
	})

	account, err := underTest.GetTracingAccount(int64(7654321))
	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestTracingAccount_List(t *testing.T) {
	underTest, err, teardown := setupTracingAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/account-management/consumption/tracing-accounts", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("list_tracing_accounts.json"))
	})

	accounts, err := underTest.ListTracingAccounts()
	assert.NoError(t, err)
	assert.Len(t, accounts, 2)
	assert.Equal(t, int32(7654321), accounts[0].AccountId)
	// an account with no cap returns null, which must stay distinguishable from zero
	assert.Nil(t, accounts[1].MaxDailyGB)
}

func TestTracingAccount_Update(t *testing.T) {
	underTest, err, teardown := setupTracingAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	tracingAccountId := int64(7654321)

	mux.HandleFunc("/v1/account-management/consumption/tracing-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(tracingAccountId, 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_tracing_account.json"))
	})

	account, err := underTest.UpdateTracingAccount(tracingAccountId, validCreateTracingAccount())
	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, int32(7654321), account.AccountId)
}

func TestTracingAccount_Delete(t *testing.T) {
	underTest, err, teardown := setupTracingAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	tracingAccountId := int64(7654321)

	mux.HandleFunc("/v1/account-management/consumption/tracing-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(tracingAccountId, 10))
		w.WriteHeader(http.StatusOK)
	})

	err = underTest.DeleteTracingAccount(tracingAccountId)
	assert.NoError(t, err)
}
