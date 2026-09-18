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

func TestTracingAccount_GetSoftLimit(t *testing.T) {
	underTest, err, teardown := setupTracingAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	tracingAccountId := int64(7654321)

	mux.HandleFunc("/v1/account-management/consumption/tracing-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(tracingAccountId, 10))
		assert.Contains(t, r.URL.String(), "soft-limit")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_tracing_account_soft_limit.json"))
	})

	softLimit, err := underTest.GetTracingAccountSoftLimit(tracingAccountId)
	assert.NoError(t, err)
	assert.NotNil(t, softLimit)
	assert.Equal(t, int32(7654321), softLimit.AccountId)
	assert.NotNil(t, softLimit.SoftLimitGB)
	assert.Equal(t, float32(5), *softLimit.SoftLimitGB)
}

// No soft limit set returns null, which must be preserved rather than collapsed to zero -
// zero is a meaningful value on its own.
func TestTracingAccount_GetSoftLimitNotSet(t *testing.T) {
	underTest, err, teardown := setupTracingAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/account-management/consumption/tracing-accounts/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_tracing_account_soft_limit_null.json"))
	})

	softLimit, err := underTest.GetTracingAccountSoftLimit(int64(7654321))
	assert.NoError(t, err)
	assert.NotNil(t, softLimit)
	assert.Nil(t, softLimit.SoftLimitGB)
}

// The endpoint is consumption-only; a subscription owner gets a 400 that must surface as an error.
func TestTracingAccount_GetSoftLimitNotConsumptionAccount(t *testing.T) {
	underTest, err, teardown := setupTracingAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/account-management/consumption/tracing-accounts/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, fixture("tracing_account_not_consumption.txt"))
	})

	softLimit, err := underTest.GetTracingAccountSoftLimit(int64(7654321))
	assert.Error(t, err)
	assert.Nil(t, softLimit)
}

func TestTracingAccount_UpdateSoftLimit(t *testing.T) {
	underTest, err, teardown := setupTracingAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	tracingAccountId := int64(7654321)

	mux.HandleFunc("/v1/account-management/consumption/tracing-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), "soft-limit")

		body, readErr := io.ReadAll(r.Body)
		assert.NoError(t, readErr)
		var sent map[string]any
		assert.NoError(t, json.Unmarshal(body, &sent))
		// the API requires the id in the body as well as the path
		assert.Equal(t, float64(7654321), sent["tracingAccountId"])
		assert.Equal(t, 12.5, sent["softLimitGB"])

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_tracing_account_soft_limit.json"))
	})

	softLimit, err := underTest.UpdateTracingAccountSoftLimit(tracingAccountId,
		tracing_accounts.UpdateTracingAccountSoftLimit{TracingAccountId: 7654321, SoftLimitGB: 12.5})
	assert.NoError(t, err)
	assert.NotNil(t, softLimit)
	assert.NotNil(t, softLimit.SoftLimitGB)
	assert.Equal(t, float32(12.5), *softLimit.SoftLimitGB)
}

// The API rejects a body id that does not match the path id with 400 WRONG_ACCOUNT_ID,
// so the client catches it before sending.
func TestTracingAccount_UpdateSoftLimitIdMismatch(t *testing.T) {
	underTest, err, teardown := setupTracingAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	called := false
	mux.HandleFunc("/v1/account-management/consumption/tracing-accounts/", func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	softLimit, err := underTest.UpdateTracingAccountSoftLimit(int64(7654321),
		tracing_accounts.UpdateTracingAccountSoftLimit{TracingAccountId: 999, SoftLimitGB: 5})
	assert.Error(t, err)
	assert.Nil(t, softLimit)
	assert.Contains(t, err.Error(), "must match the tracing account id in the path")
	assert.False(t, called, "no request should be sent when validation fails")
}

// Mirrors the API validation, which rejects a negative soft limit with 400 ILLEGAL_SOFT_LIMIT.
func TestTracingAccount_UpdateSoftLimitNegative(t *testing.T) {
	underTest, err, teardown := setupTracingAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	called := false
	mux.HandleFunc("/v1/account-management/consumption/tracing-accounts/", func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	softLimit, err := underTest.UpdateTracingAccountSoftLimit(int64(7654321),
		tracing_accounts.UpdateTracingAccountSoftLimit{TracingAccountId: 7654321, SoftLimitGB: -1})
	assert.Error(t, err)
	assert.Nil(t, softLimit)
	assert.Contains(t, err.Error(), "softLimitGB should be non-negative")
	assert.False(t, called, "no request should be sent when validation fails")
}
