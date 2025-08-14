package metrics_rollup_rules

import (
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	deleteMetricsRollupRulesServiceUrl = metricsRollupRulesServiceEndpoint + "/%s"
	deleteRollupRuleMethod             = http.MethodDelete
	deleteRollupRuleSuccess            = http.StatusOK
	deleteRollupRuleNotFound           = http.StatusNotFound
)

// DeleteRollupRule deletes a rollup rule by id
func (c *MetricsRollupRulesClient) DeleteRollupRule(rollupRuleId string) error {
	err := validateRollupRuleId(rollupRuleId)
	if err != nil {
		return err
	}
	url := fmt.Sprintf(deleteMetricsRollupRulesServiceUrl, c.BaseUrl, rollupRuleId)
	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteRollupRuleMethod,
		Url:          url,
		Body:         nil,
		SuccessCodes: []int{deleteRollupRuleSuccess},
		NotFoundCode: deleteRollupRuleNotFound,
		ResourceId:   rollupRuleId,
		ApiAction:    operationDeleteMetricsRollupRule,
		ResourceName: resourceName,
	})
	if err != nil {
		return err
	}
	return nil
}
