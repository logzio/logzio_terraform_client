package alerts_v2

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	listAlertServiceUrl    = alertsServiceEndpoint
	listAlertServiceMethod = http.MethodGet
	listAlertMethodSuccess = http.StatusOK
)

// ListAlerts returns all the alerts in an array associated with the account identified by the supplied API token, returns an error if
// any problem occurs during the API call
func (c *AlertsV2Client) ListAlerts() ([]AlertType, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listAlertServiceMethod,
		Url:          fmt.Sprintf(listAlertServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{listAlertMethodSuccess},
		NotFoundCode: http.StatusNotFound,
		ResourceId:   nil,
		ApiAction:    listAlertOperation,
		ResourceName: alertResourceName,
	})

	if err != nil {
		return nil, err
	}

	var alerts []AlertType
	err = json.Unmarshal(res, &alerts)

	if err != nil {
		return nil, err
	}

	return alerts, nil
}
