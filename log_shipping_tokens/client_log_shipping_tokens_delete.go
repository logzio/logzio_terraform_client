package log_shipping_tokens

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteLogShippingTokenServiceMethod  = http.MethodDelete
	deleteLogShippingTokenServiceUrl     = logShippingTokensServiceEndpoint + "/%d"
	deleteLogShippingTokenMethodSuccess  = http.StatusOK
	deleteLogShippingTokenMethodNotFound = http.StatusNotFound
)

// DeleteLogShippingToken deletes a log shipping token, specified by its unique id, returns an error if a problem is encountered
func (c *LogShippingTokensClient) DeleteLogShippingToken(tokenId int32) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteLogShippingTokenServiceMethod,
		Url:          fmt.Sprintf(deleteLogShippingTokenServiceUrl, c.BaseUrl, tokenId),
		Body:         nil,
		SuccessCodes: []int{deleteLogShippingTokenMethodSuccess},
		NotFoundCode: deleteLogShippingTokenMethodNotFound,
		ResourceId:   tokenId,
		ApiAction:    operationDeleteLogShippingToken,
		ResourceName: logShippingTokenResourceName,
	})

	return err
}
