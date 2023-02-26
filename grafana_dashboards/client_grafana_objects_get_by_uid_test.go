package grafana_dashboards_test

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/logzio/logzio_terraform_client/grafana_dashboards"
	"github.com/stretchr/testify/assert"
)

func TestGrafanaObjects_GetOK(t *testing.T) {
	underTest, teardown, err := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()
	dashboardId := int64(1234)

	mux.HandleFunc(dashboardsApiBasePath+"/uid/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(dashboardId, 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get.json"))
	})

	result, err := underTest.GetGrafanaDashboard(fmt.Sprint(dashboardId))
	assert.NoError(t, err)
	assert.Equal(t, result, &grafana_dashboards.GetResults{
		Dashboard: map[string]interface{}{
			"title": "getOK",
			"uid":   fmt.Sprint(dashboardId),
		},
		Meta: map[string]interface{}{
			"isStarred": true,
			"url":       "testUrl",
			"folderId":  float64(1),
			"folderUid": "testUid",
			"slug":      "testSlug",
		},
	},
	)
}

func TestGrafanaObjects_GetNotFound(t *testing.T) {
	underTest, teardown, err := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()
	dashboardId := int64(1234)

	mux.HandleFunc(dashboardsApiBasePath+"/uid/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(dashboardId, 10))
		w.WriteHeader(http.StatusNotFound)
	})

	dashboard, err := underTest.GetGrafanaDashboard(fmt.Sprint(dashboardId))
	assert.Error(t, err)
	assert.Nil(t, dashboard)
	assert.Contains(t, err.Error(), "failed with missing grafana dashboard")
}

func TestGrafanaObjects_GetError(t *testing.T) {
	underTest, teardown, err := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()
	dashboardId := int64(1234)

	mux.HandleFunc(dashboardsApiBasePath+"/uid/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(dashboardId, 10))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	})

	dashboard, err := underTest.GetGrafanaDashboard(fmt.Sprint(dashboardId))
	assert.Error(t, err)
	assert.Nil(t, dashboard)
}
