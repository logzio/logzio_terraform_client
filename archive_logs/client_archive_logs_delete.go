package archive_logs

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteArchiveLogsServiceUrl            = archiveLogsServiceEndpoint + "/%d"
	deleteArchiveLogsServiceMethod         = http.MethodDelete
	deleteArchiveLogsServiceSuccessCreated = http.StatusCreated
	deleteArchiveLogsServiceSuccess        = http.StatusOK
	deleteArchiveLogsStatusNotFound        = http.StatusNotFound
)

func (c *ArchiveLogsClient) DeleteArchiveLogs(archiveId int32) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteArchiveLogsServiceMethod,
		Url:          fmt.Sprintf(deleteArchiveLogsServiceUrl, c.BaseUrl, archiveId),
		Body:         nil,
		SuccessCodes: []int{deleteArchiveLogsServiceSuccess, deleteArchiveLogsServiceSuccessCreated},
		NotFoundCode: deleteArchiveLogsStatusNotFound,
		ResourceId:   archiveId,
		ApiAction:    deleteArchiveSettings,
		ResourceName: archiveResourceName,
	})

	return err
}
