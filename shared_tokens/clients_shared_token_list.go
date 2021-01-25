package shared_tokens

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	listServiceUrl     string = sharedTokenServiceEndpoint
	listServiceMethod  string = http.MethodGet
	listServiceSuccess int    = http.StatusOK
)

func (c *SharedTokenClient) listApiRequest(apiToken string) (*http.Request, error) {
	url := fmt.Sprintf(listServiceUrl, c.BaseUrl)
	req, err := http.NewRequest(listServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)
	return req, err
}

func (c *SharedTokenClient) listHttpRequest(req *http.Request) ([]map[string]interface{}, error) {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{listServiceSuccess}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}
	var target []map[string]interface{}
	err = json.Unmarshal(jsonBytes, &target)
	if err != nil {
		return nil, err
	}
	return target, nil
}

func (c *SharedTokenClient) ListSharedTokens() ([]SharedToken, error) {

	req, _ := c.listApiRequest(c.ApiToken)
	target, err := c.listHttpRequest(req)
	if err != nil {
		return nil, err
	}

	var sharedTokens []SharedToken
	for _, json := range target {
		sharedToken := jsonToSharedToken(json)
		sharedTokens = append(sharedTokens, sharedToken)
	}
	return sharedTokens, nil
}