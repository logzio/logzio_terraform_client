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
	err := validateCreateGrafanaAlertRule(payload)
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

func validateCreateGrafanaAlertRule(payload GrafanaAlertRule) error {
	if len(payload.Condition) == 0 {
		return fmt.Errorf("Field condition must be set!")
	}

	if payload.Data == nil {
		return fmt.Errorf("Field data must be set!")
	}

	if len(payload.ExecErrState) == 0 {
		return fmt.Errorf("Field execErrState must be set!")
	}

	if len(payload.FolderUID) == 0 {
		return fmt.Errorf("Field folderUID must be set!")
	}

	if len(payload.For) == 0 {
		return fmt.Errorf("Field for must be set!")
	}

	if len(payload.NoDataState) == 0 {
		return fmt.Errorf("Field noDataState must be set!")
	}

	if payload.OrgID == 0 {
		return fmt.Errorf("Field orgID must be set!")
	}

	if len(payload.RuleGroup) == 0 {
		return fmt.Errorf("Field ruleGroup must be set!")
	}

	if len(payload.Title) == 0 {
		return fmt.Errorf("Field title must be set!")
	}

	return nil
}
