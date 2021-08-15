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
	listSubAccountServiceUrl     string = subAccountServiceEndpoint
	listSubAccountServiceMethod  string = http.MethodGet
	listSubAccountServiceSuccess int    = http.StatusOK
)

// Returns all the sub accounts in an array associated with the account identified by the supplied API token, returns an error if
// any problem occurs during the API call
func (c *SubAccountClient) ListSubAccounts() ([]SubAccount, error) {
	req, err := c.buildListApiRequest(c.ApiToken)
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

	if !logzio_client.CheckValidStatus(resp, []int{listSubAccountServiceSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", operationListSubAccounts, resp.StatusCode, jsonBytes)
	}

	var subAccounts []SubAccount
	err = json.Unmarshal(jsonBytes, &subAccounts)
	if err != nil {
		return nil, err
	}

	return subAccounts, nil
}

func (c *SubAccountClient) buildListApiRequest(apiToken string) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(listSubAccountServiceMethod, fmt.Sprintf(listSubAccountServiceUrl, baseUrl), nil)
	if err != nil {
		return nil, err
	}
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}
