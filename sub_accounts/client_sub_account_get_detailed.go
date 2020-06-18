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
	getDetailedServiceUrl     string = subAccountServiceEndpoint + "/detailed/%d"
	getDetailedServiceMethod  string = http.MethodGet
	getDetailedServiceSuccess int    = http.StatusOK
)

func (c *SubAccountClient) getDetailedValidateRequest(id int64) (error, bool) {
	return nil, true
}

func (c *SubAccountClient) getDetailedApiRequest(apiToken string, id int64) (*http.Request, error) {

	url := fmt.Sprintf(getDetailedServiceUrl, c.BaseUrl, id)
	req, err := http.NewRequest(getDetailedServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)
	return req, err
}

func (c *SubAccountClient) getDetailedHttpRequest(req *http.Request) (map[string]interface{}, error) {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{getDetailedServiceSuccess}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}
	var target map[string]interface{}
	err = json.Unmarshal(jsonBytes, &target)
	if err != nil {
		return nil, err
	}
	return target, nil
}

func (c *SubAccountClient) getDetailedCheckResponse(response map[string]interface{}) error {
	return nil
}

func (c *SubAccountClient) GetDetailedSubAccount(id int64) (*SubAccountDetailed, error) {
	if err, ok := c.getDetailedValidateRequest(id); !ok {
		return nil, err
	}
	req, _ := c.getDetailedApiRequest(c.ApiToken, id)

	target, err := c.getDetailedHttpRequest(req)
	if err != nil {
		return nil, err
	}

	err = c.getDetailedCheckResponse(target)
	if err != nil {
		return nil, err
	}

	subAccount, err := jsonToDetailedSubAccount(target)
	if err != nil {
		return nil, err
	}
	return subAccount, nil
}
