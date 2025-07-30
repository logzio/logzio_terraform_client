package drop_metrics

import (
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	deleteDropMetricServiceUrl      = dropMetricsServiceEndpoint + "/%d"
	deleteDropMetricServiceMethod   = http.MethodDelete
	deleteDropMetricServiceSuccess  = http.StatusNoContent
	deleteDropMetricServiceNotFound = http.StatusNotFound
)

// DeleteDropMetric deletes a drop metric filter by ID, returns an error if the operation fails
func (c *DropMetricsClient) DeleteDropMetric(dropFilterId int64) error {
	if dropFilterId <= 0 {
		return fmt.Errorf("dropFilterId must be greater than 0")
	}

	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteDropMetricServiceMethod,
		Url:          fmt.Sprintf(deleteDropMetricServiceUrl, c.BaseUrl, dropFilterId),
		Body:         nil,
		SuccessCodes: []int{deleteDropMetricServiceSuccess},
		NotFoundCode: deleteDropMetricServiceNotFound,
		ResourceId:   dropFilterId,
		ApiAction:    deleteDropMetricOperation,
		ResourceName: resourceName,
	})

	return err
}
