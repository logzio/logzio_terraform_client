package unified_alerts

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	updateUnifiedAlertServiceUrl    = unifiedAlertsServiceEndpoint + "/%s/%s"
	updateUnifiedAlertServiceMethod = http.MethodPut
	updateUnifiedAlertSuccess       = http.StatusOK
	updateUnifiedAlertNotFound      = http.StatusNotFound
)

// UpdateUnifiedAlert updates an existing unified alert, based on the supplied alert type and identifier, using the parameters of the specified alert
// Returns the updated alert if successful, an error otherwise
func (c *UnifiedAlertsClient) UpdateUnifiedAlert(alertType string, alertId string, req CreateUnifiedAlert) (*UnifiedAlert, error) {
	err := validateUrlType(alertType)
	if err != nil {
		return nil, err
	}

	if len(alertId) == 0 {
		return nil, fmt.Errorf("alertId must be set")
	}

	err = validateCreateUnifiedAlertRequest(req)
	if err != nil {
		return nil, err
	}

	updateUnifiedAlertJson, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateUnifiedAlertServiceMethod,
		Url:          fmt.Sprintf(updateUnifiedAlertServiceUrl, c.BaseUrl, alertType, alertId),
		Body:         updateUnifiedAlertJson,
		SuccessCodes: []int{updateUnifiedAlertSuccess},
		NotFoundCode: updateUnifiedAlertNotFound,
		ResourceId:   alertId,
		ApiAction:    updateUnifiedAlertOperation,
		ResourceName: unifiedAlertResourceName,
	})

	if err != nil {
		return nil, err
	}

	var target UnifiedAlert
	err = json.Unmarshal(res, &target)
	if err != nil {
		return nil, err
	}

	return &target, nil
}
