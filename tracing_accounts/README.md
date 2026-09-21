# Tracing Accounts

Compatible with Logz.io's consumption tracing account management API.

These endpoints are available for **consumption** accounts only. For a subscription owner the API
responds with `400`. The API is additionally gated per account by a feature flag; when it is off the
API responds with `400 FORBIDDEN`.

To create a new consumption tracing account, on a main account.
```go
client, _ := tracing_accounts.New(apiToken, apiServerAddress)
account := tracing_accounts.CreateOrUpdateTracingAccount{
                AccountName:          "tf_client_test_tracing",
                AuthorizedAccountIds: []int32{},
                MaxDailyGB:           5,
                Email:                "some@email.test",
            }
```

| function               | func name                                                                                                                                                        |
|------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| create tracing account | `func (c *TracingAccountClient) CreateTracingAccount(createTracingAccount CreateOrUpdateTracingAccount) (*TracingAccount, error)`                                |
| update tracing account | `func (c *TracingAccountClient) UpdateTracingAccount(tracingAccountId int64, updateTracingAccount CreateOrUpdateTracingAccount) (*TracingAccount, error)`        |
| delete tracing account | `func (c *TracingAccountClient) DeleteTracingAccount(tracingAccountId int64) error`                                                                              |
| get tracing account    | `func (c *TracingAccountClient) GetTracingAccount(tracingAccountId int64) (*TracingAccount, error)`                                                              |
| list tracing accounts  | `func (c *TracingAccountClient) ListTracingAccounts() ([]TracingAccount, error)`                                                                                 |
| get soft limit         | `func (c *TracingAccountClient) GetTracingAccountSoftLimit(tracingAccountId int64) (*TracingAccountSoftLimit, error)`                                            |
| update soft limit      | `func (c *TracingAccountClient) UpdateTracingAccountSoftLimit(tracingAccountId int64, updateSoftLimit UpdateTracingAccountSoftLimit) (*TracingAccountSoftLimit, error)` |

`UpdateTracingAccountSoftLimit.TracingAccountId` must match the id in the path; the API rejects a
mismatch with `400`.
