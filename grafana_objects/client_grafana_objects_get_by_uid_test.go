package grafana_objects_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/grafana_objects"
	"github.com/stretchr/testify/assert"
)

func TestGrafanaObjects_GetOK(t *testing.T) {
	underTest, err, teardown := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/grafana/api/dashboards/uid/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fileGet, err := ioutil.ReadFile("testdata/fixtures/get.json")
		assert.NoError(t, err)

		resp := grafana_objects.GetResults{}
		err = json.Unmarshal([]byte(fileGet), &resp)
		assert.NoError(t, err)

		bytes, err := json.Marshal(resp)
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	},
	)

	result, err := underTest.Get("getOK")
	assert.NoError(t, err)
	assert.Equal(t, result, &grafana_objects.GetResults{
		Dashboard: grafana_objects.DashboardObject{
			Title: "getOK",
			Uid:   "test1",
		},
		Meta: grafana_objects.DashboardMeta{
			IsStarred: true,
			Url:       "testUrl",
			FolderId:  1,
			FolderUid: "testUid",
			Slug:      "testSlug",
		},
	},
	)
}

func TestGrafanaObjects_GetNotFound(t *testing.T) {
	underTest, err, teardown := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/grafana/api/dashboards/uid/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("Content&grafana_objects.GetResults{-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Dashboard Not found"

		bytes, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusNotFound)
		w.Write(bytes)
	},
	)

	_, err = underTest.Get("getNOK")
	assert.Error(t, err)
}

func TestGrafanaObjects_GetError(t *testing.T) {
	underTest, err, teardown := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/grafana/api/dashboards/uid/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusInternalServerError)
	},
	)

	_, err = underTest.Get("getNOK")
	assert.Error(t, err)
}
