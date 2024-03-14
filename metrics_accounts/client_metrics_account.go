package metrics_accounts

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	metricsAccountServiceEndpoint = "%s/v1/account-management/metrics-accounts"
	loggerName                    = "logzio-client"
	operationGetMetricsAccount    = "GetMetricsAccount"
	operationDeleteMetricsAccount = "DeleteMetricsAccount"
	operationListMetricsAccounts  = "ListMetricsAccounts"
	operationUpdateMetricsAccount = "UpdateMetricsAccount"
	operationCreateMetricsAccount = "CreateMetricsAccount"

	metricsAccountResourceName = "metrics account"
)

type MetricsAccountClient struct {
	*client.Client
	logger hclog.Logger
}

type CreateOrUpdateMetricsAccount struct {
	Email                 string  `json:"email,omitempty"`
	AccountName           string  `json:"accountName,omitempty"`
	PlanUts               *int32  `json:"planUts,omitempty"`
	AuthorizedAccountsIds []int32 `json:"authorizedAccountsIds,omitempty"`
}

type MetricsAccount struct {
	Id                    int32   `json:"Id"`
	AccountName           string  `json:"accountName"`
	AccountToken          string  `json:"token"`
	CreatedAt             int64   `json:"createdAt"`
	PlanUts               int32   `json:"planUts"`
	AuthorizedAccountsIds []int32 `json:"authorizedAccountsIds"`
}

type MetricsAccountCreateResponse struct {
	Id                    int32   `json:"Id"`
	AccountName           string  `json:"accountName"`
	Token                 string  `json:"token"`
	CreatedAt             int64   `json:"createdAt"`
	PlanUts               int32   `json:"planUts"`
	AuthorizedAccountsIds []int32 `json:"authorizedAccountsIds"`
}

// Creates a new entry point into the metrics-account functions, accepts the user's logz.io API token and account Id
func New(apiToken string, baseUrl string) (*MetricsAccountClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	c := &MetricsAccountClient{
		Client: client.New(apiToken, baseUrl),
		logger: hclog.New(&hclog.LoggerOptions{
			Level:      hclog.Debug,
			Name:       loggerName,
			JSONFormat: true,
		}),
	}
	return c, nil
}
