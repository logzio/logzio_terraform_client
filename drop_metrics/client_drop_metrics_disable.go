package drop_metrics

import (
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	disableDropMetricServiceUrl      = dropMetricsServiceEndpoint + "/%d/disable"
	disableDropMetricServiceMethod   = http.MethodPost
	disableDropMetricServiceSuccess  = http.StatusOK
	disableDropMetricServiceNotFound = http.StatusNotFound
)

// DisableDropMetric disables a drop metric filter by ID, returns an error if the operation fails
func (c *DropMetricsClient) DisableDropMetric(dropFilterId int64) error {
	if dropFilterId <= 0 {
		return fmt.Errorf("dropFilterId must be greater than 0")
	}

	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   disableDropMetricServiceMethod,
		Url:          fmt.Sprintf(disableDropMetricServiceUrl, c.BaseUrl, dropFilterId),
		Body:         nil,
		SuccessCodes: []int{disableDropMetricServiceSuccess, http.StatusAccepted, http.StatusCreated},
		NotFoundCode: disableDropMetricServiceNotFound,
		ResourceId:   dropFilterId,
		ApiAction:    disableDropMetricOperation,
		ResourceName: resourceName,
	})

	return err
}
