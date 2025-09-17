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

func (c *ProjectsClient) DeleteProject(folderId string) error {
	if err := validateDeleteProjectRequest(folderId); err != nil {
		return err
	}

	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteProjectMethod,
		Url:          fmt.Sprintf(projectsByNameEndpoint, c.BaseUrl, folderId),
		Body:         nil,
		SuccessCodes: []int{deleteProjectSuccess, deleteProjectOk},
		NotFoundCode: deleteProjectNotFound,
		ResourceId:   folderId,
		ApiAction:    deleteProjectOperation,
		ResourceName: projectResourceName,
	})
	if err != nil {
		return err
	}

	return nil
}
