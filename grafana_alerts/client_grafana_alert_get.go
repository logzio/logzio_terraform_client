package grafana_alerts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getGrafanaAlertServiceUrl      = grafanaAlertServiceEndpoint + "/%s"
	getGrafanaAlertServiceMethod   = http.MethodGet
	getGrafanaAlertServiceSuccess  = http.StatusOK
	getGrafanaAlertServiceNotFound = http.StatusNotFound
)

func (c *GrafanaAlertClient) GetGrafanaAlertRule(uid string) (*GrafanaAlertRule, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getGrafanaAlertServiceMethod,
		Url:          fmt.Sprintf(getGrafanaAlertServiceUrl, c.BaseUrl, uid),
		Body:         nil,
		SuccessCodes: []int{getGrafanaAlertServiceSuccess},
		NotFoundCode: getGrafanaAlertServiceNotFound,
		ResourceId:   uid,
		ApiAction:    operationGetGrafanaAlert,
		ResourceName: grafanaAlertResourceName,
	})

	if err != nil {
		return nil, err
	}

	var grafanaAlertRule GrafanaAlertRule
	err = json.Unmarshal(res, &grafanaAlertRule)
	if err != nil {
		return nil, err
	}

	return &grafanaAlertRule, nil
}
