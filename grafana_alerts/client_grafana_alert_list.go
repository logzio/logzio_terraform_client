package grafana_alerts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	listGrafanaAlertServiceUrl     = grafanaAlertServiceEndpoint
	listGrafanaAlertServiceMethod  = http.MethodGet
	listGrafanaAlertServiceSuccess = http.StatusOK
	listGrafanaAlertStatusNotFound = http.StatusNotFound
)

func (c *GrafanaAlertClient) ListGrafanaAlertRules() ([]GrafanaAlertRule, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listGrafanaAlertServiceMethod,
		Url:          fmt.Sprintf(listGrafanaAlertServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{listGrafanaAlertServiceSuccess},
		NotFoundCode: listGrafanaAlertStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationListGrafanaAlerts,
		ResourceName: grafanaAlertResourceName,
	})

	if err != nil {
		return nil, err
	}

	var alertRules []GrafanaAlertRule
	err = json.Unmarshal(res, &alertRules)
	if err != nil {
		return nil, err
	}

	return alertRules, nil
}
