# Log Shipping Tokens

Compatible with Logz.io's [Manage Log Shipping Tokens API](https://api-docs.logz.io/docs/logz/manage-log-shipping-tokens).

To create a new log shipping token:

```go
client, _ := log_shipping_tokens.New(apiToken, apiServerAddress)
token, err := client.CreateLogShippingToken(log_shipping_tokens.CreateLogShippingToken{
                Name: "client_integration_test",
            })
```

|function|func name|
|---|---|
|create log shipping token| `func (c *LogShippingTokensClient) CreateLogShippingToken(token CreateLogShippingToken) (*LogShippingToken, error)` |
|delete log shipping token| `func (c *LogShippingTokensClient) DeleteLogShippingToken(tokenId int32) error` |
|get log shipping token| `func (c *LogShippingTokensClient) GetLogShippingToken(tokenId int32) (*LogShippingToken, error)` |
|get available number of tokens| `func (c *LogShippingTokensClient) GetLogShippingLimitsToken() (*LogShippingTokensLimits, error)` |
|retrieve tokens| `func (c *LogShippingTokensClient) RetrieveLogShippingTokens(retrieveRequest RetrieveLogShippingTokensRequest) (*RetrieveLogShippingTokensResponse, error)` |
|update log shipping token| `func (c *LogShippingTokensClient) UpdateLogShippingToken(tokenId int32, token UpdateLogShippingToken) (*LogShippingToken, error)` |
