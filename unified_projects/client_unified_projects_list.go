package unified_projects

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	listProjectsMethod   = http.MethodGet
	listProjectsSuccess  = http.StatusOK
	listProjectsNotFound = http.StatusNotFound
)

func (c *ProjectsClient) ListProjects(withDashboards bool) ([]ProjectModel, error) {
	endpoint := fmt.Sprintf(projectsServiceEndpoint, c.BaseUrl)

	if withDashboards {
		u, err := url.Parse(endpoint)
		if err != nil {
			return nil, err
		}
		q := u.Query()
		q.Set("withDashboards", "true")
		u.RawQuery = q.Encode()
		endpoint = u.String()
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listProjectsMethod,
		Url:          endpoint,
		Body:         nil,
		SuccessCodes: []int{listProjectsSuccess},
		NotFoundCode: listProjectsNotFound,
		ResourceId:   nil,
		ApiAction:    listProjectsOperation,
		ResourceName: projectResourceName,
	})
	if err != nil {
		return nil, err
	}

	var result []ProjectModel
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, err
	}

	return result, nil
}
