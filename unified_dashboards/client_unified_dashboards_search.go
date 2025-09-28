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

func (c *DashboardsClient) SearchDashboards(req SearchDashboardsRequest) ([]Dashboard, error) {
	if err := validateSearchDashboardsRequest(req); err != nil {
		return nil, err
	}

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
		ResourceId:   "search",
		ApiAction:    searchDashboardsOperation,
		ResourceName: dashboardResourceName,
	})
	if err != nil {
		return nil, err
	}

	var result []Dashboard
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, err
	}

	return result, nil
}
