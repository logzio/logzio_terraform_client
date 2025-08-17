package metrics_rollup_rules

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	updateMetricsRollupRulesServiceUrl = metricsRollupRulesServiceEndpoint + "/%s"
	updateRollupRuleMethod             = http.MethodPut
	updateRollupRuleSuccess            = http.StatusOK
	updateRollupRuleNotFound           = http.StatusNotFound
)

// UpdateRollupRule updates a rollup rule by id
func (c *MetricsRollupRulesClient) UpdateRollupRule(rollupRuleId string, req CreateUpdateRollupRule) (*RollupRule, error) {
	err := validateRollupRuleId(rollupRuleId)
	if err != nil {
		return nil, err
	}
	if err = validateCreateUpdateRollupRuleRequest(req); err != nil {
		return nil, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(updateMetricsRollupRulesServiceUrl, c.BaseUrl, rollupRuleId)
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateRollupRuleMethod,
		Url:          url,
		Body:         body,
		SuccessCodes: []int{updateRollupRuleSuccess},
		NotFoundCode: updateRollupRuleNotFound,
		ResourceId:   rollupRuleId,
		ApiAction:    operationUpdateMetricsRollupRule,
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
