package drop_metrics

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	bulkDeleteDropMetricsServiceUrl      = dropMetricsServiceEndpoint + "/bulk/delete"
	bulkDeleteDropMetricsServiceMethod   = http.MethodPost
	bulkDeleteDropMetricsServiceSuccess  = http.StatusNoContent
	bulkDeleteDropMetricsServiceNotFound = http.StatusNotFound
)

// BulkDeleteDropMetrics deletes multiple drop metric filters by IDs, returns an error if the operation fails
func (c *DropMetricsClient) BulkDeleteDropMetrics(dropFilterIds []int64) error {
	if len(dropFilterIds) == 0 {
		return fmt.Errorf("dropFilterIds array cannot be empty")
	}

	// Validate all IDs
	for i, id := range dropFilterIds {
		if id <= 0 {
			return fmt.Errorf("dropFilterIds[%d] must be greater than 0", i)
		}
	}

	body, err := json.Marshal(dropFilterIds)
	if err != nil {
		return err
	}

	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   bulkDeleteDropMetricsServiceMethod,
		Url:          fmt.Sprintf(bulkDeleteDropMetricsServiceUrl, c.BaseUrl),
		Body:         body,
		SuccessCodes: []int{http.StatusOK, bulkDeleteDropMetricsServiceSuccess},
		NotFoundCode: bulkDeleteDropMetricsServiceNotFound,
		ResourceId:   nil,
		ApiAction:    bulkDeleteDropMetricOperation,
		ResourceName: resourceName,
	})

	return err
}
