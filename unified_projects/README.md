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
updatedProject, err := client.UpdateProject("system-metrics", unified_projects.Project{
    Kind: "Project",
    Metadata: unified_projects.ProjectMetadata{
        Name: "system-metrics",
    },
    Spec: unified_projects.ProjectSpec{
        Display: unified_projects.ProjectDisplay{
            Name:        "System Metrics Updated",
            Description: "Updated description",
        },
    },
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
| update | `func (c *ProjectsClient) UpdateProject(name string, project Project) (*Project, error)` |
| search | `func (c *ProjectsClient) SearchProjects(req SearchProjectsRequest) (*SearchProjectsResponse, error)` |
| delete | `func (c *ProjectsClient) DeleteProject(folderId string) error` |

## Data Types

### Request Types

- `CreateProjectRequest` - Request payload for creating a project
- `Project` - Full project specification for updates
- `SearchProjectsRequest` - Search parameters for project queries

### Response Types

- `ProjectModel` - Wrapper containing project information
- `ProjectSummary` - Basic project information
- `SearchProjectsResponse` - Paginated search results
- `DashboardListItem` - Dashboard reference in project listings

All request and response types include proper JSON tags with `omitempty` for optional fields. 