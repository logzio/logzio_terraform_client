package tracing_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	createTracingAccountServiceUrl     = tracingAccountServiceEndpoint
	createTracingAccountServiceMethod  = http.MethodPost
	createTracingAccountMethodSuccess  = http.StatusOK
	createTracingAccountStatusNotFound = http.StatusNotFound
)

// CreateTracingAccount creates a consumption tracing account, returns the created account if
// successful, an error otherwise
func (c *TracingAccountClient) CreateTracingAccount(createTracingAccount CreateOrUpdateTracingAccount) (*TracingAccount, error) {
	err := validateCreateOrUpdateTracingAccount(createTracingAccount)
	if err != nil {
		return nil, err
	}

	tracingAccountJson, err := json.Marshal(createTracingAccount)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createTracingAccountServiceMethod,
		Url:          fmt.Sprintf(createTracingAccountServiceUrl, c.BaseUrl),
		Body:         tracingAccountJson,
		SuccessCodes: []int{createTracingAccountMethodSuccess},
		NotFoundCode: createTracingAccountStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationCreateTracingAccount,
		ResourceName: tracingAccountResourceName,
	})

	if err != nil {
		return nil, err
	}

	var tracingAccount TracingAccount
	err = json.Unmarshal(res, &tracingAccount)
	if err != nil {
		return nil, err
	}

	return &tracingAccount, nil
}

func validateCreateOrUpdateTracingAccount(tracingAccount CreateOrUpdateTracingAccount) error {
	if len(tracingAccount.Email) == 0 {
		return fmt.Errorf("email must be set")
	}
	if len(tracingAccount.AccountName) == 0 {
		return fmt.Errorf("accountName must be set")
	}
	if tracingAccount.AuthorizedAccountIds == nil {
		return fmt.Errorf("authorizedAccountIds must be initialized, even without any ids")
	}
	if tracingAccount.MaxDailyGB < 0 {
		return fmt.Errorf("maxDailyGB should be non-negative")
	}
	return nil
}
