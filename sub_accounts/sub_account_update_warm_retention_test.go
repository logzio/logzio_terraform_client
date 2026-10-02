package sub_accounts_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/sub_accounts"
	"github.com/stretchr/testify/assert"
)

func TestSubAccount_UpdateWarmRetention(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	accountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/time-based-accounts/1234567/retention-details", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)

		// the wire field name is the contract with the API - assert it explicitly
		body, readErr := io.ReadAll(r.Body)
		assert.NoError(t, readErr)
		var sent map[string]any
		assert.NoError(t, json.Unmarshal(body, &sent))
		assert.Equal(t, map[string]any{"snapSearchRetentionDays": float64(3)}, sent)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_warm_retention.json"))
	})

	retentionDetails, err := underTest.UpdateWarmRetention(accountId, sub_accounts.UpdateWarmRetention{SnapSearchRetentionDays: 3})
	assert.NoError(t, err)
	assert.Len(t, retentionDetails, 3)
	assert.Equal(t, int32(1234567), retentionDetails[0].AccountId)
	assert.Equal(t, int32(7), retentionDetails[0].RetentionDays)
	assert.NotNil(t, retentionDetails[0].SnapSearchRetentionDays)
	assert.Equal(t, int32(3), *retentionDetails[0].SnapSearchRetentionDays)
	// an account that never had warm retention comes back as null and must stay nil, not 0
	assert.Nil(t, retentionDetails[1].SnapSearchRetentionDays)
}

// 0 turns warm tier off for a sub account, so it must be sent rather than dropped.
func TestSubAccount_UpdateWarmRetentionZero(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/account-management/time-based-accounts/1234567/retention-details", func(w http.ResponseWriter, r *http.Request) {
		body, readErr := io.ReadAll(r.Body)
		assert.NoError(t, readErr)
		assert.JSONEq(t, `{"snapSearchRetentionDays": 0}`, string(body))

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_warm_retention.json"))
	})

	_, err = underTest.UpdateWarmRetention(int64(1234567), sub_accounts.UpdateWarmRetention{SnapSearchRetentionDays: 0})
	assert.NoError(t, err)
}

func TestSubAccount_UpdateWarmRetentionNegative(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	called := false
	mux.HandleFunc("/v1/account-management/time-based-accounts/", func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	retentionDetails, err := underTest.UpdateWarmRetention(int64(1234567), sub_accounts.UpdateWarmRetention{SnapSearchRetentionDays: -1})
	assert.Error(t, err)
	assert.Nil(t, retentionDetails)
	assert.Contains(t, err.Error(), "snapSearchRetentionDays should be non-negative")
	assert.False(t, called, "no request should be sent when validation fails")
}

func TestSubAccount_UpdateWarmRetentionRejected(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/account-management/time-based-accounts/1234567/retention-details", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, fixture("update_warm_retention_invalid.txt"))
	})

	retentionDetails, err := underTest.UpdateWarmRetention(int64(1234567), sub_accounts.UpdateWarmRetention{SnapSearchRetentionDays: 6})
	assert.Error(t, err)
	assert.Nil(t, retentionDetails)
	assert.Contains(t, err.Error(), "INVALID_WARM_RETENTION")
}

func TestSubAccount_UpdateWarmRetentionNotFound(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/account-management/time-based-accounts/1234567/retention-details", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fixture("update_warm_retention_not_found.txt"))
	})

	retentionDetails, err := underTest.UpdateWarmRetention(int64(1234567), sub_accounts.UpdateWarmRetention{SnapSearchRetentionDays: 2})
	assert.Error(t, err)
	assert.Nil(t, retentionDetails)
}
