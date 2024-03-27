package metrics_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	createMetricsAccountServiceUrl     = metricsAccountServiceEndpoint
	createMetricsAccountServiceMethod  = http.MethodPost
	createMetricsAccountMethodSuccess  = http.StatusOK
	createMetricsAccountStatusNotFound = http.StatusNotFound
)

type FieldError struct {
	Field   string
	Message string
}

func (e FieldError) Error() string {
	return fmt.Sprintf("%v: %v", e.Field, e.Message)
}

// CreateMetricsAccount creates metrics account, return account's id & token if successful, an error otherwise
func (c *MetricsAccountClient) CreateMetricsAccount(createMetricsAccount CreateOrUpdateMetricsAccount) (*MetricsAccountCreateResponse, error) {
	err := validateCreateMetricsAccount(createMetricsAccount)
	if err != nil {
		return nil, err
	}

	MetricsAccountJson, err := json.Marshal(createMetricsAccount)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createMetricsAccountServiceMethod,
		Url:          fmt.Sprintf(createMetricsAccountServiceUrl, c.BaseUrl),
		Body:         MetricsAccountJson,
		SuccessCodes: []int{createMetricsAccountMethodSuccess},
		NotFoundCode: createMetricsAccountStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationCreateMetricsAccount,
		ResourceName: metricsAccountResourceName,
	})

	if err != nil {
		return nil, err
	}

	var reVal MetricsAccountCreateResponse
	err = json.Unmarshal(res, &reVal)
	if err != nil {
		return nil, err
	}

	return &reVal, nil
}

func validateCreateMetricsAccount(createMetricsAccount CreateOrUpdateMetricsAccount) error {
	if len(createMetricsAccount.Email) == 0 {
		return fmt.Errorf("email must be set")
	}

	if createMetricsAccount.AuthorizedAccountsIds == nil {
		return fmt.Errorf("AuthorizedAccountsIds must be initialized, even without any ids")
	}

	if *createMetricsAccount.PlanUts < 100 {
		return fmt.Errorf("PlanUts should be larger than 100")
	}
	return nil
}
