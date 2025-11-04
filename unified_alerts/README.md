# Unified Alerts

Compatible with Logz.io's unified alerts API (POC).

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

    logAlert := unified_alerts.CreateUnifiedAlert{
        Title:       "High Error Rate",
        Type:        unified_alerts.TypeLogAlert,
        Description: "Alert when error rate is too high",
        Tags:        []string{"production", "errors"},
        FolderId:    "folder-123",
        LogAlert: &unified_alerts.LogAlertConfig{
            Output: unified_alerts.LogAlertOutput{
                Recipients: unified_alerts.Recipients{
                    Emails:                  []string{"alerts@example.com"},
                    NotificationEndpointIds: []int{1, 2},
                },
                SuppressNotificationsMinutes: 5,
                Type:                         unified_alerts.OutputTypeJson,
            },
            SearchTimeFrameMinutes: 15,
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
            Schedule: unified_alerts.Schedule{
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
metricAlert := unified_alerts.CreateUnifiedAlert{
    Title:       "High CPU Usage",
    Type:        unified_alerts.TypeMetricAlert,
    Description: "Alert when CPU usage exceeds threshold",
    Tags:        []string{"infrastructure", "cpu"},
    FolderId:    "folder-456",
    MetricAlert: &unified_alerts.MetricAlertConfig{
        Severity: unified_alerts.SeverityHigh,
        Trigger: unified_alerts.MetricTrigger{
            TriggerType:            unified_alerts.TriggerTypeThreshold,
            MetricOperator:         unified_alerts.MetricOperatorAbove,
            MinThreshold:           80.0,
            SearchTimeFrameMinutes: 5,
        },
        Queries: []unified_alerts.MetricQuery{
            {
                RefId: "A",
                QueryDefinition: unified_alerts.MetricQueryDefinition{
                    DatasourceUid: "prometheus-uid",
                    PromqlQuery:   "avg(cpu_usage_percent)",
                },
            },
        },
        Recipients: unified_alerts.Recipients{
            Emails:                  []string{"ops@example.com"},
            NotificationEndpointIds: []int{3, 4},
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
    Type:        unified_alerts.TypeLogAlert,
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
- Timestamps (`updatedAt`, `createdAt`) are returned as `float64` (Unix timestamp with milliseconds as decimal)
- The API endpoint is `/poc/unified-alerts` (POC endpoint)
- All operations include comprehensive validation of required fields and enum values
- The `logAlert` fields follow the same structure as the existing alerts v2 API
- Valid values for `logAlert.output.type`: "JSON", "TABLE"
- Valid values for `metricAlert.trigger.triggerType`: "THRESHOLD", "MATH_EXPRESSION"
- Valid values for `metricAlert.trigger.metricOperator`: "ABOVE", "BELOW", "WITHIN_RANGE", "OUTSIDE_RANGE"

## Important Validation Rules

### Log Alert Query Definition
- If `shouldQueryOnAllAccounts` is `false`, then `accountIdsToQueryOn` must be a non-empty list
- If `shouldQueryOnAllAccounts` is `true`, then `accountIdsToQueryOn` can be empty or omitted


