package alerts_v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const updateAlertServiceUrl string = alertsServiceEndpoint + "/%d"
const updateAlertServiceMethod string = http.MethodPut
const updateAlertMethodSuccess int = http.StatusOK


func (c *AlertsV2Client) buildUpdateApiRequest(apiToken string, alertId int64, alert CreateAlertType) (*http.Request, error) {
	jsonBytes, err := json.Marshal(alert)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	req, err := http.NewRequest(updateAlertServiceMethod, fmt.Sprintf(updateAlertServiceUrl, baseUrl, alertId), bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Updates an existing alert, based on the supplied alert identifier, using the parameters of the specified alert
// Returns the updated alert if successful, an error otherwise
func (c *AlertsV2Client) UpdateAlert(alertId int64, alert CreateAlertType) (*AlertType, error) {
	err := validateCreateAlertRequest(alert)
	if err != nil {
		return nil, err
	}

	req, err := c.buildUpdateApiRequest(c.ApiToken, alertId, alert)
	if err != nil {
		return nil, err
	}

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{updateAlertMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "UpdateAlert", resp.StatusCode, jsonBytes)
	}

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("API call %s failed with missing alert %d, data: %s", "UpdateAlert", alertId, jsonBytes)
	}

	var target AlertType
	json.Unmarshal(jsonBytes, &target)

	return &target, nil
}
