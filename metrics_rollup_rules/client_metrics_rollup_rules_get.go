package metrics_rollup_rules

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	getMetricsRollupRulesServiceUrl = metricsRollupRulesServiceEndpoint + "/%s"
	getRollupRuleMethod             = http.MethodGet
	getRollupRuleSuccess            = http.StatusOK
	getRollupRuleNotFound           = http.StatusNotFound
)

// GetRollupRule retrieves a rollup rule by id
func (c *MetricsRollupRulesClient) GetRollupRule(rollupRuleId string) (*RollupRule, error) {
	err := validateRollupRuleId(rollupRuleId)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(getMetricsRollupRulesServiceUrl, c.BaseUrl, rollupRuleId)
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getRollupRuleMethod,
		Url:          url,
		Body:         nil,
		SuccessCodes: []int{getRollupRuleSuccess},
		NotFoundCode: getRollupRuleNotFound,
		ResourceId:   rollupRuleId,
		ApiAction:    operationGetMetricsRollupRule,
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
