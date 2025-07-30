package drop_metrics

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	createDropMetricServiceUrl      = dropMetricsServiceEndpoint
	createDropMetricServiceMethod   = http.MethodPost
	createDropMetricServiceSuccess  = http.StatusOK
	createDropMetricServiceNotFound = http.StatusNotFound
)

// CreateDropMetric creates a drop metric filter, returns the created filter if successful, an error otherwise
func (c *DropMetricsClient) CreateDropMetric(req CreateDropMetric) (*DropMetric, error) {
	if err := validateCreateDropMetricRequest(req); err != nil {
		return nil, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createDropMetricServiceMethod,
		Url:          fmt.Sprintf(createDropMetricServiceUrl, c.BaseUrl),
		Body:         body,
		SuccessCodes: []int{createDropMetricServiceSuccess},
		NotFoundCode: createDropMetricServiceNotFound,
		ResourceId:   nil,
		ApiAction:    createDropMetricOperation,
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
