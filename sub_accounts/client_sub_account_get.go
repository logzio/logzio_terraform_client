package sub_accounts

import (
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
)

const (
	getServiceUrl     string = subAccountServiceEndpoint + "/%d"
	getServiceMethod  string = http.MethodGet
	getServiceSuccess int    = http.StatusOK
)

func (c *SubAccountClient) getValidateRequest(id int64) (error, bool) {
	return nil, true
}

func (c *SubAccountClient) getApiRequest(apiToken string, id int64) (*http.Request, error) {

	url := fmt.Sprintf(getServiceUrl, c.BaseUrl, id)
	req, err := http.NewRequest(getServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)
	return req, err
}

func (c *SubAccountClient) getHttpRequest(req *http.Request) (map[string]interface{}, error) {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{getServiceSuccess}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}
	var target map[string]interface{}
	err = json.Unmarshal(jsonBytes, &target)
	if err != nil {
		return nil, err
	}
	return target, nil
}

func (c *SubAccountClient) getCheckResponse(response map[string]interface{}) error {
	return nil
}

func (c *SubAccountClient) GetSubAccount(id int64) (*SubAccount, error) {
	if err, ok := c.getValidateRequest(id); !ok {
		return nil, err
	}
	req, _ := c.getApiRequest(c.ApiToken, id)

	target, err := c.getHttpRequest(req)
	if err != nil {
		return nil, err
	}

	err = c.getCheckResponse(target)
	if err != nil {
		return nil, err
	}

	subAccount := jsonToSubAccount(target)
	return &subAccount, nil
}
