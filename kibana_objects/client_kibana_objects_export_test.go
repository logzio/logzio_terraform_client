package kibana_objects_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/kibana_objects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func exportMockHandler(t *testing.T) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintln(w, "this endpoint only supports the POST method")
			return
		}

		jsonBytes, _ := ioutil.ReadAll(r.Body)
		payload := struct {
			T string `json:"type"`
		}{}
		err := json.Unmarshal(jsonBytes, &payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "could not unmarshal request's payload")
			return
		}

		switch payload.T {
		case kibana_objects.ExportTypeSearch.String():
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("export_search.json"))
		case kibana_objects.ExportTypeDashboard.String():
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("export_dashboard.json"))
		case kibana_objects.ExportTypeVisualization.String():
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("export_visualization.json"))
		default:
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid value for type")
		}
	}
}

func TestKibanaObjects_ExportSearch(t *testing.T) {
	underTest, err, teardown := setupKibanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/kibana/export", exportMockHandler(t))

	objects, err := underTest.Export(kibana_objects.ExportTypeVisualization)
	assert.NoError(t, err)
	require.NotNil(t, objects)
	assert.Equal(t, "visualization_1", objects.Hits[0]["name"])
	assert.Equal(t, "visualization_2", objects.Hits[1]["name"])
}

func TestKibanaObjects_ExportDashboard(t *testing.T) {
	underTest, err, teardown := setupKibanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/kibana/export", exportMockHandler(t))

	objects, err := underTest.Export(kibana_objects.ExportTypeDashboard)
	assert.NoError(t, err)
	require.NotNil(t, objects)
	assert.Equal(t, 2, len(objects.Hits))
	assert.Equal(t, "dashboard_1", objects.Hits[0]["name"])
	assert.Equal(t, "dashboard_2", objects.Hits[1]["name"])
}

func TestKibanaObjects_ExportVisualization(t *testing.T) {
	underTest, err, teardown := setupKibanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/kibana/export", exportMockHandler(t))

	objects, err := underTest.Export(kibana_objects.ExportTypeVisualization)
	assert.NoError(t, err)
	require.NotNil(t, objects)
	assert.Equal(t, "visualization_1", objects.Hits[0]["name"])
	assert.Equal(t, "visualization_2", objects.Hits[1]["name"])
}

func TestKibanaObjects_ExportWithBadResponseStatus(t *testing.T) {
	underTest, err, teardown := setupKibanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/kibana/export", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "bad response")
	})

	_, err = underTest.Export(kibana_objects.ExportTypeSearch)
	assert.EqualError(t, err, "500 bad response")
}

func TestKibanaObjects_BadResponse(t *testing.T) {
	underTest, err, teardown := setupKibanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/kibana/export", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "bad response")
	})

	_, err = underTest.Export(kibana_objects.ExportTypeSearch)
	assert.EqualError(t, err, "could not unmarshal response body into KibanaObjects: invalid character 'b' looking for beginning of value")
}
