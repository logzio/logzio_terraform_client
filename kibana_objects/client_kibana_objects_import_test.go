package kibana_objects_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/kibana_objects"
	"github.com/stretchr/testify/assert"
)

func importMockHandler(t *testing.T) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintln(w, "this endpoint only supports the POST method")
			return
		}

		jsonBytes, _ := ioutil.ReadAll(r.Body)
		payload := kibana_objects.ImportPayload{}
		err := json.Unmarshal(jsonBytes, &payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "could not unmarshal request's payload")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		if payload.Hits[0]["name"] != "search_1" ||
			payload.Hits[1]["name"] != "search_2" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "was expecting data from the test fixtures")
			return
		}

		if payload.Override {
			fmt.Fprint(w, fixture("import_with_override_response.json"))
			return
		}

		fmt.Fprint(w, fixture("import_without_override_response.json"))
	}
}

func TestKibanaObjects_ImportWithoutOverride(t *testing.T) {
	underTest, err, teardown := setupKibanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/kibana/import", importMockHandler(t))

	payload, err := getImportPayloadFromFixture("import_without_override.json")
	assert.NoError(t, err)

	results, err := underTest.Import(payload)
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, &kibana_objects.ImportResults{
		Created: []string{"search_1", "search_2"},
		Updated: []string{},
		Ignored: []string{},
		Failed:  []string{},
	}, results)
}

func TestKibanaObjects_ImportWithOverride(t *testing.T) {
	underTest, err, teardown := setupKibanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/kibana/import", importMockHandler(t))

	payload, err := getImportPayloadFromFixture("import_with_override.json")
	assert.NoError(t, err)

	results, err := underTest.Import(payload)
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, &kibana_objects.ImportResults{
		Created: []string{"search_1"},
		Updated: []string{"search_2"},
		Ignored: []string{},
		Failed:  []string{},
	}, results)
}

func TestKibanaObjects_ImportWithBadResponseStatus(t *testing.T) {
	underTest, err, teardown := setupKibanaObjectsTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/kibana/import", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "bad response")
	})

	payload, err := getImportPayloadFromFixture("import_without_override.json")
	assert.NoError(t, err)

	_, err = underTest.Import(payload)
	assert.EqualError(t, err, "500 bad response")
}

func getImportPayloadFromFixture(fixtureFile string) (kibana_objects.ImportPayload, error) {
	data := fixture(fixtureFile)
	payload := kibana_objects.ImportPayload{}
	err := json.Unmarshal([]byte(data), &payload)
	if err != nil {
		return kibana_objects.ImportPayload{}, err
	}

	return payload, nil
}
