package tracing_accounts

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteTracingAccountServiceUrl      = tracingAccountServiceEndpoint + "/%d"
	deleteTracingAccountServiceMethod   = http.MethodDelete
	deleteTracingAccountServiceSuccess  = http.StatusOK
	deleteTracingAccountServiceNotFound = http.StatusNotFound
)

// DeleteTracingAccount deletes a consumption tracing account given its unique identifier,
// an error otherwise
func (c *TracingAccountClient) DeleteTracingAccount(tracingAccountId int64) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteTracingAccountServiceMethod,
		Url:          fmt.Sprintf(deleteTracingAccountServiceUrl, c.BaseUrl, tracingAccountId),
		Body:         nil,
		SuccessCodes: []int{deleteTracingAccountServiceSuccess},
		NotFoundCode: deleteTracingAccountServiceNotFound,
		ResourceId:   tracingAccountId,
		ApiAction:    operationDeleteTracingAccount,
		ResourceName: tracingAccountResourceName,
	})

	return err
}
