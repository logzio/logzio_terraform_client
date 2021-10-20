package log_shipping_tokens

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getAvailableLogShippingTokensNumberServiceUrl     = logShippingTokensServiceEndpoint + "/limits"
	getAvailableLogShippingTokensNumberServiceMethod  = http.MethodGet
	getAvailableLogShippingTokensNumberMethodSuccess  = http.StatusOK
	getAvailableLogShippingTokensNumberStatusNotFound = http.StatusNotFound
)

// GetLogShippingLimitsToken returns the number of log shipping tokens currently in use and the number of available tokens that can be enabled,
// error otherwise.
// Disabled tokens don't count against the token limit.
func (c *LogShippingTokensClient) GetLogShippingLimitsToken() (*LogShippingTokensLimits, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getAvailableLogShippingTokensNumberServiceMethod,
		Url:          fmt.Sprintf(getAvailableLogShippingTokensNumberServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{getAvailableLogShippingTokensNumberMethodSuccess},
		NotFoundCode: getAvailableLogShippingTokensNumberStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationGetLogShippingTokensLimits,
	})

	if err != nil {
		return nil, err
	}

	var limits LogShippingTokensLimits
	err = json.Unmarshal(res, &limits)
	if err != nil {
		return nil, err
	}

	return &limits, nil
}
