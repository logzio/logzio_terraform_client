package grafana_folders

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	createGrafanaFolderServiceUrl    = grafanaFolderServiceEndpoint
	createGrafanaFolderServiceMethod = http.MethodPost
	createGrafanaFolderMethodCreated = http.StatusOK
	createS3FetcherStatusNotFound    = http.StatusNotFound
)

func (c *GrafanaFolderClient) CreateGrafanaFolder(payload CreateUpdateFolder) (*GrafanaFolder, error) {
	err := validateCreateGrafanaFolder(payload)
	if err != nil {
		return nil, err
	}

	createGrafanaFolderJson, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createGrafanaFolderServiceMethod,
		Url:          fmt.Sprintf(createGrafanaFolderServiceUrl, c.BaseUrl),
		Body:         createGrafanaFolderJson,
		SuccessCodes: []int{createGrafanaFolderMethodCreated},
		NotFoundCode: createS3FetcherStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationCreateGrafanaFolder,
		ResourceName: grafanaFolderResourceName,
	})

	if err != nil {
		return nil, err
	}

	var retVal GrafanaFolder
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}

func validateCreateGrafanaFolder(payload CreateUpdateFolder) error {
	if len(payload.Title) == 0 {
		return fmt.Errorf("Field title must be set!")
	}

	return nil
}
