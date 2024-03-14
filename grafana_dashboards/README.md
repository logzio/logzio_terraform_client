# Grafana dashboards

Compatible with Logz.io's [grafana dashboards API](https://api-docs.logz.io/docs/logz/create-dashboard).
To create a Grafana Dashboard:

```go
client, err := grafana_objects.New(apiToken, apiServerAddress)
dashboardConfig := map[string]interface{}{
        "title":  "dashboard_test",
        "tags":   []string{"some", "tags"},
        "uid":    "",
        "panels": make([]interface{}, 0),
    }
    
request := grafana_dashboards.CreateUpdatePayload{
        Dashboard: dashboardConfig,
        FolderId:  1,
        Message:   "some message",
        Overwrite: true,
    }
dashboard, err := client.CreateUpdate(request)
```

| function                | func name                                                                                                                |
|-------------------------|--------------------------------------------------------------------------------------------------------------------------|
| create/update dashboard | `func (c *GrafanaObjectsClient) CreateUpdateGrafanaDashboard(payload CreateUpdatePayload) (*CreateUpdateResults, error)` |
| delete dashboard by uid | `func (c *GrafanaObjectsClient) DeleteGrafanaDashboard(objectUid string) (*DeleteResults, error)`                        |
| get dashboard by uid    | `func (c *GrafanaObjectsClient) GetGrafanaDashboard(objectUid string) (*GetResults, error)`                              |