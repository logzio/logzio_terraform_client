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
	listArchiveLogsServiceUrl     string = archiveLogsServiceEndpoint
	listArchiveLogsServiceMethod  string = http.MethodGet
	listArchiveLogsServiceSuccess int    = http.StatusOK
)

// ListArchiveLog Returns all the archives for logs in an array associated with the account identified by the supplied API token, returns an error if
// any problem occurs during the API call
func (c *ArchiveLogsClient) ListArchiveLog() ([]ArchiveLogs, error) {
	req, err := c.buildListApiRequest(c.ApiToken)
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

	if !logzio_client.CheckValidStatus(resp, []int{listArchiveLogsServiceSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", listArchivesSettings, resp.StatusCode, jsonBytes)
	}

	var archives []ArchiveLogs
	err = json.Unmarshal(jsonBytes, &archives)
	if err != nil {
		return nil, err
	}

	return archives, nil
}

func (c *ArchiveLogsClient) buildListApiRequest(apiToken string) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(listArchiveLogsServiceMethod, fmt.Sprintf(listArchiveLogsServiceUrl, baseUrl), nil)
	if err != nil {
		return nil, err
	}
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}
