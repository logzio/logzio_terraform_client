package unified_projects

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	createProjectMethod   = http.MethodPost
	createProjectSuccess  = http.StatusOK
	createProjectCreated  = http.StatusCreated
	createProjectNotFound = http.StatusNotFound
)

// CreateProject creates a project (dashboard folder). The API expects a
// Perses Project document; it is built here from the request fields.
func (c *ProjectsClient) CreateProject(req CreateProjectRequest) (*ProjectSummary, error) {
	if err := validateCreateProjectRequest(req); err != nil {
		return nil, err
	}

	displayName := req.DisplayName
	if len(displayName) == 0 {
		displayName = req.Name
	}

	body, err := json.Marshal(newProjectEnvelope(req.Name, displayName, req.Description))
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createProjectMethod,
		Url:          fmt.Sprintf(projectsServiceEndpoint, c.BaseUrl),
		Body:         body,
		SuccessCodes: []int{createProjectSuccess, createProjectCreated},
		NotFoundCode: createProjectNotFound,
		ResourceId:   req.Name,
		ApiAction:    createProjectOperation,
		ResourceName: projectResourceName,
	})
	if err != nil {
		return nil, err
	}

	return unmarshalProject(createProjectOperation, res)
}
