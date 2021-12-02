package grafana_objects_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/grafana_objects"
	"github.com/stretchr/testify/assert"
)

func TestGrafanaObjects_CreateUpdateOK(t *testing.T) {
	underTest, err, teardown := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/grafana/api/dashboards/db", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)
		var payload grafana_objects.CreateUpdatePayload
		err = json.Unmarshal(jsonBytes, &payload)
		assert.NoError(t, err)
		assert.NotNil(t, payload)

		fileGet, err := ioutil.ReadFile("testdata/fixtures/createupdate_ok_resp.json")
		var resp grafana_objects.CreateUpdateResults
		err = json.Unmarshal([]byte(fileGet), &resp)
		assert.NoError(t, err)

		bytes, err := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
		assert.NoError(t, err)
	})

	file, err := ioutil.ReadFile("testdata/fixtures/createupdate_ok.json")
	var payload grafana_objects.CreateUpdatePayload
	err = json.Unmarshal([]byte(file), &payload)
	assert.NoError(t, err)

	resp, err := underTest.CreateUpdate(payload)
	assert.NoError(t, err)
	assert.Equal(t, resp, &grafana_objects.CreateUpdateResults{
		Id:      1,
		Uid:     "test1",
		Status:  "ok",
		Version: 1,
		Url:     "testUrl",
		Slug:    "testSlug",
	},
	)
}

func TestGrafanaObjects_CreateUpdateNOK(t *testing.T) {
	underTest, err, teardown := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/grafana/api/dashboards/db", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)
		var payload grafana_objects.CreateUpdatePayload
		err = json.Unmarshal(jsonBytes, &payload)
		assert.NoError(t, err)
		assert.NotNil(t, payload)

		fileGet, err := ioutil.ReadFile("testdata/fixtures/createupdate_nok_resp.json")
		assert.NoError(t, err)

		var resp grafana_objects.CreateUpdateResults
		err = json.Unmarshal([]byte(fileGet), &resp)
		assert.NoError(t, err)

		bytes, err := json.Marshal(resp)
		assert.NoError(t, err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(bytes)
	})

	file, err := ioutil.ReadFile("testdata/fixtures/createupdate_nok.json")
	assert.NoError(t, err)

	var payload grafana_objects.CreateUpdatePayload
	err = json.Unmarshal([]byte(file), &payload)
	assert.NoError(t, err)

	_, err = underTest.CreateUpdate(payload)
	assert.Error(t, err)
}
