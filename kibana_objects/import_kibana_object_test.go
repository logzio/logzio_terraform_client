package kibana_objects_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/kibana_objects"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestKibanaObjects_ImportKibanaObject(t *testing.T) {
	underTest, teardown, err := setupKibanaObjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(importPath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target kibana_objects.KibanaObjectImportRequest
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.KibanaVersion)
			assert.NotEmpty(t, target.Hits)

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("import_response.json"))
		})

		importReq, err := getImportRequest()
		if assert.NoError(t, err) {
			importRes, err := underTest.ImportKibanaObject(importReq)
			assert.NoError(t, err)
			assert.NotNil(t, importRes)
			assert.Equal(t, 1, len(importRes.Created))
		}
	}
}

func TestKibanaObjects_ImportKibanaObjectApiFail(t *testing.T) {
	underTest, teardown, err := setupKibanaObjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(importPath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target kibana_objects.KibanaObjectImportRequest
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.KibanaVersion)
			assert.NotEmpty(t, target.Hits)

			w.WriteHeader(http.StatusInternalServerError)
		})

		importReq := getExportRequest()

		importRes, err := underTest.ExportKibanaObject(importReq)
		assert.Error(t, err)
		assert.Nil(t, importRes)
	}
}
