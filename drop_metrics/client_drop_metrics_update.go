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

// UpdateDropMetric updates a drop metric filter by ID, returns the updated filter if successful, an error otherwise
func (c *DropMetricsClient) UpdateDropMetric(dropFilterId int64, req CreateUpdateDropMetric) (*DropMetric, error) {
	if dropFilterId <= 0 {
		return nil, fmt.Errorf("dropFilterId must be greater than 0")
	}

	if err := validateCreateUpdateDropMetricRequest(req); err != nil {
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
