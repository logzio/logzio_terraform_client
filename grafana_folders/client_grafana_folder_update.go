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

func (c *GrafanaFolderClient) UpdateGrafanaFolder(update CreateUpdateFolder) error {
	err := validateUpdateGrafanaFolder(update)
	if err != nil {
		return err
	}

	// update to a new version
	update.Overwrite = true
	updateGrafanaFolderJson, err := json.Marshal(update)
	if err != nil {
		return err
	}

	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateGrafanaFolderServiceMethod,
		Url:          fmt.Sprintf(updateGrafanaFolderServiceUrl, c.BaseUrl, update.Uid),
		Body:         updateGrafanaFolderJson,
		SuccessCodes: []int{updateGrafanaFolderServiceSuccess},
		NotFoundCode: updateGrafanaFolderServiceNotFound,
		ResourceId:   update.Uid,
		ApiAction:    operationUpdateGrafanaFolder,
		ResourceName: grafanaFolderResourceName,
	})

	return err
}

func validateUpdateGrafanaFolder(payload CreateUpdateFolder) error {
	if len(payload.Title) == 0 {
		return fmt.Errorf("Field title must be set!")
	}

	if len(payload.Uid) == 0 {
		fmt.Errorf("Field uid must be set!")
	}

	return nil
}
