package tracing_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getTracingAccountSoftLimitServiceUrl      = tracingAccountServiceEndpoint + "/%d/soft-limit"
	getTracingAccountSoftLimitServiceMethod   = http.MethodGet
	getTracingAccountSoftLimitServiceSuccess  = http.StatusOK
	getTracingAccountSoftLimitServiceNotFound = http.StatusNotFound
)

// GetTracingAccountSoftLimit returns the soft limit, in GB, of a consumption tracing account given
// its unique identifier, an error otherwise.
// The owner account must be a consumption account, otherwise the API responds with 400.
func (c *TracingAccountClient) GetTracingAccountSoftLimit(tracingAccountId int64) (*TracingAccountSoftLimit, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getTracingAccountSoftLimitServiceMethod,
		Url:          fmt.Sprintf(getTracingAccountSoftLimitServiceUrl, c.BaseUrl, tracingAccountId),
		Body:         nil,
		SuccessCodes: []int{getTracingAccountSoftLimitServiceSuccess},
		NotFoundCode: getTracingAccountSoftLimitServiceNotFound,
		ResourceId:   tracingAccountId,
		ApiAction:    operationGetTracingAccountSoftLimit,
		ResourceName: tracingAccountResourceName,
	})

	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("failed with missing tracing account soft limit")
	}

	var softLimit TracingAccountSoftLimit
	err = json.Unmarshal(res, &softLimit)
	if err != nil {
		return nil, err
	}

	return &softLimit, nil
}
