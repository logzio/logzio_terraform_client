package metrics_rollup_rules

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	bulkCreateMetricsRollupRulesServiceUrl = metricsRollupRulesServiceEndpoint + "/bulk/create"
	bulkCreateRollupRulesMethod            = http.MethodPost
	bulkCreateRollupRulesSuccess           = http.StatusOK
	bulkCreateRollupRulesNotFound          = http.StatusNotFound
)

// BulkCreateRollupRules creates multiple rollup rules
func (c *MetricsRollupRulesClient) BulkCreateRollupRules(req []CreateUpdateRollupRule) ([]RollupRule, error) {
	if len(req) == 0 {
		return nil, fmt.Errorf("bulk create request should contain at least one rollup rule")
	}
	for i, r := range req {
		if err := validateCreateUpdateRollupRuleRequest(r); err != nil {
			return nil, fmt.Errorf("request[%d] validation failed: %w", i, err)
		}
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   bulkCreateRollupRulesMethod,
		Url:          fmt.Sprintf(bulkCreateMetricsRollupRulesServiceUrl, c.BaseUrl),
		Body:         body,
		SuccessCodes: []int{bulkCreateRollupRulesSuccess, http.StatusCreated},
		NotFoundCode: bulkCreateRollupRulesNotFound,
		ResourceId:   nil,
		ApiAction:    operationBulkCreateMetricsRollupRules,
		ResourceName: resourceName,
	})
	if err != nil {
		return nil, err
	}

	var result []RollupRule
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, err
	}
	return result, nil
}
