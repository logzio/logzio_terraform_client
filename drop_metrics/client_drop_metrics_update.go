package drop_metrics

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	updateDropMetricServiceUrl      = dropMetricsServiceEndpoint + "/%d"
	updateDropMetricServiceMethod   = http.MethodPut
	updateDropMetricServiceSuccess  = http.StatusOK
	updateDropMetricServiceNotFound = http.StatusNotFound
)

// UpdateDropMetric represents the request payload for updating a drop metric filter
type UpdateDropMetric struct {
	AccountId int64        `json:"accountId"`
	Active    *bool        `json:"active,omitempty"`
	Filter    FilterObject `json:"filter"`
}

// UpdateDropMetric updates a drop metric filter by ID, returns the updated filter if successful, an error otherwise
func (c *DropMetricsClient) UpdateDropMetric(dropFilterId int64, req UpdateDropMetric) (*DropMetric, error) {
	if dropFilterId <= 0 {
		return nil, fmt.Errorf("dropFilterId must be greater than 0")
	}

	if err := validateUpdateDropMetricRequest(req); err != nil {
		return nil, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateDropMetricServiceMethod,
		Url:          fmt.Sprintf(updateDropMetricServiceUrl, c.BaseUrl, dropFilterId),
		Body:         body,
		SuccessCodes: []int{updateDropMetricServiceSuccess},
		NotFoundCode: updateDropMetricServiceNotFound,
		ResourceId:   dropFilterId,
		ApiAction:    updateDropMetricOperation,
		ResourceName: resourceName,
	})

	if err != nil {
		return nil, err
	}

	var result DropMetric
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// validateUpdateDropMetricRequest validates the update drop metric request
func validateUpdateDropMetricRequest(req UpdateDropMetric) error {
	if req.AccountId <= 0 {
		return fmt.Errorf("accountId must be set and greater than 0")
	}

	if err := validateFilterObject(req.Filter); err != nil {
		return fmt.Errorf("filter validation failed: %w", err)
	}

	return nil
}
