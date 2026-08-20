package unified_dashboards

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	searchDashboardsMethod   = http.MethodPost
	searchDashboardsSuccess  = http.StatusOK
	searchDashboardsNotFound = http.StatusNotFound
)

// SearchDashboards searches dashboards account-wide and returns them flat —
// unlike unified_projects.SearchProjects, which returns the same underlying
// search grouped by folder.
//
// Filter.SearchTerm genuinely narrows the result set here (the projects
// endpoint returns every folder regardless of the term). Total is the
// unpaginated match count, so it can exceed len(Results); page through with
// Pagination. The endpoint is POST-only — GET /dashboards/search 404s.
func (c *DashboardsClient) SearchDashboards(req SearchDashboardsRequest) (*SearchDashboardsResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   searchDashboardsMethod,
		Url:          fmt.Sprintf(dashboardsSearchEndpoint, c.BaseUrl),
		Body:         body,
		SuccessCodes: []int{searchDashboardsSuccess},
		NotFoundCode: searchDashboardsNotFound,
		ResourceId:   nil,
		ApiAction:    searchDashboardsOperation,
		ResourceName: dashboardResourceName,
	})
	if err != nil {
		return nil, err
	}

	var result SearchDashboardsResponse
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, fmt.Errorf("%s: failed to unmarshal response: %w (body: %.200s)", searchDashboardsOperation, err, res)
	}

	return &result, nil
}
