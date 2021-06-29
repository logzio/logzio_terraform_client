package log_shipping_tokens

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const deleteLogShippingTokenServiceMethod string = http.MethodDelete
const deleteLogShippingTokenServiceUrl = logShippingTokensServiceEndpoint + "/%d"
const deleteLogShippingTokenMethodSuccess int = http.StatusOK
const deleteLogShippingTokenMethodNotFound int = http.StatusNotFound

func (c *LogShippingTokensClient) buildDeleteApiRequest(apiToken string, tokenId int32) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(deleteLogShippingTokenServiceMethod, fmt.Sprintf(deleteLogShippingTokenServiceUrl, baseUrl, tokenId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Delete a log shipping token, specified by it's unique id, returns an error if a problem is encountered
func (c *LogShippingTokensClient) DeleteLogShippingToken(tokenId int32) error {
	req, _ := c.buildDeleteApiRequest(c.ApiToken, tokenId)

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{deleteLogShippingTokenMethodSuccess}) {
		if resp.StatusCode == deleteLogShippingTokenMethodNotFound {
			return fmt.Errorf("API call %s failed with missing log shipping token %d, data: %s", operationDeleteLogShippingToken, tokenId, jsonBytes)
		}

		return fmt.Errorf("API call %s failed with status code %d, data: %s", operationDeleteLogShippingToken, resp.StatusCode, jsonBytes)
	}

	return nil
}
