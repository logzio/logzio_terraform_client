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

type Project struct {
	Kind     string          `json:"kind"`
	Metadata ProjectMetadata `json:"metadata"`
	Spec     ProjectSpec     `json:"spec"`
}

type ProjectMetadata struct {
	Name      string `json:"name"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Version   int    `json:"version,omitempty"`
}

type ProjectSpec struct {
	Display ProjectDisplay `json:"display"`
}

type ProjectDisplay struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type SearchProjectsRequest struct {
	Query string `json:"query,omitempty"`
	Limit int    `json:"limit,omitempty"`
	Page  int    `json:"page,omitempty"`
	Sort  string `json:"sort,omitempty"`
	Order string `json:"order,omitempty"`
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

type SearchProjectsResponse struct {
	Items []ProjectModel `json:"items"`
	Total int            `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
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

func validateUpdateProjectRequest(project Project) error {
	if len(project.Kind) == 0 {
		return fmt.Errorf("kind must be set")
	}
	if len(project.Metadata.Name) == 0 {
		return fmt.Errorf("metadata.name must be set")
	}
	if len(project.Spec.Display.Name) == 0 {
		return fmt.Errorf("spec.display.name must be set")
	}
	return nil
}

func validateSearchProjectsRequest(req SearchProjectsRequest) error {
	return nil
}

func validateDeleteProjectRequest(folderId string) error {
	if len(folderId) == 0 {
		return fmt.Errorf("folderId must be set")
	}
	return nil
}
