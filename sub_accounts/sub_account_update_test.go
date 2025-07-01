package sub_accounts_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/sub_accounts"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strconv"
	"testing"
)

func TestSubAccount_UpdateValidSubAccount(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	subAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/time-based-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(subAccountId), 10))
		jsonBytes, _ := io.ReadAll(r.Body)
		var target sub_accounts.CreateOrUpdateSubAccount
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		w.WriteHeader(http.StatusNoContent)
	})

	updateSubAccount := getCreateOrUpdateSubAccount("test@user.test")
	err = underTest.UpdateSubAccount(subAccountId, updateSubAccount)
	assert.NoError(t, err)
}

func TestSubAccount_UpdateSubAccountIdNotFound(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	subAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/time-based-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(subAccountId), 10))
		jsonBytes, _ := io.ReadAll(r.Body)
		var target sub_accounts.CreateOrUpdateSubAccount
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fixture("update_subaccount_not_fount.txt"))
	})

	updateSubAccount := getCreateOrUpdateSubAccount("test@user.test")
	err = underTest.UpdateSubAccount(subAccountId, updateSubAccount)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed with missing sub account")
}

func TestSubAccount_UpdateSubAccountWithWarmTier(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	subAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/time-based-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(subAccountId), 10))
		jsonBytes, _ := io.ReadAll(r.Body)
		var target sub_accounts.CreateOrUpdateSubAccount
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		w.WriteHeader(http.StatusNoContent)
	})

	updateSubAccount := getCreateOrUpdateSubAccount("test@user.test")
	warmRetention := int32(5)
	updateSubAccount.SnapSearchRetentionDays = &warmRetention
	err = underTest.UpdateSubAccount(subAccountId, updateSubAccount)
	assert.NoError(t, err)
}
