package drop_metrics

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	getDropMetricServiceUrl      = dropMetricsServiceEndpoint + "/%d"
	getDropMetricServiceMethod   = http.MethodGet
	getDropMetricServiceSuccess  = http.StatusOK
	getDropMetricServiceNotFound = http.StatusNotFound
)

// GetDropMetric retrieves a drop metric filter by ID, returns the filter if successful, an error otherwise
func (c *DropMetricsClient) GetDropMetric(dropFilterId int64) (*DropMetric, error) {
	if dropFilterId <= 0 {
		return nil, fmt.Errorf("dropFilterId must be greater than 0")
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getDropMetricServiceMethod,
		Url:          fmt.Sprintf(getDropMetricServiceUrl, c.BaseUrl, dropFilterId),
		Body:         nil,
		SuccessCodes: []int{getDropMetricServiceSuccess},
		NotFoundCode: getDropMetricServiceNotFound,
		ResourceId:   dropFilterId,
		ApiAction:    getDropMetricOperation,
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
