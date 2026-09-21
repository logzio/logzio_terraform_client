# Metrics Accounts

Compatible with Logz.io's [account management API](https://api-docs.logz.io/docs/logz/create-a-new-metrics-account).

To create a new metrics account, on a main account.
```go
client, _ := metrics_accounts.New(apiToken, apiServerAddress)
planUts := new(int32)
*planUts = 1000
account := metrics_accounts.CreateOrUpdateMetricsAccount{
                Email:                  "some@email.test",
                AccountName:            "tf_client_test",
                PlanUts:                planUts,
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
| get soft limit         | `func (c *MetricsAccountClient) GetMetricsAccountSoftLimit(metricsAccountId int64) (*MetricsAccountSoftLimit, error)`                           |
| update soft limit      | `func (c *MetricsAccountClient) UpdateMetricsAccountSoftLimit(metricsAccountId int64, updateSoftLimit UpdateMetricsAccountSoftLimit) (*MetricsAccountSoftLimit, error)` |

The soft limit endpoints are available for **consumption** accounts only. For a subscription
owner the API responds with `400`.
