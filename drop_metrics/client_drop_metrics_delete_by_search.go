package drop_metrics

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	deleteBySearchDropMetricsServiceUrl      = dropMetricsServiceEndpoint
	deleteBySearchDropMetricsServiceMethod   = http.MethodDelete
	deleteBySearchDropMetricsServiceSuccess  = http.StatusNoContent
	deleteBySearchDropMetricsServiceNotFound = http.StatusNotFound
)

// DeleteDropMetricsBySearch deletes drop metric filters based on search criteria, returns an error if the operation fails
func (c *DropMetricsClient) DeleteDropMetricsBySearch(req SearchDropMetricsRequest) error {
	if err := validateSearchDropMetricsRequest(req); err != nil {
		return err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteBySearchDropMetricsServiceMethod,
		Url:          fmt.Sprintf(deleteBySearchDropMetricsServiceUrl, c.BaseUrl),
		Body:         body,
		SuccessCodes: []int{deleteBySearchDropMetricsServiceSuccess},
		NotFoundCode: deleteBySearchDropMetricsServiceNotFound,
		ResourceId:   nil,
		ApiAction:    deleteBySearchOperation,
		ResourceName: resourceName,
	})

	return err
}
