package sub_accounts

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteSubAccountServiceUrl     = subAccountServiceEndpoint + "/%d"
	deleteSubAccountServiceMethod  = http.MethodDelete
	deleteSubAccountServiceSuccess = http.StatusNoContent
	deleteSubAccountMethodNotFound = http.StatusNotFound
)

// DeleteSubAccount deletes a sub account specified by its unique id, returns an error if a problem is encountered
func (c *SubAccountClient) DeleteSubAccount(subAccountId int64) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteSubAccountServiceMethod,
		Url:          fmt.Sprintf(deleteSubAccountServiceUrl, c.BaseUrl, subAccountId),
		Body:         nil,
		SuccessCodes: []int{deleteSubAccountServiceSuccess},
		NotFoundCode: deleteSubAccountMethodNotFound,
		ResourceId:   subAccountId,
		ApiAction:    operationDeleteSubAccount,
	})

	return err
}
