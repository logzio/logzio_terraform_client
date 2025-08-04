package drop_metrics

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	searchDropMetricsServiceUrl      = dropMetricsServiceEndpoint + "/search"
	searchDropMetricsServiceMethod   = http.MethodPost
	searchDropMetricsServiceSuccess  = http.StatusOK
	searchDropMetricsServiceNotFound = http.StatusNotFound
)

// SearchDropMetricsResponse represents the response wrapper for search results
type SearchDropMetricsResponse struct {
	Results    []DropMetric `json:"results"`
	Pagination *Pagination  `json:"pagination,omitempty"`
}

// SearchDropMetrics searches for drop metric filters based on criteria, returns the filters if successful, an error otherwise
func (c *DropMetricsClient) SearchDropMetrics(req SearchDropMetricsRequest) ([]DropMetric, error) {
	if err := validateSearchDropMetricsRequest(req); err != nil {
		return nil, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   searchDropMetricsServiceMethod,
		Url:          fmt.Sprintf(searchDropMetricsServiceUrl, c.BaseUrl),
		Body:         body,
		SuccessCodes: []int{searchDropMetricsServiceSuccess},
		NotFoundCode: searchDropMetricsServiceNotFound,
		ResourceId:   nil,
		ApiAction:    searchDropMetricsOperation,
		ResourceName: resourceName,
	})

	if err != nil {
		return nil, err
	}

	// Try to unmarshal as wrapped response first
	var wrapper SearchDropMetricsResponse
	if err := json.Unmarshal(res, &wrapper); err != nil {
		// Fallback to direct array unmarshal for backward compatibility
		var results []DropMetric
		if err := json.Unmarshal(res, &results); err != nil {
			return nil, err
		}
		return results, nil
	}

	return wrapper.Results, nil
}
