package unified_dashboards

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	updateDashboardMethod   = http.MethodPut
	updateDashboardSuccess  = http.StatusOK
	updateDashboardNotFound = http.StatusNotFound
)

func (c *DashboardsClient) UpdateDashboard(folderId, uid string, req UpdateDashboardRequest) (*Dashboard, error) {
	if err := validateUpdateDashboardRequest(folderId, uid, req); err != nil {
		return nil, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateDashboardMethod,
		Url:          fmt.Sprintf(dashboardByUidEndpoint, c.BaseUrl, folderId, uid),
		Body:         body,
		SuccessCodes: []int{updateDashboardSuccess},
		NotFoundCode: updateDashboardNotFound,
		ResourceId:   uid,
		ApiAction:    updateDashboardOperation,
		ResourceName: dashboardResourceName,
	})
	if err != nil {
		return nil, err
	}

	var result Dashboard
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
