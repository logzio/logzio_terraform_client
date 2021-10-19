package alerts_v2

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	enableAlertServiceUrl            = alertsServiceEndpoint + "/%d/enable"
	enableAlertServiceMethod  string = http.MethodPost
	enableAlertMethodSuccess  int    = http.StatusNoContent
	enableAlertMethodNotFound int    = http.StatusNotFound
)

// EnableAlert enables an alert given its unique identifier. Returns the alert, an error otherwise
func (c *AlertsV2Client) EnableAlert(alert AlertType) (*AlertType, error) {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   enableAlertServiceMethod,
		Url:          fmt.Sprintf(enableAlertServiceUrl, c.BaseUrl, alert.AlertId),
		Body:         nil,
		SuccessCodes: []int{enableAlertMethodSuccess},
		NotFoundCode: enableAlertMethodNotFound,
		ResourceId:   alert.AlertId,
		ApiAction:    enableAlertOperation,
	})

	if err != nil {
		return nil, err
	}

	alert.Enabled = true
	return &alert, nil
}
