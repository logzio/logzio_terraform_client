# Grafana Notification Policy

To set a Grafana notification policy tree:

```go
setupGrafanaNotificationPolicy := grafana_notification_policies.GrafanaNotificationPolicyTree{
            GroupBy:        []string{"hello-world", "alertname"},
            GroupInterval:  "5m",
            GroupWait:      "10s",
            Receiver:       grafanaDefaultReceiver,
            RepeatInterval: "5h",
            Routes: []grafana_notification_policies.GrafanaNotificationPolicy{
            {
            Receiver:       grafanaDefaultReceiver,
            ObjectMatchers: grafana_notification_policies.MatchersObj{grafana_notification_policies.MatcherObj{"hello", "=", "darkness"}},
            Continue:       true,
            },
            {
            Receiver:       grafanaDefaultReceiver,
                    ObjectMatchers: grafana_notification_policies.MatchersObj{grafana_notification_policies.MatcherObj{"my", "=~", "oldfriend.*"}},
                    Continue:       false,
                },
            },
        }
client, err := grafana_alerts.New(apiToken, server.URL)
err = client.SetupGrafanaNotificationPolicyTree(setupGrafanaNotificationPolicy)
```

| function                               | func name                                                                                                                   |
|----------------------------------------|-----------------------------------------------------------------------------------------------------------------------------|
| set grafana notification policy tree   | `func (c *GrafanaNotificationPolicyClient) SetupGrafanaNotificationPolicyTree(payload GrafanaNotificationPolicyTree) error` |
| get grafana notification policy tree   | `func (c *GrafanaNotificationPolicyClient) GetGrafanaNotificationPolicyTree() (GrafanaNotificationPolicyTree, error)`       |
| reset grafana notification policy tree | `func (c *GrafanaNotificationPolicyClient) ResetGrafanaNotificationPolicyTree() error`                                      |
