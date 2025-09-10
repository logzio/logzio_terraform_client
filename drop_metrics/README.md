# Drop Metrics Objects
Compatible with Logz.io's [Metrics Management API](https://docs.logz.io/api/).

Manages drop metric filters for your Logz.io account.

## Usage
```go
client, _ := drop_metrics.New(apiToken, baseUrl)

active := true
result, err := client.CreateDropMetric(drop_metrics.CreateUpdateDropMetric{
    AccountId: 1234,
    Name:      "my-cpu-filter",
    Active:    &active,
    DropPolicy: drop_metrics.DropPolicyBeforeProcessing, // optional, defaults to DROP_BEFORE_PROCESSING
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


// Update a filter
active := true
updateResult, err := client.UpdateDropMetric(filterId, drop_metrics.CreateUpdateDropMetric{
    AccountId: 1234,
    Name:      "updated-cpu-filter", // Optional: update the name
    Active:    &active,
    DropPolicy: drop_metrics.DropPolicyBeforeStoring,
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

// Search for filters by name
searchResults, err := client.SearchDropMetrics(drop_metrics.SearchDropMetricsRequest{
    Filter: &drop_metrics.SearchFilter{
        AccountIds: []int64{1234},
        SearchTerm: "cpu", // Optional: search filters by name
    },
    Pagination: &drop_metrics.Pagination{
        PageNumber: 0,
        PageSize:   10,
    },
})
```

## Supported Operators

Currently supported comparison operators:
- `EQ` - Equal comparison
- `NOT_EQ` - Not equal comparison
- `REGEX_MATCH` - Regular expression match
- `REGEX_NO_MATCH` - Regular expression no match

## Drop Policies

Drop policies determine when metrics are dropped:
- `DROP_BEFORE_PROCESSING` (default) - drops metrics before processing
- `DROP_BEFORE_STORING` - drops metrics after processing, before storing in Thanos

## API Functions

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