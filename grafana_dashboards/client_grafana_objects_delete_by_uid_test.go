package grafana_objects_test

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/logzio/logzio_terraform_client/grafana_objects"
	"github.com/stretchr/testify/assert"
)

func TestGrafanaObjects_DeleteOK(t *testing.T) {
	underTest, teardown, err := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	dashboardId := int64(1234)
	mux.HandleFunc(dashboardsApiBasePath+"/uid/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(dashboardId, 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("delete.json"))
	})

	result, err := underTest.Delete(fmt.Sprint(dashboardId))
	assert.NoError(t, err)
	assert.Equal(t, result, &grafana_objects.DeleteResults{
		Title:   "testDeleteOK",
		Message: "deleteOK",
		Id:      1234,
	})
}

func TestGrafanaObjects_DeleteNotFound(t *testing.T) {
	underTest, teardown, err := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	dashboardId := int64(1234)
	mux.HandleFunc(dashboardsApiBasePath+"/uid/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(dashboardId, 10))
		w.WriteHeader(http.StatusNotFound)
	})

	dashboard, err := underTest.Delete(fmt.Sprint(dashboardId))
	assert.Error(t, err)
	assert.Nil(t, dashboard)
	assert.Contains(t, err.Error(), "failed with missing grafana dashboard")
}

func TestGrafanaObjects_DeleteServerError(t *testing.T) {
	underTest, teardown, err := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	dashboardId := int64(1234)
	mux.HandleFunc(dashboardsApiBasePath+"/uid/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(dashboardId, 10))
		w.WriteHeader(http.StatusInternalServerError)
	})

	dashboard, err := underTest.Delete(fmt.Sprint(dashboardId))
	assert.Error(t, err)
	assert.Nil(t, dashboard)
}
