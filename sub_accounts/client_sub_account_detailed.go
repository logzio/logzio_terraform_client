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
	detailedServiceUrl     string = subAccountServiceEndpoint + "/detailed"
	detailedServiceMethod  string = http.MethodGet
	detailedServiceSuccess int    = http.StatusOK
)

func (c *SubAccountClient) detailedValidateRequest(id int64) (error, bool) {
	return nil, true
}

func (c *SubAccountClient) detailedApiRequest(apiToken string) (*http.Request, error) {
	url := fmt.Sprintf(detailedServiceUrl, c.BaseUrl)
	req, err := http.NewRequest(detailedServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)
	return req, err
}

func (c *SubAccountClient) detailedHttpRequest(req *http.Request) ([]map[string]interface{}, error) {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{detailedServiceSuccess}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}
	var target []map[string]interface{}
	err = json.Unmarshal(jsonBytes, &target)
	if err != nil {
		return nil, err
	}
	return target, nil
}

func (c *SubAccountClient) detailedCheckResponse(response []map[string]interface{}) error {
	return nil
}

func (c *SubAccountClient) DetailedSubAccounts() ([]SubAccount, error) {

	req, _ := c.detailedApiRequest(c.ApiToken)

	target, err := c.detailedHttpRequest(req)
	if err != nil {
		return nil, err
	}

	err = c.detailedCheckResponse(target)
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
