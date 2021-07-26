package sub_accounts_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSubAccount_ListSubAccounts(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/account-management/time-based-accounts", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("list_subaccounts.json"))
	})

	subAccounts, err := underTest.ListSubAccounts()
	assert.NoError(t, err)
	assert.NotNil(t, subAccounts)
	assert.Equal(t, 1, len(subAccounts))
}

func TestSubAccount_ListSubAccountsAPIFail(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/account-management/time-based-accounts", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("list_subaccounts_failed.txt"))
	})

	subAccounts, err := underTest.ListSubAccounts()
	assert.Error(t, err)
	assert.Nil(t, subAccounts)
}
