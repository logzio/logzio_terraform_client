package grafana_dashboards

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	grafanaObjectsDeleteDashboardsByUID         = grafanaObjectServiceEndpoint + "/uid/%s"
	grafanaObjectsDeleteDashboardsByUIDMethod   = http.MethodDelete
	grafanaObjectsDeleteDashboardsByUIDSuccess  = http.StatusOK
	grafanaObjectsDeleteDashboardsByUIDNotFound = http.StatusNotFound
)

// DeleteGrafanaDashboard allows deleting Grafana objects configuration.
// https://docs.logz.io/api/#operation/deleteDashboarById
func (c *GrafanaObjectsClient) DeleteGrafanaDashboard(objectUid string) (*DeleteResults, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   grafanaObjectsDeleteDashboardsByUIDMethod,
		Url:          fmt.Sprintf(grafanaObjectsDeleteDashboardsByUID, c.BaseUrl, objectUid),
		Body:         nil,
		SuccessCodes: []int{grafanaObjectsDeleteDashboardsByUIDSuccess},
		NotFoundCode: grafanaObjectsDeleteDashboardsByUIDNotFound,
		ApiAction:    dashboardDelete,
		ResourceId:   objectUid,
		ResourceName: dashboardResourceName,
	})

	if err != nil {
		return nil, err
	}

	var results DeleteResults
	err = json.Unmarshal(res, &results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}
