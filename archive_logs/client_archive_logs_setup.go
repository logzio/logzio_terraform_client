package archive_logs

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	setupArchiveLogsServiceUrl           = archiveLogsServiceEndpoint
	setupArchiveLogsServiceMethod string = http.MethodPost
	setupArchiveLogsStatusSuccess int    = http.StatusOK
	setupArchiveLogsStatusCreated int    = http.StatusCreated
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

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   setupArchiveLogsServiceMethod,
		Url:          fmt.Sprintf(setupArchiveLogsServiceUrl, c.BaseUrl),
		Body:         archiveLogsJson,
		SuccessCodes: []int{setupArchiveLogsStatusSuccess, setupArchiveLogsStatusCreated},
		NotFoundCode: http.StatusNotFound,
		ResourceId:   nil,
		ApiAction:    setupArchiveApi,
	})

	if err != nil {
		return nil, err
	}

	var retVal ArchiveLogs
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}
