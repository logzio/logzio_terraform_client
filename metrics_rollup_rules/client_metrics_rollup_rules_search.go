package metrics_rollup_rules

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	searchMetricsRollupRulesServiceUrl = metricsRollupRulesServiceEndpoint + "/search"
	searchRollupRulesMethod            = http.MethodPost
	searchRollupRulesSuccess           = http.StatusOK
	searchRollupRulesNotFound          = http.StatusNotFound
)

// SearchRollupRules retrieves rollup rules according to search
func (c *MetricsRollupRulesClient) SearchRollupRules(req SearchRollupRulesRequest) ([]RollupRule, error) {
	if err := validateSearchRollupRuleRequest(req); err != nil {
		return nil, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   searchRollupRulesMethod,
		Url:          fmt.Sprintf(searchMetricsRollupRulesServiceUrl, c.BaseUrl),
		Body:         body,
		SuccessCodes: []int{searchRollupRulesSuccess},
		NotFoundCode: searchRollupRulesNotFound,
		ResourceId:   nil,
		ApiAction:    operationSearchMetricsRollupRules,
		ResourceName: resourceName,
	})
	if err != nil {
		return nil, err
	}

	var wrapper SearchRollupRulesResponse
	if err := json.Unmarshal(res, &wrapper); err != nil {
		// Fallback to direct array unmarshal for backward compatibility
		var results []RollupRule
		if err := json.Unmarshal(res, &results); err != nil {
			return nil, err
		}
		return results, nil
	}

	return wrapper.Results, nil
}
