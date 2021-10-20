package log_shipping_tokens

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getLogShippingTokenServiceUrl     = logShippingTokensServiceEndpoint + "/%d"
	getLogShippingTokenServiceMethod  = http.MethodGet
	getLogShippingTokenMethodSuccess  = http.StatusOK
	getLogShippingTokenMethodNotFound = http.StatusNotFound
)

// GetLogShippingToken returns a log shipping token given its unique identifier, an error otherwise
func (c *LogShippingTokensClient) GetLogShippingToken(tokenId int32) (*LogShippingToken, error) {

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getLogShippingTokenServiceMethod,
		Url:          fmt.Sprintf(getLogShippingTokenServiceUrl, c.BaseUrl, tokenId),
		Body:         nil,
		SuccessCodes: []int{getLogShippingTokenMethodSuccess},
		NotFoundCode: getLogShippingTokenMethodNotFound,
		ResourceId:   tokenId,
		ApiAction:    operationGetLogShippingToken,
	})

	if err != nil {
		return nil, err
	}

	var token LogShippingToken
	err = json.Unmarshal(res, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
