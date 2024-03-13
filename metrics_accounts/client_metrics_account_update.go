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
	updateMetricsAccountServiceSuccess  = 200
	updateMetricsAccountServiceNotFound = http.StatusNotFound
)

func (c *MetricsAccountClient) UpdateMetricsAccount(metricsAccountId int64, updateMetricsAccount CreateOrUpdateMetricsAccount) error {
	err := validateUpdateSubAccount(updateMetricsAccount)
	if err != nil {
		return err
	}

	updateSubAccountJson, err := json.Marshal(updateMetricsAccount)
	if err != nil {
		return err
	}

	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateMetricsAccountServiceMethod,
		Url:          fmt.Sprintf(updateMetricsAccountServiceUrl, c.BaseUrl, metricsAccountId),
		Body:         updateSubAccountJson,
		SuccessCodes: []int{updateMetricsAccountServiceSuccess},
		NotFoundCode: updateMetricsAccountServiceNotFound,
		ResourceId:   metricsAccountId,
		ApiAction:    operationUpdateMetricsAccount,
		ResourceName: metricsAccountResourceName,
	})

	return err
}

func validateUpdateSubAccount(updateSubAccount CreateOrUpdateMetricsAccount) error {
	if len(updateSubAccount.AccountName) == 0 {
		return fmt.Errorf("account name must be set")
	}

	if updateSubAccount.PlanUts < 0 {
		return fmt.Errorf("PlanUts should be >=100 or empty")
	}

	return nil
}
