# Drop Metrics Objects
Compatible with Logz.io's [drop metrics API](https://api-docs.logz.io/docs/logz/drop-metrics).

Drop metrics provide a solution for filtering out metrics before they are indexed in your account to help lower costs and reduce account volume.

## Usage

```go
client, _ := drop_metrics.New(apiToken, apiServerAddress)

// Create a drop metric filter with new comparison operators
enabled := true
dropMetric, err := client.CreateDropMetric(drop_metrics.CreateDropMetric{
    AccountId: 1234,
    Enabled:   &enabled,
    Filter: drop_metrics.FilterObject{
        Operator: "and",
        Expression: []drop_metrics.FilterExpression{
            {
                Name:             "__name__",
                Value:            "cpu.*",
                ComparisonFilter: "regex_match",  // New regex support
            },
            {
                Name:             "environment",
                Value:            "test",
                ComparisonFilter: "not_eq",       // New not equals operator
            },
        },
    },
})

// Update an existing drop metric
updatedMetric, err := client.UpdateDropMetric(dropMetric.Id, drop_metrics.UpdateDropMetric{
    AccountId: 1234,
    Enabled:   &enabled,
    Filter: drop_metrics.FilterObject{
        Operator: "and",
        Expression: []drop_metrics.FilterExpression{
            {
                Name:             "__name__",
                Value:            "memory.*",
                ComparisonFilter: "regex_no_match", // New regex no match operator
            },
        },
    },
})

// Bulk delete multiple drop metrics
err = client.BulkDeleteDropMetrics([]int64{1, 2, 3})
```

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