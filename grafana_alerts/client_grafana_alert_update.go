package grafana_alerts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	updateGrafanaAlertServiceUrl      = grafanaAlertServiceEndpoint + "/%s"
	updateGrafanaAlertServiceMethod   = http.MethodPut
	updateGrafanaAlertServiceSuccess  = http.StatusOK
	updateGrafanaAlertServiceNotFound = http.StatusNotFound
)

func (c *GrafanaAlertClient) UpdateGrafanaAlertRule(payload GrafanaAlertRule) error {
	err := validateGrafanaAlertRuleCreateUpdate(payload, true)
	if err != nil {
		return err
	}

	updateGrafanaAlertRuleJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateGrafanaAlertServiceMethod,
		Url:          fmt.Sprintf(updateGrafanaAlertServiceUrl, c.BaseUrl, payload.Uid),
		Body:         updateGrafanaAlertRuleJson,
		SuccessCodes: []int{updateGrafanaAlertServiceSuccess},
		NotFoundCode: updateGrafanaAlertServiceNotFound,
		ResourceId:   payload.Uid,
		ApiAction:    operationUpdateGrafanaAlert,
		ResourceName: grafanaAlertResourceName,
	})

	return err
}
