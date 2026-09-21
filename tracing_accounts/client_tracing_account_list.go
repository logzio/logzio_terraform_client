package tracing_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	listTracingAccountServiceUrl     = tracingAccountServiceEndpoint
	listTracingAccountServiceMethod  = http.MethodGet
	listTracingAccountServiceSuccess = http.StatusOK
	listTracingAccountStatusNotFound = http.StatusNotFound
)

// ListTracingAccounts returns all the consumption tracing accounts associated with the account
// identified by the supplied API token, an error otherwise
func (c *TracingAccountClient) ListTracingAccounts() ([]TracingAccount, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listTracingAccountServiceMethod,
		Url:          fmt.Sprintf(listTracingAccountServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{listTracingAccountServiceSuccess},
		NotFoundCode: listTracingAccountStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationListTracingAccounts,
		ResourceName: tracingAccountResourceName,
	})

	if err != nil {
		return nil, err
	}

	var tracingAccounts []TracingAccount
	err = json.Unmarshal(res, &tracingAccounts)
	if err != nil {
		return nil, err
	}

	return tracingAccounts, nil
}
