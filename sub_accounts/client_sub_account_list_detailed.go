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
	listDetailedSubAccountsServiceUrl     string = subAccountServiceEndpoint + "/detailed"
	listDetailedSubAccountsServiceMethod  string = http.MethodGet
	listDetailedSubAccountsServiceSuccess int    = http.StatusOK
)

// Returns all the detailed sub accounts in an array associated with the account identified by the supplied API token, returns an error if
// any problem occurs during the API call
func (c *SubAccountClient) ListDetailedSubAccounts() ([]DetailedSubAccount, error) {
	req, _ := c.buildListDetailedApiRequest(c.ApiToken)
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{listDetailedSubAccountsServiceSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", operationListDetailedSubAccounts, resp.StatusCode, jsonBytes)
	}

	var subAccounts []DetailedSubAccount
	err = json.Unmarshal(jsonBytes, &subAccounts)
	if err != nil {
		return nil, err
	}

	return subAccounts, nil
}

func (c *SubAccountClient) buildListDetailedApiRequest(apiToken string) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(listDetailedSubAccountsServiceMethod, fmt.Sprintf(listDetailedSubAccountsServiceUrl, baseUrl), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}
