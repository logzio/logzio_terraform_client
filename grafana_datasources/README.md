# Grafana Datasource

To get Grafana Datasource for specific account:

```go
	client, err := grafana_datasources.New(apiToken, test_utils.GetLogzIoBaseUrl())
    datasource, err := client.GetForAccount("my-metrics-account-name")
```

| function                   | func name                                                                                         |
|----------------------------|---------------------------------------------------------------------------------------------------|
| get datasource for account | `func (c *GrafanaDatasourceClient) GetForAccount(accountName string) (*GrafanaDataSource, error)` |