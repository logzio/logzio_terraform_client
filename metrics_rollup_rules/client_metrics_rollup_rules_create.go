package metrics_rollup_rules

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	createMetricsRollupRuleServiceUrl = metricsRollupRulesServiceEndpoint
	createRollupRuleMethod            = http.MethodPost
	createRollupRuleSuccess           = http.StatusOK
	createRollupRuleNotFound          = http.StatusNotFound
)

// CreateRollupRule creates a rollup rule
func (c *MetricsRollupRulesClient) CreateRollupRule(req CreateUpdateRollupRule) (*RollupRule, error) {
	if err := validateCreateUpdateRollupRuleRequest(req); err != nil {
		return nil, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createRollupRuleMethod,
		Url:          fmt.Sprintf(createMetricsRollupRuleServiceUrl, c.BaseUrl),
		Body:         body,
		SuccessCodes: []int{createRollupRuleSuccess},
		NotFoundCode: createRollupRuleNotFound,
		ResourceId:   nil,
		ApiAction:    operationCreateMetricsRollupRule,
		ResourceName: resourceName,
	})
	if err != nil {
		return nil, err
	}

	var result RollupRule
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
