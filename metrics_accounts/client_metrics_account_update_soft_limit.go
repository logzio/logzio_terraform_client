package metrics_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	updateMetricsAccountSoftLimitServiceUrl      = metricsAccountServiceEndpoint + "/%d/soft-limit"
	updateMetricsAccountSoftLimitServiceMethod   = http.MethodPut
	updateMetricsAccountSoftLimitServiceSuccess  = http.StatusOK
	updateMetricsAccountSoftLimitServiceNotFound = http.StatusNotFound
)

// UpdateMetricsAccountSoftLimit sets the soft limit, in unique time series, of a consumption
// metrics account and returns the updated soft limit, an error otherwise.
// The owner account must be a consumption account, otherwise the API responds with 400.
func (c *MetricsAccountClient) UpdateMetricsAccountSoftLimit(metricsAccountId int64, updateSoftLimit UpdateMetricsAccountSoftLimit) (*MetricsAccountSoftLimit, error) {
	err := validateUpdateMetricsAccountSoftLimit(updateSoftLimit)
	if err != nil {
		return nil, err
	}

	updateSoftLimitJson, err := json.Marshal(updateSoftLimit)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateMetricsAccountSoftLimitServiceMethod,
		Url:          fmt.Sprintf(updateMetricsAccountSoftLimitServiceUrl, c.BaseUrl, metricsAccountId),
		Body:         updateSoftLimitJson,
		SuccessCodes: []int{updateMetricsAccountSoftLimitServiceSuccess},
		NotFoundCode: updateMetricsAccountSoftLimitServiceNotFound,
		ResourceId:   metricsAccountId,
		ApiAction:    operationUpdateMetricsAccountSoftLimit,
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

func validateUpdateMetricsAccountSoftLimit(updateSoftLimit UpdateMetricsAccountSoftLimit) error {
	if updateSoftLimit.SoftLimitUniqueMetrics < 0 {
		return fmt.Errorf("softLimitUniqueMetrics should be non-negative")
	}
	return nil
}
