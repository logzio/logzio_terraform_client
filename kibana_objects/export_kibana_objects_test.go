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

func TestKibanaObjects_ExportKibanaObject(t *testing.T) {
	underTest, teardown, err := setupKibanaObjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(exportPath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target kibana_objects.KibanaObjectExportRequest
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.Contains(t,
				[]kibana_objects.ExportType{
					kibana_objects.ExportTypeSearch,
					kibana_objects.ExportTypeDashboard,
					kibana_objects.ExportTypeVisualization},
				target.Type)

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("export_response.json"))
		})

		exportReq := getExportRequest()

		exportRes, err := underTest.ExportKibanaObject(exportReq)
		assert.NoError(t, err)
		assert.Equal(t, "4.0.0-beta3", exportRes.KibanaVersion)
		assert.Equal(t, 2, len(exportRes.Hits))
	}
}

func TestKibanaObjects_ExportKibanaObjectApiFail(t *testing.T) {
	underTest, teardown, err := setupKibanaObjectsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(exportPath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target kibana_objects.KibanaObjectExportRequest
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.Contains(t,
				[]kibana_objects.ExportType{
					kibana_objects.ExportTypeSearch,
					kibana_objects.ExportTypeDashboard,
					kibana_objects.ExportTypeVisualization},
				target.Type)

			w.WriteHeader(http.StatusInternalServerError)
		})

		exportReq := getExportRequest()

		exportRes, err := underTest.ExportKibanaObject(exportReq)
		assert.Error(t, err)
		assert.Nil(t, exportRes)
	}
}
