package metrics_accounts_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/metrics_accounts"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestMetricsAccount_UpdateValidMetricsAccount(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	subAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(subAccountId), 10))
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target metrics_accounts.CreateOrUpdateMetricsAccount
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		w.WriteHeader(200) //updateMetricsAccountServiceSuccess
	})

	updateSubAccount := getCreateOrUpdateMetricsAccount("test@user.test")
	err = underTest.UpdateMetricsAccount(subAccountId, updateSubAccount)
	assert.NoError(t, err)
}

func TestMetricsAccount_UpdateMetricsAccountIdNotFound(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	subAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(subAccountId), 10))
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target metrics_accounts.CreateOrUpdateMetricsAccount
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fixture("update_metrics_account_not_fount.txt"))
	})

	updateSubAccount := getCreateOrUpdateMetricsAccount("test@user.test")
	err = underTest.UpdateMetricsAccount(subAccountId, updateSubAccount)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed with missing metrics account")
}
