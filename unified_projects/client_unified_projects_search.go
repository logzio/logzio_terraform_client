package unified_projects

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	searchProjectsMethod   = http.MethodPost
	searchProjectsSuccess  = http.StatusOK
	searchProjectsNotFound = http.StatusNotFound
)

func (c *ProjectsClient) SearchProjects(req SearchProjectsRequest) (*SearchProjectsResponse, error) {
	if err := validateSearchProjectsRequest(req); err != nil {
		return nil, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   searchProjectsMethod,
		Url:          fmt.Sprintf(projectsSearchEndpoint, c.BaseUrl),
		Body:         body,
		SuccessCodes: []int{searchProjectsSuccess},
		NotFoundCode: searchProjectsNotFound,
		ResourceId:   "search",
		ApiAction:    searchProjectsOperation,
		ResourceName: projectResourceName,
	})
	if err != nil {
		return nil, err
	}

	var result SearchProjectsResponse
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
