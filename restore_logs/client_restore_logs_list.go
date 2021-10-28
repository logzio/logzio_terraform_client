package restore_logs

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
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
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listRestoreLogsServiceMethod,
		Url:          fmt.Sprintf(listRestoreLogsServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{listRestoreLogsServiceSuccess},
		NotFoundCode: http.StatusNotFound,
		ResourceId:   nil,
		ApiAction:    listRestoreOperations,
		ResourceName: restoreResourceName,
	})

	if err != nil {
		return nil, err
	}

	var restoreOperations []RestoreOperation
	err = json.Unmarshal(res, &restoreOperations)
	if err != nil {
		return nil, err
	}

	return restoreOperations, nil
}
