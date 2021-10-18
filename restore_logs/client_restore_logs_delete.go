package restore_logs

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	deleteRestoreLogsLogsServiceUrl = restoreLogsServiceEndpoint + "/%d"
	deleteRestoreLogsServiceMethod  = http.MethodDelete
	deleteRestoreLogsServiceSuccess = http.StatusOK
	deleteRestoreLogsMethodNotFound = http.StatusNotFound
)

// DeleteRestoreOperation aborts a restore process before its completion
func (c *RestoreClient) DeleteRestoreOperation(restoreId int32) (*RestoreOperation, error) {
	req, err := c.buildDeleteApiRequest(c.ApiToken, restoreId)
	if err != nil {
		return nil, err
	}
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{deleteRestoreLogsServiceSuccess}) {
		if resp.StatusCode == deleteRestoreLogsMethodNotFound {
			return nil, fmt.Errorf("API call %s failed with missing restore operation %d, data: %s", deleteRestoreOperation, restoreId, jsonBytes)
		}

		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", deleteRestoreOperation, resp.StatusCode, jsonBytes)
	}

	var restoreOperation RestoreOperation
	err = json.Unmarshal(jsonBytes, &restoreOperation)
	if err != nil {
		return nil, err
	}

	return &restoreOperation, nil
}

func (c *RestoreClient) buildDeleteApiRequest(apiToken string, restoreId int32) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(deleteRestoreLogsServiceMethod, fmt.Sprintf(deleteRestoreLogsLogsServiceUrl, baseUrl, restoreId), nil)
	if err != nil {
		return nil, err
	}
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}
