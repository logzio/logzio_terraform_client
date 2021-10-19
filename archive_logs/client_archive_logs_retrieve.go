package archive_logs

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	retrieveArchiveLogsServiceUrl            = archiveLogsServiceEndpoint + "/%d"
	retrieveArchiveLogsServiceMethod  string = http.MethodGet
	retrieveArchiveLogsMethodSuccess  int    = http.StatusOK
	retrieveArchiveLogsMethodNotFound int    = http.StatusNotFound
)

// RetrieveArchiveLogsSetting gets an archive id and returns its settings or an error
func (c *ArchiveLogsClient) RetrieveArchiveLogsSetting(archiveId int32) (*ArchiveLogs, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   retrieveArchiveLogsServiceMethod,
		Url:          fmt.Sprintf(retrieveArchiveLogsServiceUrl, c.BaseUrl, archiveId),
		Body:         nil,
		SuccessCodes: []int{retrieveArchiveLogsMethodSuccess},
		NotFoundCode: retrieveArchiveLogsMethodNotFound,
		ResourceId:   archiveId,
		ApiAction:    retrieveArchiveSettings,
	})

	if err != nil {
		return nil, err
	}

	var archive ArchiveLogs
	err = json.Unmarshal(res, &archive)
	if err != nil {
		return nil, err
	}

	return &archive, nil
}
