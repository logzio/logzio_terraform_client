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

func (c *ProjectsClient) UpdateProject(name string, project Project) (*Project, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("name must be set")
	}
	if err := validateUpdateProjectRequest(project); err != nil {
		return nil, err
	}

	body, err := json.Marshal(project)
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

	var result Project
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
