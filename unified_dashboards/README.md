# Unified Dashboards

Compatible with Logz.io's unified dashboards API (`/perses-public/api/v1`).

Provides endpoints for managing unified (Perses-based) dashboards: CRUD operations plus move-between-folders.

The wire contract below was verified against the live api.logz.io gateway on 2026-08-14.

## The dashboard document

`Doc` is the full **Perses Dashboard document** — the API requires the envelope and rejects bare field maps with "Dashboard name is required":

```go
doc := map[string]any{
    "kind":     "Dashboard",
    "metadata": map[string]any{"name": "cpu-usage"},
    "spec": map[string]any{
        "display":  map[string]any{"name": "CPU Usage"},
        "duration": "1h",
        "panels":   map[string]any{},
        "layouts":  []any{},
    },
}
```

Note the asymmetry with the projects API: dashboards nest the document under a `"doc"` key, while projects send the bare envelope — both shapes are live-verified.

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

// List just one folder's dashboards
folderDashboards, err := client.ListFolderDashboards(folderId)

// Search dashboards by name, paginated
results, err := client.SearchDashboards(unified_dashboards.SearchDashboardsRequest{
    Filter:     &unified_dashboards.SearchDashboardsFilter{SearchTerm: "CPU"},
    Pagination: &unified_dashboards.SearchDashboardsPagination{PageNumber: 1, PageSize: 50},
})

// Update a dashboard (the PUT replaces the document; bumps Version)
updated, err := client.UpdateDashboard(folderId, created.Uid, unified_dashboards.UpdateDashboardRequest{Doc: doc})

// Move a dashboard to a different folder
moved, err := client.MoveDashboard(unified_dashboards.MoveDashboardRequest{
    DashboardId:  created.Uid,
    OldProjectId: folderId,
    NewProjectId: otherFolderId,
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
| list (one folder) | `func (c *DashboardsClient) ListFolderDashboards(folderId string) ([]Dashboard, error)` |
| search | `func (c *DashboardsClient) SearchDashboards(req SearchDashboardsRequest) (*SearchDashboardsResponse, error)` |
| update | `func (c *DashboardsClient) UpdateDashboard(folderId, uid string, req UpdateDashboardRequest) (*Dashboard, error)` |
| move | `func (c *DashboardsClient) MoveDashboard(req MoveDashboardRequest) (*MoveDashboardResponse, error)` |
| delete | `func (c *DashboardsClient) DeleteDashboard(folderId, uid string) error` |

## Data Types

### Request Types

- `CreateDashboardRequest` / `UpdateDashboardRequest` — `Doc` (the Perses Dashboard document, required)
- `MoveDashboardRequest` — `DashboardId`, `OldProjectId`, `NewProjectId` (all required; PUT — the published docs describe a different, non-working shape)
- `SearchDashboardsRequest` — `Filter{SearchTerm, CreatedBy}` + `Pagination{PageNumber, PageSize}`, both optional (POST only; `GET /dashboards/search` 404s)

### Response Types

- `Dashboard` — `Uid`, `Id`, `Name`, `ProjectId`, `Doc`, `Version`, `CreatedAt`, `UpdatedAt`, `IsPrivate` (see [Identifiers](#identifiers))
- `MoveDashboardResponse` — `Id`, the moved dashboard's **version-row** id (not its uid)
- `SearchDashboardsResponse` — `Results` (flat `[]Dashboard`), `Total` (matches before pagination), `Pagination`

## Identifiers

A dashboard has two identifiers and they are not interchangeable:

| Field | Wire key | Stable? | What it is for |
|-------|----------|---------|----------------|
| `Dashboard.Uid` | `uid` | yes — survives updates | every folder-scoped call: `GetDashboard`, `UpdateDashboard`, `DeleteDashboard`, and `MoveDashboardRequest.DashboardId`. **Persist this one.** |
| `Dashboard.Id` | `id` | no — a new version row on every update | nothing in this client takes it; mapped for completeness |

**The trap:** a brand-new dashboard is created with `id == uid == name`. They only diverge after the first update, so anything that looks solely at freshly created dashboards — provider code, or a test — cannot tell the two apart, and will keep working right up until a user edits a dashboard.

Consequences worth knowing:

- `MoveDashboardResponse.Id` is the current **version-row id**, not the uid. It is an acknowledgement; keep using the uid you passed in as `DashboardId`.
- `unified_projects.ProjectDashboard` carries **both** `Uid` and `Id`, with the same meanings as here. Use its `Uid`.
- Passing an `Id` to any folder-scoped route returns "failed with missing unified dashboard".

All of the above was settled by probing the live gateway with a version-2 dashboard (2026-08-20), not from the published docs. The integration tests now update a dashboard before asserting on its identifiers, so a regression cannot hide behind the version-1 coincidence.

## Notes

- Responses carry additional server-side fields (numeric `createdBy`/`updatedBy`, `isDeleted`) that this client deliberately does not map.
- `ListFolderDashboards` does not 404 for an unknown folder — the gateway answers `200 []`, indistinguishable from an existing empty folder. Check the folder with `unified_projects.GetProject` if you need to tell them apart.
- `SearchDashboards`'s `Filter.SearchTerm` really filters, unlike `unified_projects.SearchProjects`, which returns every folder regardless of the term. `Filter.CreatedBy` is currently ignored by the server.
