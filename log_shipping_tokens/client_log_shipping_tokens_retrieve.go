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

const retrieveLogShippingTokenServiceUrl = logShippingTokensServiceEndpoint + "/search"
const retrieveLogShippingTokenServiceMethod string = http.MethodPost
const retrieveLogShippingTokenMethodSuccess int = http.StatusOK

func (c *LogShippingTokensClient) buildRetrieveApiRequest(apiToken string, request RetrieveLogShippingTokensRequest) (*http.Request, error) {
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	req, err := http.NewRequest(retrieveLogShippingTokenServiceMethod, fmt.Sprintf(retrieveLogShippingTokenServiceUrl, baseUrl), bytes.NewBuffer(jsonBytes))

	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Returns the relevant shipping tokens, filtered, sorted and paginated as per the request, error otherwise
func (c *LogShippingTokensClient) RetrieveLogShippingTokens(retrieveRequest RetrieveLogShippingTokensRequest) (*RetrieveLogShippingTokensResponse, error) {
	err := validateRetrieveLogShippingTokensRequest(retrieveRequest)
	if err != nil {
		return nil, err
	}

	req, err := c.buildRetrieveApiRequest(c.ApiToken, retrieveRequest)
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

	if !logzio_client.CheckValidStatus(resp, []int{retrieveLogShippingTokenMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", operationRetrieveLogShippingTokens, resp.StatusCode, jsonBytes)
	}

	var target RetrieveLogShippingTokensResponse
	err = json.Unmarshal(jsonBytes, &target)
	if err != nil {
		return nil, err
	}

	return &target, nil
}
