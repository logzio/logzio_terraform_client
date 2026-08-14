package unified_projects

import (
	"fmt"

	"github.com/logzio/logzio_terraform_client/client"
)

const (
	projectsServiceEndpoint = "%s/perses-public/api/v1/projects"
	projectsByNameEndpoint  = "%s/perses-public/api/v1/projects/%s"
	projectsSearchEndpoint  = "%s/perses-public/api/v1/projects/search"

	projectResourceName = "unified project"

	createProjectOperation  = "CreateUnifiedProject"
	getProjectOperation     = "GetUnifiedProject"
	listProjectsOperation   = "ListUnifiedProjects"
	updateProjectOperation  = "UpdateUnifiedProject"
	searchProjectsOperation = "SearchUnifiedProjects"

	deleteProjectOperation = "DeleteUnifiedProject"
)

type ProjectsClient struct {
	*client.Client
}

// Request types
type CreateProjectRequest struct {
	Name string `json:"name"`
}

type UpdateProjectRequest struct {
	DisplayName string `json:"displayName,omitempty"`
	Description string `json:"description,omitempty"`
}

// SearchProjectsRequest is encoded as query parameters (the search endpoint is a GET).
type SearchProjectsRequest struct {
	Query string // required
	Limit int    // optional, API default 1000
	Page  int    // optional
	Sort  string // optional, e.g. "asc"/"desc"
}

// Response types
type ProjectSummary struct {
	Id          string              `json:"id"`
	Name        string              `json:"name"`
	DisplayName string              `json:"displayName"`
	Description string              `json:"description,omitempty"`
	Dashboards  []DashboardListItem `json:"dashboards,omitempty"`
	CreatedAt   string              `json:"createdAt,omitempty"`
	UpdatedAt   string              `json:"updatedAt,omitempty"`
}

type DashboardListItem struct {
	Uid   string `json:"uid"`
	Title string `json:"title"`
}

type ProjectModel struct {
	Project ProjectSummary `json:"project"`
}

func New(apiToken, baseUrl string) (*ProjectsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}
	return &ProjectsClient{
		Client: client.New(apiToken, baseUrl),
	}, nil
}

// Validation helpers
func validateCreateProjectRequest(req CreateProjectRequest) error {
	if len(req.Name) == 0 {
		return fmt.Errorf("name must be set")
	}
	return nil
}

func validateUpdateProjectRequest(name string, req UpdateProjectRequest) error {
	if len(name) == 0 {
		return fmt.Errorf("name must be set")
	}
	if len(req.DisplayName) == 0 && len(req.Description) == 0 {
		return fmt.Errorf("displayName or description must be set")
	}
	return nil
}

func validateSearchProjectsRequest(req SearchProjectsRequest) error {
	if len(req.Query) == 0 {
		return fmt.Errorf("query must be set")
	}
	return nil
}

func validateDeleteProjectRequest(folderId string) error {
	if len(folderId) == 0 {
		return fmt.Errorf("folderId must be set")
	}
	return nil
}
