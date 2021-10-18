package restore_logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	initiateRestoreLogsServiceUrl           = restoreLogsServiceEndpoint
	initiateRestoreLogsServiceMethod string = http.MethodPost
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

	req, err := c.buildInitiateApiRequest(c.ApiToken, initiateRestoreJson)
	if err != nil {
		return nil, err
	}

	jsonResponse, err := logzio_client.CreateHttpRequestBytesResponse(req)
	if err != nil {
		return nil, err
	}

	var retVal RestoreOperation
	err = json.Unmarshal(jsonResponse, &retVal)
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

func (c *RestoreClient) buildInitiateApiRequest(apiToken string, jsonBytes []byte) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(initiateRestoreLogsServiceMethod, fmt.Sprintf(initiateRestoreLogsServiceUrl, baseUrl), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	logzio_client.AddHttpHeaders(apiToken, req)
	return req, err
}
