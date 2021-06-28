package log_shipping_tokens

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const createLogShippingTokenServiceUrl = LogShippingTokensServiceEndpoint
const createLogShippingTokenServiceMethod string = http.MethodPost

type FieldError struct {
	Field   string
	Message string
}

func (e FieldError) Error() string {
	return fmt.Sprintf("%v: %v", e.Field, e.Message)
}

// Create a log shipping token, return the created token if successful, an error otherwise
func (c *LogShippingTokensClient) CreateLogShippingToken(token CreateLogShippingToken) (*LogShippingToken, error) {
	err := validateCreateLogShippingTokenRequest(token)
	if err != nil {
		return nil, err
	}

	createToken, err := json.Marshal(token)
	if err != nil {
		return nil, err
	}

	req, _ := c.buildCreateApiRequest(c.ApiToken, createToken)
	jsonResponse, err := logzio_client.CreateHttpRequestBytesResponse(req)
	if err != nil {
		return nil, err
	}

	var retVal LogShippingToken
	err = json.Unmarshal(jsonResponse, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}

func (c *LogShippingTokensClient) buildCreateApiRequest(apiToken string, jsonBytes []byte) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(createLogShippingTokenServiceMethod, fmt.Sprintf(createLogShippingTokenServiceUrl, baseUrl), bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}