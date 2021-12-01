package grafana_objects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	grafanaObjectsDashboardsCreateUpdate = grafanaObjectServiceEndpoint + "/db"
)

// CreateUpdate allows the creation or update of a Grafana dashboard
// https://docs.logz.io/api/#operation/createDashboard
func (c *GrafanaObjectsClient) CreateUpdate(payload CreateUpdatePayload) (*CreateUpdateResults, error) {

	jsonBytes, err := json.Marshal(&payload)
	if err != nil {
		return nil, fmt.Errorf("could not marshal payload for Kibana import request: %w", err)
	}

	url := fmt.Sprintf(grafanaObjectsDashboardsCreateUpdate, c.BaseUrl)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, fmt.Errorf("could not create HTTP request for Grafana create/update: %w", err)
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

	var results CreateUpdateResults
	err = json.Unmarshal(jsonBytes, &results)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body into CreateUpdateResults: %w", err)
	}

	return &results, nil
}
