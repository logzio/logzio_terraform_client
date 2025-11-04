package unified_alerts

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	getUnifiedAlertServiceUrl    = unifiedAlertsServiceEndpoint + "/%s/%s"
	getUnifiedAlertServiceMethod = http.MethodGet
	getUnifiedAlertSuccess       = http.StatusOK
	getUnifiedAlertNotFound      = http.StatusNotFound
)

// GetUnifiedAlert returns a unified alert given its type and identifier, an error otherwise
func (c *UnifiedAlertsClient) GetUnifiedAlert(alertType string, alertId string) (*UnifiedAlert, error) {
	err := validateUrlType(alertType)
	if err != nil {
		return nil, err
	}

	if len(alertId) == 0 {
		return nil, fmt.Errorf("alertId must be set")
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getUnifiedAlertServiceMethod,
		Url:          fmt.Sprintf(getUnifiedAlertServiceUrl, c.BaseUrl, alertType, alertId),
		Body:         nil,
		SuccessCodes: []int{getUnifiedAlertSuccess},
		NotFoundCode: getUnifiedAlertNotFound,
		ResourceId:   alertId,
		ApiAction:    getUnifiedAlertOperation,
		ResourceName: unifiedAlertResourceName,
	})

	if err != nil {
		return nil, err
	}

	var alert UnifiedAlert
	err = json.Unmarshal(res, &alert)
	if err != nil {
		return nil, err
	}

	return &alert, nil
}
