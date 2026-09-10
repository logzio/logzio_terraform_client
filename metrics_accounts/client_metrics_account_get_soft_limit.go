package metrics_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getMetricsAccountSoftLimitServiceUrl      = metricsAccountServiceEndpoint + "/%d/soft-limit"
	getMetricsAccountSoftLimitServiceMethod   = http.MethodGet
	getMetricsAccountSoftLimitServiceSuccess  = http.StatusOK
	getMetricsAccountSoftLimitServiceNotFound = http.StatusNotFound
)

// GetMetricsAccountSoftLimit returns the soft limit, in unique time series, of a consumption
// metrics account given its unique identifier, an error otherwise.
// The owner account must be a consumption account, otherwise the API responds with 400.
func (c *MetricsAccountClient) GetMetricsAccountSoftLimit(metricsAccountId int64) (*MetricsAccountSoftLimit, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getMetricsAccountSoftLimitServiceMethod,
		Url:          fmt.Sprintf(getMetricsAccountSoftLimitServiceUrl, c.BaseUrl, metricsAccountId),
		Body:         nil,
		SuccessCodes: []int{getMetricsAccountSoftLimitServiceSuccess},
		NotFoundCode: getMetricsAccountSoftLimitServiceNotFound,
		ResourceId:   metricsAccountId,
		ApiAction:    operationGetMetricsAccountSoftLimit,
		ResourceName: metricsAccountResourceName,
	})

	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("failed with missing metrics account soft limit")
	}

	var softLimit MetricsAccountSoftLimit
	err = json.Unmarshal(res, &softLimit)
	if err != nil {
		return nil, err
	}

	return &softLimit, nil
}
