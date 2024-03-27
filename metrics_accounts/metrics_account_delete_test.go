package metrics_accounts_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestMetricsAccount_DeleteValidMetricsAccount(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	metricsAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(metricsAccountId, 10))
		w.WriteHeader(200) //deleteAccountMetricsSuccess
	})

	err = underTest.DeleteMetricsAccount(metricsAccountId)
	assert.NoError(t, err)
}

func TestMetricsAccount_DeleteValidMetricsAccountNotFound(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	metricsAccountId := int64(12345)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(metricsAccountId, 10))
		w.WriteHeader(http.StatusNoContent) //deleteMetricsAccountMethodNotFound
		fmt.Fprint(w, fixture("delete_metrics_account_failed.txt"))
	})

	err = underTest.DeleteMetricsAccount(metricsAccountId)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed with missing metrics account")
}
