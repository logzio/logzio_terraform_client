package grafana_objects

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	grafanaObjectsDashboardsByUID            = grafanaObjectServiceEndpoint + "/uid/%s"
	grafanaObjectsGetDashboardsByUIDMethod   = http.MethodGet
	grafanaObjectsGetDashboardsByUIDSuccess  = http.StatusOK
	grafanaObjectsGetDashboardsByUIDNotFound = http.StatusNotFound
)

// Get allows getting a Grafana objects configuration.
// https://docs.logz.io/api/#operation/getDashboarById
func (c *GrafanaObjectsClient) Get(objectUid string) (*GetResults, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   grafanaObjectsGetDashboardsByUIDMethod,
		Url:          fmt.Sprintf(grafanaObjectsDashboardsByUID, c.BaseUrl, objectUid),
		Body:         nil,
		SuccessCodes: []int{grafanaObjectsGetDashboardsByUIDSuccess},
		NotFoundCode: grafanaObjectsGetDashboardsByUIDNotFound,
		ApiAction:    "GetDashboard",
		ResourceId:   objectUid,
		ResourceName: "Dashboard",
	})

	if err != nil {
		return nil, err
	}

	var results GetResults
	err = json.Unmarshal(res, &results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}
