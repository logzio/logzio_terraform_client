package grafana_dashboards

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

// CreateUpdateGrafanaDashboard allows the creation or update of a Grafana dashboard
func (c *GrafanaObjectsClient) CreateUpdateGrafanaDashboard(payload CreateUpdatePayload) (*CreateUpdateResults, error) {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	var dashboardUid string
	if uid, ok := payload.Dashboard["uid"]; ok {
		dashboardUid = uid.(string)
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   grafanaObjectsCreateUpdateDashboardsByUIDMethod,
		Url:          fmt.Sprintf(grafanaObjectsDashboardsCreateUpdate, c.BaseUrl),
		Body:         payloadJson,
		SuccessCodes: []int{grafanaObjectsCreateUpdateDashboardsByUIDSuccess},
		NotFoundCode: grafanaObjectsCreateUpdateDashboardsByUIDNotFound,
		ApiAction:    dashboardCreateUpdate,
		ResourceId:   dashboardUid,
		ResourceName: dashboardResourceName,
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
