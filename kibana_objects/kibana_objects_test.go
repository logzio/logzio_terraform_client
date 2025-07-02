package kibana_objects_test

import (
	"encoding/json"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"time"

	"github.com/logzio/logzio_terraform_client/kibana_objects"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

const (
	exportPath = "/v1/kibana/export"
	importPath = "/v1/kibana/import"
)

func fixture(path string) string {
	b, err := os.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setupKibanaObjectsTest() (*kibana_objects.KibanaObjectsClient, func(), error) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, err := kibana_objects.New(apiToken, server.URL)

	return underTest, func() {
		server.Close()
	}, err
}

func setupKibanaObjectsIntegrationTest() (*kibana_objects.KibanaObjectsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := kibana_objects.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, nil
}

func getExportRequest() kibana_objects.KibanaObjectExportRequest {
	return kibana_objects.KibanaObjectExportRequest{Type: kibana_objects.ExportTypeSearch}
}

func getImportRequest() (kibana_objects.KibanaObjectImportRequest, error) {
	randomSuffix := getRandomId()
	source := `{
        "search": {
          "columns": [
            "message"
          ],
          "sort": [
            "@timestamp",
            "desc"
          ],
          "id": "tf-client-test-` + randomSuffix + `",
          "title": "tf-client-test-` + randomSuffix + `",
          "version": 1,
          "kibanaSavedObjectMeta": {
            "searchSourceJSON": "{\"highlight\":{\"pre_tags\":[\"@kibana-highlighted-field@\"],\"post_tags\":[\"@/kibana-highlighted-field@\"],\"fields\":{\"*\":{}},\"fragment_size\":2147483647},\"filter\":[],\"query\":{\"query\":\"type: tf-client-test\",\"language\":\"lucene\"},\"source\":{\"excludes\":[]},\"highlightAll\":true,\"version\":true,\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"
          }
        },
        "type": "search",
        "id": "tf-client-test-` + randomSuffix + `"
      }`
	var sourceObj map[string]interface{}
	err := json.Unmarshal([]byte(source), &sourceObj)

	hitsObj := map[string]interface{}{
		"_index":  "logzioCustomerIndex*",
		"_type":   "_doc",
		"_id":     "search:tf-client-test-" + randomSuffix,
		"_source": sourceObj,
	}

	return kibana_objects.KibanaObjectImportRequest{
		KibanaVersion: "7.2.1",
		Hits:          []map[string]interface{}{hitsObj},
	}, err
}

func getRandomId() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(rand.Intn(10000))
}
