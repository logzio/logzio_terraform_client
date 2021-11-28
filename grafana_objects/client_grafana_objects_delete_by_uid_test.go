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

func deleteMockHandler(t *testing.T) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintln(w, "this endpoint only supports the delete method")
			return
		}
		w.Header().Set("Content-Type", "application/json")

		if path.Base(r.URL.Path) == "test1" {
			fileDelete, _ := ioutil.ReadFile("testdata/delete.json")

			resp := grafana_objects.DeleteResults{}
			_ = json.Unmarshal([]byte(fileDelete), &resp)

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

func TestGrafanaObjects_Delete(t *testing.T) {
	underTest, err, teardown := setupGrafanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/grafana/api/dashboards/uid/", deleteMockHandler(t))

	_, err = underTest.Delete("test1")
	assert.NoError(t, err)
	_, err = underTest.Delete("test2")
	assert.Error(t, err)
}
