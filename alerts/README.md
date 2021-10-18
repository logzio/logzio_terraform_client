# Alerts

### Deprecation note:
This version of the Alerts API is deprecated. Use [Alerts V2] instead.

To create an alert where the type field = 'mytype' and the loglevel field = ERROR, see the logz.io docs for more info

https://support.logz.io/hc/en-us/articles/209487329-How-do-I-create-an-Alert-

```go
client, _ := alerts.New(apiToken, apiServerAddress)
alert := client.CreateAlert(alerts.CreateAlertType{
    Title:       "this is my alert",
    Description: "this is my description",
    QueryString: "loglevel:ERROR",
    Filter:      "{\"bool\":{\"must\":[{\"match\":{\"type\":\"mytype\"}}],\"must_not\":[]}}",
    Operation:   alerts.OperatorGreaterThan,
    SeverityThresholdTiers: []alerts.SeverityThresholdType{
        alerts.SeverityThresholdType{
            alerts.SeverityHigh,
            10,
        },
    },
    SearchTimeFrameMinutes:       0,
    NotificationEmails:           []interface{}{},
    IsEnabled:                    true,
    SuppressNotificationsMinutes: 0,
    ValueAggregationType:         alerts.AggregationTypeCount,
    ValueAggregationField:        nil,
    GroupByAggregationFields:     []interface{}{"my_field"},
    AlertNotificationEndpoints:   []interface{}{},
})
```

|function|func name|
|---|---|
|create alert|`func (c *AlertsClient) CreateAlert(alert CreateAlertType) (*AlertType, error)`|
|update alert|`func (c *AlertsClient) UpdateAlert(alertId int64, alert CreateAlertType) (*AlertType, error)`
|delete alert|`func (c *AlertsClient) DeleteAlert(alertId int64) error`|
|get alert (by id)|`func (c *AlertsClient) GetAlert(alertId int64) (*AlertType, error)`|
|list alerts|`func (c *AlertsClient) ListAlerts() ([]AlertType, error)`|
