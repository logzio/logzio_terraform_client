package archive_logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	updateArchiveLogsServiceUrl            = archiveLogsServiceEndpoint + "/%d"
	updateArchiveLogsServiceMethod         = http.MethodPut
	updateArchiveLogsServiceSuccess        = http.StatusOK
	updateArchiveLogsServiceSuccessCreated = http.StatusCreated
	updateArchiveLogsServiceNotFound       = http.StatusNotFound
)

// UpdateArchiveLogs updates an existing logs archive
func (c *ArchiveLogsClient) UpdateArchiveLogs(archiveId int32, updateSubAccount CreateOrUpdateArchiving) (*ArchiveLogs, error) {
	err := validateCreateOrUpdateArchiveRequest(updateSubAccount)
	if err != nil {
		return nil, err
	}

	req, err := c.buildUpdateApiRequest(c.ApiToken, archiveId, updateSubAccount)
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
	if !logzio_client.CheckValidStatus(resp, []int{updateArchiveLogsServiceSuccess, updateArchiveLogsServiceSuccessCreated}) {
		if resp.StatusCode == updateArchiveLogsServiceNotFound {
			return nil, fmt.Errorf("API call %s failed with missing archive %d, data: %s", updateArchiveSettings, archiveId, jsonBytes)
		}

		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", updateArchiveSettings, resp.StatusCode, jsonBytes)
	}

	var retVal ArchiveLogs
	err = json.Unmarshal(jsonBytes, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}

func (c *ArchiveLogsClient) buildUpdateApiRequest(apiToken string, subAccountId int32, updateSubAccount CreateOrUpdateArchiving) (*http.Request, error) {
	jsonBytes, err := json.Marshal(updateSubAccount)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	req, err := http.NewRequest(updateArchiveLogsServiceMethod, fmt.Sprintf(updateArchiveLogsServiceUrl, baseUrl, subAccountId), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}
