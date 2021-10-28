package archive_logs

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	updateArchiveLogsServiceUrl            = archiveLogsServiceEndpoint + "/%d"
	updateArchiveLogsServiceMethod         = http.MethodPut
	updateArchiveLogsServiceSuccess        = http.StatusOK
	updateArchiveLogsServiceSuccessCreated = http.StatusCreated
	updateArchiveLogsServiceNotFound       = http.StatusNotFound
)

// UpdateArchiveLogs updates an existing logs archive
func (c *ArchiveLogsClient) UpdateArchiveLogs(archiveId int32, updateArchive CreateOrUpdateArchiving) (*ArchiveLogs, error) {
	err := validateCreateOrUpdateArchiveRequest(updateArchive)
	if err != nil {
		return nil, err
	}

	updateArchiveBodyJson, err := json.Marshal(updateArchive)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateArchiveLogsServiceMethod,
		Url:          fmt.Sprintf(updateArchiveLogsServiceUrl, c.BaseUrl, archiveId),
		Body:         updateArchiveBodyJson,
		SuccessCodes: []int{updateArchiveLogsServiceSuccess, updateArchiveLogsServiceSuccessCreated},
		NotFoundCode: updateArchiveLogsServiceNotFound,
		ResourceId:   archiveId,
		ApiAction:    updateArchiveSettings,
		ResourceName: archiveResourceName,
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
