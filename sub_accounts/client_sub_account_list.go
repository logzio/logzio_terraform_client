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
	listServiceUrl     string = subAccountServiceEndpoint
	listServiceMethod  string = http.MethodGet
	listServiceSuccess int    = http.StatusOK
)

func (c *SubAccountClient) listValidateRequest(id int64) (error, bool) {
	return nil, true
}

func (c *SubAccountClient) listApiRequest(apiToken string) (*http.Request, error) {
	url := fmt.Sprintf(listServiceUrl, c.BaseUrl)
	req, err := http.NewRequest(listServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)
	return req, err
}

func (c *SubAccountClient) listHttpRequest(req *http.Request) ([]map[string]interface{}, error) {
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

func (c *SubAccountClient) listCheckResponse(response []map[string]interface{}) error {
	return nil
}

func (c *SubAccountClient) ListSubAccounts() ([]SubAccount, error) {

	req, _ := c.listApiRequest(c.ApiToken)

	target, err := c.listHttpRequest(req)
	if err != nil {
		return nil, err
	}

	err = c.listCheckResponse(target)
	if err != nil {
		return nil, err
	}

	var subAccounts []SubAccount
	for _, json := range target {
		subAccount := jsonToSubAccount(json)
		subAccounts = append(subAccounts, subAccount)
	}

	return subAccounts, nil
}
