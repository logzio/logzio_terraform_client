# Grafana dashboards

Compatible with Logz.io's [grafana dashboards API](https://docs.logz.io/api/#operation/createDashboard).
To create a Grafana Dashboard:

```go
client, err := grafana_objects.New(apiToken, apiServerAddress)
dashboardReq := grafana_objects.DashboardObject{
                Title:  "dashboard_test",
                Tags:   []string{"some", "tags"},
                Panels: make([]map[string]interface{}, 0),
            }

request := grafana_objects.CreateUpdatePayload{
                Dashboard: dashboardReq,
                FolderId:  1,
                Message:   "some message",
                Overwrite: true,
            }
dashboard, err := client.CreateUpdate(request)
```

| function                | func name                                                                                                |
|-------------------------|----------------------------------------------------------------------------------------------------------|
| create/update dashboard | `func (c *GrafanaObjectsClient) CreateUpdate(payload CreateUpdatePayload) (*CreateUpdateResults, error)` |
| delete dashboard by uid | `func (c *GrafanaObjectsClient) Delete(objectUid string) (*DeleteResults, error)`                        |
| get dashboard by uid    | `func (c *GrafanaObjectsClient) Get(objectUid string) (*GetResults, error)`                              |