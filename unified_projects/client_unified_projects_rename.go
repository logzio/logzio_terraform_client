package unified_projects

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	renameProjectMethod   = http.MethodPut
	renameProjectSuccess  = http.StatusOK
	renameProjectNotFound = http.StatusNotFound
)

type renameProjectRequest struct {
	NewProjectName string `json:"newProjectName"`
}

// RenameProject renames a project (dashboard folder), addressed by its id.
func (c *ProjectsClient) RenameProject(id string, newName string) (*ProjectSummary, error) {
	if err := validateRenameProjectRequest(id, newName); err != nil {
		return nil, err
	}

	body, err := json.Marshal(renameProjectRequest{NewProjectName: newName})
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   renameProjectMethod,
		Url:          fmt.Sprintf(projectsRenameEndpoint, c.BaseUrl, id),
		Body:         body,
		SuccessCodes: []int{renameProjectSuccess},
		NotFoundCode: renameProjectNotFound,
		ResourceId:   id,
		ApiAction:    renameProjectOperation,
		ResourceName: projectResourceName,
	})
	if err != nil {
		return nil, err
	}

	return unmarshalProject(renameProjectOperation, res)
}
