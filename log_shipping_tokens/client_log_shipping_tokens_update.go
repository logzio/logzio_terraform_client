package log_shipping_tokens

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const updateLogShippingTokenServiceUrl string = logShippingTokensServiceEndpoint + "/%d"
const updateLogShippingTokenServiceMethod string = http.MethodPut
const updateLogShippingTokenMethodSuccess int = http.StatusOK
const updateLogShippingTokenMethodNotFound int = http.StatusNotFound

func (c *LogShippingTokensClient) buildUpdateApiRequest(apiToken string, tokenId int32, token UpdateLogShippingToken) (*http.Request, error) {
	jsonBytes, err := json.Marshal(token)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	req, err := http.NewRequest(updateLogShippingTokenServiceMethod, fmt.Sprintf(updateLogShippingTokenServiceUrl, baseUrl, tokenId), bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Updates an existing log shipping token, based on the supplied token identifier, using the parameters of the specified token
// Returns the updated token if successful, an error otherwise
func (c *LogShippingTokensClient) UpdateLogShippingToken(tokenId int32, token UpdateLogShippingToken) (*LogShippingToken, error) {
	err := validateUpdateLogShippingTokenRequest(token)
	if err != nil {
		return nil, err
	}

	req, err := c.buildUpdateApiRequest(c.ApiToken, tokenId, token)
	if err != nil {
		return nil, err
	}

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{updateLogShippingTokenMethodSuccess}) {
		if resp.StatusCode == updateLogShippingTokenMethodNotFound {
			return nil, fmt.Errorf("API call %s failed with missing log shippinng token %d, data: %s", operationUpdateLogShippingToken, tokenId, jsonBytes)
		}

		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", operationUpdateLogShippingToken, resp.StatusCode, jsonBytes)
	}

	var target LogShippingToken
	err = json.Unmarshal(jsonBytes, &target)
	if err != nil {
		return nil, err
	}

	return &target, nil
}
