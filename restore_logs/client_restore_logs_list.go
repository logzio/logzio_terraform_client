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
	listRestoreLogsServiceUrl     string = restoreLogsServiceEndpoint
	listRestoreLogsServiceMethod  string = http.MethodGet
	listRestoreLogsServiceSuccess int    = http.StatusOK
)

// ListRestoreOperations Returns all the restore operations for logs in an array associated with the account identified by the supplied API token, returns an error if
// any problem occurs during the API call
func (c *RestoreClient) ListRestoreOperations() ([]RestoreOperation, error) {
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
	if !logzio_client.CheckValidStatus(resp, []int{listRestoreLogsServiceSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", listRestoreOperations, resp.StatusCode, jsonBytes)
	}

	var restoreOperations []RestoreOperation
	err = json.Unmarshal(jsonBytes, &restoreOperations)
	if err != nil {
		return nil, err
	}

	return restoreOperations, nil
}

func (c *RestoreClient) buildListApiRequest(apiToken string) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(listRestoreLogsServiceMethod, fmt.Sprintf(listRestoreLogsServiceUrl, baseUrl), nil)
	if err != nil {
		return nil, err
	}
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}
