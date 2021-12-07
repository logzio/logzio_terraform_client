package grafana_objects_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/grafana_objects"
	"github.com/stretchr/testify/assert"
)

func TestGrafanaObjects_DeleteOK(t *testing.T) {
	underTest, err, teardown := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/grafana/api/dashboards/uid/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fileDelete, err := ioutil.ReadFile("testdata/fixtures/delete.json")
		assert.NoError(t, err)

		resp := grafana_objects.DeleteResults{}
		err = json.Unmarshal([]byte(fileDelete), &resp)
		assert.NoError(t, err)

		bytes, err := json.Marshal(resp)
		assert.NoError(t, err)
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	},
	)

	result, err := underTest.Delete("getOK")
	assert.NoError(t, err)
	assert.Equal(t, result, &grafana_objects.DeleteResults{
		Title:   "testDeleteOK",
		Message: "deleteOK",
		Id:      1,
	},
	)
}

func TestGrafanaObjects_DeleteNotFound(t *testing.T) {
	underTest, err, teardown := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/grafana/api/dashboards/uid/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		w.Header().Set("Content-Type", "application/json")
		resp := grafana_objects.DeleteResults{
			Message: "Dashboard Not Found",
		}
		bytes, err := json.Marshal(resp)
		assert.NoError(t, err)

		w.WriteHeader(http.StatusNotFound)
		w.Write(bytes)
	},
	)

	_, err = underTest.Delete("getNOK")
	assert.Error(t, err)
}

func TestGrafanaObjects_DeleteServerError(t *testing.T) {
	underTest, err, teardown := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/grafana/api/dashboards/uid/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)

		w.WriteHeader(http.StatusInternalServerError)
	},
	)

	_, err = underTest.Delete("getNOK")
	assert.Error(t, err)
}
