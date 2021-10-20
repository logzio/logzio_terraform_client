package log_shipping_tokens

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	createLogShippingTokenServiceUrl      = logShippingTokensServiceEndpoint
	createLogShippingTokenServiceMethod   = http.MethodPost
	createLogShippingTokenMethodSuccess   = http.StatusOK
	createLogShippingTokenMethodCreated   = http.StatusCreated
	createLogShippingTokenMethodNoContent = http.StatusNoContent
	createLogShippingTokenStatusNotFound  = http.StatusNotFound
)

type FieldError struct {
	Field   string
	Message string
}

func (e FieldError) Error() string {
	return fmt.Sprintf("%v: %v", e.Field, e.Message)
}

// CreateLogShippingToken creates a log shipping token, returns the created token if successful, an error otherwise
func (c *LogShippingTokensClient) CreateLogShippingToken(token CreateLogShippingToken) (*LogShippingToken, error) {
	err := validateCreateLogShippingTokenRequest(token)
	if err != nil {
		return nil, err
	}

	createTokenJson, err := json.Marshal(token)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createLogShippingTokenServiceMethod,
		Url:          fmt.Sprintf(createLogShippingTokenServiceUrl, c.BaseUrl),
		Body:         createTokenJson,
		SuccessCodes: []int{createLogShippingTokenMethodSuccess, createLogShippingTokenMethodCreated, createLogShippingTokenMethodNoContent},
		NotFoundCode: createLogShippingTokenStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationCreateLogShippingToken,
	})

	if err != nil {
		return nil, err
	}

	var retVal LogShippingToken
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}
