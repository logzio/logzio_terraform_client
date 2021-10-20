package restore_logs

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getRestoreLogsServiceUrl                = restoreLogsServiceEndpoint + "/%d"
	getRestoreLogsServiceMethod      string = http.MethodGet
	getRestoreLogsMethodSuccess      int    = http.StatusOK
	getRestoreLogsLogsMethodNotFound int    = http.StatusNotFound
)

// GetRestoreOperation gets a restore operation id and returns its status or an error
func (c *RestoreClient) GetRestoreOperation(restoreId int32) (*RestoreOperation, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getRestoreLogsServiceMethod,
		Url:          fmt.Sprintf(getRestoreLogsServiceUrl, c.BaseUrl, restoreId),
		Body:         nil,
		SuccessCodes: []int{getRestoreLogsMethodSuccess},
		NotFoundCode: getRestoreLogsLogsMethodNotFound,
		ResourceId:   restoreId,
		ApiAction:    getRestoreOperation,
	})

	if err != nil {
		return nil, err
	}

	var restoreOperation RestoreOperation
	err = json.Unmarshal(res, &restoreOperation)
	if err != nil {
		return nil, err
	}

	return &restoreOperation, nil
}
