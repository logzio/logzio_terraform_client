package metrics_accounts_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestMetricsAccount_GetValidMetricsAccount(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	subAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(subAccountId), 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_metrics_account.json"))
	})

	subAccount, err := underTest.GetMetricsAccount(subAccountId)
	assert.NoError(t, err)
	assert.NotNil(t, subAccount)
	assert.Equal(t, int32(1234567), subAccount.Id)
	//assert.Empty(t, subAccount.Email)
	assert.Equal(t, "testAccount", subAccount.AccountName)
	assert.Equal(t, int32(5), subAccount.PlanUts)
	assert.Equal(t, len(subAccount.AuthorizedAccountsIds), 1)
}

func TestMetricsAccount_GetValidMetricsAccountNotFound(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	subAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(subAccountId), 10))
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fixture("get_metrics_account_not_found.txt"))
	})

	subAccount, err := underTest.GetMetricsAccount(subAccountId)
	assert.Error(t, err)
	assert.Nil(t, subAccount)
	assert.Contains(t, err.Error(), "failed with missing metrics account")
}
