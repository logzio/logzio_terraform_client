package alerts_v2

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const getAlertServiceUrl = alertsServiceEndpoint + "/%d"
const getAlertServiceMethod string = http.MethodGet
const getAlertMethodSuccess int = http.StatusOK

func (c *AlertsV2Client) buildGetApiRequest(apiToken string, alertId int64) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(getAlertServiceMethod, fmt.Sprintf(getAlertServiceUrl, baseUrl, alertId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Returns an alert given it's unique identifier, an error otherwise
func (c *AlertsV2Client) GetAlert(alertId int64) (*AlertType, error) {
	req, _ := c.buildGetApiRequest(c.ApiToken, alertId)

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{getAlertMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "GetAlert", resp.StatusCode, jsonBytes)
	}

	str := fmt.Sprintf("%s", jsonBytes)
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("API call %s failed with missing alert %d, data: %s", "GetAlert", alertId, str)
	}

	var alert AlertType
	err = json.Unmarshal(jsonBytes, &alert)
	if err != nil {
		return nil, err
	}

	return &alert, nil
}