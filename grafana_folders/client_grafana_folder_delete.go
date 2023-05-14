package grafana_folders

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteGrafanaFolderServiceUrl     = grafanaFolderServiceEndpoint + "/%s"
	deleteGrafanaFolderServiceMethod  = http.MethodDelete
	deleteGrafanaFolderServiceSuccess = http.StatusOK
	deleteGrafanaFolderNotFound       = http.StatusNotFound
)

func (c *GrafanaFolderClient) DeleteGrafanaFolder(uid string) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteGrafanaFolderServiceMethod,
		Url:          fmt.Sprintf(deleteGrafanaFolderServiceUrl, c.BaseUrl, uid),
		Body:         nil,
		SuccessCodes: []int{deleteGrafanaFolderServiceSuccess},
		NotFoundCode: deleteGrafanaFolderNotFound,
		ResourceId:   uid,
		ApiAction:    operationDeleteGrafanaFolder,
		ResourceName: grafanaFolderResourceName,
	})

	return err
}
