package tracing_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getTracingAccountServiceUrl      = tracingAccountServiceEndpoint + "/%d"
	getTracingAccountServiceMethod   = http.MethodGet
	getTracingAccountServiceSuccess  = http.StatusOK
	getTracingAccountServiceNotFound = http.StatusNotFound
)

// GetTracingAccount returns a consumption tracing account given its unique identifier, an error otherwise
func (c *TracingAccountClient) GetTracingAccount(tracingAccountId int64) (*TracingAccount, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getTracingAccountServiceMethod,
		Url:          fmt.Sprintf(getTracingAccountServiceUrl, c.BaseUrl, tracingAccountId),
		Body:         nil,
		SuccessCodes: []int{getTracingAccountServiceSuccess},
		NotFoundCode: getTracingAccountServiceNotFound,
		ResourceId:   tracingAccountId,
		ApiAction:    operationGetTracingAccount,
		ResourceName: tracingAccountResourceName,
	})

	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("failed with missing tracing account")
	}

	var tracingAccount TracingAccount
	err = json.Unmarshal(res, &tracingAccount)
	if err != nil {
		return nil, err
	}

	return &tracingAccount, nil
}
