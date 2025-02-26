package grafana_folders

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
	"time"
)

const (
	createGrafanaFolderServiceUrl     = grafanaFolderServiceEndpoint
	createGrafanaFolderServiceMethod  = http.MethodPost
	createGrafanaFolderMethodCreated  = http.StatusOK
	createGrafanaFolderStatusNotFound = http.StatusNotFound
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
		NotFoundCode: createGrafanaFolderStatusNotFound,
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

	retVal.Updated = normalizeTimestamp(retVal.Updated)
	retVal.Created = normalizeTimestamp(retVal.Created)

	return &retVal, nil
}

func validateCreateGrafanaFolder(payload CreateUpdateFolder) error {
	if len(payload.Title) == 0 {
		return fmt.Errorf("Field title must be set!")
	}

	return nil
}

func normalizeTimestamp(timestamp string) string {
	parsedTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return timestamp
	}
	return parsedTime.Format(time.RFC3339)
}
