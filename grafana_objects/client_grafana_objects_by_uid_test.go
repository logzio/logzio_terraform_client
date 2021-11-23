package grafana_objects_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"testing"

	"github.com/logzio/logzio_terraform_client/grafana_objects"
	"github.com/stretchr/testify/assert"
)

func getMockHandler(t *testing.T) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintln(w, "this endpoint only supports the GET method")
			return
		}
		w.Header().Set("Content-Type", "application/json")

		if path.Base(r.URL.Path) == "test1" {
			fileGet, _ := ioutil.ReadFile("testdata/get.json")

			resp := grafana_objects.GetResults{}
			_ = json.Unmarshal([]byte(fileGet), &resp)

			bytes, _ := json.Marshal(resp)
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		}

		if path.Base(r.URL.Path) == "test2" {
			resp := make(map[string]string)
			resp["message"] = "Dashboard Not found"
			bytes, _ := json.Marshal(resp)
			w.WriteHeader(http.StatusNotFound)
			w.Write(bytes)
		}
	}
}

func TestGrafanaObjects_Get(t *testing.T) {
	underTest, err, teardown := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/grafana/api/dashboards/uid/", getMockHandler(t))

	_, err = underTest.Get("test1")
	assert.NoError(t, err)
	_, err = underTest.Get("test2")
	assert.Error(t, err)
}
