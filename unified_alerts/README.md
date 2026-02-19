# Unified Alerts

Compatible with Logz.io's unified alerts v2 API (`/v2/unified-alerts`).

This package provides a unified interface for managing both log-based and metric-based alerts through a single API.

## Usage Examples

### Create a Log Alert

```go
package main

import (
    "fmt"
    "github.com/logzio/logzio_terraform_client/unified_alerts"
)

func main() {
    client, err := unified_alerts.New(apiToken, baseUrl)
    if err != nil {
        panic(err)
    }

    enabled := true
    logAlert := unified_alerts.CreateUnifiedAlert{
        Title:       "High Error Rate",
        Description: "Alert when error rate is too high",
        Tags:        []string{"production", "errors"},
        LinkedPanel: &unified_alerts.LinkedPanel{
            FolderId: "folder-123",
        },
        Enabled: &enabled,
        Recipients: &unified_alerts.Recipients{
            Emails:                  []string{"alerts@example.com"},
            NotificationEndpointIds: []int{1, 2},
        },
        AlertConfiguration: &unified_alerts.AlertConfiguration{
            Type:                         unified_alerts.TypeLogAlert,
            SuppressNotificationsMinutes: 5,
            AlertOutputTemplateType:      unified_alerts.OutputTypeJson,
            SearchTimeFrameMinutes:       15,
            SubComponents: []unified_alerts.SubComponent{
                {
                    QueryDefinition: unified_alerts.QueryDefinition{
                        Query: "loglevel:ERROR",
                        Aggregation: unified_alerts.Aggregation{
                            AggregationType: unified_alerts.AggregationTypeCount,
                        },
                        ShouldQueryOnAllAccounts: true,
                    },
                    Trigger: unified_alerts.SubComponentTrigger{
                        Operator: unified_alerts.OperatorGreaterThan,
                        SeverityThresholdTiers: map[string]float32{
                            unified_alerts.SeverityHigh: 100.0,
                            unified_alerts.SeverityInfo: 50.0,
                        },
                    },
                    Output: unified_alerts.SubComponentOutput{
                        ShouldUseAllFields: true,
                    },
                },
            },
            Schedule: &unified_alerts.Schedule{
                CronExpression: "* ",
                Timezone:       "UTC",
            },
        },
    }

    result, err := client.CreateUnifiedAlert(unified_alerts.UrlTypeLogs, logAlert)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Created log alert with ID: %s\n", result.Id)
}
```

### Create a Metric Alert

```go
threshold := 80.0
metricAlert := unified_alerts.CreateUnifiedAlert{
    Title:       "High CPU Usage",
    Description: "Alert when CPU usage exceeds threshold",
    Tags:        []string{"infrastructure", "cpu"},
    LinkedPanel: &unified_alerts.LinkedPanel{
        FolderId:    "folder-456",
        DashboardId: "dashboard-789",
        PanelId:     "panel-012",
    },
    Recipients: &unified_alerts.Recipients{
        Emails:                  []string{"ops@example.com"},
        NotificationEndpointIds: []int{3, 4},
    },
    AlertConfiguration: &unified_alerts.AlertConfiguration{
        Type:     unified_alerts.TypeMetricAlert,
        Severity: unified_alerts.SeverityHigh,
        Trigger: &unified_alerts.MetricAlertTrigger{
            Type: unified_alerts.TriggerTypeThreshold,
            Condition: &unified_alerts.TriggerCondition{
                OperatorType: unified_alerts.OperatorTypeAbove,
                Threshold:    &threshold,
            },
        },
        Queries: []unified_alerts.MetricQuery{
            {
                RefId: "A",
                QueryDefinition: unified_alerts.MetricQueryDefinition{
                    AccountId:   12345,
                    PromqlQuery: "avg(cpu_usage_percent)",
                },
            },
        },
    },
}

result, err := client.CreateUnifiedAlert(unified_alerts.UrlTypeMetrics, metricAlert)
```

### Get an Alert

```go
// For log alerts
alert, err := client.GetUnifiedAlert(unified_alerts.UrlTypeLogs, "alert-id-123")

// For metric alerts
alert, err := client.GetUnifiedAlert(unified_alerts.UrlTypeMetrics, "alert-id-456")
```

### Update an Alert

```go
updatedAlert := unified_alerts.CreateUnifiedAlert{
    Title:       "Updated Alert Title",
    // ... rest of alert configuration
}

result, err := client.UpdateUnifiedAlert(unified_alerts.UrlTypeLogs, "alert-id-123", updatedAlert)
```

### Delete an Alert

```go
// Returns the deleted alert details
deletedAlert, err := client.DeleteUnifiedAlert(unified_alerts.UrlTypeLogs, "alert-id-123")
```

## API Methods

| Function | Signature |
|----------|-----------|
| Create alert | `func (c *UnifiedAlertsClient) CreateUnifiedAlert(alertType string, req CreateUnifiedAlert) (*UnifiedAlert, error)` |
| Get alert | `func (c *UnifiedAlertsClient) GetUnifiedAlert(alertType string, alertId string) (*UnifiedAlert, error)` |
| Update alert | `func (c *UnifiedAlertsClient) UpdateUnifiedAlert(alertType string, alertId string, req CreateUnifiedAlert) (*UnifiedAlert, error)` |
| Delete alert | `func (c *UnifiedAlertsClient) DeleteUnifiedAlert(alertType string, alertId string) (*UnifiedAlert, error)` |

## Notes

- The `alertType` parameter in all operations should be either "logs" or "metrics"
- Alert IDs are strings
- Timestamps (`updatedAt`, `createdAt`) are returned as `int64` (seconds since epoch)
- The API endpoint is `/v2/unified-alerts`
- All operations include comprehensive validation of required fields and enum values
- Alert configuration is unified under `alertConfiguration` with a `type` field (`LOG_ALERT` or `METRIC_ALERT`)
- Recipients are at the top level, not nested inside alert-type-specific config
- Panel references use the `linkedPanel` object instead of flat fields
- Valid values for `alertConfiguration.alertOutputTemplateType`: "JSON", "TABLE"
- Valid values for `alertConfiguration.trigger.type`: "threshold", "math"
- Valid values for `alertConfiguration.trigger.condition.operatorType`: "above", "below", "within_range", "outside_range"

## Important Validation Rules

### Log Alert Query Definition
- If `shouldQueryOnAllAccounts` is `false`, then `accountIdsToQueryOn` must be a non-empty list
- If `shouldQueryOnAllAccounts` is `true`, then `accountIdsToQueryOn` can be empty or omitted

### Metric Alert Trigger
- For `threshold` triggers: `condition` must be set with a valid `operatorType`
- For `math` triggers: `expression` must be non-empty
- `queries` must not be empty for metric alerts
