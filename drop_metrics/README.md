# Drop Metrics Objects
Compatible with Logz.io's [Metrics Management API](https://docs.logz.io/api/#tag/Drop-Filters-For-Metrics).

Manages drop metric filters for your Logz.io account.

## Usage
```go
client, _ := drop_metrics.New(apiToken, baseUrl)

active := true
result, err := client.CreateDropMetric(drop_metrics.CreateDropMetric{
    AccountId: 1234,
    Active:    &active,
    Filter: drop_metrics.FilterObject{
        Operator: drop_metrics.OperatorAnd,
        Expression: []drop_metrics.FilterExpression{
            {
                Name:             "__name__",
                Value:            "CpuUsage", 
                ComparisonFilter: drop_metrics.ComparisonEq,
            },
        },
    },
})

// Bulk create multiple filters
active := true
bulkResult, err := client.BulkCreateDropMetrics([]drop_metrics.CreateDropMetric{
    {
        AccountId: 1234,
        Active:    &active,
        Filter: drop_metrics.FilterObject{
            Operator: drop_metrics.OperatorAnd,
            Expression: []drop_metrics.FilterExpression{
                {
                    Name:             "__name__",
                    Value:            "MemoryUsage",
                    ComparisonFilter: drop_metrics.ComparisonEq,
                },
                {
                    Name:             "environment", 
                    Value:            "production",
                    ComparisonFilter: drop_metrics.ComparisonEq,
                },
            },
        },
    },
})

// Update a filter
active := true
updateResult, err := client.UpdateDropMetric(filterId, drop_metrics.UpdateDropMetric{
    AccountId: 1234,
    Active:    &active,
    Filter: drop_metrics.FilterObject{
        Operator: drop_metrics.OperatorAnd,
        Expression: []drop_metrics.FilterExpression{
            {
                Name:             "__name__",
                Value:            "UpdatedMetricName",
                ComparisonFilter: drop_metrics.ComparisonEq,
            },
        },
    },
})
```

## Supported Operators

Currently supported comparison operators:
- `EQ` - Equal comparison

Additional operators (`NOT_EQ`, `REGEX_MATCH`, `REGEX_NO_MATCH`) will be available in future API versions.

| Function | Signature |
|----------|-----------|
| create drop metric | `func (c *DropMetricsClient) CreateDropMetric(req CreateDropMetric) (*DropMetric, error)` |
| bulk create drop metrics | `func (c *DropMetricsClient) BulkCreateDropMetrics(requests []CreateDropMetric) ([]DropMetric, error)` |
| get drop metric | `func (c *DropMetricsClient) GetDropMetric(dropFilterId int64) (*DropMetric, error)` |
| update drop metric | `func (c *DropMetricsClient) UpdateDropMetric(dropFilterId int64, req UpdateDropMetric) (*DropMetric, error)` |
| search drop metrics | `func (c *DropMetricsClient) SearchDropMetrics(req SearchDropMetricsRequest) ([]DropMetric, error)` |
| enable drop metric | `func (c *DropMetricsClient) EnableDropMetric(dropFilterId int64) error` |
| disable drop metric | `func (c *DropMetricsClient) DisableDropMetric(dropFilterId int64) error` |
| delete drop metric | `func (c *DropMetricsClient) DeleteDropMetric(dropFilterId int64) error` |
| bulk delete drop metrics | `func (c *DropMetricsClient) BulkDeleteDropMetrics(dropFilterIds []int64) error` | 