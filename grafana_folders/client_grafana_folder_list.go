package grafana_folders

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	listGrafanaFolderServiceUrl     = grafanaFolderServiceEndpoint
	listGrafanaFolderServiceMethod  = http.MethodGet
	listGrafanaFolderServiceSuccess = http.StatusOK
	listGrafanaFolderStatusNotFound = http.StatusNotFound
)

func (c *GrafanaFolderClient) ListGrafanaFolders() ([]GrafanaFolder, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listGrafanaFolderServiceMethod,
		Url:          fmt.Sprintf(listGrafanaFolderServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{listGrafanaFolderServiceSuccess},
		NotFoundCode: listGrafanaFolderStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationListGrafanaFolders,
		ResourceName: grafanaFolderResourceName,
	})

	if err != nil {
		return nil, err
	}

	var folders []GrafanaFolder
	err = json.Unmarshal(res, &folders)
	if err != nil {
		return nil, err
	}

	return folders, nil
}
