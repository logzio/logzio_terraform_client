# Unified Projects

Compatible with Logz.io's [unified projects API](https://api-docs.logz.io/docs/logz/get-dashboard-folder-by-name).

Provides endpoints for managing projects (folders) in the unified dashboard system, including CRUD operations, search, and project management features.

## Usage

```go
client, _ := unified_projects.New(apiToken, baseUrl)

// List all projects
projects, err := client.ListProjects(false)

// List projects with dashboard information
projects, err := client.ListProjects(true)

// Get a specific project by name
project, err := client.GetProject("system-metrics")

// Create a new project
result, err := client.CreateProject(unified_projects.CreateProjectRequest{
    Name: "new-project",
})

// Update a project
updatedProject, err := client.UpdateProject("system-metrics", unified_projects.UpdateProjectRequest{
    DisplayName: "System Metrics Updated",
    Description: "Updated description",
})

// Search projects
searchResults, err := client.SearchProjects(unified_projects.SearchProjectsRequest{
    Query: "system",
    Limit: 10,
    Page:  1,
})

// Delete a project
err = client.DeleteProject("project-id")
```

## Functions

| Function | Signature |
|----------|-----------|
| list | `func (c *ProjectsClient) ListProjects(withDashboards bool) ([]ProjectModel, error)` |
| get | `func (c *ProjectsClient) GetProject(name string) (*ProjectSummary, error)` |
| create | `func (c *ProjectsClient) CreateProject(req CreateProjectRequest) (*ProjectSummary, error)` |
| update | `func (c *ProjectsClient) UpdateProject(name string, req UpdateProjectRequest) (*ProjectSummary, error)` |
| search | `func (c *ProjectsClient) SearchProjects(req SearchProjectsRequest) ([]ProjectSummary, error)` |
| delete | `func (c *ProjectsClient) DeleteProject(folderId string) error` |

## Data Types

### Request Types

- `CreateProjectRequest` - Request payload for creating a project
- `UpdateProjectRequest` - Display name and/or description to update on a project
- `SearchProjectsRequest` - Search parameters for project queries (encoded as query parameters; the search endpoint is a GET)

### Response Types

- `ProjectModel` - Wrapper containing project information
- `ProjectSummary` - Basic project information
- `DashboardListItem` - Dashboard reference in project listings

`CreateProjectRequest`, `UpdateProjectRequest`, and the response types include proper JSON tags with `omitempty` for optional fields. `SearchProjectsRequest` has no JSON tags since it is encoded as query parameters rather than a request body.