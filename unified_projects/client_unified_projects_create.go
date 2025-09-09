package unified_projects

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	createProjectMethod   = http.MethodPost
	createProjectSuccess  = http.StatusCreated
	createProjectNotFound = http.StatusNotFound
)

func (c *ProjectsClient) CreateProject(req CreateProjectRequest) (*ProjectSummary, error) {
	if err := validateCreateProjectRequest(req); err != nil {
		return nil, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createProjectMethod,
		Url:          fmt.Sprintf(projectsServiceEndpoint, c.BaseUrl),
		Body:         body,
		SuccessCodes: []int{createProjectSuccess},
		NotFoundCode: createProjectNotFound,
		ResourceId:   nil,
		ApiAction:    createProjectOperation,
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
