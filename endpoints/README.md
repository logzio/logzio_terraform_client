# Endpoints

Compatible with Logz.io's [endpoints API](https://docs.logz.io/api/#tag/Manage-notification-endpoints).

For each type of endpoint there is a different structure, below you can find an example for creating a Slack endpoint.
For more info, see: https://docs.logz.io/api/#tag/Manage-notification-endpoints or check our endpoints tests for more examples.

```go
client, _ := endpoints.New(apiToken, apiServerAddress)
endpoint, err := client.CreateEndpoint(endpoints.CreateOrUpdateEndpoint{
                Title:         "New endpoint",
                Description:   "this is a description",
                Type:          "slack",
                Url:           "https://jsonplaceholder.typicode.com/todos/1",
            })
```

|function|func name|
|---|---|
|create endpoint| `func (c *EndpointsClient) CreateEndpoint(endpoint CreateOrUpdateEndpoint) (*CreateOrUpdateEndpointResponse, error)` |
|delete endpoint| `func (c *EndpointsClient) DeleteEndpoint(endpointId int64) error` |
|get endpoint| `func (c *EndpointsClient) GetEndpoint(endpointId int64) (*Endpoint, error)` |
|list endpoints| `func (c *EndpointsClient) ListEndpoints() ([]Endpoint, error)` |
|update endpoint| `func (c *EndpointsClient) UpdateEndpoint(id int64, endpoint CreateOrUpdateEndpoint) (*CreateOrUpdateEndpointResponse, error)` |
