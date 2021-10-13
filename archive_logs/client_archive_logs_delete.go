package archive_logs

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	deleteArchiveLogsServiceUrl     = archiveLogsServiceEndpoint + "/%d"
	deleteArchiveLogsServiceMethod  = http.MethodDelete
	deleteArchiveLogsServiceSuccessCreated = http.StatusCreated
	deleteArchiveLogsServiceSuccess = http.StatusOK
	deleteArchiveLogsMethodNotFound = http.StatusNotFound
)

func (c *ArchiveLogsClient) DeleteArchiveLogs(archiveId int32) error {
	req, err := c.buildDeleteApiRequest(c.ApiToken, archiveId)
	if err != nil {
		return err
	}
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{deleteArchiveLogsServiceSuccess, deleteArchiveLogsServiceSuccessCreated}) {
		if resp.StatusCode == deleteArchiveLogsMethodNotFound {
			return fmt.Errorf("API call %s failed with missing archive %d, data: %s", deleteArchiveSettings, archiveId, jsonBytes)
		}

		return fmt.Errorf("API call %s failed with status code %d, data: %s", deleteArchiveSettings, resp.StatusCode, jsonBytes)
	}

	return nil
}

func (c *ArchiveLogsClient) buildDeleteApiRequest(apiToken string, archiveId int32) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(deleteArchiveLogsServiceMethod, fmt.Sprintf(deleteArchiveLogsServiceUrl, baseUrl, archiveId), nil)
	if err != nil {
		return nil, err
	}
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}
