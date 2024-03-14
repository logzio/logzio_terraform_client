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

	metricsAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(metricsAccountId), 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_metrics_account.json"))
	})

	metricsAccount, err := underTest.GetMetricsAccount(metricsAccountId)
	assert.NoError(t, err)
	assert.NotNil(t, metricsAccount)
	assert.Equal(t, int32(1234567), metricsAccount.Id)
	assert.Equal(t, "testAccount", metricsAccount.AccountName)
	assert.Equal(t, int32(5), metricsAccount.PlanUts)
	assert.Equal(t, len(metricsAccount.AuthorizedAccountsIds), 1)
}

func TestMetricsAccount_GetValidMetricsAccountNotFound(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	metricsAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(metricsAccountId), 10))
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fixture("get_metrics_account_not_found.txt"))
	})

	metricsAccount, err := underTest.GetMetricsAccount(metricsAccountId)
	assert.Error(t, err)
	assert.Nil(t, metricsAccount)
	assert.Contains(t, err.Error(), "failed with missing metrics account")
}
