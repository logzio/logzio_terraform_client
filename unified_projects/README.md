# Unified Projects (Dashboard Folders)

Compatible with Logz.io's unified dashboards folders API (`/perses-public/api/v1/projects`).

Provides endpoints for managing projects (dashboard folders) in the unified (Perses-based) dashboard system: CRUD operations, search, and rename.

The wire contract below was verified against the live api.logz.io gateway on 2026-08-14; where the published API docs disagree (create/update body shape, addressing, search method), the live behavior documented here wins.

## Addressing

All single-project operations address the project by its **`id`** (the `id` field returned from `CreateProject`, `ListProjects`, `SearchProjects`, or `GetProject`). The project's `metadata.name` is a separate identity used inside the Perses document; passing it as the path parameter yields a 404.

## Usage

```go
client, _ := unified_projects.New(apiToken, baseUrl)

// Create a project (folder). The client builds the Perses Project document for you.
created, err := client.CreateProject(unified_projects.CreateProjectRequest{
    Name:        "my-project",          // identity (metadata.name)
    DisplayName: "My Project",          // optional; defaults to Name
    Description: "team dashboards",     // optional
})

// Get a project by id
project, err := client.GetProject(created.Id)

// List all projects (optionally with each project's dashboards)
items, err := client.ListProjects(true)

// Update a project — the PUT replaces the Perses document, so send the full
// desired state; an empty Description clears any existing description.
updated, err := client.UpdateProject(created.Id, unified_projects.UpdateProjectRequest{
    Name:        "my-project",
    DisplayName: "My Project (renamed)",
    Description: "updated description",
})

// Search: a dashboard search grouped by folder — SearchTerm matches
// dashboards, every folder is returned with its matching dashboards nested,
// and Total counts the matching dashboards.
resp, err := client.SearchProjects(unified_projects.SearchProjectsRequest{
    Filter:     &unified_projects.SearchProjectsFilter{SearchTerm: "cpu"},
    Pagination: &unified_projects.SearchProjectsPagination{PageNumber: 1, PageSize: 100},
})

// Rename a project
renamed, err := client.RenameProject(created.Id, "my-project-renamed")

// Delete a project by id
err = client.DeleteProject(created.Id)
```

## Functions

| Function | Signature |
|----------|-----------|
| create | `func (c *ProjectsClient) CreateProject(req CreateProjectRequest) (*ProjectSummary, error)` |
| get | `func (c *ProjectsClient) GetProject(id string) (*ProjectSummary, error)` |
| list | `func (c *ProjectsClient) ListProjects(withDashboards bool) ([]ProjectListItem, error)` |
| update | `func (c *ProjectsClient) UpdateProject(id string, req UpdateProjectRequest) (*ProjectSummary, error)` |
| search | `func (c *ProjectsClient) SearchProjects(req SearchProjectsRequest) (*SearchProjectsResponse, error)` |
| rename | `func (c *ProjectsClient) RenameProject(id string, newName string) (*ProjectSummary, error)` |
| delete | `func (c *ProjectsClient) DeleteProject(id string) error` |

## Data Types

### Request Types

- `CreateProjectRequest` — `Name` (required), `DisplayName`, `Description`; the client wraps them in the Perses Project envelope the API requires.
- `UpdateProjectRequest` — `Name` (required, the Perses `metadata.name`), `DisplayName` (required), `Description`; the PUT replaces the whole document.
- `SearchProjectsRequest` — `Filter{SearchTerm, CreatedBy}` + `Pagination{PageNumber, PageSize}` (POST body; verified live — the published docs describe a different, non-working GET shape).
- Rename sends `{"newProjectName": ...}` on the wire (the docs' `newName` is silently ignored by the server).

### Response Types

- `ProjectSummary` — `Id`, `Name` (display name), `Doc` (the raw Perses Project document), `CreatedAt`, `UpdatedAt`; `MetadataName()` returns the Perses identity for read-modify-write updates.
- `ProjectListItem` — one list/search entry: `Project` plus its `Dashboards`.
- `SearchProjectsResponse` — `Results`, `Total`, `Pagination`.

## Notes

- Responses carry additional server-side fields (`entityId`, numeric `createdBy`/`updatedBy`, `isDeleted` as either a boolean or 0/1) that this client deliberately does not map.
