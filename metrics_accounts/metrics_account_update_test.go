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

func TestMetricsAccount_UpdateValidMetricsAccount(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	metricsAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			metricsAccount := &metrics_accounts.MetricsAccount{
				AccountName: "testAccount",
				Id:          int32(metricsAccountId),
				PlanUts:     100,
			}
			jsonBytes, _ := json.Marshal(metricsAccount)
			w.Write(jsonBytes)
			return
		}
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(metricsAccountId), 10))
		jsonBytes, _ := io.ReadAll(r.Body)
		var target metrics_accounts.CreateOrUpdateMetricsAccount
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		w.WriteHeader(200) //updateMetricsAccountServiceSuccess
	})

	updateMetricsAccount := getCreateOrUpdateMetricsAccount("test@user.test")
	err = underTest.UpdateMetricsAccount(metricsAccountId, updateMetricsAccount)
	assert.NoError(t, err)
}

func TestMetricsAccount_UpdateOnlySomeParamsCorrectly(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	metricsAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			metricsAccount := &metrics_accounts.MetricsAccount{
				AccountName: "testAccount",
				Id:          int32(metricsAccountId),
				PlanUts:     100,
			}
			jsonBytes, _ := json.Marshal(metricsAccount)
			w.Write(jsonBytes)
			return
		}
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(metricsAccountId), 10))
		jsonBytes, _ := io.ReadAll(r.Body)
		var target metrics_accounts.CreateOrUpdateMetricsAccount
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		w.WriteHeader(200) //updateMetricsAccountServiceSuccess
	})

	updateMetricsAccount := getCreateOrUpdateMetricsAccount("test@user.test")
	updateMetricsAccount.PlanUts = nil
	err = underTest.UpdateMetricsAccount(metricsAccountId, updateMetricsAccount)
	assert.NoError(t, err)
}

func TestMetricsAccount_UpdateMetricsAccountIdNotFound(t *testing.T) {
	underTest, err, teardown := setupMetricsAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	metricsAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/metrics-accounts/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			metricsAccount := &metrics_accounts.MetricsAccount{
				AccountName: "testAccount",
				Id:          int32(metricsAccountId),
				PlanUts:     100,
			}
			jsonBytes, _ := json.Marshal(metricsAccount)
			w.Write(jsonBytes)
			return
		}
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(metricsAccountId), 10))
		jsonBytes, _ := io.ReadAll(r.Body)
		var target metrics_accounts.CreateOrUpdateMetricsAccount
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fixture("update_metrics_account_not_fount.txt"))
	})

	updateMetricsAccount := getCreateOrUpdateMetricsAccount("test@user.test")
	err = underTest.UpdateMetricsAccount(metricsAccountId, updateMetricsAccount)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed with missing metrics account")
}
