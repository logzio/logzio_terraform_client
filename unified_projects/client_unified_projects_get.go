package unified_projects

import (
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	getProjectMethod   = http.MethodGet
	getProjectSuccess  = http.StatusOK
	getProjectNotFound = http.StatusNotFound
)

// GetProject returns a project (dashboard folder) by its id.
func (c *ProjectsClient) GetProject(id string) (*ProjectSummary, error) {
	if err := validateGetProjectRequest(id); err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getProjectMethod,
		Url:          fmt.Sprintf(projectsByIdEndpoint, c.BaseUrl, id),
		Body:         nil,
		SuccessCodes: []int{getProjectSuccess},
		NotFoundCode: getProjectNotFound,
		ResourceId:   id,
		ApiAction:    getProjectOperation,
		ResourceName: projectResourceName,
	})
	if err != nil {
		return nil, err
	}

	return unmarshalProject(getProjectOperation, res)
}
