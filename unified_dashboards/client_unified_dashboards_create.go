package unified_dashboards

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	createDashboardMethod   = http.MethodPost
	createDashboardSuccess  = http.StatusCreated
	createDashboardNotFound = http.StatusNotFound
)

func (c *DashboardsClient) CreateDashboard(folderId string, req CreateDashboardRequest) (*Dashboard, error) {
	if err := validateCreateDashboardRequest(folderId, req); err != nil {
		return nil, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createDashboardMethod,
		Url:          fmt.Sprintf(dashboardsCreateEndpoint, c.BaseUrl, folderId),
		Body:         body,
		SuccessCodes: []int{createDashboardSuccess, http.StatusOK},
		NotFoundCode: createDashboardNotFound,
		ResourceId:   folderId,
		ApiAction:    createDashboardOperation,
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
