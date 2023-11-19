package grafana_datasources_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGrafanaDatasourceClient_GetForAccount(t *testing.T) {
	underTest, teardown, err := setupGrafanaDatasourceTest()
	defer teardown()
	metricsAccountName := "my-metrics-account"

	mux.HandleFunc(fmt.Sprintf("/v1/grafana/api/datasources/name/%s/summary", metricsAccountName), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_for_account.json"))
	})

	grafanaDatasource, err := underTest.GetForAccount(metricsAccountName)
	assert.NoError(t, err)
	assert.NotNil(t, grafanaDatasource)
	assert.Equal(t, metricsAccountName, grafanaDatasource.Name)
	assert.Equal(t, int64(1234), grafanaDatasource.Id)
	assert.Equal(t, "some-uid", grafanaDatasource.Uid)
	assert.Equal(t, "elasticsearch", grafanaDatasource.Type)
	assert.Equal(t, "567890", grafanaDatasource.Database)
}

func TestGrafanaDatasourceClient_GetForAccountInternalServerError(t *testing.T) {
	underTest, teardown, err := setupGrafanaDatasourceTest()
	defer teardown()
	metricsAccountName := "my-metrics-account"

	mux.HandleFunc(fmt.Sprintf("/v1/grafana/api/datasources/name/%s/summary", metricsAccountName), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
	})

	grafanaDatasource, err := underTest.GetForAccount(metricsAccountName)
	assert.Error(t, err)
	assert.Nil(t, grafanaDatasource)
}
