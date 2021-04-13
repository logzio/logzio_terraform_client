package alerts_v2

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
	"strings"
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
	req, _ := c.buildEnableApiRequest(c.ApiToken, alert.AlertId)

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{enableAlertMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "EnableAlert", resp.StatusCode, jsonBytes)
	}

	str := fmt.Sprintf("%s", jsonBytes)
	if strings.Contains(str, fmt.Sprintf("alert id %d not found", alert.AlertId)) {
		return nil, fmt.Errorf("API call %s failed with missing alert %d, data: %s", "EnableAlert", alert.AlertId, str)
	}

	return &alert, nil
}
