package sub_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	listSubAccountServiceUrl     = subAccountServiceEndpoint
	listSubAccountServiceMethod  = http.MethodGet
	listSubAccountServiceSuccess = http.StatusOK
	listSubAccountStatusNotFound = http.StatusNotFound
)

// ListSubAccounts returns all the sub-accounts in an array associated with the account identified by the supplied API token, returns an error if
// any problem occurs during the API call
func (c *SubAccountClient) ListSubAccounts() ([]SubAccount, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listSubAccountServiceMethod,
		Url:          fmt.Sprintf(listSubAccountServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{listSubAccountServiceSuccess},
		NotFoundCode: listSubAccountStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationListSubAccounts,
		ResourceName: subAccountResourceName,
	})

	if err != nil {
		return nil, err
	}

	var subAccounts []SubAccount
	err = json.Unmarshal(res, &subAccounts)
	if err != nil {
		return nil, err
	}

	return subAccounts, nil
}
