package unified_dashboards

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	moveDashboardMethod   = http.MethodPut
	moveDashboardSuccess  = http.StatusOK
	moveDashboardNotFound = http.StatusNotFound
)

// MoveDashboard relocates a dashboard to a different folder.
func (c *DashboardsClient) MoveDashboard(req MoveDashboardRequest) (*MoveDashboardResponse, error) {
	if err := validateMoveDashboardRequest(req); err != nil {
		return nil, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   moveDashboardMethod,
		Url:          fmt.Sprintf(dashboardsMoveEndpoint, c.BaseUrl),
		Body:         body,
		SuccessCodes: []int{moveDashboardSuccess},
		NotFoundCode: moveDashboardNotFound,
		ResourceId:   req.DashboardId,
		ApiAction:    moveDashboardOperation,
		ResourceName: dashboardResourceName,
	})
	if err != nil {
		return nil, err
	}

	var result MoveDashboardResponse
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, fmt.Errorf("%s: failed to unmarshal response: %w (body: %.200s)", moveDashboardOperation, err, res)
	}
	if len(result.Id) == 0 {
		return nil, fmt.Errorf("%s succeeded but the response contained no dashboard id (body: %.200s)", moveDashboardOperation, res)
	}

	return &result, nil
}
