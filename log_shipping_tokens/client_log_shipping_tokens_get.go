package log_shipping_tokens

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const getLogShippingTokenServiceUrl = logShippingTokensServiceEndpoint + "/%d"
const getLogShippingTokenServiceMethod string = http.MethodGet
const getLogShippingTokenMethodSuccess int = http.StatusOK
const getLogShippingTokenMethodNotFound int = http.StatusNotFound

func (c *LogShippingTokensClient) buildGetApiRequest(apiToken string, tokenId int32) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(getLogShippingTokenServiceMethod, fmt.Sprintf(getLogShippingTokenServiceUrl, baseUrl, tokenId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Returns a log shipping token given it's unique identifier, an error otherwise
func (c *LogShippingTokensClient) GetLogShippingToken(tokenId int32) (*LogShippingToken, error) {
	req, _ := c.buildGetApiRequest(c.ApiToken, tokenId)
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{getLogShippingTokenMethodSuccess}) {
		if resp.StatusCode == getLogShippingTokenMethodNotFound {
			return nil, fmt.Errorf("API call %s failed with missing log shipping token %d, data: %s", operationGetLogShippingToken, tokenId, jsonBytes)
		}

		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", operationGetLogShippingToken, resp.StatusCode, jsonBytes)
	}

	var token LogShippingToken
	err = json.Unmarshal(jsonBytes, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
