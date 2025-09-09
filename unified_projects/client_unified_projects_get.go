package unified_projects

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	getProjectMethod   = http.MethodGet
	getProjectSuccess  = http.StatusOK
	getProjectNotFound = http.StatusNotFound
)

func (c *ProjectsClient) GetProject(name string) (*ProjectSummary, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("name must be set")
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getProjectMethod,
		Url:          fmt.Sprintf(projectsByNameEndpoint, c.BaseUrl, name),
		Body:         nil,
		SuccessCodes: []int{getProjectSuccess},
		NotFoundCode: getProjectNotFound,
		ResourceId:   name,
		ApiAction:    getProjectOperation,
		ResourceName: projectResourceName,
	})
	if err != nil {
		return nil, err
	}

	var result ProjectSummary
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
