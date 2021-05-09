package alerts_v2

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const enableAlertServiceUrl = alertsServiceEndpoint + "/%d/enable"
const enableAlertServiceMethod string = http.MethodPost
const enableAlertMethodSuccess int = http.StatusNoContent

func (c *AlertsV2Client) buildEnableApiRequest(apiToken string, alertId int64) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(enableAlertServiceMethod, fmt.Sprintf(enableAlertServiceUrl, baseUrl, alertId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Enables an alert given it's unique identifier. Returns the alert, an error otherwise
func (c *AlertsV2Client) EnableAlert(alert AlertType) (*AlertType, error) {
	return c.EnableOrDisableAlert(alert, true)
}
