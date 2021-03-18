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

type ImportPayload struct {
	ExportResults
	Override bool
}

type ImportResults struct {
	// Name of Kibana objects that were created
	Created []string `json:"created"`

	// Names of the Kibana objects that were overwritten. Objects are shown here only if override was true in the import call.
	Updated []string `json:"updated"`

	// Names of the Kibana objects that were not overwritten. Objects are shown here only if override was false in the import call.
	Ignored []string `json:"ignored"`

	// Names of the Kibana objects that could not be created, updated, or ignored.
	Failed []string `json:"failed"`
}

// Import allows import of the Kibana objects configuration.
// https://docs.logz.io/api/#operation/exportSavedObjects
func (c *KibanaObjectsClient) Import(payload ImportPayload) (*ImportResults, error) {
	jsonBytes, err := json.Marshal(&payload)
	if err != nil {
		return nil, fmt.Errorf("could not marshal payload for Kibana import request: %w", err)
	}

	url := fmt.Sprintf(kibanaObjectsImportServiceEndpoint, c.BaseUrl)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, fmt.Errorf("could not create HTTP request for Kibana import: %w", err)
	}
	logzio_client.AddHttpHeaders(c.ApiToken, req)

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send HTTP request to Logz.io: %w", err)
	}
	defer resp.Body.Close()

	jsonBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	if !logzio_client.CheckValidStatus(resp, []int{http.StatusOK}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}

	var results ImportResults
	err = json.Unmarshal(jsonBytes, &results)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body into ImportResults: %w", err)
	}

	return &results, nil
}
