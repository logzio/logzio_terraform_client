package archive_logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	setupArchiveLogsServiceUrl           = archiveLogsServiceEndpoint
	setupArchiveLogsServiceMethod string = http.MethodPost
)

// SetupArchive receives archive settings, and sends it as api request to logz.io
// returns an error if occurred
func (c *ArchiveLogsClient) SetupArchive(createArchive CreateOrUpdateArchiving) (*ArchiveLogs, error) {
	err := validateCreateOrUpdateArchiveRequest(createArchive)
	if err != nil {
		return nil, err
	}

	archiveLogsJson, err := json.Marshal(createArchive)
	if err != nil {
		return nil, err
	}

	req, err := c.buildSetupApiRequest(c.ApiToken, archiveLogsJson)
	if err != nil {
		return nil, err
	}

	jsonResponse, err := logzio_client.CreateHttpRequestBytesResponse(req)
	if err != nil {
		return nil, err
	}

	var retVal ArchiveLogs
	err = json.Unmarshal(jsonResponse, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}

func (c *ArchiveLogsClient) buildSetupApiRequest(apiToken string, jsonBytes []byte) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(setupArchiveLogsServiceMethod, fmt.Sprintf(setupArchiveLogsServiceUrl, baseUrl), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	logzio_client.AddHttpHeaders(apiToken, req)
	return req, err
}
