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
	NewName string `json:"newName"`
}

// RenameProject renames a project (dashboard folder), addressed by its folder id.
func (c *ProjectsClient) RenameProject(folderId string, newName string) (*ProjectSummary, error) {
	if err := validateRenameProjectRequest(folderId, newName); err != nil {
		return nil, err
	}

	body, err := json.Marshal(renameProjectRequest{NewName: newName})
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   renameProjectMethod,
		Url:          fmt.Sprintf(projectsRenameEndpoint, c.BaseUrl, folderId),
		Body:         body,
		SuccessCodes: []int{renameProjectSuccess},
		NotFoundCode: renameProjectNotFound,
		ResourceId:   folderId,
		ApiAction:    renameProjectOperation,
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
