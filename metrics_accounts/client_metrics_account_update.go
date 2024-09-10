package metrics_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	updateMetricsAccountServiceUrl      = metricsAccountServiceEndpoint + "/%d"
	updateMetricsAccountServiceMethod   = http.MethodPut
	updateMetricsAccountServiceSuccess  = http.StatusOK
	updateMetricsAccountServiceNotFound = http.StatusNotFound
)

func (c *MetricsAccountClient) UpdateMetricsAccount(metricsAccountId int64, updateMetricsAccount CreateOrUpdateMetricsAccount) error {
	err := validateUpdateMetricsAccount(updateMetricsAccount)
	if err != nil {
		return err
	}
	currentAccount, err := c.GetMetricsAccount(metricsAccountId)
	if err != nil {
		return err
	}
	if currentAccount.AccountName == updateMetricsAccount.AccountName {
		updateMetricsAccount.AccountName = ""
	}
	updateMetricsAccountJson, err := json.Marshal(updateMetricsAccount)
	if err != nil {
		return err
	}

	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateMetricsAccountServiceMethod,
		Url:          fmt.Sprintf(updateMetricsAccountServiceUrl, c.BaseUrl, metricsAccountId),
		Body:         updateMetricsAccountJson,
		SuccessCodes: []int{updateMetricsAccountServiceSuccess},
		NotFoundCode: updateMetricsAccountServiceNotFound,
		ResourceId:   metricsAccountId,
		ApiAction:    operationUpdateMetricsAccount,
		ResourceName: metricsAccountResourceName,
	})

	return err
}

func validateUpdateMetricsAccount(updateMetricsAccount CreateOrUpdateMetricsAccount) error {
	if updateMetricsAccount.PlanUts != nil && *updateMetricsAccount.PlanUts < 100 {
		return fmt.Errorf("PlanUts should be larger than 100 or empty")
	}
	return nil
}
