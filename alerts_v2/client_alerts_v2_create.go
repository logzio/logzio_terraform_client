package alerts_v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)


const createAlertServiceUrl = alertsServiceEndpoint
const createAlertServiceMethod string = http.MethodPost

type FieldError struct {
	Field   string
	Message string
}

func (e FieldError) Error() string {
	return fmt.Sprintf("%v: %v", e.Field, e.Message)
}

// Create an alert, return the created alert if successful, an error otherwise
func (c *AlertsV2Client) CreateAlert(alert CreateAlertType) (*AlertType, error) {
	err := validateCreateAlertRequest(alert)
	if err != nil {
		return nil, err
	}

	createAlert, err := json.Marshal(alert)
	if err != nil {
		return nil, err
	}

	req, _ := c.buildCreateApiRequest(c.ApiToken, createAlert)
	jsonResponse, err := logzio_client.CreateHttpRequestBytesResponse(req)
	if err != nil {
		return nil, err
	}

	var retVal AlertType
	err = json.Unmarshal(jsonResponse, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}

func (c *AlertsV2Client) buildCreateApiRequest(apiToken string, jsonBytes []byte) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(createAlertServiceMethod, fmt.Sprintf(createAlertServiceUrl, baseUrl), bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}