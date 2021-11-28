package grafana_objects

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
)

type DeleteResults struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Id      int    `json:"message"`
}

// Delete allows deleting Grafana objects configuration.
// https://docs.logz.io/api/#operation/deleteDashboarById
func (c *GrafanaObjectsClient) Delete(objectUid string) (*DeleteResults, error) {
	url := fmt.Sprintf(grafanaObjectsDashboardsByUID, c.BaseUrl, objectUid)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create HTTP request for Grafana import: %w", err)
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

	var results DeleteResults
	err = json.Unmarshal(jsonBytes, &results)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body into DeleteResults: %w", err)
	}

	return &results, nil
}
