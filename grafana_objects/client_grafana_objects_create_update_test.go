package grafana_objects_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/grafana_objects"
	"github.com/stretchr/testify/assert"
)

func createUpdateMockHandler(t *testing.T) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintln(w, "this endpoint only supports the POST method")
			return
		}
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		payload := grafana_objects.CreateUpdatePayload{}
		err := json.Unmarshal(jsonBytes, &payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "could not unmarshal request's payload")
			return
		}
		w.Header().Set("Content-Type", "application/json")

		if payload.Dashboard["uid"] == "test1" {
			fileGet, err := ioutil.ReadFile("testdata/createupdate_ok_resp.json")
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Unable to open resp file")
				return
			}
			resp := grafana_objects.CreateUpdateResults{}
			err = json.Unmarshal([]byte(fileGet), &resp)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Unable to unmarshall resp file")
				return
			}

			bytes, _ := json.Marshal(resp)
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		}

		if payload.Dashboard["uid"] == "test2" {
			fileGet, err := ioutil.ReadFile("testdata/createupdate_nok_resp.json")
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Unable to open resp file")
				return
			}
			resp := make(map[string]string)
			err = json.Unmarshal([]byte(fileGet), &resp)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Unable to unmarshall resp file")
				return
			}

			bytes, _ := json.Marshal(resp)
			w.WriteHeader(http.StatusNotFound)
			w.Write(bytes)
		}
	}
}

func TestGrafanaObjects_CreateUpdateOK(t *testing.T) {
	underTest, err, teardown := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/grafana/api/dashboards/db", createUpdateMockHandler(t))

	file, _ := ioutil.ReadFile("testdata/createupdate_ok.json")
	payload := grafana_objects.CreateUpdatePayload{}
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

	mux.HandleFunc("/v1/grafana/api/dashboards/db", createUpdateMockHandler(t))

	file, _ := ioutil.ReadFile("testdata/createupdate_nok.json")
	payload := grafana_objects.CreateUpdatePayload{}
	err = json.Unmarshal([]byte(file), &payload)
	assert.NoError(t, err)

	_, err = underTest.CreateUpdate(payload)
	assert.Error(t, err)
}
