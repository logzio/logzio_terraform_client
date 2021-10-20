package sub_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getDetailedSubAccountServiceUrl      = subAccountServiceEndpoint + "/detailed/%d"
	getDetailedSubAccountServiceMethod   = http.MethodGet
	getDetailedSubAccountServiceSuccess  = http.StatusOK
	getDetailedSubAccountServiceNotFound = http.StatusNotFound
)

// GetDetailedSubAccount returns a detailed sub-account given its unique identifier, an error otherwise
func (c *SubAccountClient) GetDetailedSubAccount(subAccountId int64) (*DetailedSubAccount, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getDetailedSubAccountServiceMethod,
		Url:          fmt.Sprintf(getDetailedSubAccountServiceUrl, c.BaseUrl, subAccountId),
		Body:         nil,
		SuccessCodes: []int{getDetailedSubAccountServiceSuccess},
		NotFoundCode: getDetailedSubAccountServiceNotFound,
		ResourceId:   subAccountId,
		ApiAction:    operationGetDetailedSubAccount,
	})

	if err != nil {
		return nil, err
	}

	var subAccount DetailedSubAccount
	err = json.Unmarshal(res, &subAccount)
	if err != nil {
		return nil, err
	}

	return &subAccount, nil
}
