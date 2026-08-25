package unified_projects

import (
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	deleteProjectMethod   = http.MethodDelete
	deleteProjectSuccess  = http.StatusNoContent
	deleteProjectOk       = http.StatusOK
	deleteProjectNotFound = http.StatusNotFound
)

// DeleteProject deletes a project (dashboard folder), addressed by its id.
func (c *ProjectsClient) DeleteProject(id string) error {
	if err := validateDeleteProjectRequest(id); err != nil {
		return err
	}

	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteProjectMethod,
		Url:          fmt.Sprintf(projectsByIdEndpoint, c.BaseUrl, id),
		Body:         nil,
		SuccessCodes: []int{deleteProjectSuccess, deleteProjectOk},
		NotFoundCode: deleteProjectNotFound,
		ResourceId:   id,
		ApiAction:    deleteProjectOperation,
		ResourceName: projectResourceName,
	})
	if err != nil {
		return err
	}

	return nil
}
