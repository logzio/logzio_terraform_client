package log_shipping_tokens

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	retrieveLogShippingTokenServiceUrl     = logShippingTokensServiceEndpoint + "/search"
	retrieveLogShippingTokenServiceMethod  = http.MethodPost
	retrieveLogShippingTokenMethodSuccess  = http.StatusOK
	retrieveLogShippingTokenStatusNotFound = http.StatusNotFound
)

// RetrieveLogShippingTokens returns the relevant shipping tokens, filtered, sorted and paginated as per the request, error otherwise
func (c *LogShippingTokensClient) RetrieveLogShippingTokens(retrieveRequest RetrieveLogShippingTokensRequest) (*RetrieveLogShippingTokensResponse, error) {
	err := validateRetrieveLogShippingTokensRequest(retrieveRequest)
	if err != nil {
		return nil, err
	}

	searchTokensJson, err := json.Marshal(retrieveRequest)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   retrieveLogShippingTokenServiceMethod,
		Url:          fmt.Sprintf(retrieveLogShippingTokenServiceUrl, c.BaseUrl),
		Body:         searchTokensJson,
		SuccessCodes: []int{retrieveLogShippingTokenMethodSuccess},
		NotFoundCode: retrieveLogShippingTokenStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationRetrieveLogShippingTokens,
	})

	if err != nil {
		return nil, err
	}

	var target RetrieveLogShippingTokensResponse
	err = json.Unmarshal(res, &target)
	if err != nil {
		return nil, err
	}

	return &target, nil
}
