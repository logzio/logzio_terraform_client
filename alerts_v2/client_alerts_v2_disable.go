package alerts_v2

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
	"strings"
)

const disableAlertServiceUrl = alertsServiceEndpoint + "/%d/enable"
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
	req, _ := c.buildDisableApiRequest(c.ApiToken, alert.AlertId)

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{disableAlertMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "DisableAlert", resp.StatusCode, jsonBytes)
	}

	str := fmt.Sprintf("%s", jsonBytes)
	if strings.Contains(str, fmt.Sprintf("alert id %d not found", alert.AlertId)) {
		return nil, fmt.Errorf("API call %s failed with missing alert %d, data: %s", "DisableAlert", alert.AlertId, str)
	}

	return &alert, nil
}
