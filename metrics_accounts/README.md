# Metrics sub-accounts

Compatible with Logz.io's [account management API](https://docs.logz.io/api/#tag/Manage-metrics-account).

To create a new metrics account, on a main account.
```go
client, _ := metrics_accounts.New(apiToken, apiServerAddress)
account := metrics_accounts.CreateOrUpdateMetricsAccount{
                Email:                  "some@email.test",
                AccountName:            "tf_client_test",
                PlanUts:                1000,
                AuthorizedAccountsIds: []int32{},
            }
```

| function               | func name                                                                                                                                       |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| create metrics account | `func (c *MetricsAccountClient) CreateMetricsAccount(createMetricsAccount CreateOrUpdateMetricsAccount) (*MetricsAccountCreateResponse, error)` |
| update metrics account | `func (c *MetricsAccountClient) UpdateMetricsAccount(metricsAccountId int64, updateMetricsAccount CreateOrUpdateMetricsAccount) error`          |
| delete metrics account | `func (c *MetricsAccountClient) DeleteMetricsAccount(metricsAccountId int64) error`                                                             |
| get metrics account    | `func (c *MetricsAccountClient) GetMetricsAccount(metricsAccountId int64) (*MetricsAccount, error)`                                             |
| list metrics accounts  | `func (c *MetricsAccountClient) ListMetricsAccounts() ([]MetricsAccount, error)`                                                                |
