package unified_alerts

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	createUnifiedAlertServiceUrl     = unifiedAlertsServiceEndpoint + "/%s"
	createUnifiedAlertServiceMethod  = http.MethodPost
	createUnifiedAlertServiceSuccess = http.StatusOK
	createUnifiedAlertServiceCreated = http.StatusCreated
	createUnifiedAlertNotFound       = http.StatusNotFound
)

// CreateUnifiedAlert creates a unified alert, returns the created alert if successful, an error otherwise
func (c *UnifiedAlertsClient) CreateUnifiedAlert(alertType string, req CreateUnifiedAlert) (*UnifiedAlert, error) {
	err := validateUrlType(alertType)
	if err != nil {
		return nil, err
	}

	err = validateCreateUnifiedAlertRequest(req)
	if err != nil {
		return nil, err
	}

	createUnifiedAlertJson, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createUnifiedAlertServiceMethod,
		Url:          fmt.Sprintf(createUnifiedAlertServiceUrl, c.BaseUrl, alertType),
		Body:         createUnifiedAlertJson,
		SuccessCodes: []int{createUnifiedAlertServiceSuccess, createUnifiedAlertServiceCreated},
		NotFoundCode: createUnifiedAlertNotFound,
		ResourceId:   nil,
		ApiAction:    createUnifiedAlertOperation,
		ResourceName: unifiedAlertResourceName,
	})

	if err != nil {
		return nil, err
	}

	var retVal UnifiedAlert
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}
