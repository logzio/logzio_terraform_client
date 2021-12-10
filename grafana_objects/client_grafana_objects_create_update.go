package grafana_objects

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	grafanaObjectsDashboardsCreateUpdate              = grafanaObjectServiceEndpoint + "/db"
	grafanaObjectsCreateUpdateDashboardsByUIDMethod   = http.MethodPost
	grafanaObjectsCreateUpdateDashboardsByUIDSuccess  = http.StatusOK
	grafanaObjectsCreateUpdateDashboardsByUIDNotFound = http.StatusNotFound
)

// CreateUpdate allows the creation or update of a Grafana dashboard
// https://docs.logz.io/api/#operation/createDashboard
func (c *GrafanaObjectsClient) CreateUpdate(payload CreateUpdatePayload) (*CreateUpdateResults, error) {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   grafanaObjectsCreateUpdateDashboardsByUIDMethod,
		Url:          fmt.Sprintf(grafanaObjectsDashboardsCreateUpdate, c.BaseUrl),
		Body:         payloadJson,
		SuccessCodes: []int{grafanaObjectsCreateUpdateDashboardsByUIDSuccess},
		NotFoundCode: grafanaObjectsCreateUpdateDashboardsByUIDNotFound,
		ApiAction:    "CreateUpdate",
		ResourceId:   payload.Dashboard.Id,
		ResourceName: "Dashboard",
	})

	if err != nil {
		return nil, err
	}

	var results CreateUpdateResults
	err = json.Unmarshal(res, &results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}
