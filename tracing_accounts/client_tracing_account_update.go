package tracing_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	updateTracingAccountServiceUrl      = tracingAccountServiceEndpoint + "/%d"
	updateTracingAccountServiceMethod   = http.MethodPut
	updateTracingAccountServiceSuccess  = http.StatusOK
	updateTracingAccountServiceNotFound = http.StatusNotFound
)

// UpdateTracingAccount updates a consumption tracing account, returns the updated account if
// successful, an error otherwise
func (c *TracingAccountClient) UpdateTracingAccount(tracingAccountId int64, updateTracingAccount CreateOrUpdateTracingAccount) (*TracingAccount, error) {
	err := validateCreateOrUpdateTracingAccount(updateTracingAccount)
	if err != nil {
		return nil, err
	}

	tracingAccountJson, err := json.Marshal(updateTracingAccount)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateTracingAccountServiceMethod,
		Url:          fmt.Sprintf(updateTracingAccountServiceUrl, c.BaseUrl, tracingAccountId),
		Body:         tracingAccountJson,
		SuccessCodes: []int{updateTracingAccountServiceSuccess},
		NotFoundCode: updateTracingAccountServiceNotFound,
		ResourceId:   tracingAccountId,
		ApiAction:    operationUpdateTracingAccount,
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
