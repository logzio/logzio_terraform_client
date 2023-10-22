# Grafana Alert

To create a Grafana alert:

```go
data := grafana_alerts.GrafanaAlertQuery{
    DatasourceUid: "__expr__",
    Model:         json.RawMessage(`{"conditions":[{"evaluator":{"params":[0,0],"type":"gt"},"operator":{"type":"and"},"query":{"params":[]},"reducer":{"params":[],"type":"avg"},"type":"query"}],"datasource":{"type":"__expr__","uid":"__expr__"},"expression":"1 == 1","hide":false,"intervalMs":1000,"maxDataPoints":43200,"refId":"A","type":"math"}`),
    RefId:         "A",
    RelativeTimeRange: grafana_alerts.RelativeTimeRangeObj{
    From: 0,
    To:   0,
    },
}

createGrafanaAlert := grafana_alerts.GrafanaAlertRule{
    Annotations:  map[string]string{"key_test": "value_test"},
    Condition:    "A",
    Data:         []*grafana_alerts.GrafanaAlertQuery{&data},
    FolderUID:    os.Getenv(envGrafanaFolderUid),
    NoDataState:  grafana_alerts.NoDataOk,
    ExecErrState: grafana_alerts.ErrOK,
    OrgID:        1,
    RuleGroup:    "rule_group_1",
    Title:        "test_alert",
    For:          int64(3),
}

client, err := grafana_alerts.New(apiToken, server.URL)
grafanaAlert, err := client.CreateGrafanaAlertRule(createGrafanaAlert)
```

| function            | func name                                                                                                  |
|---------------------|------------------------------------------------------------------------------------------------------------|
| create alert        | `func (c *GrafanaAlertClient) CreateGrafanaAlertRule(payload GrafanaAlertRule) (*GrafanaAlertRule, error)` |
| update alert        | `func (c *GrafanaAlertClient) UpdateGrafanaAlertRule(payload GrafanaAlertRule) error`                      |
| delete alert by uid | `func (c *GrafanaAlertClient) DeleteGrafanaAlertRule(uid string) error`                                    |
| get alert by uid    | `func (c *GrafanaAlertClient) GetGrafanaAlertRule(uid string) (*GrafanaAlertRule, error)`                  |
| list alerts         | `func (c *GrafanaAlertClient) ListGrafanaAlertRules() ([]GrafanaAlertRule, error)`                         |
