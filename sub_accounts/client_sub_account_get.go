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
	getSubAccountServiceUrl      string = subAccountServiceEndpoint + "/%d"
	getSubAccountServiceMethod   string = http.MethodGet
	getSubAccountServiceSuccess  int    = http.StatusOK
	getSubAccountServiceNotFound int    = http.StatusNotFound
)

// Returns a sub account given it's unique identifier, an error otherwise
func (c *SubAccountClient) GetSubAccount(subAccountId int64) (*SubAccount, error) {
	req, _ := c.buildGetApiRequest(c.ApiToken, subAccountId)
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{getSubAccountServiceSuccess}) {
		if resp.StatusCode == getSubAccountServiceNotFound {
			return nil, fmt.Errorf("API call %s failed with missing sub account %d, data: %s", operationGetSubAccount, subAccountId, jsonBytes)
		}

		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", operationGetSubAccount, resp.StatusCode, jsonBytes)
	}

	var subAccount SubAccount
	err = json.Unmarshal(jsonBytes, &subAccount)
	if err != nil {
		return nil, err
	}

	return &subAccount, nil
}

func (c *SubAccountClient) buildGetApiRequest(apiToken string, subAccountId int64) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(getSubAccountServiceMethod, fmt.Sprintf(getSubAccountServiceUrl, baseUrl, subAccountId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}
