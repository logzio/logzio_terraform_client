# Unified Dashboards

Compatible with Logz.io's [unified dashboards API](https://api-docs.logz.io/docs/logz/create-a-new-dashboard).

Provides endpoints for managing unified dashboards, including CRUD operations.

## Usage

```go
client, _ := unified_dashboards.New(apiToken, baseUrl)

// List all dashboards
dashboards, err := client.ListDashboards()

// Get a specific dashboard by folder and UID
dashboard, err := client.GetDashboard("project-1", "dashboard-1")

// Create a new dashboard
result, err := client.CreateDashboard("project-1", unified_dashboards.CreateDashboardRequest{
    Doc: map[string]interface{}{
        "title": "CPU Usage Dashboard",
        "panels": []interface{}{
            map[string]interface{}{
                "id":    1,
                "title": "CPU Usage",
                "type":  "graph",
            },
        },
    },
})

// Update a dashboard
updatedDashboard, err := client.UpdateDashboard("project-1", "dashboard-1", unified_dashboards.UpdateDashboardRequest{
    Doc: map[string]interface{}{
        "title":       "Updated Dashboard",
        "description": "Updated description",
    },
})

// Move a dashboard to a different folder
movedDashboard, err := client.MoveDashboard(unified_dashboards.MoveDashboardRequest{
    Uid:            "dashboard-1",
    TargetFolderId: "project-2",
})

// Delete a dashboard
err = client.DeleteDashboard("project-1", "dashboard-1")
```

## Functions

| Function | Signature |
|----------|-----------|
| list | `func (c *DashboardsClient) ListDashboards() ([]Dashboard, error)` |
| get | `func (c *DashboardsClient) GetDashboard(folderId, uid string) (*Dashboard, error)` |
| create | `func (c *DashboardsClient) CreateDashboard(folderId string, req CreateDashboardRequest) (*Dashboard, error)` |
| update | `func (c *DashboardsClient) UpdateDashboard(folderId, uid string, req UpdateDashboardRequest) (*Dashboard, error)` |
| move | `func (c *DashboardsClient) MoveDashboard(req MoveDashboardRequest) (*MoveDashboardResponse, error)` |
| delete | `func (c *DashboardsClient) DeleteDashboard(folderId, uid string) error` |

## Data Types

### Request Types

- `CreateDashboardRequest` - Request payload for creating a dashboard
- `UpdateDashboardRequest` - Request payload for updating a dashboard
- `MoveDashboardRequest` - Request payload for moving a dashboard to a different folder

### Response Types

- `Dashboard` - Dashboard definition with metadata
- `MoveDashboardResponse` - Response from moving a dashboard

All request and response types include proper JSON tags with `omitempty` for optional fields.

## Notes

- The `Doc` field in dashboard requests/responses uses `map[string]interface{}` to support flexible panel configurations 