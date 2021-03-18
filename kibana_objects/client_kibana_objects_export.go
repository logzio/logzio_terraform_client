package kibana_objects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
)

type exportType string

// Enums for exportType
const (
	ExportTypeSearch        exportType = "search"
	ExportTypeVisualization exportType = "visualization"
	ExportTypeDashboard     exportType = "dashboard"
)

func (s exportType) String() string {
	return string(s)
}

type ExportResults struct {
	KibanaVersion string                   `json:"kibanaVersion"`
	Hits          []map[string]interface{} `json:"hits"`
}

type ExportPayload struct {
	Type exportType
}

// Export allows export of the Kibana objects configuration.
// https://docs.logz.io/api/#operation/exportSavedObjects
func (c *KibanaObjectsClient) Export(t exportType) (*ExportResults, error) {
	payload, err := json.Marshal(&ExportPayload{t})
	if err != nil {
		return nil, fmt.Errorf("could not marshal payload for Kibana export request: %w", err)
	}

	url := fmt.Sprintf(kibanaObjectsExportServiceEndpoint, c.BaseUrl)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("could not create HTTP request for Kibana export: %w", err)
	}
	logzio_client.AddHttpHeaders(c.ApiToken, req)

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send HTTP request to Logz.io: %w", err)
	}
	defer resp.Body.Close()

	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	if !logzio_client.CheckValidStatus(resp, []int{http.StatusOK}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}

	var results ExportResults
	err = json.Unmarshal(jsonBytes, &results)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body into KibanaObjects: %w", err)
	}

	return &results, nil
}
