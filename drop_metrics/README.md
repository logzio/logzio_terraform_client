# Drop Metrics Objects
Compatible with Logz.io's [drop metrics API](https://api-docs.logz.io/docs/logz/drop-metrics).

Drop metrics provide a solution for filtering out metrics before they are indexed in your account to help lower costs and reduce account volume.

## Usage

```go
client, _ := drop_metrics.New(apiToken, apiServerAddress)

// Create a simple drop metric filter
enabled := true
dropMetric, err := client.CreateDropMetric(drop_metrics.CreateDropMetric{
    AccountId: 1234,
    Enabled:   &enabled,
    Filter: drop_metrics.FilterObject{
        Operator: drop_metrics.OperatorAnd,
        Expression: []drop_metrics.FilterExpression{
            {
                Name:             "__name__",
                Value:            "CpuUsage",
                ComparisonFilter: drop_metrics.ComparisonEq,
            },
            {
                Name:             "service",
                Value:            "metrics-service",
                ComparisonFilter: drop_metrics.ComparisonEq,
            },
        },
    },
})

// Create filter for multiple conditions
dropMetric, err = client.CreateDropMetric(drop_metrics.CreateDropMetric{
    AccountId: 1234,
    Enabled:   &enabled,
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
})

// Search for drop metrics
searchResults, err := client.SearchDropMetrics(drop_metrics.SearchDropMetricsRequest{
    Filter: &drop_metrics.SearchFilter{
        AccountIds: []int64{1234},
        Enabled:    &enabled,
    },
    Pagination: &drop_metrics.Pagination{
        PageNumber: 1,
        PageSize:   10,
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
| delete by search | `func (c *DropMetricsClient) DeleteDropMetricsBySearch(req SearchDropMetricsRequest) error` |
| bulk delete drop metrics | `func (c *DropMetricsClient) BulkDeleteDropMetrics(dropFilterIds []int64) error` | 