package metrics_accounts_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/metrics_accounts"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strconv"
	"testing"
)

func TestMetricsAccount_GetSoftLimit(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	metricsAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(metricsAccountId, 10))
		assert.Contains(t, r.URL.String(), "soft-limit")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_metrics_account_soft_limit.json"))
	})

	softLimit, err := underTest.GetMetricsAccountSoftLimit(metricsAccountId)
	assert.NoError(t, err)
	assert.NotNil(t, softLimit)
	assert.Equal(t, int32(1234567), softLimit.MetricsAccountId)
	assert.NotNil(t, softLimit.SoftLimitUniqueMetrics)
	assert.Equal(t, int32(1000), *softLimit.SoftLimitUniqueMetrics)
}

// An account with no soft limit set returns null, which must be preserved rather than
// collapsed to zero - zero is a meaningful value on its own.
func TestMetricsAccount_GetSoftLimitNotSet(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	metricsAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_metrics_account_soft_limit_null.json"))
	})

	softLimit, err := underTest.GetMetricsAccountSoftLimit(metricsAccountId)
	assert.NoError(t, err)
	assert.NotNil(t, softLimit)
	assert.Nil(t, softLimit.SoftLimitUniqueMetrics)
}

func TestMetricsAccount_GetSoftLimitNotFound(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	metricsAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fixture("get_metrics_account_not_found.txt"))
	})

	softLimit, err := underTest.GetMetricsAccountSoftLimit(metricsAccountId)
	assert.Error(t, err)
	assert.Nil(t, softLimit)
}

// The endpoint is consumption-only; a subscription owner gets a 400 that must surface as an error.
func TestMetricsAccount_GetSoftLimitNotConsumptionAccount(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	metricsAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, fixture("metrics_account_soft_limit_not_consumption.txt"))
	})

	softLimit, err := underTest.GetMetricsAccountSoftLimit(metricsAccountId)
	assert.Error(t, err)
	assert.Nil(t, softLimit)
}

func TestMetricsAccount_UpdateSoftLimit(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	metricsAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(metricsAccountId, 10))
		assert.Contains(t, r.URL.String(), "soft-limit")

		// the wire field name is the contract with the API - assert it explicitly
		body, readErr := io.ReadAll(r.Body)
		assert.NoError(t, readErr)
		var sent map[string]any
		assert.NoError(t, json.Unmarshal(body, &sent))
		assert.Equal(t, float64(2500), sent["softLimitUniqueMetrics"])

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_metrics_account_soft_limit.json"))
	})

	softLimit, err := underTest.UpdateMetricsAccountSoftLimit(metricsAccountId,
		metrics_accounts.UpdateMetricsAccountSoftLimit{SoftLimitUniqueMetrics: 2500})
	assert.NoError(t, err)
	assert.NotNil(t, softLimit)
	assert.Equal(t, int32(1234567), softLimit.MetricsAccountId)
	assert.NotNil(t, softLimit.SoftLimitUniqueMetrics)
	assert.Equal(t, int32(2500), *softLimit.SoftLimitUniqueMetrics)
}

// Zero is accepted by the API and must not be rejected client side. It is NOT a removal: it is
// stored as-is, and because the limiter only applies a soft limit greater than the plan UTS, a
// stored 0 is silently ignored rather than capping the account at zero.
func TestMetricsAccount_UpdateSoftLimitZeroIsStoredNotARemoval(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_metrics_account_soft_limit.json"))
	})

	_, err = underTest.UpdateMetricsAccountSoftLimit(int64(1234567),
		metrics_accounts.UpdateMetricsAccountSoftLimit{SoftLimitUniqueMetrics: 0})
	assert.NoError(t, err)
}

// Mirrors the API validation, which rejects a negative soft limit with 400 INVALID_SOFT_LIMIT.
func TestMetricsAccount_UpdateSoftLimitNegative(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	called := false
	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	softLimit, err := underTest.UpdateMetricsAccountSoftLimit(int64(1234567),
		metrics_accounts.UpdateMetricsAccountSoftLimit{SoftLimitUniqueMetrics: -1})
	assert.Error(t, err)
	assert.Nil(t, softLimit)
	assert.Contains(t, err.Error(), "softLimitUniqueMetrics should be non-negative")
	assert.False(t, called, "no request should be sent when validation fails")
}

func TestMetricsAccount_UpdateSoftLimitNotConsumptionAccount(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, fixture("metrics_account_soft_limit_not_consumption.txt"))
	})

	softLimit, err := underTest.UpdateMetricsAccountSoftLimit(int64(1234567),
		metrics_accounts.UpdateMetricsAccountSoftLimit{SoftLimitUniqueMetrics: 1000})
	assert.Error(t, err)
	assert.Nil(t, softLimit)
}
