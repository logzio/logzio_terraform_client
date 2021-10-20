package sub_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	listDetailedSubAccountsServiceUrl     = subAccountServiceEndpoint + "/detailed"
	listDetailedSubAccountsServiceMethod  = http.MethodGet
	listDetailedSubAccountsServiceSuccess = http.StatusOK
	listDetailedSubAccountStatusNotFound  = http.StatusNotFound
)

// ListDetailedSubAccounts returns all the detailed sub-accounts in an array associated with the account identified by the supplied API token, returns an error if
// any problem occurs during the API call
func (c *SubAccountClient) ListDetailedSubAccounts() ([]DetailedSubAccount, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listDetailedSubAccountsServiceMethod,
		Url:          fmt.Sprintf(listDetailedSubAccountsServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{listDetailedSubAccountsServiceSuccess},
		NotFoundCode: listDetailedSubAccountStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationListSubAccounts,
	})

	if err != nil {
		return nil, err
	}

	var subAccounts []DetailedSubAccount
	err = json.Unmarshal(res, &subAccounts)
	if err != nil {
		return nil, err
	}

	return subAccounts, nil
}
