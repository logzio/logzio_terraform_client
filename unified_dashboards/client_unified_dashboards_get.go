package unified_dashboards

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	getDashboardMethod   = http.MethodGet
	getDashboardSuccess  = http.StatusOK
	getDashboardNotFound = http.StatusNotFound
)

// GetDashboard returns a dashboard by its folder id and uid.
func (c *DashboardsClient) GetDashboard(folderId, uid string) (*Dashboard, error) {
	if err := validateGetDashboardRequest(folderId, uid); err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getDashboardMethod,
		Url:          fmt.Sprintf(dashboardByUidEndpoint, c.BaseUrl, folderId, uid),
		Body:         nil,
		SuccessCodes: []int{getDashboardSuccess},
		NotFoundCode: getDashboardNotFound,
		ResourceId:   uid,
		ApiAction:    getDashboardOperation,
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
