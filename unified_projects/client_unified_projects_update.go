package unified_projects

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	updateProjectMethod   = http.MethodPut
	updateProjectSuccess  = http.StatusOK
	updateProjectNotFound = http.StatusNotFound
)

// UpdateProject updates a project's (dashboard folder's) display name and/or description.
// The project is addressed by its name.
func (c *ProjectsClient) UpdateProject(name string, req UpdateProjectRequest) (*ProjectSummary, error) {
	if err := validateUpdateProjectRequest(name, req); err != nil {
		return nil, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateProjectMethod,
		Url:          fmt.Sprintf(projectsByNameEndpoint, c.BaseUrl, name),
		Body:         body,
		SuccessCodes: []int{updateProjectSuccess},
		NotFoundCode: updateProjectNotFound,
		ResourceId:   name,
		ApiAction:    updateProjectOperation,
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
