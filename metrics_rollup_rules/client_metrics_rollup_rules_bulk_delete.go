package metrics_rollup_rules

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	bulkDeleteMetricsRollupRulesServiceUrl = metricsRollupRulesServiceEndpoint + "/bulk/delete"
	bulkDeleteRollupRulesMethod            = http.MethodPost
	bulkDeleteRollupRulesSuccess           = http.StatusNoContent
	bulkDeleteRollupRulesNotFound          = http.StatusNotFound
)

// BulkDeleteRollupRules deletes rollup rules by ids
func (c *MetricsRollupRulesClient) BulkDeleteRollupRules(ruleIds []string) error {
	if len(ruleIds) == 0 {
		return fmt.Errorf("ruleIds to bulk delete must not be empty")
	}

	body, err := json.Marshal(ruleIds)
	if err != nil {
		return err
	}

	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   bulkDeleteRollupRulesMethod,
		Url:          fmt.Sprintf(bulkDeleteMetricsRollupRulesServiceUrl, c.BaseUrl),
		Body:         body,
		SuccessCodes: []int{bulkDeleteRollupRulesSuccess, http.StatusOK},
		NotFoundCode: bulkDeleteRollupRulesNotFound,
		ResourceId:   nil,
		ApiAction:    operationBulkDeleteMetricsRollupRules,
		ResourceName: resourceName,
	})
	if err != nil {
		return err
	}
	return nil
}
