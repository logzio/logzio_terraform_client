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

type CreateUpdatePayload struct {
	Dashboard interface{} `json:"dashboard"`
	FolderId  int         `json:"folderId"`
	FolderUid int         `json:"folderUid"`
	Message   string      `json:"message"`
	Overwrite bool        `json:"overwrite"`
}

type CreateUpdateResults struct {
	Id      int    `json:"id"`
	Uid     string `json:"uid"`
	Status  string `json:"status"`
	Version int    `json:"Version"`
	Url     string `json:"url"`
	Slug    string `json:"slug"`
}

// Get allows getting a Grafana objects configuration.
// https://docs.logz.io/api/#operation/getDashboarById
func (c *GrafanaObjectsClient) CreateUpdate(payload CreateUpdatePayload) (*CreateUpdateResults, error) {

	jsonBytes, err := json.Marshal(&payload)
	if err != nil {
		return nil, fmt.Errorf("could not marshal payload for Kibana import request: %w", err)
	}

	url := fmt.Sprintf(grafanaObjectsDashboardsCreateUpdate, c.BaseUrl)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBytes))
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
		return nil, fmt.Errorf("could not unmarshal response body into GetResults: %w", err)
	}

	return &results, nil
}
