package drop_metrics

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	bulkCreateDropMetricServiceUrl      = dropMetricsServiceEndpoint + "/bulk/create"
	bulkCreateDropMetricServiceMethod   = http.MethodPost
	bulkCreateDropMetricServiceSuccess  = http.StatusOK
	bulkCreateDropMetricServiceNotFound = http.StatusNotFound
)

// BulkCreateDropMetrics creates multiple drop metric filters in bulk, returns the created filters if successful, an error otherwise
func (c *DropMetricsClient) BulkCreateDropMetrics(requests []CreateDropMetric) ([]DropMetric, error) {
	if len(requests) == 0 {
		return nil, fmt.Errorf("requests array cannot be empty")
	}

	// Validate each request
	for i, req := range requests {
		if err := validateCreateDropMetricRequest(req); err != nil {
			return nil, fmt.Errorf("request[%d] validation failed: %w", i, err)
		}
	}

	body, err := json.Marshal(requests)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   bulkCreateDropMetricServiceMethod,
		Url:          fmt.Sprintf(bulkCreateDropMetricServiceUrl, c.BaseUrl),
		Body:         body,
		SuccessCodes: []int{bulkCreateDropMetricServiceSuccess},
		NotFoundCode: bulkCreateDropMetricServiceNotFound,
		ResourceId:   nil,
		ApiAction:    bulkCreateDropMetricOperation,
		ResourceName: resourceName,
	})

	if err != nil {
		return nil, err
	}

	var results []DropMetric
	if err := json.Unmarshal(res, &results); err != nil {
		return nil, err
	}

	return results, nil
}
