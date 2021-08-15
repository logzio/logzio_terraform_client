package sub_accounts_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSubAccount_ListDetailedSubAccounts(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/account-management/time-based-accounts/detailed", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("list_subaccounts_detailed.json"))
	})

	subAccounts, err := underTest.ListDetailedSubAccounts()
	assert.NoError(t, err)
	assert.NotNil(t, subAccounts)
	assert.Equal(t, 1, len(subAccounts))
}

func TestSubAccount_ListDetailedSubAccountsAPIFail(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/account-management/time-based-accounts/detailed", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("list_subaccounts_detailed_failed.txt"))
	})

	subAccounts, err := underTest.ListDetailedSubAccounts()
	assert.Error(t, err)
	assert.Nil(t, subAccounts)
}
