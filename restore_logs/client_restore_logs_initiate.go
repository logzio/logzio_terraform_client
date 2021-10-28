package restore_logs

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	initiateRestoreLogsServiceUrl            = restoreLogsServiceEndpoint
	initiateRestoreLogsServiceMethod  string = http.MethodPost
	initiateRestoreLogsServiceSuccess        = http.StatusOK
)

// InitiateRestoreOperation initiates a new operation to restore data from a specific time frame
func (c *RestoreClient) InitiateRestoreOperation(initiateRestore InitiateRestore) (*RestoreOperation, error) {
	err := validateInitiateRestoreOperation(initiateRestore)
	if err != nil {
		return nil, err
	}

	initiateRestoreJson, err := json.Marshal(initiateRestore)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   initiateRestoreLogsServiceMethod,
		Url:          fmt.Sprintf(initiateRestoreLogsServiceUrl, c.BaseUrl),
		Body:         initiateRestoreJson,
		SuccessCodes: []int{initiateRestoreLogsServiceSuccess},
		NotFoundCode: http.StatusNotFound,
		ResourceId:   nil,
		ApiAction:    initiateRestoreOperation,
		ResourceName: restoreResourceName,
	})

	if err != nil {
		return nil, err
	}

	var retVal RestoreOperation
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}

func validateInitiateRestoreOperation(initiateRestore InitiateRestore) error {
	if len(initiateRestore.AccountName) == 0 {
		return fmt.Errorf("AccountName must be set")
	}

	if initiateRestore.StartTime == 0 {
		return fmt.Errorf("StartTime must be set")
	}

	if initiateRestore.EndTime == 0 {
		return fmt.Errorf("EndTime must be set")
	}

	return nil
}
