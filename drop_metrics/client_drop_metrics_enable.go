package drop_metrics

import (
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	enableDropMetricServiceUrl      = dropMetricsServiceEndpoint + "/%d/enable"
	enableDropMetricServiceMethod   = http.MethodPost
	enableDropMetricServiceSuccess  = http.StatusOK
	enableDropMetricServiceNotFound = http.StatusNotFound
)

// EnableDropMetric enables a drop metric filter by ID, returns an error if the operation fails
func (c *DropMetricsClient) EnableDropMetric(dropFilterId int64) error {
	if dropFilterId <= 0 {
		return fmt.Errorf("dropFilterId must be greater than 0")
	}

	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   enableDropMetricServiceMethod,
		Url:          fmt.Sprintf(enableDropMetricServiceUrl, c.BaseUrl, dropFilterId),
		Body:         nil,
		SuccessCodes: []int{enableDropMetricServiceSuccess},
		NotFoundCode: enableDropMetricServiceNotFound,
		ResourceId:   dropFilterId,
		ApiAction:    enableDropMetricOperation,
		ResourceName: resourceName,
	})

	return err
}
