package metrics_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getMetricsAccountServiceUrl      = metricsAccountServiceEndpoint + "/%d"
	getMetricsAccountServiceMethod   = http.MethodGet
	getMetricsAccountServiceSuccess  = http.StatusOK
	getMetricsAccountServiceNotFound = http.StatusNotFound
)

// GetMetricsAccount returns a metrics account given its unique identifier, an error otherwise
func (c *MetricsAccountClient) GetMetricsAccount(metricsAccountId int64) (*MetricsAccount, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getMetricsAccountServiceMethod,
		Url:          fmt.Sprintf(getMetricsAccountServiceUrl, c.BaseUrl, metricsAccountId),
		Body:         nil,
		SuccessCodes: []int{getMetricsAccountServiceSuccess},
		NotFoundCode: getMetricsAccountServiceNotFound,
		ResourceId:   metricsAccountId,
		ApiAction:    operationGetMetricsAccount,
		ResourceName: metricsAccountResourceName,
	})

	if err != nil {
		return nil, err
	}

	var subAccount MetricsAccount
	err = json.Unmarshal(res, &subAccount)
	if err != nil {
		return nil, err
	}

	return &subAccount, nil
}
