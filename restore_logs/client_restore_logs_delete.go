package restore_logs

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
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
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteRestoreLogsServiceMethod,
		Url:          fmt.Sprintf(deleteRestoreLogsLogsServiceUrl, c.BaseUrl, restoreId),
		Body:         nil,
		SuccessCodes: []int{deleteRestoreLogsServiceSuccess},
		NotFoundCode: deleteRestoreLogsMethodNotFound,
		ResourceId:   restoreId,
		ApiAction:    deleteRestoreOperation,
		ResourceName: restoreResourceName,
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
