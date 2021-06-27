package log_shipping_tokens

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const getAvailableLogShippingTokensNumberServiceUrl = logShippingTokensServiceEndpoint + "/limits"
const getAvailableLogShippingTokensNumberServiceMethod string = http.MethodGet
const getAvailableLogShippingTokensNumberMethodSuccess int = http.StatusOK

func (c *LogShippingTokensClient) buildGetAvailableNumberApiRequest(apiToken string) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(getAvailableLogShippingTokensNumberServiceMethod, fmt.Sprintf(getAvailableLogShippingTokensNumberServiceUrl, baseUrl), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Returns the number of log shipping tokens currently in use and the number of available tokens that can be enabled,
// error otherwise.
// Disabled tokens don't count against the token limit.
func (c *LogShippingTokensClient) GetLogShippingLimitsToken() (*LogShippingTokensLimits, error) {
	req, _ := c.buildGetAvailableNumberApiRequest(c.ApiToken)
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{getAvailableLogShippingTokensNumberMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", operationGetLogShippingTokensLimits, resp.StatusCode, jsonBytes)
	}

	var limits LogShippingTokensLimits
	err = json.Unmarshal(jsonBytes, &limits)
	if err != nil {
		return nil, err
	}

	return &limits, nil
}