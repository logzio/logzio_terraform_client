package unified_alerts

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteUnifiedAlertServiceMethod = http.MethodDelete
	deleteUnifiedAlertServiceUrl    = unifiedAlertsServiceEndpoint + "/%s/%s"
	deleteUnifiedAlertSuccess       = http.StatusOK
	deleteUnifiedAlertNotFound      = http.StatusNotFound
)

// DeleteUnifiedAlert deletes a unified alert specified by its type and unique id, returns the deleted alert response if successful, an error otherwise
func (c *UnifiedAlertsClient) DeleteUnifiedAlert(alertType string, alertId string) (*UnifiedAlert, error) {
	err := validateUrlType(alertType)
	if err != nil {
		return nil, err
	}

	if len(alertId) == 0 {
		return nil, fmt.Errorf("alertId must be set")
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteUnifiedAlertServiceMethod,
		Url:          fmt.Sprintf(deleteUnifiedAlertServiceUrl, c.BaseUrl, alertType, alertId),
		Body:         nil,
		SuccessCodes: []int{deleteUnifiedAlertSuccess},
		NotFoundCode: deleteUnifiedAlertNotFound,
		ResourceId:   alertId,
		ApiAction:    deleteUnifiedAlertOperation,
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
