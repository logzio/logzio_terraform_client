# Metrics Roll-Up Rules

Compatible with Logz.io's metrics roll-up rules API.

Manage metrics rollup rules for your Logz.io account.

## Usage

```go
client, _ := metrics_rollup_rules.New(apiToken, baseUrl)

// Create a GAUGE rollup rule (requires rollupFunction)
result, err := client.CreateRollupRule(metrics_rollup_rules.CreateUpdateRollupRule{
    AccountId:               1234,
    Name:                    "my-cpu-rollup",  (no length limit)
    MetricName:              "cpu_usage",
    MetricType:              metrics_rollup_rules.MetricTypeGauge,
    RollupFunction:          metrics_rollup_rules.AggLast, 
    LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
    Labels:                  []string{"label1"},
})

counterResult, err := client.CreateRollupRule(metrics_rollup_rules.CreateUpdateRollupRule{
    AccountId:               1234,
    Name:                    "my-counter-rollup",
    MetricName:              "request_count",
    MetricType:              metrics_rollup_rules.MetricTypeCounter,
    LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
    Labels:                  []string{"endpoint"},
})

// Update a rollup rule
updateResult, err := client.UpdateRollupRule(ruleId, metrics_rollup_rules.CreateUpdateRollupRule{
    Name:                    "updated-cpu-rollup",
    MetricName:              "cpu_usage_updated",
    MetricType:              metrics_rollup_rules.MetricTypeGauge,
    RollupFunction:          metrics_rollup_rules.AggMax,
    LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
    Labels:                  []string{"label1", "label2"},
})

// Search for rollup rules by name
searchResults, err := client.SearchRollupRules(metrics_rollup_rules.SearchRollupRulesRequest{
    Filter: &metrics_rollup_rules.SearchFilter{
        AccountIds: []int64{1234},
        SearchTerm: "cpu", // Optional: search rules by name
    },
    Pagination: &metrics_rollup_rules.Pagination{
        PageNumber: 0,
        PageSize:   10,
    },
})
```

> [!NOTE]
> Supported metric types include: GAUGE, COUNTER, DELTA_COUNTER, CUMULATIVE_COUNTER, and MEASUREMENT.

> [!NOTE]
> The `rollupFunction` field is only supported for GAUGE and MEASUREMENT metric types. For other metric types (COUNTER, DELTA_COUNTER, CUMULATIVE_COUNTER), the rollupFunction should be omitted.

> [!NOTE]
> For MEASUREMENT metric type, only the following aggregation functions are supported: SUM, MIN, MAX, COUNT, SUMSQ, MEAN, and LAST.

> [!NOTE]
> The `update` method does not support updating the `AccountId` and `metricName` fields. If you need to change either, you must delete the existing rule and create a new one.


| Function | Signature |
|----|-----|
| create | `func (c *MetricsRollupRulesClient) CreateRollupRule(req CreateUpdateRollupRule) (*RollupRule, error)` |
| get | `func (c *MetricsRollupRulesClient) GetRollupRule(rollupRuleId string) (*RollupRule, error)` |
| update | `func (c *MetricsRollupRulesClient) UpdateRollupRule(rollupRuleId string, req CreateUpdateRollupRule) (*RollupRule, error)` |
| delete | `func (c *MetricsRollupRulesClient) DeleteRollupRule(rollupRuleId string) error` |
| bulk create | `func (c *MetricsRollupRulesClient) BulkCreateRollupRules(req []CreateUpdateRollupRule) ([]RollupRule, error)` |
| bulk delete | `func (c *MetricsRollupRulesClient) BulkDeleteRollupRules(ruleIds []string) error` |
| search | `func (c *MetricsRollupRulesClient) SearchRollupRules(req SearchRollupRulesRequest) ([]RollupRule, error)` |
