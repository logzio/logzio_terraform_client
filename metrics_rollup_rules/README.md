# Metrics Roll-Up Rules

Compatible with Logz.io's metrics roll-up rules API.

Manage metrics rollup rules for you Logz.io account.

## Usage

```go
client, _ := metrics_rollup_rules.New(apiToken, baseUrl)
result, err := client.CreateRollupRule(metrics_rollup_rules.CreateUpdateRollupRule{
    AccountId:               1234,
    MetricName:              "cpu_usage",
    MetricType:              metrics_rollup_rules.MetricTypeGauge,
    RollupFunction:          metrics_rollup_rules.AggLast,
    LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
    Labels:                  []string{"label1"},
})
```

> [!NOTE]
> The `update` method does not support updating the `AccountId` and `metricName` fields. If you need to change either, you must delete the existing rule and create a new one.

| Function | Signature |
|----|-----|
| create | `func (c *MetricsRollupRulesClient) CreateRollupRule(req CreateUpdateRollupRule) (*RollupRule, error)` |
| get | `func (c *MetricsRollupRulesClient) GetRollupRule(rollupRuleId string) (*RollupRule, error)` |
| update | `func (c *MetricsRollupRulesClient) UpdateRollupRule(rollupRuleId string, req CreateUpdateRollupRule) (*RollupRule, error)` |
| delete | `func (c *MetricsRollupRulesClient) DeleteRollupRule(rollupRuleId string) error` |
| bulk create | `func (c *MetricsRollupRulesClient) BulkCreateRollupRules(req []CreateRollupRule) ([]RollupRule, error)` |
| bulk delete | `func (c *MetricsRollupRulesClient) BulkDeleteRollupRules(ruleIds []string) error` |
| search | `func (c *MetricsRollupRulesClient) SearchRollupRules(req SearchRollupRulesRequest) ([]RollupRule, error)` |
