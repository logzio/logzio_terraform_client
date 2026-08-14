package unified_dashboards

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	listDashboardsMethod   = http.MethodGet
	listDashboardsSuccess  = http.StatusOK
	listDashboardsNotFound = http.StatusNotFound
)

func (c *DashboardsClient) ListDashboards() ([]Dashboard, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listDashboardsMethod,
		Url:          fmt.Sprintf(dashboardsListEndpoint, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{listDashboardsSuccess},
		NotFoundCode: listDashboardsNotFound,
		ResourceId:   nil,
		ApiAction:    listDashboardsOperation,
		ResourceName: dashboardResourceName,
	})
	if err != nil {
		return nil, err
	}

	var result []Dashboard
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, fmt.Errorf("%s: failed to unmarshal response: %w (body: %.200s)", listDashboardsOperation, err, res)
	}

	return result, nil
}
