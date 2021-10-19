package alerts_v2

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	createAlertServiceUrl = alertsServiceEndpoint
	createAlertServiceMethod = http.MethodPost
	createAlertServiceSuccess = http.StatusOK
	createAlertServiceCreated = http.StatusCreated
	createAlertServiceNoContent = http.StatusNoContent
	createAlertServiceStatusNotFound = http.StatusNotFound
)

type FieldError struct {
	Field   string
	Message string
}

func (e FieldError) Error() string {
	return fmt.Sprintf("%v: %v", e.Field, e.Message)
}

// CreateAlert creates an alert, returns the created alert if successful, an error otherwise
func (c *AlertsV2Client) CreateAlert(alert CreateAlertType) (*AlertType, error) {
	err := validateCreateAlertRequest(alert)
	if err != nil {
		return nil, err
	}

	createAlertJson, err := json.Marshal(alert)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createAlertServiceMethod,
		Url:          fmt.Sprintf(createAlertServiceUrl, c.BaseUrl),
		Body:         createAlertJson,
		SuccessCodes: []int{createAlertServiceSuccess, createAlertServiceCreated, createAlertServiceNoContent},
		NotFoundCode: createAlertServiceStatusNotFound,
		ResourceId:   nil,
		ApiAction:    createAlertOperation,
	})

	if err != nil {
		return nil, err
	}

	var retVal AlertType
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}
