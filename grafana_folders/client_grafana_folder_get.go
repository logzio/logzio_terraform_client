package grafana_folders

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getGrafanaFolderServiceUrl      = grafanaFolderServiceEndpoint + "/%s"
	getGrafanaFolderServiceMethod   = http.MethodGet
	getGrafanaFolderServiceSuccess  = http.StatusOK
	getGrafanaFolderServiceNotFound = http.StatusNotFound
)

func (c *GrafanaFolderClient) GetGrafanaFolder(uid string) (*GrafanaFolder, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getGrafanaFolderServiceMethod,
		Url:          fmt.Sprintf(getGrafanaFolderServiceUrl, c.BaseUrl, uid),
		Body:         nil,
		SuccessCodes: []int{getGrafanaFolderServiceSuccess},
		NotFoundCode: getGrafanaFolderServiceNotFound,
		ResourceId:   uid,
		ApiAction:    operationGetGrafanaFolder,
		ResourceName: grafanaFolderResourceName,
	})

	if err != nil {
		return nil, err
	}

	var grafanaFolder GrafanaFolder
	err = json.Unmarshal(res, &grafanaFolder)
	if err != nil {
		return nil, err
	}

	return &grafanaFolder, nil
}
