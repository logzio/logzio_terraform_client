package grafana_folders

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	updateGrafanaFolderServiceUrl      = grafanaFolderServiceEndpoint + "/%s"
	updateGrafanaFolderServiceMethod   = http.MethodPut
	updateGrafanaFolderServiceSuccess  = http.StatusOK
	updateGrafanaFolderServiceNotFound = http.StatusNotFound
)

func (c *GrafanaFolderClient) UpdateGrafanaFolder(uid string, update CreateUpdateFolder) error {
	updateGrafanaFolderJson, err := json.Marshal(update)
	if err != nil {
		return err
	}

	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateGrafanaFolderServiceMethod,
		Url:          fmt.Sprintf(updateGrafanaFolderServiceUrl, c.BaseUrl, uid),
		Body:         updateGrafanaFolderJson,
		SuccessCodes: []int{updateGrafanaFolderServiceSuccess},
		NotFoundCode: updateGrafanaFolderServiceNotFound,
		ResourceId:   uid,
		ApiAction:    operationUpdateGrafanaFolder,
		ResourceName: grafanaFolderResourceName,
	})

	return err
}
