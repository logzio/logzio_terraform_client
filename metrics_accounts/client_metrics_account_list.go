package metrics_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	listMetricsAccountServiceUrl     = metricsAccountServiceEndpoint
	listMetricsAccountServiceMethod  = http.MethodGet
	listMetricsAccountServiceSuccess = http.StatusOK
	listMetricsAccountStatusNotFound = http.StatusNotFound
)

// ListMetricsAccounts returns all the metrics accounts in an array associated with the account identified by the supplied API token, returns an error if
// any problem occurs during the API call
func (c *MetricsAccountClient) ListMetricsAccounts() ([]MetricsAccount, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listMetricsAccountServiceMethod,
		Url:          fmt.Sprintf(listMetricsAccountServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{listMetricsAccountServiceSuccess},
		NotFoundCode: listMetricsAccountStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationListMetricsAccounts,
		ResourceName: metricsAccountResourceName,
	})

	if err != nil {
		return nil, err
	}

	var subAccounts []MetricsAccount
	err = json.Unmarshal(res, &subAccounts)
	if err != nil {
		return nil, err
	}

	return subAccounts, nil
}
