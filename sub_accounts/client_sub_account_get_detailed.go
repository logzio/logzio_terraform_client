package sub_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	getDetailedSubAccountServiceUrl      string = subAccountServiceEndpoint + "/detailed/%d"
	getDetailedSubAccountServiceMethod   string = http.MethodGet
	getDetailedSubAccountServiceSuccess  int    = http.StatusOK
	getDetailedSubAccountServiceNotFound int    = http.StatusNotFound
)

func (c *SubAccountClient) GetDetailedSubAccount(subAccountId int64) (*DetailedSubAccount, error) {
	req, err := c.buildGetDetailedApiRequest(c.ApiToken, subAccountId)
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
	if !logzio_client.CheckValidStatus(resp, []int{getDetailedSubAccountServiceSuccess}) {
		if resp.StatusCode == getDetailedSubAccountServiceNotFound {
			return nil, fmt.Errorf("API call %s failed with missing sub account %d, data: %s", operationGetDetailedSubAccount, subAccountId, jsonBytes)
		}

		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", operationGetDetailedSubAccount, resp.StatusCode, jsonBytes)
	}

	var subAccount DetailedSubAccount
	err = json.Unmarshal(jsonBytes, &subAccount)
	if err != nil {
		return nil, err
	}

	return &subAccount, nil
}

func (c *SubAccountClient) buildGetDetailedApiRequest(apiToken string, subAccountId int64) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(getDetailedSubAccountServiceMethod, fmt.Sprintf(getDetailedSubAccountServiceUrl, baseUrl, subAccountId), nil)
	if err != nil {
		return nil, err
	}
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}
