package archive_logs

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	listArchiveLogsServiceUrl     string = archiveLogsServiceEndpoint
	listArchiveLogsServiceMethod  string = http.MethodGet
	listArchiveLogsServiceSuccess int    = http.StatusOK
)

// ListArchiveLog Returns all the archives for logs in an array associated with the account identified by the supplied API token, returns an error if
// any problem occurs during the API call
func (c *ArchiveLogsClient) ListArchiveLog() ([]ArchiveLogs, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listArchiveLogsServiceMethod,
		Url:          fmt.Sprintf(listArchiveLogsServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{listArchiveLogsServiceSuccess},
		NotFoundCode: http.StatusNotFound,
		ResourceId:   nil,
		ApiAction:    listArchivesSettings,
		ResourceName: archiveResourceName,
	})

	if err != nil {
		return nil, err
	}

	var archives []ArchiveLogs
	err = json.Unmarshal(res, &archives)
	if err != nil {
		return nil, err
	}

	return archives, nil
}
