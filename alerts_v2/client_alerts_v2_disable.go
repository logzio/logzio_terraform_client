package alerts_v2

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const disableAlertServiceUrl = alertsServiceEndpoint + "/%d/disable"
const disableAlertServiceMethod string = http.MethodPost
const disableAlertMethodSuccess int = http.StatusNoContent

func (c *AlertsV2Client) buildDisableApiRequest(apiToken string, alertId int64) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(disableAlertServiceMethod, fmt.Sprintf(disableAlertServiceUrl, baseUrl, alertId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Disables an alert given it's unique identifier. Returns the alert, an error otherwise
func (c *AlertsV2Client) DisableAlert(alert AlertType) (*AlertType, error) {
	return c.EnableOrDisableAlert(alert, false)
}
