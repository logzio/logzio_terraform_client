package alerts_v2

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteAlertServiceMethod = http.MethodDelete
	deleteAlertServiceUrl = alertsServiceEndpoint + "/%d"
	deleteAlertMethodSuccess int = http.StatusOK
	deleteAlertMethodNotFound int = http.StatusNotFound
)

// DeleteAlert deletes an alert specified by its unique id, returns an error if a problem is encountered
func (c *AlertsV2Client) DeleteAlert(alertId int64) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteAlertServiceMethod,
		Url:          fmt.Sprintf(deleteAlertServiceUrl, c.BaseUrl, alertId),
		Body:         nil,
		SuccessCodes: []int{deleteAlertMethodSuccess},
		NotFoundCode: deleteAlertMethodNotFound,
		ResourceId:   alertId,
		ApiAction:    deleteAlertOperation,
	})

	return err
}
