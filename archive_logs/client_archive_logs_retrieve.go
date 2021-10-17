package archive_logs

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	retrieveArchiveLogsServiceUrl            = archiveLogsServiceEndpoint + "/%d"
	retrieveArchiveLogsServiceMethod  string = http.MethodGet
	retrieveArchiveLogsMethodSuccess  int    = http.StatusOK
	retrieveArchiveLogsMethodNotFound int    = http.StatusNotFound
)

// RetrieveArchiveLogsSetting gets an archive id and returns its settings or an error
func (c *ArchiveLogsClient) RetrieveArchiveLogsSetting(archiveId int32) (*ArchiveLogs, error) {
	req, err := c.buildRetrieveApiRequest(c.ApiToken, archiveId)
	if err != nil {
		return nil, err
	}
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{retrieveArchiveLogsMethodSuccess}) {
		if resp.StatusCode == retrieveArchiveLogsMethodNotFound {
			return nil, fmt.Errorf("API call %s failed with missing archive %d, data: %s", retrieveArchiveSettings, archiveId, jsonBytes)
		}

		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", retrieveArchiveSettings, resp.StatusCode, jsonBytes)
	}

	var archive ArchiveLogs
	err = json.Unmarshal(jsonBytes, &archive)
	if err != nil {
		return nil, err
	}

	return &archive, nil
}

func (c *ArchiveLogsClient) buildRetrieveApiRequest(apiToken string, archiveId int32) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(retrieveArchiveLogsServiceMethod, fmt.Sprintf(retrieveArchiveLogsServiceUrl, baseUrl, archiveId), nil)
	if err != nil {
		return nil, err
	}
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}
