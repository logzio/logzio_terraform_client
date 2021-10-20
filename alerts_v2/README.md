# Alerts V2
Compatible with Logz.io's [alerts API](https://docs.logz.io/api/#tag/Alerts).

Logz.io alerts use a Kibana search query to continuously scan your logs and alert you when a certain set of conditions is met. The simplest alerts can use a simple search query or a particular filter, but others can be quite complex and involve several conditions with varying thresholds.

To create an alert where the type field = 'mytype' and the loglevel field = ERROR, see the logz.io docs for more info

https://support.logz.io/hc/en-us/articles/209487329-How-do-I-create-an-Alert-

```go
client, _ := alerts_v2.New(apiToken, apiServerAddress)
alertQuery := alerts_v2.AlertQuery{
		Query:                    "loglevel:ERROR",
		Aggregation:              alerts_v2.AggregationObj{AggregationType: alerts_v2.AggregationTypeCount},
		ShouldQueryOnAllAccounts: true,
	}

	trigger := alerts_v2.AlertTrigger{
		Operator:               alerts_v2.OperatorEquals,
		SeverityThresholdTiers: map[string]float32{alerts_v2.SeverityHigh: 10, alerts_v2.SeverityInfo: 5},
	}
	
	subComponent := alerts_v2.SubAlert{
		QueryDefinition: alertQuery,
		Trigger:         trigger,
		Output:          alerts_v2.SubAlertOutput{},
	}

	createAlertType := alerts_v2.CreateAlertType{
		Title:                  "test create alert",
		Description:            "this is my description",
		Tags:                   []string{"some", "words"},
		Output:                 alerts_v2.AlertOutput{},
		SubComponents:          []alerts_v2.SubAlert{subComponent},
		Correlations:           alerts_v2.SubAlertCorrelation{},
		Enabled:                strconv.FormatBool(true),
	}

alert := client.CreateAlert(createAlertType)
```

|function|func name|
|---|---|
| Create alert | `func (c *AlertsV2Client) CreateAlert(alert CreateAlertType) (*AlertType, error)` |
| Delete alert | `func (c *AlertsV2Client) DeleteAlert(alertId int64) error` |
| Disable alert | `func (c *AlertsV2Client) DisableAlert(alert AlertType) (*AlertType, error)` |
| Enable alert | `func (c *AlertsV2Client) EnableAlert(alert AlertType) (*AlertType, error)` |
| Get alert | `func (c *AlertsV2Client) GetAlert(alertId int64) (*AlertType, error)` |
| List alerts | `func (c *AlertsV2Client) ListAlerts() ([]AlertType, error)` |
| Update alert | `func (c *AlertsV2Client) UpdateAlert(alertId int64, alert CreateAlertType) (*AlertType, error)` |
