package log_shipping_tokens

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	updateLogShippingTokenServiceUrl     = logShippingTokensServiceEndpoint + "/%d"
	updateLogShippingTokenServiceMethod  = http.MethodPut
	updateLogShippingTokenMethodSuccess  = http.StatusOK
	updateLogShippingTokenMethodNotFound = http.StatusNotFound
)

// UpdateLogShippingToken updates an existing log shipping token, based on the supplied token identifier, using the parameters of the specified token
// Returns the updated token if successful, an error otherwise
func (c *LogShippingTokensClient) UpdateLogShippingToken(tokenId int32, token UpdateLogShippingToken) (*LogShippingToken, error) {
	err := validateUpdateLogShippingTokenRequest(token)
	if err != nil {
		return nil, err
	}

	updateTokenJson, err := json.Marshal(token)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateLogShippingTokenServiceMethod,
		Url:          fmt.Sprintf(updateLogShippingTokenServiceUrl, c.BaseUrl, tokenId),
		Body:         updateTokenJson,
		SuccessCodes: []int{updateLogShippingTokenMethodSuccess},
		NotFoundCode: updateLogShippingTokenMethodNotFound,
		ResourceId:   tokenId,
		ApiAction:    operationUpdateLogShippingToken,
	})

	if err != nil {
		return nil, err
	}

	var target LogShippingToken
	err = json.Unmarshal(res, &target)
	if err != nil {
		return nil, err
	}

	return &target, nil
}
