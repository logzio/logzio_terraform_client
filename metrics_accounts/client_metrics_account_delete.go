package metrics_accounts

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteMetricsAccountServiceUrl     = metricsAccountServiceEndpoint + "/%d"
	deleteMetricsAccountServiceMethod  = http.MethodDelete
	deleteMetricsAccountServiceSuccess = http.StatusOK
	deleteMetricsAccountMethodNotFound = http.StatusNoContent
)

// DeleteMetricsAccount deletes a metrics account specified by its unique id, returns an error if a problem is encountered
func (c *MetricsAccountClient) DeleteMetricsAccount(metricsAccountId int64) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteMetricsAccountServiceMethod,
		Url:          fmt.Sprintf(deleteMetricsAccountServiceUrl, c.BaseUrl, metricsAccountId),
		Body:         nil,
		SuccessCodes: []int{deleteMetricsAccountServiceSuccess},
		NotFoundCode: deleteMetricsAccountMethodNotFound,
		ResourceId:   metricsAccountId,
		ApiAction:    operationDeleteMetricsAccount,
		ResourceName: metricsAccountResourceName,
	})

	return err
}
