package alerts_v2

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	updateAlertServiceUrl     = alertsServiceEndpoint + "/%d"
	updateAlertServiceMethod  = http.MethodPut
	updateAlertMethodSuccess  = http.StatusOK
	updateAlertMethodNotFound = http.StatusNotFound
)

// UpdateAlert updates an existing alert, based on the supplied alert identifier, using the parameters of the specified alert
// Returns the updated alert if successful, an error otherwise
func (c *AlertsV2Client) UpdateAlert(alertId int64, alert CreateAlertType) (*AlertType, error) {
	err := validateCreateAlertRequest(alert)
	if err != nil {
		return nil, err
	}

	updateAlertJson, err := json.Marshal(alert)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateAlertServiceMethod,
		Url:          fmt.Sprintf(updateAlertServiceUrl, c.BaseUrl, alertId),
		Body:         updateAlertJson,
		SuccessCodes: []int{updateAlertMethodSuccess},
		NotFoundCode: updateAlertMethodNotFound,
		ResourceId:   alert,
		ApiAction:    updateAlertOperation,
	})

	if err != nil {
		return nil, err
	}

	var target AlertType
	err = json.Unmarshal(res, &target)
	if err != nil {
		return nil, err
	}

	return &target, nil
}
