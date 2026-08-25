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

// UpdateProject updates a project (dashboard folder), addressed by its id.
// The API expects the full Perses Project document and replaces it, so the
// request must carry the complete desired state (see UpdateProjectRequest).
func (c *ProjectsClient) UpdateProject(id string, req UpdateProjectRequest) (*ProjectSummary, error) {
	if err := validateUpdateProjectRequest(id, req); err != nil {
		return nil, err
	}

	body, err := json.Marshal(newProjectEnvelope(req.Name, req.DisplayName, req.Description))
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateProjectMethod,
		Url:          fmt.Sprintf(projectsByIdEndpoint, c.BaseUrl, id),
		Body:         body,
		SuccessCodes: []int{updateProjectSuccess},
		NotFoundCode: updateProjectNotFound,
		ResourceId:   id,
		ApiAction:    updateProjectOperation,
		ResourceName: projectResourceName,
	})
	if err != nil {
		return nil, err
	}

	return unmarshalProject(updateProjectOperation, res)
}
