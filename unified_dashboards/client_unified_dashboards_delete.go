package unified_dashboards

import (
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	deleteDashboardMethod   = http.MethodDelete
	deleteDashboardSuccess  = http.StatusNoContent
	deleteDashboardOk       = http.StatusOK
	deleteDashboardNotFound = http.StatusNotFound
)

func (c *DashboardsClient) DeleteDashboard(folderId, uid string) error {
	if err := validateDeleteDashboardRequest(folderId, uid); err != nil {
		return err
	}

	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteDashboardMethod,
		Url:          fmt.Sprintf(dashboardByUidEndpoint, c.BaseUrl, folderId, uid),
		Body:         nil,
		SuccessCodes: []int{deleteDashboardSuccess, deleteDashboardOk},
		NotFoundCode: deleteDashboardNotFound,
		ResourceId:   uid,
		ApiAction:    deleteDashboardOperation,
		ResourceName: dashboardResourceName,
	})
	if err != nil {
		return err
	}

	return nil
}
