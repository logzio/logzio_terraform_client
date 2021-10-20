package sub_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getSubAccountServiceUrl      = subAccountServiceEndpoint + "/%d"
	getSubAccountServiceMethod   = http.MethodGet
	getSubAccountServiceSuccess  = http.StatusOK
	getSubAccountServiceNotFound = http.StatusNotFound
)

// GetSubAccount returns a sub account given its unique identifier, an error otherwise
func (c *SubAccountClient) GetSubAccount(subAccountId int64) (*SubAccount, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getSubAccountServiceMethod,
		Url:          fmt.Sprintf(getSubAccountServiceUrl, c.BaseUrl, subAccountId),
		Body:         nil,
		SuccessCodes: []int{getSubAccountServiceSuccess},
		NotFoundCode: getSubAccountServiceNotFound,
		ResourceId:   subAccountId,
		ApiAction:    operationGetSubAccount,
	})

	if err != nil {
		return nil, err
	}

	var subAccount SubAccount
	err = json.Unmarshal(res, &subAccount)
	if err != nil {
		return nil, err
	}

	return &subAccount, nil
}
