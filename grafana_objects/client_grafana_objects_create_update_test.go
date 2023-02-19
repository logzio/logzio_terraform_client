package grafana_objects_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/grafana_objects"
	"github.com/stretchr/testify/assert"
)

func TestGrafanaObjects_CreateUpdateOK(t *testing.T) {
	underTest, teardown, err := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	createDashboard := getCreateUpdateDashboard()
	createDashboard.Dashboard.Title += "_create"
	createDashboard.Dashboard.Uid += "test1"

	mux.HandleFunc(dashboardsApiBasePath+"/db", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		var target grafana_objects.CreateUpdatePayload
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		assert.NotNil(t, target.Dashboard)
		assert.Equal(t, createDashboard, target)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("createupdate_ok_resp.json"))
	})

	resp, err := underTest.CreateUpdate(createDashboard)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotZero(t, resp.Id)
	assert.NotEmpty(t, resp.Uid)
	assert.Equal(t, grafana_objects.GrafanaSuccessStatus, resp.Status)
}

func TestGrafanaObjects_CreateUpdateNOKPreconditionFailed(t *testing.T) {
	underTest, teardown, err := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	createDashboard := getCreateUpdateDashboard()
	createDashboard.Dashboard.Title += "_create_412"
	createDashboard.Dashboard.Uid += "_create_412"

	mux.HandleFunc(dashboardsApiBasePath+"/db", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		var target grafana_objects.CreateUpdatePayload
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusPreconditionFailed)
		fmt.Fprint(w, fixture("createupdate_nok_resp_412.json"))
	})

	dashboard, err := underTest.CreateUpdate(createDashboard)
	assert.Error(t, err)
	assert.Nil(t, dashboard)
}

func TestGrafanaObjects_CreateUpdateNOKNotFound(t *testing.T) {
	underTest, teardown, err := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	createDashboard := getCreateUpdateDashboard()
	createDashboard.Dashboard.Title += "_update_404"
	createDashboard.Dashboard.Uid += "_update_404"

	mux.HandleFunc(dashboardsApiBasePath+"/db", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		var target grafana_objects.CreateUpdatePayload
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)

		w.WriteHeader(http.StatusNotFound)
	})

	dashboard, err := underTest.CreateUpdate(createDashboard)
	assert.Error(t, err)
	assert.Nil(t, dashboard)
	assert.Contains(t, err.Error(), "failed with missing grafana dashboard")
}

func TestGrafanaObjects_CreateUpdateNOKApiFail(t *testing.T) {
	underTest, teardown, err := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	createDashboard := getCreateUpdateDashboard()
	createDashboard.Dashboard.Title += "_create_500"
	createDashboard.Dashboard.Uid += "_create_500"

	mux.HandleFunc(dashboardsApiBasePath+"/db", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		var target grafana_objects.CreateUpdatePayload
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)

		w.WriteHeader(http.StatusInternalServerError)
	})

	dashboard, err := underTest.CreateUpdate(createDashboard)
	assert.Error(t, err)
	assert.Nil(t, dashboard)
}
