package tracing_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	updateTracingAccountSoftLimitServiceUrl      = tracingAccountServiceEndpoint + "/%d/soft-limit"
	updateTracingAccountSoftLimitServiceMethod   = http.MethodPut
	updateTracingAccountSoftLimitServiceSuccess  = http.StatusOK
	updateTracingAccountSoftLimitServiceNotFound = http.StatusNotFound
)

// UpdateTracingAccountSoftLimit sets the soft limit, in GB, of a consumption tracing account and
// returns the updated soft limit, an error otherwise.
// The owner account must be a consumption account, otherwise the API responds with 400.
func (c *TracingAccountClient) UpdateTracingAccountSoftLimit(tracingAccountId int64, updateSoftLimit UpdateTracingAccountSoftLimit) (*TracingAccountSoftLimit, error) {
	err := validateUpdateTracingAccountSoftLimit(tracingAccountId, updateSoftLimit)
	if err != nil {
		return nil, err
	}

	updateSoftLimitJson, err := json.Marshal(updateSoftLimit)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateTracingAccountSoftLimitServiceMethod,
		Url:          fmt.Sprintf(updateTracingAccountSoftLimitServiceUrl, c.BaseUrl, tracingAccountId),
		Body:         updateSoftLimitJson,
		SuccessCodes: []int{updateTracingAccountSoftLimitServiceSuccess},
		NotFoundCode: updateTracingAccountSoftLimitServiceNotFound,
		ResourceId:   tracingAccountId,
		ApiAction:    operationUpdateTracingAccountSoftLimit,
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

func validateUpdateTracingAccountSoftLimit(tracingAccountId int64, updateSoftLimit UpdateTracingAccountSoftLimit) error {
	// the API rejects a body id that does not match the path id, so catch it before sending
	if int64(updateSoftLimit.TracingAccountId) != tracingAccountId {
		return fmt.Errorf("tracingAccountId in the request must match the tracing account id in the path")
	}
	if updateSoftLimit.SoftLimitGB < 0 {
		return fmt.Errorf("softLimitGB should be non-negative")
	}
	return nil
}
