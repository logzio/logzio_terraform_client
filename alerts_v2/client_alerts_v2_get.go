package alerts_v2

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getAlertServiceUrl            = alertsServiceEndpoint + "/%d"
	getAlertServiceMethod  string = http.MethodGet
	getAlertMethodSuccess  int    = http.StatusOK
	getAlertMethodNotFound int    = http.StatusNotFound
)

// GetAlert returns an alert given itss unique identifier, an error otherwise
func (c *AlertsV2Client) GetAlert(alertId int64) (*AlertType, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getAlertServiceMethod,
		Url:          fmt.Sprintf(getAlertServiceUrl, c.BaseUrl, alertId),
		Body:         nil,
		SuccessCodes: []int{getAlertMethodSuccess},
		NotFoundCode: getAlertMethodNotFound,
		ResourceId:   alertId,
		ApiAction:    getAlertOperation,
		ResourceName: alertResourceName,
	})

	if err != nil {
		return nil, err
	}

	var alert AlertType
	err = json.Unmarshal(res, &alert)
	if err != nil {
		return nil, err
	}

	return &alert, nil
}
