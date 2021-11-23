package grafana_objects

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
)

type GetResults struct {
	Dashboard map[string]interface{} `json:"dashboard"`
	Meta      DashboardMeta          `json:"meta"`
}

type DeleteResults struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Id      int    `json:"message"`
}

type DashboardObject struct {
	Id            int      `json:"id"`
	Uid           string   `json:"uid"`
	Title         string   `json:"title"`
	Tags          []string `json:"tags"`
	Timezone      string   `json:"timezone"`
	SchemaVersion int      `json:"schemaVersion"`
	Version       int      `json:"Version"`
}

type DashboardMeta struct {
	IsStarred bool   `json:"isStarred"`
	Url       string `json:"url"`
	FolderId  int    `json:"folderId"`
	FolderUid string `json:"folderUid"`
	Slug      string `json:"slug"`
}

// Get allows getting a Grafana objects configuration.
// https://docs.logz.io/api/#operation/getDashboarById
func (c *GrafanaObjectsClient) Get(objectUid string) (*GetResults, error) {
	url := fmt.Sprintf(grafanaObjectsDashboardsByUID, c.BaseUrl, objectUid)
	req, err := http.NewRequest(http.MethodGet, url, nil)
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

	var results GetResults
	err = json.Unmarshal(jsonBytes, &results)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body into GetResults: %w", err)
	}

	return &results, nil
}

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
