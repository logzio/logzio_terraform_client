package unified_projects

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	searchProjectsMethod   = http.MethodGet
	searchProjectsSuccess  = http.StatusOK
	searchProjectsNotFound = http.StatusNotFound
)

// SearchProjects searches projects (dashboard folders) by name.
func (c *ProjectsClient) SearchProjects(req SearchProjectsRequest) ([]ProjectSummary, error) {
	if err := validateSearchProjectsRequest(req); err != nil {
		return nil, err
	}

	u, err := url.Parse(fmt.Sprintf(projectsSearchEndpoint, c.BaseUrl))
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("query", req.Query)
	if req.Limit > 0 {
		q.Set("limit", strconv.Itoa(req.Limit))
	}
	if req.Page > 0 {
		q.Set("page", strconv.Itoa(req.Page))
	}
	if len(req.Sort) > 0 {
		q.Set("sort", req.Sort)
	}
	u.RawQuery = q.Encode()

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   searchProjectsMethod,
		Url:          u.String(),
		Body:         nil,
		SuccessCodes: []int{searchProjectsSuccess},
		NotFoundCode: searchProjectsNotFound,
		ResourceId:   "search",
		ApiAction:    searchProjectsOperation,
		ResourceName: projectResourceName,
	})
	if err != nil {
		return nil, err
	}

	var result []ProjectSummary
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, err
	}

	return result, nil
}
