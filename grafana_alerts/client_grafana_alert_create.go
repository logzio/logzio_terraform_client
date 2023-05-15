package grafana_alerts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	createGrafanaAlertServiceUrl     = grafanaAlertServiceEndpoint
	createGrafanaAlertServiceMethod  = http.MethodPost
	createGrafanaAlertMethodCreated  = http.StatusCreated
	createGrafanaAlertStatusNotFound = http.StatusNotFound
)

func (c *GrafanaAlertClient) CreateGrafanaAlertRule(payload GrafanaAlertRule) (*GrafanaAlertRule, error) {
	err := validateGrafanaAlertRule(payload)
	if err != nil {
		return nil, err
	}

	createGrafanaAlertRuleJson, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createGrafanaAlertServiceMethod,
		Url:          fmt.Sprintf(createGrafanaAlertServiceUrl, c.BaseUrl),
		Body:         createGrafanaAlertRuleJson,
		SuccessCodes: []int{createGrafanaAlertMethodCreated},
		NotFoundCode: createGrafanaAlertStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationCreateGrafanaAlert,
		ResourceName: grafanaAlertResourceName,
	})

	if err != nil {
		return nil, err
	}

	var retVal GrafanaAlertRule
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}
