package alerts_v2

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	disableAlertServiceUrl                   = alertsServiceEndpoint + "/%d/disable"
	disableAlertServiceMethod         string = http.MethodPost
	disableAlertServiceSuccess        int    = http.StatusNoContent
	disableAlertServiceStatusNotFound        = http.StatusNotFound
)

// DisableAlert disables an alert given its unique identifier. Returns the alert, an error otherwise
func (c *AlertsV2Client) DisableAlert(alert AlertType) (*AlertType, error) {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   disableAlertServiceMethod,
		Url:          fmt.Sprintf(disableAlertServiceUrl, c.BaseUrl, alert.AlertId),
		Body:         nil,
		SuccessCodes: []int{disableAlertServiceSuccess},
		NotFoundCode: disableAlertServiceStatusNotFound,
		ResourceId:   alert.AlertId,
		ApiAction:    disableAlertOperation,
	})

	if err != nil {
		return nil, err
	}

	alert.Enabled = false
	return &alert, nil
}
