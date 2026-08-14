# Unified Dashboards

Compatible with Logz.io's unified dashboards API (`/perses-public/api/v1`).

Provides endpoints for managing unified (Perses-based) dashboards: CRUD operations plus move-between-folders.

The wire contract below was verified against the live api.logz.io gateway on 2026-08-14.

## The dashboard document

`Doc` is the full **Perses Dashboard document** — the API requires the envelope and rejects bare field maps with "Dashboard name is required":

```go
doc := map[string]interface{}{
    "kind":     "Dashboard",
    "metadata": map[string]interface{}{"name": "cpu-usage"},
    "spec": map[string]interface{}{
        "display":  map[string]interface{}{"name": "CPU Usage"},
        "duration": "1h",
        "panels":   map[string]interface{}{},
        "layouts":  []interface{}{},
    },
}
```

`folderId` is the containing project's **`id`** (see the unified_projects package), not its name.

## Usage

```go
client, _ := unified_dashboards.New(apiToken, baseUrl)

// Create a dashboard in a folder
created, err := client.CreateDashboard(folderId, unified_dashboards.CreateDashboardRequest{Doc: doc})

// Get a dashboard by folder id and uid
dashboard, err := client.GetDashboard(folderId, created.Uid)

// List all dashboards in the account
dashboards, err := client.ListDashboards()

// Update a dashboard (the PUT replaces the document; bumps Version)
updated, err := client.UpdateDashboard(folderId, created.Uid, unified_dashboards.UpdateDashboardRequest{Doc: doc})

// Move a dashboard to a different folder
moved, err := client.MoveDashboard(unified_dashboards.MoveDashboardRequest{
    Uid:            created.Uid,
    TargetFolderId: otherFolderId,
})

// Delete a dashboard
err = client.DeleteDashboard(folderId, created.Uid)
```

## Functions

| Function | Signature |
|----------|-----------|
| create | `func (c *DashboardsClient) CreateDashboard(folderId string, req CreateDashboardRequest) (*Dashboard, error)` |
| get | `func (c *DashboardsClient) GetDashboard(folderId, uid string) (*Dashboard, error)` |
| list | `func (c *DashboardsClient) ListDashboards() ([]Dashboard, error)` |
| update | `func (c *DashboardsClient) UpdateDashboard(folderId, uid string, req UpdateDashboardRequest) (*Dashboard, error)` |
| move | `func (c *DashboardsClient) MoveDashboard(req MoveDashboardRequest) (*MoveDashboardResponse, error)` |
| delete | `func (c *DashboardsClient) DeleteDashboard(folderId, uid string) error` |

## Data Types

### Request Types

- `CreateDashboardRequest` / `UpdateDashboardRequest` — `Doc` (the Perses Dashboard document, required)
- `MoveDashboardRequest` — `Uid`, `TargetFolderId` (both required)

### Response Types

- `Dashboard` — `Id`, `Uid` (the stable identifier), `Name`, `ProjectId`, `Doc`, `Version`, `CreatedAt`, `UpdatedAt`, `IsPrivate`
- `MoveDashboardResponse` — `Uid`

## Notes

- `Uid` is the stable dashboard identifier; `Id` is a version-row id that changes on update.
- The move endpoint (`POST /dashboards/move`) is documented but was not yet deployed on the public gateway as of 2026-08-14; `MoveDashboard` returns a 404 error until it ships.
- Responses carry additional server-side fields (numeric `createdBy`/`updatedBy`, `isDeleted`) that this client deliberately does not map.
